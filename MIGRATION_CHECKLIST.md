# Floyd → Pi Migration Checklist

**Generated:** 2026-03-23  
**Decision:** Adopt Pi as Floyd replacement (MIT license, no forking required)  
**Pi version:** v0.61.1  
**Pi package:** `@mariozechner/pi-coding-agent` (homebrew)

---

## Architecture Mapping

| Floyd Concept | Pi Equivalent | Notes |
|---|---|---|
| `floyd.json` (single config) | `settings.json` + `models.json` + `auth.json` | Split into 3 files |
| `go:embed` prompt templates | Extensions (`before_agent_start`) + `AGENTS.md` + Skills + Prompt Templates | Fully user-editable |
| `FLOYD.md` (project context) | `AGENTS.md` (context file, walks up from cwd) | Direct equivalent |
| `.floyd/.supercache` (JSON state) | Pi session tree (JSONL) + custom extension state | Pi handles persistence natively |
| SuperFloyd binary (separate identity) | Pi project-scoped `AGENTS.md` + `settings.json` per directory | No separate binary needed |
| `mode.sh` (safe/balanced/beast) | Extension with env vars or `/prompt:mode` templates | Pi parallelism is already internal |
| Compiled-in tool descriptions | Built-in tools (read/edit/write/bash/grep/find) | Pi has these natively |
| LSP integration (gopls sidebar) | Extension (requires building) | Only real feature gap |
| `internal/agent/tools/*.md` (custom tools) | MCP servers + custom tool extensions | Port tool descriptions as MCP or extension |
| Floyd Skills (`extensibility/skills/`) | Pi Skills (`~/.pi/agent/skills/` or `.agents/skills/`) | Mostly compatible format |

---

## Phase 1: Identity & Prompting (Day 1)

This is the core of the migration. In Floyd, your prompting is baked into the Go binary via `go:embed`. In Pi, everything is user-editable at the filesystem level — no recompilation needed.

### 1.1 System Prompt — Floyd General Agent

**Source:** `/Volumes/Storage/floyd/internal/agent/templates/floyd-general.md.tpl` (33 lines)  
**Currently:** Compiled into binary via `go:embed` in `prompts.go:13`  
**Pi target:** `~/.pi/agent/extensions/floyd-identity/index.ts`

Action: Create a Pi extension that hooks `before_agent_start` and prepends the Floyd identity, boot summary protocol, and behavioral rules to `event.systemPrompt`.

Key content to migrate:
- Boot Summary protocol (3-line format: Active project, Last known status, Current intent)
- "No ceremony" / "No conversational filler" directive
- Lab tool references (remap to `floyd-persistent-lab` MCP)
- MCP tools injection block (Pi handles this automatically via configured servers)

### 1.2 System Prompt — SuperFloyd (Coding Agent)

**Source:** `/Volumes/Storage/floyd/internal/agent/templates/superfloyd-coder.md.tpl` (48 lines)  
**Currently:** Compiled into binary via `go:embed` in `prompts.go:16`  
**Pi target:** `~/.pi/agent/extensions/superfloyd-identity/index.ts`

Key content to migrate:
- "Elite force-multiplier" identity and code quality gates (4-gate checklist)
- "Every claim cites code evidence (path:line)"
- Fetch Anti-Bot rule (pivot to MCP on 403)
- "DEEP WORK" / "AUTONOMOUS MODE" directive
- `gofumpt` formatting rule → Pi already formats with appropriate formatters

**Pi alternative approach:** Instead of a separate extension, use project-scoped `AGENTS.md` files. When working in a coding project, the `AGENTS.md` contains the SuperFloyd coding rules. When working in general projects, it contains the Floyd general rules. This is the idiomatic Pi pattern.

### 1.3 Protocol Kernel

**Source:** `/Volumes/Storage/floyd/internal/agent/templates/protocol_kernel.md.tpl` (68 lines)  
**Currently:** Compiled into binary, always Layer 1 in the prompt stack  
**Pi target:** `~/.pi/agent/extensions/protocol-kernel/index.ts` or `~/.pi/agent/AGENTS.md` (global context)

Key content to migrate:
- EXECUTION MODEL (no pleasantries, no summaries, execute immediately)
- MODE SELECTOR (DEBUG / BUILD / EXPLORE classification)
- DEBUG MODE GATES (3-gate sequential debugging protocol)
- BUILD MODE GATES (5-rule build discipline)
- TOOL DISCIPLINE (large file handling, pivot strategy)
- ERROR RECOVERY (one fix per failure, then report)
- CONTEXT TRUST (cache is evidence not truth)
- CONTINUITY (checkpoint to persistent state)

