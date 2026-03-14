# Cache Token Accounting Fix - Comprehensive Plan

```
┌──────────────────────────────────────────────────────────────────────────────┐
│  Document: CACHE_TOKEN_ACCOUNTING_FIX.md                                    │
│  Created: 2026-02-19                                                        │
│  Status: Ready for Implementation                                           │
│  Priority: Critical - Affects core TUI functionality                        │
└──────────────────────────────────────────────────────────────────────────────┘
```

---

## I. PROBLEM STATEMENT

### Symptom
Floyd's TUI is auto-compacting (summarizing) sessions at approximately 100k tokens despite:
- Having a 200k context window model (GLM-5 via Floyd Harness)
- Content caching being implemented and active
- Expectation of compaction only when approaching ~180k tokens

### Impact
- Sessions are being summarized prematurely
- Context is lost unnecessarily
- User experience degrades before actual context limits are reached

---

## II. ROOT CAUSE ANALYSIS

### A. The Auto-Summarization Logic

**Location:** `/Volumes/Storage/floyd-main/internal/agent/agent.go`

**Constants (lines 50-53):**
```go
largeContextWindowThreshold = 200_000
largeContextWindowBuffer    = 20_000
smallContextWindowRatio     = 0.2
```

**Threshold Calculation (lines 437-458):**
```go
StopWhen: []fantasy.StopCondition{
    func(_ []fantasy.StepResult) bool {
        cw := int64(largeModel.CatwalkCfg.ContextWindow)
        if largeModel.ModelCfg.ContextWindow > 0 {
            cw = largeModel.ModelCfg.ContextWindow
        }
        tokens := currentSession.CompletionTokens + currentSession.PromptTokens  // <-- ISSUE HERE
        remaining := cw - tokens
        var threshold int64
        if cw >= largeContextWindowThreshold {
            threshold = largeContextWindowBuffer  // 20,000
        } else {
            threshold = int64(float64(cw) * smallContextWindowRatio)
        }
        if (remaining <= threshold) && !a.disableAutoSummarize {
            shouldSummarize = true
            return true
        }
        return false
    },
}
```

### B. Where PromptTokens Gets Its Value

**Location:** `/Volumes/Storage/floyd-main/internal/agent/agent.go` (lines 1001-1018)

```go
func (a *sessionAgent) updateSessionUsage(model Model, session *session.Session, usage fantasy.Usage, overrideCost *float64) {
    // ... cost calculation ...
    session.CompletionTokens = usage.OutputTokens
    session.PromptTokens = usage.InputTokens + usage.CacheReadTokens  // <-- CACHE READ INCLUDED
}
```

