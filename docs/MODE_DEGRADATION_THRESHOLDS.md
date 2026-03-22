# Model Degradation Points & Thresholds

## Overview

Even with token accumulation fixed, the model can still degrade due to context pressure, output bloat, and environmental factors. This document maps all degradation points.

---

## 1. Proactive Compaction (Context Bloat)

**Location:** `internal/agent/agent.go:287-343`

**Trigger:** Tool outputs + dynamic context exceed 85% of context window

```
┌─────────────────────┬────────────────────────────────────────┐
│ Condition          │ Threshold                                │
├─────────────────────┼────────────────────────────────────────┤
│ Total Est. Tokens   │ > 85% of context window                 │
│ Large Model (200K)  │ > 170,000 tokens                        │
│ Standard Model      │ > 85% of context window                │
│ Tool Output         │ > 30,000 chars (gets truncated)         │
└─────────────────────┴────────────────────────────────────────┘
```

**Effect:**
- Tool outputs truncated to `truncLen` characters
- `[SYSTEM CAUTION]` warning injected
- Model receives truncated context

**Detection Code:**
```go
// agent.go:315-317
maxAllowedTokens := int(float64(cw) * 0.85)
if totalEstTokens > maxAllowedTokens {
    slog.Warn("Context bloat detected. Proactively compacting tool outputs.")
}
```

---

## 2. Auto-Summarization (Session Token Limit)

**Location:** `internal/agent/agent.go:507-521`

**Trigger:** Session cumulative tokens approach context limit

```
┌──────────────────────────┬──────────────────────────────────┐
│ Condition               │ Threshold                        │
├──────────────────────────┼──────────────────────────────────┤
│ Large Context (≥200K)   │ remaining ≤ fixed buffer tokens    │
│ Small Context           │ remaining ≤ ratio of window        │
│ Default Buffer          │ 8,000 tokens (large context)     │
│ Ratio (small)           │ 8% of context window             │
└──────────────────────────┴──────────────────────────────────┘
```

**Effect:**
- Triggers automatic context summarization
- Older messages compressed into summary

**Detection Code:**
```go
// agent.go:507-521
tokens := currentSession.CompletionTokens + currentSession.PromptTokens
remaining := cw - tokens
if (remaining <= threshold) && !a.disableAutoSummarize {
    shouldSummarize = true
}
```

---

## 3. SuperFloyd Prompt Degradation

**Location:** `internal/cmd/superfloyd_resilience.go:301-316`  
**Also:** `internal/ui/model/ui.go:874-878`

**Trigger:** Prompt size exceeds soft limit (SuperFloyd only)

```
┌──────────────────────────┬──────────────────────────────────┐
│ Condition               │ Threshold                        │
├──────────────────────────┼──────────────────────────────────┤
│ Prompt Size             │ > 12,000 runes (chars)           │
│ Mode                    │ SuperFloyd only                  │
│ Degradation Controls   │ Enabled (default: ON)            │
└──────────────────────────┴──────────────────────────────────┘
```

**Effect:**
- Prompt truncated to 12,000 runes
- Interactive negotiation prompt: "Accept truncation for stability? (Y/n)"
- `[superfloyd-auto-stabilize]` notice appended

**Detection Code:**
```go
// superfloyd_resilience.go:301-316
maxRunes := 12000
if len(runes) > maxRunes {
    fmt.Printf("[superfloyd-warn] Prompt degradation detected: size %d exceeds soft limit %d\n", len(runes), maxRunes)
    // Truncate with warning
}
```

---

## 4. Environmental Performance Degradation

**Location:** `internal/cmd/superfloyd_resilience.go:318-336`

**Trigger:** System performance metrics indicate stress (SuperFloyd only)

```
┌──────────────────────────┬──────────────────────────────────┐
│ Metric                  │ Threshold                        │
├──────────────────────────┼──────────────────────────────────┤
│ Avg Response Time       │ ≥ 22,000 ms (22 seconds)         │
│ Avg Tokens/Session      │ ≥ 90,000 tokens                  │
│ Min Sessions for Stats  │ ≥ 10 sessions                    │
└──────────────────────────┴──────────────────────────────────┘
```

