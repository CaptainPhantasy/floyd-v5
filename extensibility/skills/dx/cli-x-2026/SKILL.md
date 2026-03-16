---
name: cli-x-2026
description: Build skill for production-grade Ink (React for CLI) implementations. Use this skill when adding/altering Ink UI components, implementing streaming updates (LLM tokens, tool logs), designing user flows (command palette, file picker, permissions), establishing UI/state contracts, improving UX polish (gradients, syntax highlighting, loaders), or reducing terminal noise. Transforms CLI feature requests into dashboard-first layouts with agentic streaming UX, performance rules, and fault-tolerant async patterns.
---

# CLI-X 2026: Ink Dashboard + Agentic UX Best Practices

## Skill Mission

Transform any CLI feature request into a production-grade Ink implementation plan with:

- Dashboard-first layout (no scroll-of-death)
- Agentic streaming UX (thinking, planning, tool execution)
- Performance + accessibility rules
- Fault-tolerant async + cancellation
- Composable components + stable contracts

---

## When to Use This Skill

Invoke this skill when a task involves:

- Adding/altering Ink UI components, layouts, overlays, keybindings
- Implementing streaming updates (LLM tokens, tool logs, state changes)
- Designing user flows (command palette, file picker, permissions prompt)
- Establishing UI/state contracts between engine/runtime/UI
- Improving UX polish (gradients, syntax highlighting, skeleton loaders)
- Reducing terminal noise, improving clarity, or adding observability

---

## Required Inputs

When invoked, the caller provides:

| Input | Description |
|-------|-------------|
| **Feature name** | e.g., "Command Palette", "Thinking Panel", "Split Pane Logs" |
| **User goal** | What the user is trying to accomplish |
| **Runtime facts** | Node version, Ink version, renderer constraints, target OS |
| **Event sources** | IPC events, engine events, tool outputs, logs |
| **Constraints** | Must-not-break list, perf constraints, no network, etc. |

**If any input is missing**, assume conservative defaults and output an "Assumptions" section.

---

## Mandatory Output Format

This skill always outputs the following sections in order:

### A) UX Contract (What the user experiences)

- Screen layout (panes, header/footer, overlays)
- Keybindings (global + contextual)
- Status semantics (what "busy/done/error/paused" looks like)
- Interruption semantics (Ctrl+C / Esc behavior)

### B) State Model (What your code maintains)

- State shape (minimal, normalized, typed)
- State transitions (what events cause what changes)
- Persistence (what survives restarts, where stored, when written)

### C) Rendering Plan (What components exist)

- Component list with responsibilities
- Props contracts (typed)
- Where side effects live (effects vs pure render)
- Update frequency controls (throttles/debounces)

### D) Async + Cancellation Rules

- How to cancel streams, tool runs, and spinners
- How to handle timeouts, retries, backoff
- How to prevent double-runs / race conditions

### E) Accessibility + Readability Rules (Non-negotiable)

- Contrast, color-blind safety, dim usage, glyph choices
- Keyboard-only completeness
- Terminal width constraints + responsive layout rules

### F) Error Handling + Recovery

- Error display pattern (where + how)
- "Recover" actions (retry, open logs, copy error)
- Safe fallback mode (no fancy UI, still usable)

### G) Definition of Done Checklist

- Functional checks
- Visual checks
- Performance checks
- Failure-mode checks

---

## 2026 Ink CLI Best Practices

### 4.1 Dashboard-First Layout (Kill the Scroll-of-Death)

**Default layout rule:** Persistent panes + sticky status + controlled log surfaces.

```
┌─────────────────────────────────────────────────────┐
│  Header: Product + Session + Environment            │
├──────────────────────┬──────────────────────────────┤
│                      │                              │
│  Main Pane           │  Side Pane (optional)        │
│  (user focus)        │  - Navigation                │
│                      │  - Agents/Tools              │
│                      │  - Queue                     │
├──────────────────────┴──────────────────────────────┤
│  Footer: Hotkeys + Mode + Network/LLM Status        │
└─────────────────────────────────────────────────────┘
```

**Responsive rules:**

| Width | Behavior |
|-------|----------|
| < 90 cols | Collapse side pane into overlay drawer |
| < 28 rows | Reduce chrome; prioritize main content; truncate logs |

---

### 4.2 Agentic UI: Make Thinking + Acting Visible

**Minimum agentic surfaces:**

1. **Plan Checklist:** Steps rendered as live checklist
   - States: queued → running → done/failed
   - Show: step name, icon, optional duration

2. **Tool Timeline:** Tool name + duration + status + expandable output

3. **Streaming Output:** Tokens/logs stream into bounded viewport (not infinite scroll)

4. **Collapsible "Reasoning/Notes":** Explanations and rationale (not raw chain-of-thought)

**Anti-noise rules:**

- Never stream everything everywhere
- Default: show summaries; provide expand-to-view details

---

### 4.3 Hyper-Interactivity (fzf feel, but integrated)

- Searchable select for anything list-like (files, commands, tools, agents)
- Instant filter as user types (per-keystroke update)
- Multi-select with visible selection markers + count
- Command palette as primary navigation (Ctrl+P or /)

---

### 4.4 Visual Polish (Not Gimmicky)

- **Gradients:** Headers and "success moments" only, not body text
- **Syntax highlighting:** Required for code output
- **Skeleton loaders:** Preferred over blank gaps during async

**Rizz Rule:** Polish must improve clarity—no visual effect if it reduces readability.

---

### 4.5 DX-First Behaviors

