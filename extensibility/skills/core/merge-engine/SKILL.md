---
name: merge-engine
description: Resolve concurrent file edits from multiple agents using operational transform to prevent conflicts and data loss
category: core
version: "2.0.0"
---

# Merge Engine

> Resolve concurrent file edits from multiple agents using operational transform to prevent conflicts and data loss.

## When to Use
- WHEN `mode=BUILD` and two or more agents have edited the same file simultaneously
- WHEN `mode=DEBUG` and a git merge conflict needs intelligent resolution beyond textual diff3
- WHEN applying multiple patches to the same file in sequence and earlier patches shift line numbers

## Actions
`'merge' | 'transform' | 'conflict-report'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| base | string | yes | The original file content before any edits (file path or inline content) |
| edits | object[] | yes | Array of edit operations: `{agent: string, file: string, changes: string}` |
| strategy | string | no | Conflict resolution strategy: `ours`, `theirs`, `manual`, `semantic` (default: semantic) |

## Execution Pipeline

### Step 1: Collect Base State
Use `view` to read the `base` file. If `base` is inline content, use it directly. Record the base SHA or timestamp for traceability.

### Step 2: Apply Operational Transform
For each edit in `edits`, compute the transform against the base and against all previously applied edits. Use `mcp_floyd-safe-ops_safe_refactor` to apply changes atomically with rollback support. Track line offset adjustments as each edit shifts subsequent positions.

### Step 3: Resolve Conflicts
If two edits modify the same region:
- **semantic**: analyze the intent of each edit and merge the semantic changes
- **ours**: keep the first agent's version
- **theirs**: keep the second agent's version
- **manual**: flag for human review with both versions preserved in comments

### Step 4: Verify
Use `bash` to run the project's build command on the merged result. If it fails, use `mcp_floyd-safe-ops_safe_refactor` to rollback and report.

## Output Shape
```json
{
  "file": "string — merged file path",
  "edits_applied": 3,
  "conflicts_detected": 1,
  "conflicts_resolved": 1,
  "resolution_strategy": "string — strategy used per conflict",
  "build_check": "PASS | FAIL",
  "manual_review_required": false
}
```

## Failure Modes
- IF the merged result fails to compile: rollback to base and report all conflicting edit pairs for manual resolution
- IF an edit references a line that no longer exists (stale offset): re-read the file, locate the nearest matching context, and apply at the corrected position

## Examples
```json
{
  "action": "merge",
  "base": "internal/engine/planner.go",
  "edits": [
    {"agent": "agent-1", "changes": "added error handling to Evaluate"},
    {"agent": "agent-2", "changes": "renamed variable in Evaluate"}
  ],
  "strategy": "semantic"
}
```