### C. The Core Bug

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  PromptTokens = InputTokens + CacheReadTokens                              │
│                                                                             │
│  StopWhen calculates: tokens = CompletionTokens + PromptTokens             │
│                                                                             │
│  Result: Cached content (which is REUSED and doesn't consume NEW context)  │
│          is counted toward the context window limit                        │
└─────────────────────────────────────────────────────────────────────────────┘
```

### D. Why This Causes ~100k Compaction

With Anthropic/compatible content caching:
- Cached tokens are "read" from cache (CacheReadTokens)
- These tokens DON'T consume new context window space at the API level
- But Floyd counts them as if they do

**Example:**
```
Actual API context consumed:  80k fresh tokens
CacheReadTokens reported:      80k cached tokens
─────────────────────────────────────────────────
PromptTokens stored:          160k (80+80)
CompletionTokens:              20k
─────────────────────────────────────────────────
Total "used" per Floyd:       180k
Context window:               200k
Remaining:                     20k <= threshold

→ TRIGGERS SUMMARIZATION at 90% "used"
→ But API only actually consumed ~50% of context
```

---

## III. SECONDARY ISSUES (Why a Simple Fix is Insufficient)

### Issue 1: Inconsistent Token Accounting

Two different calculations are used for `PromptTokens`:

| Location | Calculation | Purpose |
|----------|-------------|---------|
| Line 976 (title generation) | `InputTokens + CacheCreationTokens` | Saved to DB |
| Line 1017 (streaming) | `InputTokens + CacheReadTokens` | In-memory session |

These are **different values** being stored in the same field. This causes:
- Database state differs from runtime state
- Inconsistency after session restoration

### Issue 2: UI Displays Inflated Percentage

**Header (`internal/ui/model/header.go` line 124):**
```go
percentage = (float64(session.CompletionTokens+session.PromptTokens) / float64(contextWindow)) * 100
```

**Sidebar (`internal/ui/model/sidebar.go` line 58):**
```go
ContextUsed: m.session.CompletionTokens + m.session.PromptTokens,
```

Both include `CacheReadTokens` in the displayed usage, showing:
- "80% context used" when actually only 40% is consumed
- Misleading users about actual context state

### Issue 3: No Separate Cache Token Tracking

The `Session` struct has no `CacheReadTokens` field:
- Cannot distinguish cached vs fresh tokens
- Cannot display accurate "actual vs cached" breakdown
- Cannot make informed summarization decisions

---

## IV. COMPREHENSIVE FIX PLAN

### A. Database Schema Change

**New Migration:** `internal/db/migrations/20260219000000_add_cache_tokens.sql`

```sql
-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions ADD COLUMN cache_read_tokens INTEGER NOT NULL DEFAULT 0 CHECK (cache_read_tokens >= 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions DROP COLUMN cache_read_tokens;
-- +goose StatementEnd
```

### B. Session Struct Update

**File:** `internal/session/session.go`

```go
type Session struct {
    ID               string
    ParentSessionID  string
    Title            string
    MessageCount     int64
    PromptTokens     int64  // Fresh input tokens only (no cache)
    CompletionTokens int64
    CacheReadTokens  int64  // NEW: Cached tokens read
    SummaryMessageID string
    Cost             float64
    Todos            []Todo
    CreatedAt        int64
    UpdatedAt        int64
}
```

### C. Token Accounting Fix

**File:** `internal/agent/agent.go`

**Line 1017 - updateSessionUsage:**
```go
// BEFORE:
session.PromptTokens = usage.InputTokens + usage.CacheReadTokens

// AFTER:
session.PromptTokens = usage.InputTokens
session.CacheReadTokens = usage.CacheReadTokens
```

**Line 976 - title generation (for consistency):**
```go
// BEFORE:
promptTokens := resp.TotalUsage.InputTokens + resp.TotalUsage.CacheCreationTokens

// AFTER:
promptTokens := resp.TotalUsage.InputTokens
// Note: CacheCreationTokens is a one-time cost, not ongoing context consumption
```

### D. StopWhen Threshold Fix

**File:** `internal/agent/agent.go` (lines 444-445)

```go
// BEFORE:
tokens := currentSession.CompletionTokens + currentSession.PromptTokens

// AFTER:
// Only count fresh tokens - cached content doesn't consume new context
freshTokens := currentSession.CompletionTokens + currentSession.PromptTokens
tokens := freshTokens
```

**Note:** After the accounting fix, `PromptTokens` no longer includes `CacheReadTokens`, so this line can remain unchanged semantically. But for clarity:

```go
// Calculate actual context consumption (excluding cached content)
freshPromptTokens := currentSession.PromptTokens  // Now excludes cache
freshTokens := currentSession.CompletionTokens + freshPromptTokens
remaining := cw - freshTokens
```

### E. UI Display Fixes

**File:** `internal/ui/model/header.go` (line 124)

```go
// BEFORE:
percentage = (float64(session.CompletionTokens+session.PromptTokens) / float64(contextWindow)) * 100

// AFTER:
// Show actual context consumption (fresh tokens only)
freshTokens := session.CompletionTokens + session.PromptTokens
percentage = (float64(freshTokens) / float64(contextWindow)) * 100
```

**File:** `internal/ui/model/sidebar.go` (line 58)

```go
// BEFORE:
ContextUsed: m.session.CompletionTokens + m.session.PromptTokens,

// AFTER:
ContextUsed: m.session.CompletionTokens + m.session.PromptTokens,  // Now accurate
```

### F. SQL Query Updates

**File:** `internal/db/sql/sessions.sql`

Add `cache_read_tokens` to all relevant queries:
- `UpdateSession`
- `UpdateSessionTitleAndUsage`
- `GetSessionByID`
- `ListSessions`

### G. sqlc Regeneration

After schema and query changes:
```bash
cd /Volumes/Storage/floyd-main
sqlc generate
```

This regenerates:
- `internal/db/models.go`
- `internal/db/sessions.sql.go`

---

## V. FILES TO MODIFY (Complete List)

| # | File | Change Type |
|---|------|-------------|
| 1 | `internal/db/migrations/20260219000000_add_cache_tokens.sql` | NEW |
| 2 | `internal/db/sql/sessions.sql` | MODIFY |
| 3 | `internal/session/session.go` | MODIFY |
| 4 | `internal/agent/agent.go` (3 locations) | MODIFY |
| 5 | `internal/ui/model/header.go` | MODIFY |
| 6 | `internal/ui/model/sidebar.go` | MODIFY |
| 7 | `internal/db/models.go` | REGENERATE (sqlc) |
| 8 | `internal/db/sessions.sql.go` | REGENERATE (sqlc) |

---

## VI. IMPLEMENTATION ORDER

```
┌─────────────────────────────────────────────────────────────────────────────┐
│  Phase 1: Database Migration                                                │
├─────────────────────────────────────────────────────────────────────────────┤
│  1. Create migration file                                                   │
│  2. Run: goose up (or equivalent migration tool)                           │
│  3. Verify column added: sqlite3 .sqlite ".schema sessions"                │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│  Phase 2: SQL Queries                                                       │
├─────────────────────────────────────────────────────────────────────────────┤
│  4. Update internal/db/sql/sessions.sql with new column                    │
│  5. Run: sqlc generate                                                      │
│  6. Verify generated code compiles                                          │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│  Phase 3: Session Struct & Logic                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│  7. Update internal/session/session.go                                      │
│  8. Update internal/agent/agent.go (token accounting + threshold)          │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│  Phase 4: UI Updates                                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│  9. Update internal/ui/model/header.go                                      │
│  10. Update internal/ui/model/sidebar.go                                    │
└─────────────────────────────────────────────────────────────────────────────┘
                                    │
                                    ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│  Phase 5: Build & Test                                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│  11. go build ./...                                                         │
│  12. go test ./...                                                          │
│  13. Manual test: Run TUI, verify context % display                         │
│  14. Manual test: Confirm no premature compaction                           │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## VII. ROLLBACK STRATEGY

### If Migration Succeeds But Code Fails

```bash
# Revert code changes
git checkout -- internal/

# Goose down migration
goose -dir internal/db/migrations sqlite3 .sqlite down
```

### If Everything Fails

```bash
# Full revert
git checkout -- .
goose -dir internal/db/migrations sqlite3 .sqlite down

# Rebuild
go build ./...
```

### Backup Before Starting

```bash
# Create backup branch
cd /Volumes/Storage/floyd-main
git checkout -b backup/pre-cache-token-fix-2026-02-19

# Or create tarball
tar -czvf ~/floyd-main-backup-$(date +%Y%m%d).tar.gz .
```

---

## VIII. VERIFICATION CHECKLIST

After implementation, verify:

- [ ] Database has `cache_read_tokens` column
- [ ] `go build ./...` succeeds
- [ ] `go test ./...` passes
- [ ] TUI launches without errors
- [ ] Context percentage shows ~50% when 80k fresh + 80k cached
- [ ] Summarization triggers at ~180k fresh tokens (not 100k)
- [ ] Existing sessions load correctly (cache_read_tokens defaults to 0)
- [ ] New sessions track cache tokens accurately

---

## IX. WHY THIS IS THE CORRECT FIX

### 1. Addresses Root Cause
- Separates cache reads from fresh input
- Threshold calculation uses actual context consumption

### 2. No Hidden Issues Left
- UI displays accurate information
- DB stores consistent data
- Token accounting is uniform throughout

### 3. Backwards Compatible
- New column has DEFAULT 0
- Existing sessions work unchanged
- No data migration required

### 4. Observable Improvement
- Users will see accurate context %
- Compaction will occur at correct threshold
- Debug logs will show accurate token breakdowns

---

## X. ALTERNATIVE APPROACHES CONSIDERED

### A. Simple Patch (REJECTED)
Just change line 444 to exclude cache without tracking separately.

**Why rejected:**
- UI still shows wrong percentage
- DB still stores inconsistent data
- No visibility into cache effectiveness

### B. Remove Cache Token Addition (REJECTED)
Change line 1017 to only use `InputTokens`.

**Why rejected:**
- Loses visibility into cache hit rate
- Can't measure caching effectiveness
- UI still shows lower than actual usage

### C. Keep Current Logic, Increase Threshold (REJECTED)
Change `largeContextWindowBuffer` to 100,000.

**Why rejected:**
- Masks the bug, doesn't fix it
- Breaks for non-cached sessions
- Magic number with no semantic meaning

---

## XI. APPENDIX: CODE REFERENCES

### fantasy.Usage Struct
**File:** `internal/ai/types.go` (or external `charm.land/fantasy` package)

```go
type Usage struct {
    InputTokens         int `json:"input_tokens"`
    OutputTokens        int `json:"output_tokens"`
    CacheReadTokens     int `json:"cache_read_tokens,omitempty"`
    CacheCreationTokens int `json:"cache_creation_tokens,omitempty"`
    TotalTokens         int `json:"total_tokens,omitempty"`
}
```

### Current Session Schema
**File:** `internal/db/migrations/20250424200609_initial.sql`

```sql
CREATE TABLE IF NOT EXISTS sessions (
    id TEXT PRIMARY KEY,
    parent_session_id TEXT,
    title TEXT NOT NULL,
    message_count INTEGER NOT NULL DEFAULT 0,
    prompt_tokens INTEGER NOT NULL DEFAULT 0,
    completion_tokens INTEGER NOT NULL DEFAULT 0,
    cost REAL NOT NULL DEFAULT 0.0,
    updated_at INTEGER NOT NULL,
    created_at INTEGER NOT NULL
);
```

---

## XII. METADATA

```
┌─────────────────┬──────────────────────────────────────────────────────┐
│ Field           │ Value                                                │
├─────────────────┼──────────────────────────────────────────────────────┤
│ Author          │ FLOYD (GLM-5 via Floyd Harness)                      │
│ Created         │ 2026-02-19T13:45:00Z                                 │
│ Project         │ floyd-main                                           │
│ Issue           │ Premature TUI auto-compaction                        │
│ Root Cause      │ CacheReadTokens counted toward context limit         │
│ Fix Type        │ Database migration + logic correction                │
│ Risk Level      │ Low (additive change, backwards compatible)         │
└─────────────────┴──────────────────────────────────────────────────────┘
```

---

*This document is the authoritative reference for the cache token accounting fix. All implementation should follow this plan exactly.*
