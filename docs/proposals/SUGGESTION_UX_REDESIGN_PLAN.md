# Floyd Suggestion UX Redesign Plan

> Concrete implementation plan for prompt suggestions in the Floyd editor
> Date: 2026-03-19
> Status: Proposed

---

## Executive Summary

Floyd should expose **one visible suggestion experience** in the prompt editor, backed by **two AI-centric behaviors**:

1. **Passive suggestion** — Floyd offers a next-prompt suggestion when confidence is high.
2. **On-demand suggestion** — the user can explicitly request a suggestion at any time.

Prompt history should not act like a competing intelligence layer. It should be used as a **personalization signal** for AI suggestion ranking, tone, and pattern inference, but not shown as the default visible ghost text source unless the user explicitly asks to reuse a prior prompt.

This plan keeps the current accept flow, fixes the mental model, and adds a user-invoked suggestion path.

---

## Current Verified State

### Verified behavior in code

- Passive AI follow-up suggestions are already generated after a successful response in [internal/ui/model/ui.go](internal/ui/model/ui.go#L3037-L3054).
- The prompt editor currently stores AI suggestions separately from rendered ghost completions:
  - `m.aiSuggestion`
  - `m.commandSuggestion`
- Placeholder/opaque AI suggestion rendering is controlled in [internal/ui/model/ui.go](internal/ui/model/ui.go#L834-L842).
- Accepting the visible suggestion is handled in [internal/ui/model/ui.go](internal/ui/model/ui.go#L2940-L2948).
- Editor keybindings live in [internal/ui/model/keys.go](internal/ui/model/keys.go#L5-L20) and [internal/ui/model/keys.go](internal/ui/model/keys.go#L89-L158).

### Current design issue

The current architecture mixes:
- AI-originated follow-up suggestions
- history-based prompt continuations

This creates product ambiguity because users perceive both as “the suggestion,” even though they are sourced differently and have different quality expectations.

---

## Product Goal

Create a single mental model:

> “Floyd suggests the best next prompt it can. I can accept it, ignore it, or ask for one explicitly.”

That means:
- one visible suggestion channel
- one accept mechanism
- one explicit request mechanism
- optional metadata about source only if needed for debugging or power users

---

## Recommended UX Model

## Mode A — Passive Suggestion

Floyd automatically proposes a next prompt when it has enough confidence.

### User experience
- Suggestion appears as faint ghost text in the prompt editor.
- User accepts it with the normal accept key.
- User sends it with `Enter` as usual.
- If the user starts typing along the same prefix, the remainder stays visible.
- If the user diverges, the suggestion disappears.

### Rationale
This preserves the fast path for high-confidence assistance and matches the behavior the user already expects.

---

## Mode B — On-Demand Suggestion

If no passive suggestion is shown, or if the user wants help anyway, Floyd should generate one on request.

### User experience
- User presses a dedicated key.
- Floyd requests a suggested next prompt from the same AI suggestion pipeline.
- The returned suggestion appears in the same ghost-text channel.
- User can accept it with the same accept key.

### Rationale
This solves the most important missing capability:

> “If the AI has not given a suggestion, but I want one, I should be able to ask for it at any time.”

---

## History Should Be Repositioned

History should become an **input to ranking**, not the main visible suggestion source.

### Good use of history
- infer preferred tone
- detect repeated workflows
- personalize likely phrasing
- help AI rank alternatives

### Bad use of history
- directly surface stale prior prompts as the primary suggestion
- compete visually with AI suggestion
- treat prior wording as best practice

### Recommendation
Change history behavior from:
- **visible fallback ghost text**

to:
- **internal ranking signal**
- or explicit recall feature, such as “reuse recent prompt”

---

## Proposed Interaction Contract

### Accept suggestion
Keep the current multi-key accept path:
- backtick
- `Ctrl+Y`
- `Ctrl+]`

These are currently wired in [internal/ui/model/keys.go](internal/ui/model/keys.go#L154-L157).

### Request suggestion
Add a dedicated editor binding:
- **Primary proposal:** `Ctrl+E`
- **Fallback proposal:** `Ctrl+;` if terminal testing shows `Ctrl+E` is problematic

### Why `Ctrl+E`
- currently unused in the editor keymap
- mnemonic: “expand” / “engage AI” / “explicit suggestion”
- does not collide with `Tab`, `Ctrl+J`, `Enter`, `/`, `@`, `Ctrl+P`, `Ctrl+G`, `Ctrl+T`, `Ctrl+S`, `Ctrl+L`, `Ctrl+M`, `Ctrl+R`, or `Ctrl+F`

### Required caveat
Because terminal key delivery varies across iTerm2, Ghostty, Apple Terminal, and macOS betas, the request-suggestion key should be validated empirically in at least:
- iTerm2
- Terminal.app
- Ghostty if supported

---

## Proposed State Model

Replace the conceptual model:
- `aiSuggestion`
- `historySuggestion`
- `commandSuggestion`

with this product model:

```go
type SuggestionSource string

const (
    SuggestionSourceNone      SuggestionSource = "none"
    SuggestionSourcePassiveAI SuggestionSource = "passive_ai"
    SuggestionSourceDemandAI  SuggestionSource = "on_demand_ai"
    SuggestionSourceHistory   SuggestionSource = "history"
)

type EditorSuggestion struct {
    FullText   string
    Source     SuggestionSource
    Confidence float64
    Requested  bool
}
```

### UX rule
Only one `EditorSuggestion` is active at a time.

### Priority
1. on-demand AI
2. passive AI
3. history only if explicitly enabled or requested

---

## Implementation Plan

## Phase 1 — Normalize the current suggestion pipeline

### Goal
Keep a single visible ghost-text channel and stop treating history as a peer to AI.

### Changes
- Introduce explicit suggestion source tracking in the UI model.
- Route all rendered ghost text through one structure.
- Keep existing accept behavior unchanged.

### Files
- [internal/ui/model/ui.go](internal/ui/model/ui.go)
- [internal/ui/model/keys.go](internal/ui/model/keys.go)

### Done when
- Suggestion source is tracked explicitly.
- Help text can describe one visible suggestion system.

---

## Phase 2 — Add on-demand AI suggestion

### Goal
Allow the user to request a suggestion at any time.

### Changes
- Add new editor keybinding `RequestSuggestion`.
- Add UI handler for request.
- Call `AgentCoordinator.SuggestFollowup(...)` or a sibling method that works without requiring a just-completed run.

### Recommended API shape
```go
SuggestPrompt(ctx context.Context, sessionID, currentEditorText string) (string, error)
```

### Why a new method
`SuggestFollowup()` currently assumes a successful prior response context in [internal/ui/model/ui.go](internal/ui/model/ui.go#L3049-L3054). An explicit “suggest from current editor state” method better matches the requested behavior.

### Files
- [internal/ui/model/keys.go](internal/ui/model/keys.go)
- [internal/ui/model/ui.go](internal/ui/model/ui.go)
- likely coordinator/agent files behind `AgentCoordinator`

### Done when
- With no passive suggestion visible, user can press the request key and receive one.
- Returned suggestion uses the same accept key path.

---

## Phase 3 — Demote history to personalization

### Goal
Stop showing history as the default fallback ghost suggestion.

### Changes
- Remove default history ghost-text fallback from `updateCommandSuggestion()`.
- Keep prompt history navigation with Up/Down unchanged.
- Optionally add an explicit “reuse recent prompt” command later.

### Files
- [internal/ui/model/ui.go](internal/ui/model/ui.go#L2873-L2915)
- [internal/ui/model/history.go](internal/ui/model/history.go)

### Done when
- History no longer silently appears as the main suggestion.
- Up/Down still works for deliberate recall.

---

## Phase 4 — Optional provenance badge

### Goal
Reduce user confusion without visual clutter.

### Changes
Optionally render a tiny suffix label in the editor status/help area:
- `AI`
- `Ask AI`
- `History`

### Rule
Keep this subtle and optional. The primary UX should still feel like one feature.

---

## Acceptance Criteria

### Passive AI
- After an agent response, Floyd may render a suggestion.
- User can accept it into the editor.
- If the user types or dictates a matching prefix, the remainder stays visible.
- If the user diverges, it disappears.

### On-demand AI
- User can request a suggestion without waiting for passive generation.
- Returned suggestion appears in the same visible channel.
- Same accept key path works.

### History behavior
- History no longer silently overrides or competes with AI suggestion.
- Up/Down recall remains available.

### Terminal resilience
- Accept and request bindings verified in iTerm2 and Terminal.app.
- If a binding is not reliably delivered, the fallback is documented and tested.

---

## Recommended Keybinding Summary

### Accept visible suggestion
- backtick
- `Ctrl+Y`
- `Ctrl+]`

### Request a new AI suggestion
- proposed: `Ctrl+E`
- fallback candidate after terminal testing: `Ctrl+;`

### Why not these
- `Tab`: already changes focus
- `Ctrl+J`: newline
- `Enter`: send
- `/`: commands/add file
- plain letters: unsafe and conflict-prone

---

## Suggested Rollout Order

1. land explicit suggestion source tracking
2. add on-demand AI request binding and API
3. remove default visible history fallback
4. update docs/help/UI hints
5. terminal compatibility test matrix

---

## Final Recommendation

The best design is **not** two competing user-visible services.

The best design is:
- **one suggestion feature**
- **two AI behaviors**: passive and on-demand
- **history as silent personalization input**
- **history recall as deliberate user action, not default suggestion output**

That is the cleanest mental model and the most aligned with best-practice assistance.