- Graceful Ctrl+C: cancel → save state → clean exit message
- Deep links: Clickable file paths and docs links where supported
- Copyable output: "press c to copy summary" patterns

---

## Architecture Patterns

### 5.1 Separate Engine State from UI State

| State Type | Definition | Examples |
|------------|------------|----------|
| **Engine State** | Facts from orchestration/runtime | Agent status, tool runs, logs, progress |
| **UI State** | Presentation state | Overlays open/closed, focused pane, selection indices, search query |

**Rule:** UI never mutates engine state directly. Request actions via a command dispatcher.

---

### 5.2 Event-Driven Rendering

- Treat engine outputs as append-only event stream
- Reduce events into current view model (Redux/Zustand reducer pattern)
- Render from view model
- Keep bounded ring buffer for logs/events (avoid memory blowup)

---

### 5.3 Bounded Viewports for All Streams

Unbounded streams eventually ruin UX.

| Stream Type | Bounding Strategy |
|-------------|-------------------|
| Logs | Ring buffer (size N) |
| Tokens | Chunked, truncated view with "expand" |
| Tool outputs | Summaries + detail handle |

---

### 5.4 Throttle Render Pressure

Ink apps stutter under high-frequency updates.

- Throttle UI updates (30-60 FPS max)
- Batch events into frames (accumulate → reduce → render)
- Memoize lists, virtualize if needed

---

## Accessibility + Readability Rules (Non-Negotiable)

### Color Rules

- Don't rely on color alone—pair with icons/symbols/labels
- High-contrast defaults
- Avoid "gray on black" for primary content

### Consistent Semantics

| Icon | Meaning |
|------|---------|
| ✅ | success |
| ⚠ | warning |
| ✖ | error |
| ⏳ | busy |
| ⏸ | paused |

### Keyboard Completeness

- Esc closes overlays
- Tab cycles focus zones
- Arrow keys + Enter always work in lists
- Provide hints in footer

### Width Safety

- Never assume > 120 cols
- Truncate with affordances: `…` + "press e to expand"

---

## Error Handling & Recovery

### Error Surface Design

Errors appear in dedicated, consistent "Alerts" zone (top-right or overlay).

Each error includes:
- Short message (human readable)
- Error code (stable identifier)
- Suggested action (retry/open logs/diagnose)
- Details toggle

### Recovery Actions

1. Retry last operation
2. Open logs viewport
3. Copy error report block
4. Return to safe mode UI (minimal rendering)

---

## Testing & Quality Gates

### Minimum Test Set

| Test Type | Purpose |
|-----------|---------|
| Snapshot tests | Core layouts (header/footer/panes) |
| Reducer tests | Event → view-model transitions |
| Keybinding tests | Overlay open/close, focus cycle |
| Stream stress test | 10k events, bounded memory |
| Cancellation test | Ctrl+C during tool run and streaming |

### Performance Gates

- No unbounded arrays in long sessions
- UI responsive under high event throughput
- No flicker loops (layout stable under updates)

---

## Recommended "Modern Ink Stack"

**Core**
- Ink + React

**Input + Navigation**
- Text input
- Select input (single + multi)
- Fuzzy search

**Visual Polish**
- Gradient text (headers only)
- Spinner
- Syntax highlighting
- Markdown rendering
- Link rendering (deep links)

**State + Event Reduction**
- Lightweight store (Zustand or reducer pattern)
- Event ring buffer utility

---

## Drop-In Patterns (Reference Implementations)

### Agent Plan Checklist (Concept)

```typescript
type StepStatus = 'queued' | 'running' | 'done' | 'failed';

interface Step {
  id: string;
  label: string;
  status: StepStatus;
  duration?: number;
}
```

- Render: stable list order; avoid reflow on updates
- Each row: icon + label + optional time

### Stream Viewport (Concept)

Maintain bounded buffer of lines. Render last N lines in fixed-height box.

Controls:
- `e`: expand to full-screen overlay
- `f`: follow mode on/off
- `c`: clear viewport (UI-only, not engine)

### Command Palette (Concept)

Overlay with search input + filtered commands list.

```typescript
interface Command {
  id: string;
  title: string;
  description: string;
  shortcut?: string;
  handler: () => void;
}
```

- Filter on title + keywords
- Enter executes, Esc closes

---

## Safety & Permissions

If your CLI can mutate files:

1. Show "what will change" before applying patches
2. Require explicit consent for destructive actions
3. Persist allow/deny decisions with scope (command, path, repo)
4. Provide "review diff" affordance for patch operations

---

## Definition of Done (Global)

A feature is "done" only if:

- [ ] Works with keyboard-only interaction
- [ ] Degrades gracefully under narrow terminals
- [ ] Survives high-frequency events without freezing
- [ ] Cancels cleanly (no orphan processes, no corrupted state)
- [ ] Has an error path that doesn't brick the app
- [ ] Does not introduce unbounded memory growth

---

## Skill Invocation Template

Copy/paste when calling this skill:

```
SKILL: CLI-X 2026 (Ink Dashboard + Agentic UX Best Practices)

Feature:
User Goal:
Runtime Facts:
Event Sources:
Constraints:

Deliver:
A) UX Contract
B) State Model
C) Rendering Plan
D) Async + Cancellation Rules
E) Accessibility + Readability Rules
F) Error Handling + Recovery
G) Definition of Done Checklist
```

---

## Related Skills

- **ink-performance-auditor** — Use this skill after UI design to validate performance under high event throughput
- **mcp-builder** — Use when designing MCP tools that this UI will interact with
- **chrome-extension-bridge** — Use when adding browser integration to the CLI