**Recommendation:** Put this in `~/.pi/agent/AGENTS.md` as your global identity file. Pi loads `AGENTS.md` into context automatically for every session. This is the closest equivalent to Floyd's "always Layer 1" behavior.

### 1.4 Floyd Protocol Documentation

**Source:** `/Volumes/Storage/floyd/internal/agent/templates/floyd_protocol.md.tpl` (55 lines)  
**Currently:** Compiled into binary — describes the 3-layer architecture  
**Pi target:** This is internal documentation about how Floyd works. **Do not migrate** — it describes Floyd's architecture, not Pi's. Create an equivalent `PI_PROTOCOL.md` reference if needed, but Pi's system is extensions-based, not layered.

### 1.5 Project Context — FLOYD.md

**Source (general):** `/Volumes/Storage/floyd/FLOYD.md` (89 lines)  
**Source (SuperFloyd):** `/Volumes/Storage/floyd/SuperFloyd/FLOYD.md` (89 lines)  
**Currently:** User-editable project file, loaded via `go:embed` context injection  
**Pi target:** `AGENTS.md` in project root

Pi's `AGENTS.md` walks up from `cwd` to git root (or filesystem root), exactly like Floyd's `FLOYD.md`. Direct file-level equivalent.

Content to migrate from `FLOYD.md`:
- Project sovereignty rules (local memory, isolation, boot protocol)
- Core operational rules (read before edit, surgical edits, code integrity)
- Visual & documentation standards (Mermaid diagrams, `file:line` references, box tables)
- Domain-specific rules (sqlc, go fmt, table-driven tests)
- The `.supercache` → Pi session tree mapping (Pi handles natively, can be mentioned as "session history")

**Note:** Floyd's `FLOYD.md` references `.floyd/.supercache` for persistent state. Pi uses JSONL session trees with branching. The "boot protocol" concept maps to Pi's session resumption — just reference the session tree instead.

### 1.6 Deterministic Prompt Templates (SuperFloyd)

**Source (compiled):** `/Volumes/Storage/floyd/internal/agent/templates/deterministic/` (21 files, `00_` through `20_`)  
**Source (standalone):** `/Volumes/Storage/floyd/SuperFloyd/DETERMINISTIC_PROMPT_FRAMEWORK/` (20 template files + 3 meta files)  
**Currently:** Compiled into binary; standalone copies in SuperFloyd directory  
**Pi target:** `~/.pi/agent/prompts/` as slash-command prompt templates

Pi's prompt templates are invoked with `/name` in the editor. These 20 deterministic templates map directly:

| Floyd Template | Pi Prompt Template | Invoke As |
|---|---|---|
| `00_FULL_STACK_DETERMINISTIC_MASTER_PROMPT.md` | `~/.pi/agent/prompts/deterministic-master.md` | `/deterministic-master` |
| `01_Deterministic_Task_Kickoff.md` | `~/.pi/agent/prompts/task-kickoff.md` | `/task-kickoff` |
| `02_Complex_Implementation_Orchestration.md` | `~/.pi/agent/prompts/complex-impl.md` | `/complex-impl` |
| `03_Failure_Driven_Debugging.md` | `~/.pi/agent/prompts/debug-failure.md` | `/debug-failure` |
| `04_Regression_Bug_Hunt.md` | `~/.pi/agent/prompts/regression-hunt.md` | `/regression-hunt` |
| `05_Multi_File_Refactor.md` | `~/.pi/agent/prompts/multi-refactor.md` | `/multi-refactor` |
| `06_Test_First_Fix.md` | `~/.pi/agent/prompts/test-first-fix.md` | `/test-first-fix` |
| `07_Stability_Hardening.md` | `~/.pi/agent/prompts/hardening.md` | `/hardening` |
| `08_Release_Readiness_Audit.md` | `~/.pi/agent/prompts/release-audit.md` | `/release-audit` |
| `09_Commit_Preparation_and_Intent_Check.md` | `~/.pi/agent/prompts/commit-prep.md` | `/commit-prep` |
| `10_PR_Summary_and_TestPlan.md` | `~/.pi/agent/prompts/pr-summary.md` | `/pr-summary` |
| `11_Context_Window_Protection.md` | `~/.pi/agent/prompts/context-protect.md` | `/context-protect` |
| `12_Incident_Response_Production.md` | `~/.pi/agent/prompts/incident-response.md` | `/incident-response` |
| `13_Architecture_Decision_Record.md` | `~/.pi/agent/prompts/adr.md` | `/adr` |
| `14_Safe_Exploration_Tradeoff.md` | `~/.pi/agent/prompts/safe-explore.md` | `/safe-explore` |
| `15_Handoff_Continuity.md` | `~/.pi/agent/prompts/handoff.md` | `/handoff` |
| `16_10X_Execution_Proof_Enforcer.md` | `~/.pi/agent/prompts/proof-enforcer.md` | `/proof-enforcer` |
| `17_10X_Reality_Benchmark_Parity.md` | `~/.pi/agent/prompts/benchmark-parity.md` | `/benchmark-parity` |
| `18_State_Drift_Detector.md` | `~/.pi/agent/prompts/drift-detector.md` | `/drift-detector` |
| `19_Zero_Loss_Context_Handoff.md` | `~/.pi/agent/prompts/context-handoff.md` | `/context-handoff` |
| `20_Surgical_Change_MinRisk.md` | `~/.pi/agent/prompts/surgical-change.md` | `/surgical-change` |

