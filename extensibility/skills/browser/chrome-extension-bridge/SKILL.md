---
name: chrome-extension-bridge
description: Build skill for creating secure Chrome extension to local CLI bridges. Use this skill when building an extension that triggers local CLI tools (run tests, apply patches, lint, search), showing agent/tool status in the browser, adding one-click repo actions from GitHub/GitLab pages, or establishing secure local authorization between browser and CLI. Outputs architecture choice, threat/permission models, message contracts, MV3 extension blueprint, CLI bridge design, UX flows, test plans, and hard merge gates. Repo-agnostic: works for Floyd or any tool-oriented CLI.
---

# Chrome Extension Bridge Builder (Floyd to Chrome to CLI Tools)

## Skill Mission

Turn "add a Chrome extension that links to the CLI for tools" into a secure, shippable architecture plus implementation plan with minimal friction and strong permissions.

**Repo-agnostic:** Works for Floyd or any tool-oriented CLI.

---

## What This Skill Produces (Guaranteed)

When invoked, this skill outputs:

- **Architecture choice** — Native Messaging vs Local WebSocket/HTTP bridge
- **Threat model plus permission model** — Deny-by-default, user consent, allowlists
- **Extension design** — MV3, service worker, content scripts, UI surfaces
- **Local bridge design** — CLI daemon/bridge, auth handshake, request routing
- **Message contract** — Request/response schemas, event streaming, error codes
- **UX flows** — Connect, approve, run tool, view logs, revoke access
- **Test plan** — Unit, integration, security tests, failure-mode simulations
- **Definition of Done** — Hard merge gates

---

## When to Invoke

Use this skill whenever you need:

- A Chrome extension that triggers local CLI tools (run tests, apply patches, lint, search)
- A browser UI that shows agent/tool status from the CLI
- One-click repo actions from GitHub/GitLab pages or local web apps
- Secure local authorization between browser and CLI

---

## Required Inputs

Caller provides (best effort; skill will assume defaults if missing):

### A) Feature Intent

| Field | Description |
|-------|-------------|
| **What the extension should do** | Commands/tools needed |
| **Which pages it acts on** | GitHub, localhost app, any site |
| **Content script access needed** | Page access or just popup UI |

### B) Local CLI Capabilities

| Field | Description |
|-------|-------------|
| **Tools available** | Read-only vs write/exec/network |
| **Long-running daemon** | Exists or must be added |

### C) Security and UX Constraints

| Field | Description |
|-------|-------------|
| **Offline required** | Must work without network |
| **Per-action approval** | Explicit approval per action |
| **Protected paths/commands** | Restricted operations |

**Default if missing:** Maximum safety and least privilege.

---

## Mandatory Output Format

This skill always outputs:

### A) Integration Option Selection

Option 1: Chrome Native Messaging
Option 2: Local Bridge Server (localhost WebSocket/HTTP)

With recommendation and rationale.

### B) Threat Model

Top risks plus mitigations.

### C) Permission Model

Scopes, allowlists, consent prompts, revocation.

### D) Message Contract

Typed schemas for request/response/events.

### E) Chrome Extension Blueprint (MV3)

Manifest permissions, service worker, popup, content scripts.

### F) CLI Bridge Blueprint

Process model, auth handshake, request routing, logging.

### G) UX Flows

Connect, run tool, watch progress, review result, error handling, revoke.

### H) Tests and Verification

Smoke, integration, security, and failure-mode tests.

### I) Definition of Done

Hard merge gates.

---

## Integration Options (2026 Best Practice)

### Option 1 (Recommended for Highest Security): Chrome Native Messaging

**What it is:** Extension talks to a local "native host" executable via stdin/stdout.

**Pros:**
- Best security posture (no network port)
- No CORS headaches
- Harder for websites to attack

**Cons:**
- Installation involves registering a native host manifest
- Platform-specific packaging steps

**Use when:** Running tools that can write files/execute commands.

### Option 2: Localhost Bridge Server (HTTP/WebSocket)

**What it is:** CLI runs a local server (127.0.0.1) and the extension connects.

**Pros:**
- Easier dev loop
- Simple streaming over WS/SSE

**Cons:**
- Must secure against CSRF, rogue local requests
- Port discovery plus firewall issues

**Use when:** Read-only tooling or you have a strong auth story (token plus origin restrictions).

---

## Threat Model (Non-Negotiable)

These are the default risks that must be addressed:

| Risk | Description |
|------|-------------|
| **Unauthorized tool execution** | A site tries to trigger commands |
| **Privilege escalation** | Extension gets broad host permissions |
| **Data exfiltration** | Repo secrets, files, env vars leak |
| **Command injection** | Unvalidated inputs reach shell |
| **Replay attacks** | Captured messages re-run destructive actions |
| **Confused deputy** | Extension executes something user did not intend |

**Mitigations must include:**

- Deny-by-default permission scopes
- Explicit user consent
- Strict schema validation
- Allowlists (commands/paths/domains)
- Per-request request_id, expiry, nonces
- Safe output redaction

---

## Permission Model (Scopes plus Consent)

### Scopes (Example)

| Scope | Description |
|-------|-------------|
| `repo.read` | Read file, list files, search |
| `repo.write` | Apply patch, create file |
| `run.tests` | Run test suite |
| `run.commands` | Allowlisted commands only |
| `git.read` / `git.write` | Git operations |
| `net.fetch` | Allowlisted domains |

### Consent Rules