**Effect:**
- `[superfloyd-warn] Environmental degradation detected`
- Interactive prompt: "Enable strict concise mode?"
- Append instruction: `[superfloyd-auto-stabilize] Use concise responses...`

**Detection Code:**
```go
// superfloyd_resilience.go:344-365
if stats.AvgResponseTimeMs >= 22000 || stats.Total.AvgTokensPerSession >= 90000 {
    return true // Trigger stabilize mode
}
```

---

## 5. Retry Budget Exhaustion

**Location:** `internal/cmd/superfloyd_resilience.go:368-379`

**Trigger:** Too many failures in recent time window

```
┌──────────────────────────┬──────────────────────────────────┐
│ Condition               │ Threshold                        │
├──────────────────────────┼──────────────────────────────────┤
│ Failures in 1 hour      │ ≥ 6 failures                     │
│ Mode                    │ SuperFloyd only                  │
└──────────────────────────┴──────────────────────────────────┘
```

**Effect:**
- Returns error: `retry budget exceeded: X failures in last hour; stabilize before retrying`
- Hard stop on further retries

**Detection Code:**
```go
// superfloyd_resilience.go:368-379
state.Failures = keepRecentFailures(state.Failures, now-3600)
if len(state.Failures) >= 6 {
    return fmt.Errorf("retry budget exceeded: %d failures in last hour", len(state.Failures))
}
```

---

## 6. Consistency Lock Violation

**Location:** `internal/cmd/superfloyd_resilience.go:269-289`

**Trigger:** Boot contract file (FLOYD.md) has drifted from expected content

```
┌──────────────────────────┬──────────────────────────────────┐
│ Check                   │ Expected                         │
├──────────────────────────┼──────────────────────────────────┤
│ FLOYD.md exists         │ Yes                              │
│ Contains "FLOYD"       │ Yes                              │
│ SuperFloyd binary      │ Yes                              │
└──────────────────────────┴──────────────────────────────────┘
```

**Effect:**
- Returns error: `consistency lock failed: boot contract drift in FLOYD.md`
- Prevents session start until resolved

---

## Environment Variable Controls

All degradation controls can be toggled via environment variables:

```bash
# Enable/Disable Quality Gates (default: ON)
export SUPERFLOYD_QUALITY_GATES=true

# Enable/Disable Degradation Controls (default: ON)
export SUPERFLOYD_DEGRADATION_CONTROLS=true

# Enable/Disable Consistency Lock (default: ON)
export SUPERFLOYD_CONSISTENCY_LOCK=true

# Enable/Disable Auto-Stabilize (default: ON)
export SUPERFLOYD_AUTO_STABILIZE=true
```

Set to `0`, `false`, or or `off` to disable.

---

## Quick Reference Table

```
┌───────────────────────────────┬─────────────────┬─────────────────────────────────┐
│ Degradation Type              │ Mode            │ Primary Threshold               │
├───────────────────────────────┼─────────────────┼─────────────────────────────────┤
│ Tool Output Bloat            │ Both            │ > 30,000 chars                  │
│ Context Window (85%)         │ Both            │ 85% of max context             │
│ Auto-Summarize (large)       │ Both            │ Remaining ≤ 8,000 tokens       │
│ Auto-Summarize (small)       │ Both            │ Remaining ≤ 8% of window       │
│ Prompt Size                  │ SuperFloyd      │ > 12,000 runes                 │
│ Response Latency             │ SuperFloyd      │ ≥ 22,000 ms avg                │
│ Session Tokens               │ SuperFloyd      │ ≥ 90,000 avg                   │
│ Retry Failures               │ SuperFloyd      │ ≥ 6 per hour                   │
│ Consistency Lock             │ SuperFloyd      │ FLOYD.md drift                 │
└───────────────────────────────┴─────────────────┴─────────────────────────────────┘
```

---

## Summary

The model degrades when:

1. **Token accumulation is correct** but context fills up → triggers auto-summarization
2. **Tool outputs are huge** → gets truncated, model loses context
3. **Running SuperFloyd** with large prompts → truncation + negotiation
4. **System is under stress** (high latency, high token usage) → concise mode enforced
5. **Too many failures** → retry budget exhausted, hard stop

The key insight: **Tokens are counted correctly now, but degradation happens when those tokens fill the context window or when outputs bloat the context.**