**Action:** Copy each file to `~/.pi/agent/prompts/` with a clean name, add the frontmatter `description` field for autocomplete. Pi supports `$1`, `$2` positional args and `$@` for all arguments in templates.

---

## Phase 2: Configuration & Providers (Day 1)

### 2.1 Provider Configuration

**Floyd source:** `/Volumes/Storage/floyd/floyd.json` → `providers` block  
**SuperFloyd source:** `/Volumes/Storage/floyd/SuperFloyd/floyd-config-backup.json` → `providers` + `models` blocks  
**Pi target:** `~/.pi/agent/settings.json` (default provider/model) + `~/.pi/agent/models.json` (custom providers) + `~/.pi/agent/auth.json` (API keys)

**Already configured in Pi:**
- ✅ `zai` provider (GLM-5/GLM-5-turbo) — built-in, `defaultProvider: "zai"`, `defaultModel: "glm-5-turbo"`
- ✅ `anthropic` provider — auth in `auth.json`
- ✅ `openrouter` provider — auth in `auth.json`
- ✅ `mistral-openai` — custom provider in `models.json`
- ✅ `deepseek` — custom provider in `models.json`
- ✅ `opencode` provider — auth in `auth.json`

**Needs migration:**
- [ ] `z-ai` provider (Claude Sonnet 4 via Z.AI Anthropic proxy)  
  Floyd config: `type: "anthropic"`, `base_url: "https://api.z.ai/api/anthropic"`  
  Pi target: Add to `~/.pi/agent/models.json` as custom provider with `api: "anthropic"` and custom `baseUrl`

### 2.2 Model Parameters

**Floyd source:** `floyd.json` → `providers.zai.models[0]`  
**Pi target:** `~/.pi/agent/settings.json` or `~/.pi/agent/models.json`

