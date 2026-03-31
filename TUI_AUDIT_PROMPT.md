# TASK: Full-Scope Floyd Harness Audit — Salvage-or-Rebuild Decision

## MODE: EXPLORE — Evidence-Only Forensic Audit

## STAKES
This audit determines whether ~1000 hours of development is salvageable or requires a ground-up rebuild. Every verdict must be provably correct. A false "working" that hides a broken code path, or a false "broken" that condemns functioning code, could cost months of misdirected effort.

## ZERO TRUST CONSTRAINT
- Do NOT assume any feature works because code exists.
- Do NOT declare "working" without tracing the COMPLETE code path from trigger → handler → effect → persistence.
- Do NOT skip a feature because it looks correct at first glance.
- Do NOT group features. Audit each individually.
- If you cannot trace a complete path, the verdict is BROKEN with the exact break point.
- You MUST read the actual code at each step. Grepping for function names alone is insufficient — you must read the function body.

## OUTPUT FORMAT (mandatory for every feature)

```
### [N.N] [Feature Name]
**Trigger:** [exact key/action/API call that initiates it]
**Expected Behavior:** [what should happen]
**Code Path:**
1. [file:line] — [what happens]
2. [file:line] — [next step]
3. [file:line] — [terminal effect or break point]
**Verdict:** WORKING | BROKEN | PARTIALLY WORKING | DEAD CODE
**Evidence:** [the exact observation that proves the verdict]
**Break Point:** [if not working — exact file:line and why]
**Severity:** CRITICAL | HIGH | MEDIUM | LOW
```

## EXECUTION RULES
- Work through sections sequentially. Complete one section before starting the next.
- Use `grep` to locate, then `read` with line ranges to verify. Never dump files >100 lines.
- After completing each section, write a checkpoint summary to `.floyd/.supercache`.
- Do NOT fix anything. Audit only.
- If a function calls another function, follow the call. Do not stop at the interface boundary.

---

# SECTION 1: AGENT RUNTIME (the core)

Source: `internal/agent/`

## 1.1 Prompt Assembly Pipeline
Trace the COMPLETE path from user input to LLM API call:
1. Where does `coderPrompt()` get called? What template does it select?
2. How does `prompt.NewPrompt` build the system prompt? What data does `.Build()` inject?
3. Where is the protocol kernel (`protocolKernelTmpl`) injected? What position in the message array?
4. Where is the role template injected? Does it carry the cache_control breakpoint?
5. Where is dynamic context (date, git, supercache) injected? Is it AFTER the cache breakpoint?
6. Where are MCP tool instructions appended? Before or after cache control?
7. Where are context files (FLOYD.md) injected? What template variable?
8. What is the FINAL message order sent to the LLM? List all messages in order with their roles.

## 1.2 StopWhen / Summarization Trigger
1. What field is read for context pressure? (`InputTokens`, `TotalTokens`, or cumulative?)
2. What happens when the provider returns `InputTokens: 0` AND `TotalTokens: 0`?
3. What is the hard cap? Is it applied before or after the threshold check?
4. Does the StopWhen function LOG anything? At what level? Will it appear in default logs?
5. When `shouldSummarize` is true, what happens next? Trace into `Summarize()`.
6. After summarization, when does requeue fire? What condition gates it?
7. What prevents infinite requeue loops?

## 1.3 Token Accounting
1. Where are tokens accumulated on the session? (`updateSessionUsage`)
2. Are these cumulative (all turns) or per-step?
3. What fields does the session store? (`PromptTokens`, `CompletionTokens`, `CacheReadTokens`)
4. Does the DB schema actually have the `cache_read_tokens` column?
5. Where does the UI read token data for the header percentage?
6. Is the UI percentage using cumulative or per-step tokens?
7. What happens if the percentage exceeds 100%? Is it clamped?

## 1.4 Tool Execution Loop
1. How does the agent receive tool calls from the LLM response?
2. How does it dispatch to the correct tool handler?
3. Where is the tool result sent back to the LLM?
4. What happens if a tool errors? Is the error sent back as a tool result?
5. Where does selfheal (go build check) fire? Is it in the tool result path?
6. What is the tool output truncation logic? Two-phase compaction — trace the thresholds.

