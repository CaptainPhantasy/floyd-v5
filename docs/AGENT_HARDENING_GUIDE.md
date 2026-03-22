# Agent Hardening Guide — March 2026

Reference documentation for the production hardening applied to `internal/agent/`
during the Universal Deep Code Review session (commit `46ea220` + this session's changes).

---

## Architecture Summary

```
coordinator.Run(sessionID, prompt)
      │
      └─▶ sessionAgent.Run(SessionAgentCall)
                │
                ├─[IsSessionBusy?] → messageQueue.Set → nil (enqueued)
                │
                ├─ PrepareStep (per LLM step):
                │       drain messageQueue into history
                │       estimate tokens (len/4 heuristic)
                │       compact tool outputs >7500 tokens
                │
                ├─[StopWhen: remaining ≤ threshold] → shouldSummarize = true
                │
                ├─[err != nil] → close tool calls, AddFinish, return error
                │
                └─[shouldSummarize == true]
                        │
                        ├─ Summarize(sessionID)
                        │       LLM generates summary message
                        │       SummaryMessageID persisted to session
                        │
                        └─[agent had active tool calls]
                                │
                                └─ nextInterruptedSessionCall(call)
                                        ├─[count ≥ 5] HARD STOP error
                                        └─ wrapInterruptedSessionPrompt (idempotent)
                                              messageQueue.Set(nextCall)
                                              dequeue → Run(firstQueued)
```

---

## Hardening Changes

### 1. SwarmCoordinator — Goroutine-Safe Session Counter

**Affected file:** `internal/agent/swarm_coordinator.go`

**Problem:** `sessionCounter` was a plain `int` incremented without synchronization.
`Submit()` launches `go executeTask()`, so concurrent submissions raced on the variable.

**Fix:** Replaced with `sync/atomic.Int64`.

```go
// Before (data race under -race):
var sessionCounter int
func nextSessionCounter() int {
    sessionCounter++
    return sessionCounter
}

// After (goroutine-safe):
var sessionCounter atomic.Int64
func nextSessionCounter() int64 {
    return sessionCounter.Add(1)
}
```

**Impact:** Session IDs are now safe to generate concurrently. The format remains
`swarm-sess-<N>` where N is monotonically increasing.

---

### 2. Provider Retry Observability

**Affected file:** `internal/agent/agent.go`

**Problem:** The `OnRetry` callback was a no-op. Provider retries (rate-limits, 429s,
transient 500s) happened silently with no log evidence.

**Fix:** Added `slog.Warn` logging on every retry.

```go
OnRetry: func(err *fantasy.ProviderError, delay time.Duration) {
    slog.Warn("Provider retry triggered",
        "session_id", call.SessionID,
        "delay", delay.String(),
        "status_code", err.StatusCode,
        "error", err.Message,
    )
},
```

**What you'll see in logs:**

```
WARN Provider retry triggered session_id=sess-abc delay=2s status_code=429 error=rate limit exceeded
```

Look for repeated `Provider retry triggered` lines to diagnose provider instability.

---

### 3. Configurable Anthropic Thinking Budget

**Affected file:** `internal/agent/coordinator.go`

**Problem:** The Anthropic `budget_tokens` for extended thinking was hardcoded at 2000,
with no way for operators to raise or lower it per model.

**Fix:** The resolver now reads `model.provider_options.budget_tokens` (int or float64)
from the model config, falling back to 2000 if absent.

**Configuration example** (in your `floyd.json` or config file):

```json
{
  "models": {
    "large": {
      "model": "claude-opus-4-5",
      "provider": "anthropic",
      "think": true,
      "provider_options": {
        "budget_tokens": 8000
      }
    }
  }
}
```

**Resolver logic:**

| `provider_options.budget_tokens` value | Effective budget |
|----------------------------------------|-----------------|
| Not set / nil                          | 2000 (default)  |
| `int` value, e.g. `8000`              | 8000            |
| `float64` value (JSON decode), e.g. `4096.0` | 4096     |
| Any other type                         | 2000 (fallback) |

**Recommendation:** For Claude models with 200K context configured, set `budget_tokens`
between 4000–16000 for complex reasoning tasks.

---

### 4. Stale Backup File Removed

`internal/agent/agent.go.backup` (pre-patch snapshot from 2025-03-13) was deleted
from the source tree. It was not git-tracked but occupied 41 KB in the working
directory and could mislead contributors examining the agent package.

---

### 5. UI/Config Context Propagation Cleanup (`context.TODO()` → `context.Background()`)

**Affected files:**

- `internal/ui/dialog/sessions.go`
- `internal/ui/dialog/rename_session.go`
- `internal/ui/model/ui.go`
- `internal/config/copilot.go`

**Problem:** `context.TODO()` was used in production call sites, which is intended as
a temporary marker and weakens context semantics in shipped code.

**Fix:** Replaced all production `context.TODO()` usages with `context.Background()`.
This is the correct default in Bubble Tea update handlers where no cancelable parent
context is available.

**Result:** zero remaining `context.TODO()` usages in non-test code under `internal/`.

---

### 6. DiffView Unknown Layout Crash Guard

**Affected file:** `internal/ui/diffview/diffview.go`

**Problem:** `DiffView.String()` panicked in the default branch for unknown layouts:

```go
panic("unknown diffview layout")
```

Because `layoutUnified = iota + 1`, the zero value (`layout == 0`) is invalid and
could crash the TUI if a `DiffView` was zero-value initialized.

**Fix:** Replaced panic with a safe fallback render string:

```go
return style.Render("[diffview: unsupported layout]")
```

**Result:** unknown layout values no longer terminate the process.

---

## Token Budget Thresholds

Auto-summarization fires when the following conditions are met:

| Context Window Size | Threshold formula | Example (128K) |
|--------------------|--------------------|----------------|
| ≥ 128K tokens (large) | `remaining ≤ 8000` | Fires at 120K used |
| < 128K tokens (small) | `remaining ≤ cw × 5%` | Fires at 95% used |
| Any model | Hard cap at 200K (even if model supports 1M+) | Always |

The 200K hard cap prevents ultra-long-context models from degrading silently before
Floyd's summarization kicks in.

---

## Running the Production Test Suite

All hardening tests live in `internal/agent/production_hardening_test.go`.
They are designed to run on the target machine — no CI/CD infrastructure required.

```bash
# Run all hardening tests with the race detector (required):
go test -race ./internal/agent -run 'TestHardening' -v

# Full combined suite (hardening + all prior invariant tests):
go test -race ./internal/agent -run 'TestHardening|TestWrap|TestNext|TestReplay|TestToken' -v

# Verify the package compiles and vets clean:
go build ./internal/agent/...
go vet ./internal/agent/...
```

Expected exit code: `0` for all commands.

### Test inventory

| Test | Issue covered |
|------|---------------|
| `TestHardeningSwarmCounterConcurrentUniqueness` | Issue 1 — data race, 200 goroutines |
| `TestHardeningSwarmSessionIDFormat` | Issue 1 — format + uniqueness under concurrency |
| `TestHardeningOnRetryHandlerNoPanic` | Issue 2 — retry callback non-nil, no panic |
| `TestHardeningThinkBudgetResolutionFallback` (5 subtests) | Issue 3 — budget resolver |
| `TestHardeningStaleBackupFileAbsent` | Issue 4 — backup file absent |
| `TestHardeningNoContextTODOInProduction` | Issue 5 — zero `context.TODO()` in production code under `internal/` |
| `TestHardeningNoTODOsInProductionCode` | Regression — zero `// TODO` or `// FIXME` in production `.go` files |
| `TestHardeningDiffviewUnknownLayoutNoPanic` | Issue 6 — unknown diff layout no longer panics |
| `TestHardeningDiffviewValidLayoutsRender` | Issue 6 — valid layouts still render correctly |
| `TestHardeningNoPanicLiteralInDiffView` | Issue 6 — source guard against panic re-introduction |

---

## Issue Ledger (closed)

| # | Severity | File | Finding | Fix |
|---|----------|------|---------|-----|
| 1 | CRITICAL | `swarm_coordinator.go:174-181` | Data race on package-level `int` counter in concurrent goroutines | `sync/atomic.Int64` |
| 2 | MEDIUM | `agent.go:464` | `OnRetry` no-op swallowed all provider retries silently | `slog.Warn` with session ID, status code, delay |
| 3 | MEDIUM | `coordinator.go:308` | `budget_tokens` hardcoded at 2000; no operator override path | Resolver from `model.ModelCfg.ProviderOptions["budget_tokens"]` with int/float64 fallback |
| 4 | LOW | `internal/agent/agent.go.backup` | Stale 41KB pre-patch backup in source tree | Deleted |
| 5 | LOW | `coordinator.go:106` | `// TODO` architectural note | Reworded to full-sentence comment |
| 6 | LOW | `coordinator.go:531` | `// TODO` architectural note | Reworded to full-sentence doc comment |
| 7 | MEDIUM | `internal/ui/**`, `internal/config/copilot.go` | Production `context.TODO()` usage in runtime paths | Replaced with `context.Background()` |
| 8 | MEDIUM | `internal/ui/diffview/diffview.go:227` | Panic on unknown/zero-value layout | Replaced panic with safe fallback render |

All 8 items: **CLOSED**.

---

## Environment Constraints (target machine)

| Constraint | Value |
|-----------|-------|
| Go version | 1.25.5 |
| OS / Arch | darwin/arm64 (Apple Silicon) |
| Kernel | XNU 25.4.0 |
| Module path | `github.com/legacy-ai/floyd` |
| Active branch | `v5.5.0-trimmed` |
| Test infra | Local `go test -race`, no external CI |
| LLM test cassettes | VCR in `internal/agent/testdata/` (network-free replay) |
