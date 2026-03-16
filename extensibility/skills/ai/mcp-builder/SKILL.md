---
name: mcp-builder
description: Build skill for MCP (Model Context Protocol) tool servers and clients. Use this skill when designing or adding a new MCP server (Patch, Runner, Browser, Git), adding tools to existing servers, standardizing tool schemas and safety controls, wiring MCP into a CLI runtime, or creating shared tool contracts. Outputs complete implementation-ready packages: tool catalog spec, transport plan, server blueprint, safety gate, test harness, and integration hooks. Framework-agnostic (Node/Rust/Python) with validation-first and least privilege standards.
---

# MCP BUILDER (2026) — Servers, Tools, Transport, and Safety Contracts

## Skill Mission

Turn any "we need an MCP tool for X" request into a safe, testable, production-ready MCP server plus tool spec that can be wired into a CLI/agent runtime without surprises.

---

## What This Skill Produces (Guaranteed)

When invoked, MCP BUILDER outputs a complete, implementation-ready package including:

- **Tool Catalog Spec** — Names, schemas, permissions, rate limits
- **Transport Plan** — stdio vs HTTP/SSE vs websockets; reconnect rules
- **Server Blueprint** — Routing, validation, logging, errors
- **Safety Gate** — Policy, protected paths, allow/deny persistence hooks
- **Test Harness** — Smoke tests plus contract tests plus failure simulations
- **Integration Hooks** — Client call shape plus event emission expectations

**Framework-agnostic:** Node, Rust, or Python are all acceptable. Assumes 2026-grade standards: validation-first, least privilege, deterministic tooling.

---

## When to Invoke

Call MCP BUILDER when you need to:

- Design or add a new MCP server (Patch server, Runner server, Browser server, Git server)
- Add a new tool to an existing MCP server
- Standardize tool schemas and safety controls across agents
- Wire MCP into a CLI runtime (stream results, progress, permission prompts)
- Create a shared "tool contract" that multiple agents must follow

---

## Required Inputs

Caller provides these fields (best effort — skill will proceed with assumptions if missing):

### A) Context

| Field | Description |
|-------|-------------|
| **Project Type** | CLI / agent runtime / both |
| **Execution environment** | OS targets, sandbox constraints, CI constraints |

### B) Tooling Need

| Field | Description |
|-------|-------------|
| **Tool goal** | What action the tool performs |
| **Risk level** | low / medium / high (file writes, network, exec are high) |
| **Target resources** | Filesystem paths, repo root, commands allowed, URLs allowed |

### C) Integration

| Field | Description |
|-------|-------------|
| **Client** | Which runtime calls MCP (manager agent, UI, automation) |
| **Event bus** | How progress/errors are surfaced (IPC, logs, UI overlay) |

### Default Assumptions (if not provided)

- Multi-agent environment
- Tools may be invoked repeatedly
- Tight validation required
- Permission prompts exist

---

## Mandatory Output Format

MCP BUILDER always outputs these sections in order:

### A) Server Role and Scope

What this server is responsible for; explicit exclusions.

### B) Tool Catalog

For each tool:
- name, description
- input_schema (typed shape)
- output_schema
- side_effects (none/read/write/exec/network)
- permissions_required
- rate_limit plus concurrency_limit
- idempotency notes (safe retries?)
- failure_modes (top 3)

### C) Transport and Lifecycle

Transport choice (stdio/HTTP/SSE/etc.) and why; connection lifecycle; reconnect strategy; heartbeat/healthcheck.

### D) Safety Gate (Non-Negotiable)

Validation pipeline; protected paths; command allowlist; network allowlist; permission escalation flow.

### E) Observability

Structured logs; metrics; trace hooks.

### F) Tests and Verification

Smoke tests; contract tests; chaos tests; security tests.

### G) Integration Notes

Client call pattern; event payloads; error mapping.

### H) Definition of Done Checklist

Hard gates for merging.

---

## 2026 MCP Best Practices (Core Rules)

### 4.1 Tool Contracts Must Be Deterministic

- Inputs are fully typed and validated
- Outputs are structured; free-form text only as message fields
- Errors have stable code plus message plus optional details

### 4.2 Least Privilege By Default

- Server starts in deny-by-default mode
- Tools declare their side effects explicitly
- Any write/exec/network requires explicit permissions

### 4.3 Make Tools Idempotent Where Possible

- Support `dry_run: true` for write tools (patch/apply) if feasible
- Include `request_id` so clients can safely retry
- Prefer atomic operations (write temp, then rename)

### 4.4 Progress Is First-Class

Every long-running tool must support:
- Progress events (0–100 or step-based)
- Stream output chunks (logs)
- Final structured result

### 4.5 Concurrency Controls Are Mandatory

- Global concurrency cap per server
- Per-tool concurrency cap
- Backpressure behavior defined (queue vs reject)

---

## MCP Server Archetypes (Recommended)

Use separate servers for isolation:

### 5.1 Patch Server (high-risk write)

- `apply_diff`, `create_file`, `move_file`, `delete_file` (often restricted)
- Must support `dry-run` plus protected paths