## 1.5 GLM Reasoning Anchor
1. Where does `isGLMModel()` check fire?
2. What does `buildReasoningAnchor()` extract from conversation history?
3. Where is the anchor injected in the message array?
4. Is it injected AFTER the cache breakpoint? (It must be — it's dynamic)

## 1.6 Ralph Loop Integration
1. Where does `ralphLoop.Check()` fire in the Run() method?
2. Does it fire AFTER queued message processing?
3. What happens if both a queued message AND ralph loop are active?
4. Is the Ralph state file created/read correctly?
5. Are the slash commands installed on boot?

## 1.7 Agent Tool (Sub-agent)
1. How does the agent spawn sub-agents?
2. Does the sub-agent share the same session?
3. How are sub-agent results returned to the parent?

---

# SECTION 2: DATABASE LAYER

Source: `internal/db/`, `internal/db/sql/`, `internal/db/migrations/`

## 2.1 Schema Integrity
For each SQL file in `internal/db/sql/`:
1. List every query that references a column.
2. Cross-reference against the migration files to verify every column exists.
3. Specifically check: does `cache_read_tokens` exist in the schema? Which migration adds it? Is there a query that references it without a migration?

## 2.2 Migration System
1. How are migrations applied? Automatically on connect? Manually?
2. What does `ensureColumns` do? Which columns does it add?
3. What does the `DROP COLUMN name` migration do? Is it safe on all SQLite versions?
4. Is there a migration ordering guarantee?

## 2.3 Session Persistence
1. Trace: session creation → DB insert → session retrieval
2. Is the `title` field nullable? What happens if it's NULL?
3. What happens when `UpdateTitleAndUsage` is called with zero tokens?

---

# SECTION 3: CONFIGURATION SYSTEM

Source: `internal/config/`

## 3.1 Config Loading
1. Where is the config loaded from? Which files, in what order?
2. What are the `defaultContextPaths`? Are they correct?
3. What is `GlobalSkillsDirs()`? Does it return paths that exist?
4. What is the `SkillsPaths` default? Does it include `extensibility/skills/`?
5. What is `NormalizeRuntimeProfile()`? Does it correctly select Floyd vs SuperFloyd?

## 3.2 Provider Configuration
1. How are LLM providers configured?
2. Where is the ZAI/GLM provider registered?
3. What is the `hyper` provider type? How does it map to real providers?
4. Are API keys loaded from environment variables or config files?

## 3.3 Model Selection
1. How does `GetModelContextWindow()` work?
2. What happens if the model is not in `provider.json`? Fallback?
3. What is the `ContextWindow` override path?

---

# SECTION 4: TUI — KEYBINDINGS AND CONFLICTS

Source: `internal/ui/model/keys.go`, `internal/ui/model/ui.go`

## 4.1 Complete Keybinding Map
For EVERY binding in `keys.go`:
1. Key(s) assigned
2. Context/state it's active in
3. Handler function in `ui.go`

## 4.2 Conflict Detection
For every key that appears in more than one binding:
1. Are the bindings in different states? (No conflict)
2. Are they in the same state? (TRUE CONFLICT — one blocks the other)
3. Specifically investigate:
   - `ctrl+f` — "add image" vs "add attachment" vs file picker
   - `ctrl+j` — "newline" vs navigation
   - `ctrl+e` — "suggest now" — does anything else use this?
   - `/` — "add file" vs "commands"
   - `` ` `` — "accept suggestion" vs literal backtick in code
   - `tab` — multiple contexts
   - `esc` — multiple contexts

Produce a CONFLICT TABLE with verdicts.

---

# SECTION 5: TUI — COMMAND PALETTE

Source: `internal/ui/dialog/commands.go`, `internal/ui/model/ui.go`

For EVERY `NewCommandItem` in `commands.go`:
1. Action type
2. Handler case in `ui.go` — does it exist?
3. Does the handler do something observable?
4. If it opens a dialog — does the dialog ID have a constructor?

Known items to deeply trace:
- `rename_session` → ActionRenameSession — REPORTED BROKEN
- `summarize` → ActionSummarize
- `export_session` → ActionExportSession
- `agent_library` → dialog exists?
- `skills_library` → dialog exists?
- `plugins_library` → dialog exists?
- `init` → ActionInitializeProject
- `config_audit` → dialog exists?

---

# SECTION 6: TUI — AI SUGGESTION / GHOST TEXT

Source: `internal/ui/model/ui.go`, `internal/ui/model/history.go`, `internal/agent/`

## 6.1 Suggestion Generation
1. What function generates suggestions?
2. Is it an LLM call or pattern matching from history?
3. What triggers it? Timer? Keystroke? Focus change?
4. What data does it use as input? Full conversation history? Last N messages?
5. How much weight does history have vs current context?

## 6.2 Suggestion Display
1. Where is ghost text rendered?
2. What style is applied? Is it visually distinct from real text?
3. Does it update when the user types?
4. Does it clear on backspace?

## 6.3 Accept Suggestion
1. Trace: key press → handler → text insertion → suggestion clear
2. Does backtick (`` ` ``) conflict with typing code?
3. Do `ctrl+y` and `ctrl+]` work as alternatives?

