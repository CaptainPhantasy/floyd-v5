---
name: bloom-sentinel
description: Global deduplication preventing agent swarm from repeating failed code paths via probabilistic bloom filter
category: core
version: "2.0.0"
---

# Bloom Sentinel

> Global deduplication preventing agent swarm from repeating failed code paths via probabilistic bloom filter.

## When to Use
- WHEN `mode=DEBUG` and an agent is about to retry a fix that has already failed with the same approach
- WHEN `mode=BUILD` and multiple agents are working in parallel and need to avoid duplicate work on the same files
- WHEN task orchestration needs to track which code paths have been attempted and their outcomes

## Actions
`'check' | 'register' | 'reset'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| path | string | yes | The code path, file, or approach fingerprint to check/register |
| outcome | string | no | Result of attempting this path: `success`, `failure`, `timeout` (required for register) |
| scope | string | no | Deduplication scope: `session`, `task`, `global` (default: session) |

## Execution Pipeline

### Step 1: Check
Before attempting a fix or approach, call `check` with the `path` fingerprint. The fingerprint should encode the file, function, and fix strategy (e.g., `internal/parser/lexer.go:parseIdentifier:extract-function`). Use `manage_scratchpad` to query the bloom filter state. If the path is marked as `failure`, skip it and try a different approach.

### Step 2: Register
After attempting a fix, call `register` with the `path` and `outcome`. Store the result in the scratchpad under a bloom-sentinel key. Use `mcp_floyd-supercache_cache_store` for persistent tracking across sessions.

### Step 3: Reset
When a task is completed or a new task begins, call `reset` to clear the bloom filter for the current `scope`. This prevents stale failure records from blocking valid approaches in a new context.

## Output Shape
```json
{
  "path": "string — the fingerprint queried",
  "exists": true,
  "outcome": "failure | success | timeout | null",
  "attempt_count": 2,
  "last_attempt": "2026-03-22T10:30:00Z",
  "recommended_action": "SKIP — this path has failed 2 times"
}
```

## Failure Modes
- IF the bloom filter returns a false positive (path was not actually tried): allow override with explicit `force=true` parameter
- IF the scratchpad is unavailable: fall back to in-memory tracking within the current session only

## Examples
```json
{
  "action": "check",
  "path": "internal/parser/lexer.go:parseIdentifier:extract-function",
  "scope": "task"
}
```
