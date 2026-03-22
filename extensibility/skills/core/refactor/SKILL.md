---
name: refactor
description: Safe, incremental code refactoring with structural verification and rollback capability
category: core
version: "2.0.0"
---

# Refactor

> Safe, incremental code refactoring with structural verification and rollback capability.

## When to Use
- WHEN `mode=BUILD` and code needs restructuring without behavior change
- WHEN `mode=DEBUG` and a fix requires extracting, renaming, or reorganizing code
- WHEN code quality metrics (complexity, duplication, coupling) exceed acceptable thresholds

## Actions
`'analyze' | 'plan' | 'execute' | 'verify'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| target | string | yes | File path, function name, or module to refactor |
| strategy | string | yes | Refactoring type: extract-function, rename, inline, split-module, move-file, simplify-conditional |
| preserve_behavior | boolean | no | Whether to run tests after each step (default: true) |
| tests | string[] | no | Specific test commands to run for verification |

## Execution Pipeline

### Step 1: Analyze
Use `view` to read the target code. Use `list_symbols` to map all function/class signatures in the file. Use `grep` to find all callers of the target function across the codebase. Assess the blast radius — how many files reference this symbol.

### Step 2: Plan
Write the refactoring plan to the scratchpad via `manage_scratchpad`. The plan must list:
1. Each structural change in order
2. Files affected by each change
3. The test command that verifies each step preserves behavior

Use `mcp_floyd-safe-ops_impact_simulate` if available to simulate the change and detect import breakage before applying.

### Step 3: Execute
Apply changes incrementally using `edit` or `multiedit`. For each step in the plan:
- Make the structural change
- If `preserve_behavior=true`, run the test command via `bash`
- If tests fail, use `edit` to revert just this step and reassess

### Step 4: Verify
Run the full test suite via `bash`. Use `grep` to confirm no dangling references to old symbol names remain. Use `bash` to run the linter. Report the complete list of files modified.

## Output Shape
```json
{
  "strategy": "string — the refactoring type applied",
  "files_modified": ["string — paths of files changed"],
  "symbols_renamed": [{"old": "string", "new": "string"}],
  "test_result": "PASS | FAIL",
  "steps_executed": 3
}
```

## Failure Modes
- IF a test fails mid-refactor: revert the last step with `edit`, re-analyze the dependency, and adjust the plan
- IF a linter reports new warnings after rename: update all call sites using `grep` + `multiedit` in batch
- IF the blast radius exceeds 10 files: consider splitting into multiple smaller refactoring passes

## Examples
```json
{
  "action": "plan",
  "target": "internal/parser/lexer.go:parseIdentifier",
  "strategy": "extract-function",
  "preserve_behavior": true,
  "tests": ["go test ./internal/parser/..."]
}
```