## 6.4 Request Suggestion (Manual)
1. Trace: `ctrl+e` → handler → suggestion generation → display
2. Does it fire an API call or is it a no-op?
3. Does it work when no suggestion is currently shown?

## 6.5 Auto-Suggestion Timing
1. Is there a debounce/delay?
2. Does it respect focus state (only when editor is focused)?
3. Does it fire during tool execution? (It shouldn't)

---

# SECTION 7: TUI — DIALOG LIFECYCLE

Source: `internal/ui/dialog/`

For EVERY dialog ID:
- `api_key_input` — opens? handles input? persists key?
- `arguments` — opens? collects args? returns result?
- `commands` — opens? filters? dispatches?
- `filepicker` — opens? navigates? selects file?
- `models` — opens? lists models? selects?
- `oauth` — opens? initiates flow? stores token?
- `permissions` — opens? prompts? records decision?
- `quit` — opens? confirms? exits?
- `rename_session` — opens? accepts input? saves?
- `session` — opens? lists? selects? deletes?

For each: trace open → input → submit → close → result handled by parent.

---

# SECTION 8: TUI — STATE MACHINE

Source: `internal/ui/model/ui.go`

States: `uiOnboarding`, `uiInitialize`, `uiLanding`, `uiChat`

For each state:
1. What transitions INTO it?
2. What transitions OUT?
3. What keybindings are active?
4. What is rendered?
5. Are there dead-end states?
6. Are there unreachable states?
7. What happens if the LLM provider is not configured? Which state?

---

# SECTION 9: TUI — CHAT RENDERING

Source: `internal/ui/chat/`

1. How are assistant messages rendered? Markdown? Raw?
2. How are tool calls rendered? Collapsed? Expanded?
3. How are tool results rendered? Truncated?
4. How are thinking blocks rendered? Hidden? Shown?
5. How is streaming handled? Incremental updates?
6. What happens with very long messages? Scroll? Truncate?
7. Is the diff view functional? Does it render side-by-side?

---

# SECTION 10: TUI — TERMINAL INTEGRATION

Source: `internal/ui/terminal/`

1. Does `ctrl+t` open a terminal?
2. Does the terminal receive stdin?
3. Does the terminal display stdout/stderr?
4. Can focus switch between terminal and editor?
5. Does the terminal persist across sessions?

---

# SECTION 11: SKILL DISCOVERY AND LOADING

Source: `internal/skills/`, `internal/extensibility/`

1. What paths does `Discover()` walk?
2. Does it find `SKILL.md` files in `extensibility/skills/core/*/`?
3. How does `Parse()` handle malformed frontmatter? Does it crash or warn?
4. Where does `ToPromptXML()` inject skills into the system prompt?
5. How many skills are discoverable from the default paths?
6. Are the 5 broken skill files in `~/.config/floyd/skills/` causing cascading failures or just warnings?

---

# SECTION 12: MCP INTEGRATION

Source: `internal/agent/tools/mcp/`

1. How are MCP servers initialized? What config file?
2. What happens when an MCP server fails to start? Does it block other servers?
3. How are MCP tools registered with the agent?
4. How are MCP prompts loaded?
5. Which MCP servers are currently failing? (from logs)
6. Does a failing MCP server prevent the agent from running?

---

# SECTION 13: CLI COMMANDS

Source: `internal/cmd/`

For each top-level command:
1. `floyd run` — does non-interactive mode work?
2. `floyd stats` — does it generate HTML? Open browser?
3. `floyd models` — does it list available models?
4. `floyd exec` — does it execute commands?
5. `floyd exec-bg` — background execution lifecycle
6. `floyd login` — OAuth flow
7. `floyd projects` — project management
8. `floyd prompt` — prompt listing/rendering
9. `floyd codebase` — analysis commands
10. `floyd schema` — schema management
11. `floyd file` — file operations
12. `floyd dirs` — directory info
13. `floyd logs` — log viewing
14. `floyd lab` — lab management
15. `floyd stats` — usage statistics

For each: does the command parse args correctly, connect to the DB, and produce output?

---

# SECTION 14: SESSION LIFECYCLE (end-to-end)

1. New session creation → DB insert → UI display
2. Session resumption → message loading → context reconstruction
3. Session summarization → summary prompt → message compaction
4. Session export → markdown generation → file write
5. Session deletion → DB cleanup → UI update
6. Session title generation → LLM call → DB update → UI update
7. Session fork → new session with context → DB insert

---

# SECTION 15: PERMISSION SYSTEM

Source: `internal/permission/`

1. How are tool permissions granted?
2. What is the "Yolo mode" (skip all permissions)?
3. Are permissions persisted? Where?
4. Do permission profiles work?
5. What happens when a permission is denied mid-tool-execution?

---

# SECTION 16: BUILD AND RELEASE INFRASTRUCTURE

Source: `.github/workflows/`, `Taskfile.yaml`, `main.go`

1. Does `go build ./...` produce working binaries?
2. Does `go test -race ./...` pass? (Expected: yes, 21/21)
3. Does goreleaser config exist? Is it valid?
4. Are both `floyd` and `superfloyd` binaries produced?
5. Is macOS code signing configured?
6. What version is the binary reporting?

---

# SECTION 17: DEAD CODE AND ORPHANS

1. Are there any `.go` files in the repo that are not imported by anything?
2. Are there templates in `internal/agent/templates/deterministic/` that are compiled into the binary? Or are they dead?
3. Are there any TODO/FIXME/HACK comments that indicate known incomplete work?
4. Is `internal/agent/agent.go.backup` still present? (Should have been deleted)
5. Are there any files in the repo root that shouldn't be there?

---

# FINAL DELIVERABLE

After completing ALL 17 sections, produce:

## Summary Table
```
| # | Section | Features Audited | Working | Broken | Partial | Dead |
```

## Keybinding Conflict Table
```
| Key | Binding 1 | Binding 2 | Same State? | Verdict |
```

## Command Palette Wiring Table
```
| Command | Action Type | Handler Exists | Handler Works | Dialog Exists |
```

## Critical Failures List
Every CRITICAL and HIGH severity finding, ordered by impact.

## Architecture Assessment
Answer these questions with evidence:
1. Is the agent runtime (prompt assembly, tool execution, summarization) architecturally sound?
2. Is the TUI event handling architecture sound, or is it fundamentally tangled?
3. Is the database layer stable, or are there schema drift issues?
4. Is the config system coherent, or are there conflicting/redundant paths?
5. How many features are genuinely broken vs how many are wiring issues?

## Final Verdict
One of exactly three options, with justification:

- **SALVAGEABLE** — The architecture is sound. Broken features are isolated wiring issues fixable in <40 hours without structural changes.
- **DEEP REPAIR** — The architecture is partially sound but has systemic wiring failures across multiple subsystems. 40-120 hours of focused repair required. List the subsystems that need rework.
- **REBUILD RECOMMENDED** — The architecture itself is the source of failures. Attempting repair would require replacing core subsystems, making a fresh build more efficient. Identify specifically which architectural decisions are unsalvageable and what should be preserved.

The verdict MUST cite the ratio of working:broken features, the nature of the failures (surface wiring vs structural), and whether the failures are concentrated in one subsystem or distributed across the codebase.