- First connection requires explicit pairing
- High-risk scopes require per-action approval by default
- Allow user to "Always allow this tool for this repo" with clear revocation

### Revocation

- One-click revoke in extension settings and CLI settings
- Tokens invalidated immediately

---

## Message Contract (Typed JSON)

All messages follow a strict envelope:

### Request Envelope

```json
{
  "protocol": "floyd-bridge/1",
  "request_id": "string (uuid)",
  "origin": {
    "extension_id": "string",
    "tab_url": "string (optional)",
    "repo_hint": "string (optional)"
  },
  "action": "string (e.g., tool.run, repo.read)",
  "scope": "string (one of allowed scopes)",
  "params": "object (validated per action)",
  "timestamp_ms": "number",
  "nonce": "string"
}
```

### Response Envelope

```json
{
  "request_id": "string",
  "ok": "boolean",
  "code": "string (stable error/success code)",
  "message": "string (human readable)",
  "data": "object (optional)",
  "redactions": "string[] (what got scrubbed)"
}
```

### Event Envelope (Streaming)

```json
{
  "event": "tool.start | tool.progress | tool.log | tool.end | tool.error",
  "request_id": "string",
  "payload": "object (must be bounded; log chunk max size)"
}
```

---

## Chrome Extension Blueprint (MV3)

### Required Components

1. **manifest.json (MV3)** — minimum permissions only
2. **action (popup)** — quick actions and status
3. **background service worker** — owns connection lifecycle
4. **optional content_scripts** — for GitHub/GitLab integration

### Service Worker

- Owns connection lifecycle to the local bridge
- Queues requests if bridge is down
- Maintains "paired state" and active repo context

### Popup UI

- "Connect to Floyd" status
- Quick actions (Run tests, Lint, Open CLI, Apply patch)
- Last run status plus link to logs
- Permissions plus revoke

### Content Script (Optional)

- Adds "Run Floyd Tool" buttons on repo pages
- Extracts repo metadata (owner/name, branch) safely
- Sends intents to service worker (never calls bridge directly)

### MV3 Constraints

- Service worker can suspend; reconnection must be resilient
- Store minimal state in chrome.storage (no secrets in plaintext)

---

## CLI Bridge Blueprint (Local Side)

### Process Model

Pick one:

1. **On-demand host process** — Native messaging spawns per request
2. **Long-running daemon** — Better for streaming and monitoring

### Authentication (Mandatory)

- Pairing step yields a short secret token plus repo binding
- Requests must include nonce plus timestamp; token must be validated
- Token stored in OS keychain if available; otherwise encrypted local store

### Routing

1. Map action to tool handler
2. Validate schema, enforce scope, enforce allowlists
3. Execute
4. Emit progress/log events

### Safety Rules

- Never run arbitrary shell
- Allowlisted commands only
- Protected paths always denied (e.g., .env*, ssh keys)
- Redact secrets in output (common patterns plus configurable)

---

## UX Flows (Required)

### Flow 1: First-Time Connect (Pairing)

1. User clicks "Connect to Floyd"
2. Extension shows pairing code
3. User confirms in CLI ("Allow this extension?")
4. Bridge returns token bound to extension_id plus repo root

### Flow 2: Run a Tool

1. User clicks "Run Tests"
2. Extension sends request with scope run.tests
3. CLI bridge prompts if needed (or extension shows approval UI)
4. Tool runs; extension streams progress/logs
5. Completion summary shown plus "Open in Floyd Monitor" action

### Flow 3: Permission Denied

- Clear reason plus what scope is missing
- One-click "Request access" (does not auto-grant)

### Flow 4: Revoke Access

- Extension setting: revoke token
- CLI also supports revoke-all

---

## Testing and Verification

### Extension Tests

- Contract tests for request envelope creation
- UI tests for connect/run/revoke flows
- Service worker reconnection plus queue behavior

### Bridge Tests

- Schema validation rejects unknown fields
- Protected path denial
- Allowlist command enforcement
- Replay prevention (nonce reuse rejected)
- Output redaction test

### End-to-End

- Simulate tool.run with streamed logs
- Disconnect mid-run to recover gracefully
- Long session stability (no memory growth)

---

## Definition of Done (Hard Gates)

This feature is "merge-ready" only if:

- [ ] Chosen integration option is implemented end-to-end
- [ ] Deny-by-default permissions are enforced
- [ ] Requests are schema-validated and authenticated
- [ ] High-risk actions require explicit consent
- [ ] Logs/events are bounded and redacted
- [ ] Revoke works instantly
- [ ] End-to-end tests include failure-mode simulations

---

## Skill Invocation Template

```
SKILL: CHROME EXTENSION BRIDGE BUILDER

Goal:
Pages/Contexts (GitHub, localhost, any site):
Tool Actions Needed:
Risk Level (read/write/exec/network):
Preferred Integration Option (Native Messaging / Local Bridge / Unsure):
Constraints (offline, no daemon, etc.):

Return:
A) Integration Option Selection
B) Threat Model
C) Permission Model
D) Message Contract
E) Chrome Extension Blueprint (MV3)
F) CLI Bridge Blueprint
G) UX Flows
H) Tests and Verification
I) Definition of Done
```

---

## Related Skills

- **cli-x-2026** — Use this skill when designing the browser extension UI
- **ink-performance-auditor** — Use this skill when the bridge causes high event throughput
- **mcp-builder** — Use this skill when exposing CLI tools via MCP through the browser bridge