Floyd-specific model config:
- `context_window: 204800` — Pi uses model registry defaults (GLM-5 reports 200K+ via Catwalk, but Z.AI reports 1M — Pi will respect model's self-reported window)
- `default_max_tokens: 32768` — Pi handles natively
- `temperature: 0.1` — Pi handles natively
- `can_reason: false` — Pi handles natively
- `extra_body.thinking.type: "disabled"` — Pi handles via `thinkingLevel` setting

**Currently in Pi:** `defaultThinkingLevel: "high"` — may need adjustment to match Floyd's `thinking: disabled` for GLM models.

### 2.3 Floyd Rules → Pi AGENTS.md

**Floyd source:** `/Volumes/Storage/floyd/floyd.json` → `rules` array (6 rules)  
**Pi target:** Project-level `AGENTS.md` (for Floyd repo specifically) or global `~/.pi/agent/AGENTS.md`

Rules to migrate:
1. `NO_CHAT` — Be concise, no filler → Already covered by Protocol Kernel
2. `DATABASE` — sqlc workflow → Project-specific `AGENTS.md` in Floyd repo
3. `FRONTEND` — UI layer split → Project-specific `AGENTS.md` in Floyd repo
4. `TESTING` — `go test ./...` → Project-specific `AGENTS.md`
5. `STYLE` — `go fmt`, no dot imports, table-driven tests → Project-specific `AGENTS.md`
6. `ARCHITECTURE` — Go CLI + Bubble Tea + MCP → Project-specific `AGENTS.md`

**Recommendation:** Rules 2-6 are Floyd-repo-specific. They belong in an `AGENTS.md` placed in `/Volumes/Storage/floyd/` (or wherever you work on the Floyd codebase). Rule 1 (NO_CHAT) is global and belongs in `~/.pi/agent/AGENTS.md`.

### 2.4 SuperFloyd Model: Small Model

**Floyd source:** `/Volumes/Storage/floyd/SuperFloyd/floyd-config-backup.json` → `models.small`  
**Currently:** `glm-4.5-air` with 131072 context window  
**Pi target:** Pi supports model cycling with `Ctrl+P`. GLM-4.5-air may need to be added as a custom model in `models.json` under the `zai` provider, or it may already be discoverable.

---

## Phase 3: MCP Servers (Day 1-2)

### 3.1 Already Configured in Pi

These are already in `~/.pi/agent/settings.json` and working:

| Server | Pi Config | Status |
|---|---|---|
| `floyd-persistent-lab` | stdio → `/Volumes/Storage/floyd/floyd-mcp --mcp` | ✅ Active |
| `open-anvil` | stdio → `node /Volumes/Storage/A-TEAM/open-anvil/mcp-server/server.js` | ✅ Active |
| `zai-mcp-server` | SSE → `https://api.z.ai/api/mcp/zai/mcp` | ✅ Active |

### 3.2 SuperFloyd-Only MCP Servers (Needs Decision)

These are configured in SuperFloyd but NOT in current Pi. Decide which to bring over:

| Server | Type | Endpoint | Recommendation |
|---|---|---|---|
| `lab-lead` | stdio | `/Volumes/Storage/MCP/lab-lead-server/dist/index.js` | **Migrate** — lab orchestration |
| `floyd-runner` | stdio | `/Volumes/Storage/FLOYD_CLI/dist/mcp/runner-server.js` | **Evaluate** — Floyd-specific CLI runner |
| `floyd-git` | stdio | `/Volumes/Storage/FLOYD_CLI/dist/mcp/git-server.js` | **Skip** — Pi has built-in git via bash |
| `floyd-explorer` | stdio | `/Volumes/Storage/FLOYD_CLI/dist/mcp/explorer-server.js` | **Evaluate** — may duplicate `find`/`grep` |
| `floyd-patch` | stdio | `/Volumes/Storage/FLOYD_CLI/dist/mcp/patch-server.js` | **Evaluate** — Pi has built-in edit |
| `floyd-devtools` | stdio | `/Volumes/Storage/MCP/floyd-devtools-server/dist/index.js` | **Evaluate** — depends on Go tooling needs |
| `floyd-supercache` | stdio | `/Volumes/Storage/MCP/floyd-supercache-server/dist/index.js` | **Skip** — Pi has native session persistence |
| `floyd-safe-ops` | stdio | `/Volumes/Storage/MCP/floyd-safe-ops-server/dist/index.js` | **Evaluate** — safety gating |
| `floyd-terminal` | stdio | `/Volumes/Storage/MCP/floyd-terminal-server/dist/index.js` | **Skip** — different path than Pi's `floyd-persistent-lab` |
| `gemini-tools` | stdio | `/Volumes/Storage/MCP/gemini-tools-server/dist/index.js` | **Migrate** — useful multi-model tooling |
| `pattern-crystallizer-v2` | stdio | `/Volumes/Storage/MCP/pattern-crystallizer-v2/dist/index.js` | **Evaluate** — may be replaced by skills |
| `context-singularity-v2` | stdio | `/Volumes/Storage/MCP/context-singularity-v2/dist/index.js` | **Skip** — Pi handles compaction natively |
| `hivemind-v2` | stdio | `/Volumes/Storage/MCP/hivemind-v2/dist/index.js` | **Evaluate** — multi-agent orchestration |
| `omega-v2` | stdio | `/Volumes/Storage/MCP/omega-v2/dist/index.js` | **Evaluate** — unknown purpose, review first |
| `novel-concepts` | stdio | `/Volumes/Storage/MCP/novel-concepts-server/dist/index.js` | ✅ Already in Floyd main config |
| `web-search-prime` | streamable-http | `https://api.z.ai/api/mcp/web_search_prime/mcp` | **Migrate** — search capability |
| `web-reader` | streamable-http | `https://api.z.ai/api/mcp/web_reader/mcp` | **Migrate** — web scraping |
| `zread` | streamable-http | `https://api.z.ai/api/mcp/zread/mcp` | **Migrate** — Z.AI reading |
| `4_5v_mcp` | streamable-http | `https://api.z.ai/api/mcp/4_5v_mcp/mcp` | **Evaluate** — GLM-4.5V vision? |
| `rube` | streamable-http | `https://rube.app/mcp` | **Evaluate** — third-party service |

### 3.3 Floyd Main MCP Servers

**Source:** `/Volumes/Storage/floyd/floyd.json` → `mcp` block

| Server | Config | Status |
|---|---|---|
| `floyd-terminal` | stdio → `mcp-servers/floyd-terminal/dist/index.js` | Already in Pi as `floyd-persistent-lab` |
| `novel-concepts` | stdio → `/Volumes/Storage/MCP/novel-concepts-server/dist/index.js` | **Add to Pi settings.json** |

### 3.4 MCP Transport Mapping

Floyd SuperFloyd uses `streamable-http` for Z.AI MCP servers. Pi supports:
- `stdio` — local Node.js/Python servers
- `sse` — Server-Sent Events (already configured for `zai-mcp-server`)
- `streamable-http` — supported in Pi's MCP client

Add streamable-http servers to `settings.json`:
```json
{
  "mcpServers": {
    "web-search-prime": {
      "url": "https://api.z.ai/api/mcp/web_search_prime/mcp",
      "transport": "streamable-http",
      "headers": { "Authorization": "Bearer $ZAI_MCP_TOKEN" }
    }
  }
}
```

---

## Phase 4: Skills Migration (Day 2-3)

### 4.1 Floyd Skills Inventory

**Source:** `/Volumes/Storage/floyd/extensibility/skills/` (22 categories, 157 .md files)  
**Pi target:** `~/.pi/agent/skills/` (global) or `.agents/skills/` (project-scoped)

Pi skills use `SKILL.md` with YAML frontmatter. Floyd skills may need minor format conversion. The majority are plain Markdown instructions the agent reads on-demand.

**Categories to review:**
- `core/` — likely overlaps with Pi built-in tools
- `debugging/` — migration candidates (DEBUG gates are already in Protocol Kernel)
- `development/` — migration candidates
- `git/` — likely redundant (Pi has bash + git)
- `build/` — project-specific (sqlc, go build — belongs in project `AGENTS.md`)
- `design/` — creative tools, probably reusable
- `ai/`, `browser/`, `data/`, `media/` — evaluate individually
- `linting/` — has `SKILL.md` format already! (`lint-fix-go/SKILL.md`, `dependency-unused-cleanup/SKILL.md`)
- `refactoring/` — migration candidates

### 4.2 Already-Compatible Skills

These Floyd skills already use Pi-compatible `SKILL.md` format:
- `/Volumes/Storage/floyd/extensibility/skills/linting/lint-fix-go/SKILL.md`
- `/Volumes/Storage/floyd/extensibility/skills/linting/dependency-unused-cleanup/SKILL.md`
- `/Volumes/Storage/floyd/extensibility/skills/design/theme-factory/SKILL.md`

**Action:** Copy these directories directly into `~/.pi/agent/skills/`.

### 4.3 SuperFloyd Deterministic Framework as Skills

The 20 deterministic prompt templates (Phase 1.6) can optionally also be registered as skills for progressive disclosure:
- Skill description shows in system prompt
- Full template loads on-demand when agent calls `/skill:name`
- This keeps the system prompt small while making templates available

---

## Phase 5: SuperFloyd Identity & Mode System (Day 2-3)

### 5.1 Separate Identity Without Separate Binary

Floyd uses a separate compiled binary (`superfloyd`) with `isSuperFloyd()` detection in `logo.go`. Pi doesn't need this.

**Pi approach options:**

**Option A — Project-scoped `AGENTS.md` (Recommended)**
- In coding project directories: `AGENTS.md` contains SuperFloyd coding rules
- In general project directories: `AGENTS.md` contains Floyd general rules
- Pi automatically loads the right `AGENTS.md` based on `cwd`
- No separate binary, no mode switching

**Option B — Extension with `/mode` command**
- Create `~/.pi/agent/extensions/floyd-modes/index.ts`
- Register a `/mode` command that toggles behavioral profile
- Sets status indicator in footer
- Switches system prompt additions for the session

**Option C — Prompt template switching**
- `/prompt:superfloyd` loads coding agent template
- `/prompt:floyd` loads general agent template
- Templates set behavior for the current session via instructions

### 5.2 Mode System (safe/balanced/beast)

**Floyd source:** `/Volumes/Storage/floyd/SuperFloyd/mode.sh` — sets env vars `SUPERFLOYD_QUALITY_GATES`, `SUPERFLOYD_MAX_PARALLEL`, etc.  
**Pi target:** Extension or prompt template

Floyd modes control:
- `MAX_PARALLEL` (6/12/24) — Pi manages parallelism internally
- `QUALITY_GATES` (on/off) — Protocol Kernel handles this
- `DEGRADATION_CONTROLS` — Pi doesn't need this (no session death)
- `CONSISTENCY_LOCK` — Pi's session tree handles this
- `AUTOSTABILIZE` — Pi doesn't need this (no session death)

**Recommendation:** The mode system was largely a workaround for Floyd's context management bugs (Bug #10, #13). Pi solves these by design. The "quality gates" concept maps to the Protocol Kernel's DEBUG/BUILD gates, which are always active. You likely don't need a mode system in Pi.

---

## Phase 6: Tool Descriptions & Custom Tools (Day 2-3)

### 6.1 Floyd Custom Tool Descriptions

**Source:** `/Volumes/Storage/floyd/internal/agent/tools/*.md` (24 tool description files)  
**Pi equivalent:** Pi has built-in tools (read, write, edit, bash, grep, find). Custom tools are either MCP servers or extension-registered tools.

**Mapping:**

| Floyd Tool | Pi Equivalent | Action |
|---|---|---|
| `view.md` | Built-in `read` tool | None — Pi has this |
| `edit.md` | Built-in `edit` tool | None — Pi has this |
| `write.md` | Built-in `write` tool | None — Pi has this |
| `bash.tpl` | Built-in `bash` tool | None — Pi has this |
| `grep.md` | Built-in `grep` (bash) | None — Pi has this |
| `glob.md` | Built-in `find` (bash) | None — Pi has this |
| `ls.md` | Built-in `ls` (bash) | None — Pi has this |
| `fetch.md` | MCP `open-anvil` | Already configured |
| `web_search.md` | MCP `web-search-prime` | Add to Pi config |
| `web_fetch.md` | MCP `web-reader` | Add to Pi config |
| `multiedit.md` | Built-in `edit` (multiple calls) | None — Pi batches edits |
| `smart_replace.md` | Built-in `edit` | None — Pi has find-and-replace |
| `download.md` | Bash `curl` | None — trivial |
| `project_map.md` | Bash `find` + `tree` | None — trivial |
| `todos.md` | Extension (optional) | Build if needed |
| `get_active_diff.md` | Bash `git diff` | None — trivial |
| `apply_unified_diff.md` | Bash `git apply` | None — trivial |
| `diagnostics.md` | LSP extension (future) | Build if needed |
| `list_symbols.md` | LSP extension (future) | Build if needed |
| `references.md` | LSP extension (future) | Build if needed |
| `lsp_restart.md` | LSP extension (future) | Build if needed |
| `manage_scratchpad.md` | Extension (optional) | Build if needed |
| `job_kill.md` | Bash `kill` | None — trivial |
| `job_output.md` | Bash output capture | None — trivial |
| `sourcegraph.md` | MCP or bash | Evaluate |

### 6.2 Tool Templates with Special Behavior

**Source:** `/Volumes/Storage/floyd/internal/agent/tools/bash.tpl`  
Floyd's bash tool template has special self-healing build checks for Go projects (auto-runs `go build` after `.go` edits).  

**Pi equivalent:** Build as an extension that hooks `after_tool_call` and checks for Go build success after edit operations. Optional — Pi already has the concept of tool result validation.

---

## Phase 7: LSP Integration (Day 3-5, Optional)

### 7.1 gopls Configuration

**Floyd source:** `/Volumes/Storage/floyd/floyd.json` → `lsp.gopls` (full config with analyses, codelenses, hints, staticcheck)  
**Pi target:** Extension that wraps gopls LSP client

Floyd has a rich gopls integration:
- `gofumpt` formatting
- Code lenses (gc_details, generate, run_govulncheck, test, tidy, upgrade_dependency)
- Hints (assignVariableTypes, compositeLiteralFields, etc.)
- Analyses (nilness, unusedparams, unusedvariable, etc.)
- Staticcheck enabled
- Semantic tokens for syntax highlighting

**Pi does not have built-in LSP support.** This is the only significant feature gap.

**Options:**
1. **Build a Pi extension** (~2-3 days) that connects to gopls via stdio and surfaces diagnostics as widgets or status indicators
2. **Use gopls manually** via bash (`gopls check`, `gopls diagnostics`) when needed
3. **Drop it** — for general coding tasks, Pi's built-in tools + model knowledge are often sufficient

---

## Phase 8: UI & Branding (Day 3-5)

### 8.1 ASCII Art

**5 ASCII art files** — user handles location directly.

**Pi targets:**
- `ctx.ui.setHeader()` — custom header component
- `ctx.ui.setFooter()` — custom footer/HUD
- `ctx.ui.setWidget()` — persistent widgets above/below editor
- Overlay components — floating panels

### 8.2 Logo / Branding

**Source:** `/Volumes/Storage/floyd/internal/ui/logo/logo.go` (524 lines)  
Contains:
- Floyd wordmark (braille-based, compact)
- SuperFloyd wordmark (block-char "SUPERFLOYD" + "FLOYD'S LABS")  
- SmallRender (sidebar variant)
- SidebarRender (gradient sidebar logo)
- PersistentBar (6-row ASCII bar)

These are compiled into the binary with gradient coloring via lipgloss. In Pi, you'd render them as custom TUI components with theme colors.

### 8.3 HUD / Navigation Bar

Pi's TUI system supports:
- **Custom footer** (`setFooter`) — persistent bottom bar with model, git branch, token stats
- **Status indicator** (`setStatus`) — persistent status text in footer
- **Widgets** (`setWidget`) — content above/below editor
- **Overlays** — floating panels with anchor positioning (left, right, center)

Build a `~/.pi/agent/extensions/floyd-hud/index.ts` extension.

### 8.4 Theme

**Source:** Floyd uses Charm lipgloss styling (primary, secondary, background colors)  
**Pi target:** `~/.pi/agent/themes/floyd-dark.json` — JSON theme file with color tokens

Pi themes define colors for: text, accent, muted, borders, messages, tools, diffs, markdown, syntax highlighting, thinking levels, and more.

---

## Phase 9: Persistent State & Session (Day 1, Ongoing)

### 9.1 Supercache → Pi Sessions

**Floyd source:** `.floyd/.supercache` (JSON file with project_name, last_status, last_intent, last_updated)  
**Pi equivalent:** Pi's JSONL session tree with `id`/`parentId`/`leaf` pointers

Pi handles session persistence natively. The "boot summary" concept (reading supercache on init) maps to Pi's session resumption — the agent sees the conversation history when a session is resumed.

### 9.2 CHIMERA Session Export

**Source:** `/Volumes/Storage/CHIMERA/.floyd/exports/session-export-2026-03-22-195359.md` (5948 lines)  
**Note:** This is a dead session that died from Bug #10/#13. Valuable as reference material but does not need migration.

---

## Phase 10: Cleanup & Validation (Day 5)

### 10.1 Validation Checklist

- [ ] Pi starts with zai/glm-5-turbo without errors
- [ ] `AGENTS.md` loads and behavioral rules are followed (no pleasantries, boot summary)
- [ ] Deterministic prompt templates accessible via `/` autocomplete
- [ ] All needed MCP servers connect successfully
- [ ] Session compaction triggers correctly (verify with long session)
- [ ] Custom header/footer renders correctly
- [ ] Theme applies correctly
- [ ] Skills load and are accessible via `/skill:name`
- [ ] Floyd repo `AGENTS.md` project rules apply when working in `/Volumes/Storage/floyd/`

### 10.2 Known Gaps After Migration

| Feature | Status | Mitigation |
|---|---|---|
| LSP/gopls diagnostics | Not in Pi | Build extension or use bash |
| Go self-healing build checks | Not in Pi | Build `after_tool_call` extension |
| Separate SuperFloyd binary identity | Not needed | Project-scoped `AGENTS.md` |
| Mode system (safe/balanced/beast) | Not needed | Pi handles parallelism natively |
| `.supercache` persistent JSON state | Different model | Pi sessions provide equivalent |
| `floyd_protocol.md.tpl` architecture reference | Floyd-specific | Create Pi equivalent if needed |

---

## Quick Reference: File Locations

### Floyd Sources (read-only, for migration)
```
/Volumes/Storage/floyd/
├── floyd.json                                          # Main config (providers, rules, MCP, LSP)
├── FLOYD.md                                            # Project context (89 lines)
├── assets/FLOYD_ASCII.txt                              # ASCII art #1
├── internal/
│   ├── agent/
│   │   ├── templates/
│   │   │   ├── protocol_kernel.md.tpl                  # Layer 1: Execution model (68 lines)
│   │   │   ├── floyd-general.md.tpl                    # Layer 2a: General agent (33 lines)
│   │   │   ├── superfloyd-coder.md.tpl                 # Layer 2b: Coding agent (48 lines)
│   │   │   ├── floyd_protocol.md.tpl                   # Architecture reference (55 lines)
│   │   │   ├── summary.md                              # Compaction template (48 lines)
│   │   │   ├── initialize.md.tpl                       # First-run template (48 lines)
│   │   │   ├── task.md.tpl                             # Task sub-agent template (8 lines)
│   │   │   └── deterministic/                          # 20 deterministic templates
│   │   │       ├── 00_FULL_STACK_DETERMINISTIC_MASTER_PROMPT.md
│   │   │       ├── 01_ through 20_ ...
│   │   │       └── README.md
│   │   ├── prompts.go                                  # go:embed declarations
│   │   └── tools/*.md                                  # 24 tool description files
│   └── ui/logo/logo.go                                 # All ASCII art (compiled, 524 lines)
├── extensibility/skills/                               # 22 categories, 157 .md files
│   ├── linting/lint-fix-go/SKILL.md                    # Pi-compatible format
│   ├── linting/dependency-unused-cleanup/SKILL.md      # Pi-compatible format
│   └── design/theme-factory/SKILL.md                   # Pi-compatible format
└── SuperFloyd/
    ├── floyd-config-backup.json                        # SuperFloyd config (22 MCP servers)
    ├── FLOYD.md                                        # Coding project context (89 lines)
    ├── mode.sh                                         # safe/balanced/beast launcher
    └── DETERMINISTIC_PROMPT_FRAMEWORK/                 # Standalone copy of 20 templates
        ├── 00_ through 20_ ...
        └── deterministic_prompt_templates/
            ├── PROMPT_TEMPLATE.md                      # Template writing guide (295 lines)
            └── DETERMINISTIC_PROMPT_WRITING_AGENT.md   # Agent for writing templates (211 lines)
```

### Pi Targets (write to)
```
~/.pi/agent/
├── settings.json                   # Main config (providers, MCP, compaction, theme)
├── models.json                     # Custom providers (deepseek, mistral, z-ai proxy)
├── auth.json                       # API keys (already populated)
├── AGENTS.md                       # Global identity + Protocol Kernel (NEW)
├── themes/
│   └── floyd-dark.json             # Custom theme (NEW)
├── prompts/                        # Slash-command templates (NEW, 20+ files)
│   ├── deterministic-master.md
│   ├── task-kickoff.md
│   ├── complex-impl.md
│   └── ...
├── skills/                         # Migrated skills (NEW)
│   ├── lint-fix-go/SKILL.md        # Copy from Floyd
│   └── ...
└── extensions/                     # Custom extensions (NEW)
    ├── floyd-identity/index.ts     # System prompt injection
    ├── floyd-hud/index.ts          # Header/footer/HUD
    └── floyd-modes/index.ts        # Optional mode switching

<Volumes/Storage/floyd>/
└── AGENTS.md                       # Floyd repo project rules (NEW)
```

---

## Estimated Timeline

| Phase | Scope | Effort | Dependencies |
|---|---|---|---|
| **1: Prompting** | Protocol Kernel + 2 identities + 20 templates | 4-6 hours | None |
| **2: Config** | Provider, rules, model params | 1-2 hours | None |
| **3: MCP Servers** | Evaluate + configure 10-12 servers | 2-3 hours | Phase 2 |
| **4: Skills** | Review + convert 157 skills | 3-5 hours | Phase 1 |
| **5: Identity/Mode** | AGENTS.md strategy, optional mode extension | 2-3 hours | Phase 1 |
| **6: Custom Tools** | Verify tool coverage, build gaps | 1-2 hours | Phase 3 |
| **7: LSP** (optional) | Build gopls extension | 2-3 days | Phase 1 |
| **8: UI/Branding** | Header, footer, HUD, theme | 3-5 hours | Phase 1 |
| **9: Validation** | Test all migrated components | 2-3 hours | All phases |

**Working agent:** End of Phase 2 (~3 hours)  
**Full parity (no LSP):** End of Phase 8 (~2-3 days)  
**Full parity (with LSP):** End of Phase 7 (~5 days)