### 5.2 Runner Server (exec)

- `detect_project`, `run_tests`, `run_build`, `lint`, `typecheck`
- Strict allowlist of commands plus cwd constraints

### 5.3 Git Server (write-ish)

- `status`, `diff`, `commit`, `branch`, `stash`
- Commit message policy and hooks

### 5.4 Browser/Fetch Server (network)

- `fetch_url`, `search`, `download_docs`
- Strict domain allowlist

---

## Gold Standard Safety Gate Spec (Drop-In)

All tools go through the same pipeline:

```
1. Schema Validate
   Reject unknown fields.
   Reject missing required fields.
   Enforce max lengths.

2. Normalize
   Resolve paths to absolute.
   Resolve repo root.
   Trim strings.

3. Policy Evaluate
   Check side effects vs permissions.
   Check protected paths.
   Check command allowlist.
   Check network allowlist.
   Check concurrency plus rate limit.

4. Permission Decision
   If policy requires user permission:
   Return NEEDS_PERMISSION with structured prompt payload.

5. Execute
   Run tool with timeouts, streaming logs.

6. Post-Validate Output
   Ensure output schema matches.
   Ensure no secret leaks.

7. Emit Events
   Emit start/progress/end/error events for UI plus logs.
```

---

## Standard Tool Schema Template (Reusable)

Use this template for every tool in your catalog:

```
Tool: <tool_name>

Description: <clear, concise>

Side Effects: none | read | write | exec | network

Input Schema:
  request_id: string (required)
  dry_run: boolean (optional; required for write tools if feasible)
  timeout_ms: number (optional; capped)
  cwd: string (optional; must be within repo root if used)
  <tool-specific fields>...

Output Schema:
  request_id: string
  ok: boolean
  code: string (stable)
  message: string (human summary)
  data: object (tool result, structured)
  artifacts?: [] (files created/modified, if applicable)
  logs?: [] (bounded)

Failure Modes: <top 3 scenarios>

Permissions Required: <scopes>

Concurrency/Rate Limits: <caps>
```

---

## Example: MCP BUILDER Output (Patch Server Minimal Catalog)

### A) Server Role and Scope

**Role:** Apply structured changes to repo files safely.

**Excludes:** Running commands, network, git operations.

### B) Tool Catalog

#### Tool: `apply_unified_diff`

| Field | Value |
|-------|-------|
| Side Effects | write |
| Input | `request_id` (req), `diff_text` (req, max size), `dry_run` (opt, default true), `base_dir` (req) |
| Output | `request_id`, `ok`, `code`, `message`, `data.modified_files: string[]`, `data.hunks_applied: number`, `data.rejected_hunks: number` |
| Permissions | `write:repo` |
| Idempotency | Safe if same diff re-applied; tool must detect already-applied hunks |
| Failure Modes | Invalid diff, protected path blocked, conflict/hunk reject |

#### Tool: `create_file`

| Field | Value |
|-------|-------|
| Side Effects | write |
| Input | `path`, `content`, `dry_run`, `request_id` |
| Safety | Path must be under repo root; protected patterns denied |
| Output | Artifact list with file path |

#### Tool: `read_file`

| Field | Value |
|-------|-------|
| Side Effects | read |
| Input | `path`, `request_id` |
| Output | `data.content` (bounded; support truncation) |

### C) Transport and Lifecycle

Prefer **stdio** when co-located with CLI (lowest friction, best local security).

**Reconnect:** Client restarts server on crash; exponential backoff.

### D) Safety Gate

- Deny writes to `.env*`, `credentials`, keychains, SSH dirs
- Require explicit permission token for write tools
- Enforce max diff size and max files touched per request

### E) Observability

Log each call with: `tool`, `request_id`, `duration`, `ok`, `files_touched`

Emit events: `tool.start`, `tool.progress`, `tool.end`, `tool.error`

### F) Tests

1. Apply diff dry-run (no writes)
2. Apply diff to protected path (must block)
3. Create file under root (ok)
4. Create file outside root (block)

---

## Definition of Done (MCP BUILDER)

A server/tool set is "merge-ready" only if:

- [ ] Every tool has input/output schemas and stable error codes
- [ ] Safety gate is enforced before execution
- [ ] Concurrency/rate limits exist and are tested
- [ ] Progress/events are emitted for long operations
- [ ] Smoke plus contract plus security tests pass
- [ ] Documentation exists: tool catalog plus permissions matrix

---

## Skill Invocation Template

```
SKILL: MCP BUILDER (2026)

Context:
Tool or Server Needed:
Risk Level:
Target Resources (paths/commands/domains):
Client plus Event Bus:
Constraints:

Return:
A) Server Role and Scope
B) Tool Catalog
C) Transport and Lifecycle
D) Safety Gate
E) Observability
F) Tests and Verification
G) Integration Notes
H) Definition of Done
```

---

## Related Skills

- **cli-x-2026** — Use this skill when designing UI for MCP tool interactions
- **ink-performance-auditor** — Use this skill when MCP tools cause high event throughput
- **chrome-extension-bridge** — Use when building MCP bridges for browser extension access
