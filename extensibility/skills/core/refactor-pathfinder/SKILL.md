---
name: refactor-pathfinder
description: Calculates the minimum-change path from a broken code state to a target architecture using edit distance analysis
category: core
version: "2.0.0"
---

# Refactor Pathfinder

> Calculates the minimum-change path from a broken code state to a target architecture using edit distance analysis.

## When to Use
- WHEN `mode=BUILD` and a large-scale refactor has an obvious goal state but many possible paths to get there
- WHEN `mode=DEBUG` and a fix requires restructuring code, and you need the safest sequence of changes
- WHEN planning a multi-file refactor and want to minimize the number of intermediate broken states

## Actions
`'plan' | 'compare' | 'checkpoint'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| start | string | yes | Description of the current (broken/suboptimal) state — can be a file path, function signature, or architectural description |
| goal | string | yes | Description of the desired target state |
| constraints | string[] | no | Constraints: no-breaking-changes, preserve-tests, incremental-compiles |
| project_path | string | no | Root directory for the codebase (default: working directory) |

## Execution Pipeline

### Step 1: Inventory Current State
Use `view` and `list_symbols` to map the current structure of `start`. Use `grep` to find all references to the symbols being refactored. Use `mcp_floyd-devtools_dependency_analyzer` to understand coupling.

### Step 2: Define Goal State
Parse the `goal` description. If it references specific file paths or function signatures, use `view` to verify they exist (or note they must be created). Identify the set of atomic operations needed: rename, move, extract, delete, create.

### Step 3: Compute Minimum Path
Order the atomic operations into a sequence where each intermediate state compiles (if `incremental-compiles` constraint is set). Use `mcp_floyd-safe-ops_impact_simulate` to validate each step. Record the plan to the scratchpad.

### Step 4: Checkpoint
Before executing, use `mcp_floyd-git_git_commit` to create a checkpoint. Then execute steps incrementally, verifying at each boundary.

## Output Shape
```json
{
  "start_state": "string — description of the initial state",
  "goal_state": "string — description of the target state",
  "steps": [
    {
      "operation": "string — rename/move/extract/delete/create",
      "target": "string — file or symbol affected",
      "verification": "string — how to verify this step"
    }
  ],
  "total_steps": 5,
  "estimated_risk": "LOW | MEDIUM | HIGH",
  "checkpoint_commit": "string — commit SHA before execution"
}
```

## Failure Modes
- IF no valid incremental path exists: report that the refactor must be done atomically and recommend creating a feature branch
- IF a midpoint step fails to compile: halt, revert to checkpoint, and re-order the remaining steps

## Examples
```json
{
  "action": "plan",
  "start": "internal/engine/ has planner.go and executor.go mixed together",
  "goal": "internal/engine/planner/ and internal/engine/executor/ as separate packages",
  "constraints": ["incremental-compiles", "preserve-tests"]
}
```
