# Ghost Text & AI Suggestion Fix Plan

## Objective
Fix the AI suggestion and ghost text features so they work correctly with `Ctrl+E` (request suggestion) and backtick/`Ctrl+Y`/`Ctrl+]` (accept suggestion).

---

## Issue Summary

| # | Issue | Severity | File | Line |
|---|-------|----------|------|------|
| 1 | Testing bypass disables suggestions in test mode | High | ui.go | 3062 |
| 2 | Silent failure on empty response (returns nil,nil) | Medium | agent.go | 1001 |
| 3 | Cursor position check too restrictive | Medium | ui.go | 2944 |
| 4 | No multi-line suggestion support | Low | ui.go | 2892, 2939 |
| 5 | No feedback when request returns empty | Low | ui.go | 3093 |
| 6 | Passive suggestion only after AI response | Medium | ui.go | 3060 |

---

## Phase 1: Fix Core Rendering Pipeline

### 1.1 Relax Cursor Position Check
**File**: `internal/ui/model/ui.go`
**Problem**: Ghost text only shows when cursor is at end of line 0
**Fix**: Allow cursor anywhere on line 0 (not just end)

```go
// Current (line 2944):
if cur == nil || cur.Y != 0 || cur.X != len(valueRunes) {

// Fixed: Allow any X position on line 0
if cur == nil || cur.Y != 0 {
```

### 1.2 Fix Multi-line Suggestion Support
**File**: `internal/ui/model/ui.go`
**Problem**: Suggestions disabled entirely for multi-line input
**Fix**: Show suggestion suffix for first line only, ignore newlines in suffix calculation

```go
// Current (line 2892):
if strings.Contains(value, "\n") {
    m.commandSuggestion = ""
    return
}

// Fixed: Only clear for multi-line suggestions when suggestion itself spans lines
// For single-line suggestions, show suffix on first line
```

---

## Phase 2: Fix Suggestion Generation

### 2.1 Add Test Mode Override
**File**: `internal/ui/model/ui.go`
**Problem**: `testing.Testing()` bypass prevents E2E testing
**Fix**: Add environment variable override

```go
// Current (line 3062):
if !testing.Testing() {

// Fixed:
skipSuggestion := testing.Testing() && os.Getenv("FLOYD_ENABLE_TEST_SUGGESTIONS") == ""
if !skipSuggestion {
```

### 2.2 Fix Silent Failure on Empty Response
**File**: `internal/agent/agent.go`
**Problem**: Returns `nil, nil` instead of distinct error
**Fix**: Return explicit empty indicator

```go
// Current (lines 1001-1002):
if resp == nil || resp.Response.Content.Text() == "" {
    return "", nil
}

// Fixed:
if resp == nil || resp.Response.Content.Text() == "" {
    return "", ErrNoSuggestion  // Distinct sentinel error
}

// Add to errors:
var ErrNoSuggestion = errors.New("no suggestion generated")
```

### 2.3 Add Feedback for Empty Suggestions
**File**: `internal/ui/model/ui.go`
**Problem**: No user feedback when suggestion request returns nothing
**Fix**: Check for specific error and show message

```go
// Current (lines 3093-3095):
if strings.TrimSpace(suggestion) == "" {
    return util.InfoMsg{Type: util.InfoTypeInfo, Msg: "No suggestion available right now."}
}

// Already correct - verify ErrNoSuggestion propagates
if errors.Is(err, ErrNoSuggestion) {
    return util.InfoMsg{Type: util.InfoTypeInfo, Msg: "No suggestion available right now."}
}
```

---

## Phase 3: Improve Passive Suggestions

### 3.1 Enable On-Demand Generation Anytime
**File**: `internal/ui/model/ui.go`
**Problem**: Passive suggestions only after successful AI response
**Fix**: After a delay post-response OR on first message of new session, generate initial suggestion

