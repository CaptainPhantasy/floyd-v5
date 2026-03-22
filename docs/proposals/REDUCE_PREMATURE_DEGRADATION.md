# Fix Proposal: Reduce Premature Context Degradation

## Root Cause Analysis

The model degrades and becomes "unusual" **30% earlier than expected** due to a **cascade of overly aggressive degradation thresholds**. Even with token accumulation fixed, the context management system truncates and compresses context too early.

## Current Degradation Cascade

```
┌─────────────────────────────────────────────────────────────────────────────────────┐
│ Trigger                       │ Current Threshold        │ Actual Capacity          │
├─────────────────────────────────────────────────────────────────────────────────────┤
│ Tool Output Truncation         │ > 8,000 chars (~2K tokens) │ Tools can output 50K+ chars │
│ Auto-Summarize (Large Models)   │ Remaining ≤ 20,000 tokens  │ 200K window = 180K usable│
│ Auto-Summarize (Small Models)   │ Remaining < 20% of window │ Too aggressive for 128K │
│ SuperFloyd Prompt Truncation   │ > 12,000 runes          │ Should be 50K+ runes   │
└─────────────────────────────────────────────────────────────────────────────────────┘
```

## Problem 1: Tool Output Truncation is Too Aggressive

**Current Code (agent.go:327-336):**
```go
truncLen := 8000 // approx 2000 tokens
if estTokens > 2000 {
    // Truncate to 8000 chars
}
```

**Issue:** Many legitimate tool outputs (like `ri`, `find`, `grep` results) exceed 8,000 chars but are not "bloat" — they're valuable context.

**Proposed Fix:**
```go
// INCREASE: 8000 → 30000 (approximately 7500 tokens)
truncLen := 30000 
if estTokens > 7500 {  // INCREASE: 2000 → 7500
    // Truncate only when truly massive
}
```

## Problem 2: Auto-Summarization Triggers Too Early

**Current Code (agent.go:510-514):**
```go
largeContextWindowBuffer = 20_000  // Only 20K remaining triggers summarize
smallContextWindowRatio  = 0.2   // 20% remaining triggers summarize
```

**Issue:** 
- 200K context window: Summarizes when 180K used (only 20K buffer)
- 128K context window: Summarizes when 102K used (25K buffer)

This wastes 10-20% of context capacity!

**Proposed Fix:**
```go
largeContextWindowBuffer = 5_000   // Only 5K remaining (more aggressive)
smallContextWindowRatio  = 0.05    // 5% remaining (more aggressive)
```

**Rationale:** Summarization should be a last resort, not a routine optimization.

## Problem 3: SuperFloyd Prompt Truncation Too Low

**Current Code (superfloyd_resilience.go:302):**
```go
maxRunes := 12000  // Only 12K chars
```

**Issue:** Complex prompts with code samples, context, and instructions easily exceed 12K chars.

**Proposed Fix:**
```go
maxRunes := 50000  // 50K chars (approximately 12K tokens)
```

**Rationale:** Modern models can handle 50K chars of prompt without degradation.

---

## Implementation Plan

### File 1: `internal/agent/agent.go`

```go
// Lines 51-52: Increase buffer thresholds
largeContextWindowBuffer = 5_000   // Was 20_000
smallContextWindowRatio  = 0.05    // Was 0.2

// Lines 325-327: Increase truncation thresholds
truncLen := 30000   // Was 8000
if estTokens > 7500 {  // Was 2000
```

### File 2: `internal/cmd/superfloyd_resilience.go`

```go
// Line 302: Increase prompt size limit
maxRunes := 50000  // Was 12000
```

---

## Expected Impact

```
┌─────────────────────────────────┬────────────────────────────────────────┐
│ Metric                        │ Before                             │ After            │
├─────────────────────────────────┼────────────────────────────────────────┤
│ Tool output preserved          │ ~2K tokens                         │ ~7.5K tokens     │
│ Context utilization (200K)     │ 90% (summarize at 180K)            │ 97.5% (summarize at 195K) │
│ Context utilization (128K)     │ 80% (summarize at 102K)            │ 95% (summarize at 121K) │
│ SuperFloyd prompt size         │ 12K chars                          │ 50K chars         │
│ Degradation frequency          │ High                               │ Low               │
└─────────────────────────────────┴────────────────────────────────────────┘
```

This fix ensures agents use **95-97.5%** of context window before summarization instead of **80-90%**.