```go
// In sendMessage completion handler, add:
if sessionID != "" && !m.hasAiSuggestion() {
    // Generate initial suggestion for new sessions
    go func() {
        time.Sleep(2 * time.Second)  // Brief delay for UX
        suggestion, _ := m.com.App.AgentCoordinator.SuggestPrompt(ctx, sessionID, "")
        if suggestion != "" {
            m.send(aiSuggestionMsg(suggestion))
        }
    }()
}
```

---

## Phase 4: Key Binding Verification

### 4.1 Document Known Terminal Limitations
**File**: `docs/guides/TERMINAL_KEYBINDINGS.md` (new)

Document which terminals reliably deliver:
- `Ctrl+E` for suggestion request
- Backtick for accept
- `Ctrl+Y` / `Ctrl+]` as alternatives

### 4.2 Add Fallback Key Option
**File**: `internal/ui/model/keys.go`

If `Ctrl+E` proves problematic, `Ctrl+;` is proposed fallback per SUGGESTION_UX_REDESIGN_PLAN.

```go
// Optional: Add alternative
km.Editor.RequestSuggestion = key.NewBinding(
    key.WithKeys("ctrl+e", "ctrl+;"),
    key.WithHelp("ctrl+e/ctrl+;", "suggest now"),
)
```

---

## Phase 5: Testing

### 5.1 Add UI Model Tests
**File**: `internal/ui/model/suggestion_test.go` (new)

```go
func TestCommandSuggestionSuffix_CursorPosition(t *testing.T)
func TestCommandSuggestionSuffix_MultiLine(t *testing.T)
func TestUpdateCommandSuggestion_Priority(t *testing.T)
```

### 5.2 Add Integration Test
**File**: `internal/agent/agent_test.go`

```go
func TestSuggestPrompt_EmptyResponse(t *testing.T)
func TestSuggestFollowup_WithSession(t *testing.T)
```

### 5.3 Manual Testing Checklist
- [ ] Start new session, verify `Ctrl+E` shows suggestion
- [ ] Type partial prefix, verify ghost text matches
- [ ] Press backtick to accept suggestion
- [ ] Press `Ctrl+Y` as alternative accept
- [ ] Verify history fallback works when no AI suggestion
- [ ] Test in Ghostty, iTerm2, Terminal.app

---

## Files to Modify

| File | Changes |
|------|---------|
| `internal/ui/model/ui.go` | Fix cursor check, multi-line, test bypass, feedback |
| `internal/agent/agent.go` | Add ErrNoSuggestion, fix empty response |
| `internal/ui/model/keys.go` | Optional: add fallback key |
| `docs/guides/TERMINAL_KEYBINDINGS.md` | Document key compatibility (new) |
| `internal/ui/model/suggestion_test.go` | Add tests (new) |
| `internal/agent/agent_test.go` | Add suggestion tests |

---

## Verification Checklist

Before marking complete:
- [ ] `go build ./...` succeeds
- [ ] `go test ./internal/ui/...` passes
- [ ] `go test ./internal/agent/...` passes
- [ ] Manual test in terminal confirms `Ctrl+E` triggers suggestion
- [ ] Manual test confirms backtick/Ctrl+Y accepts ghost text
- [ ] No regressions in existing functionality

---

## Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Key binding not delivered by terminal | Medium | High | Already have alternatives (Ctrl+Y, Ctrl+]) |
| Empty responses still silent | Low | Low | Added ErrNoSuggestion |
| Test mode bypass breaking CI | Low | Medium | Added env var override |
| Multi-line rendering edge case | Low | Low | Conservative fix, clear suggestions with actual newlines |

---

## Implementation Order

1. Phase 1 (Rendering) - Immediate visual improvement
2. Phase 2 (Generation) - Core functionality fixes
3. Phase 3 (Passive) - Enhanced UX
4. Phase 4 (Documentation) - Low risk, informational
5. Phase 5 (Testing) - Ensure no regressions
