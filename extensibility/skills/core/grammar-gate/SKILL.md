---
name: grammar-gate
description: Validate generated code against project-specific structural rules, naming conventions, and idiomatic patterns
category: core
version: "2.0.0"
---

# Grammar Gate

> Validate generated code against project-specific structural rules, naming conventions, and idiomatic patterns.

## When to Use
- WHEN `mode=BUILD` after generating or modifying code to verify it follows project conventions
- WHEN `mode=DEBUG` when a generated patch fails linting or formatting checks
- WHEN onboarding to a new codebase and needing to understand its structural constraints

## Actions
`'validate' | 'lint' | 'fix'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| code | string | yes | The code snippet or file path to validate |
| rules | string[] | no | Specific rules to check (default: all project rules) |
| auto_fix | boolean | no | Whether to automatically apply fixes for violations (default: false) |

## Execution Pipeline

### Step 1: Detect Project Grammar
Use `glob` to find project configuration files (`.eslintrc`, `.golangci.yml`, `.prettierrc`, `pyproject.toml`). Use `view` to read the relevant config and extract the active rules. If no config exists, fall back to language-default idiomatic patterns.

### Step 2: Validate
Run the project's configured linter via `mcp_floyd-runner_lint` or `bash` with the appropriate lint command. If `rules` is specified, scope the check to only those rules. Capture all violations with file, line, and rule name.

### Step 3: Fix (if auto_fix=true)
For each violation, apply the fix using `edit` or `multiedit`. Re-run the linter to confirm zero violations. If a fix introduces a new violation, revert and report the conflict.

## Output Shape
```json
{
  "valid": true,
  "violations": [
    {
      "file": "string — file path",
      "line": 42,
      "rule": "string — rule name",
      "severity": "ERROR | WARNING | INFO",
      "message": "string — human-readable description",
      "auto_fixed": true
    }
  ],
  "total_violations": 3,
  "files_checked": 5
}
```

## Failure Modes
- IF no linter is configured: fall back to built-in checks (naming conventions, indentation, missing error handling) using `grep` pattern matching
- IF auto_fix creates a loop (fixing one violation creates another): halt after 3 iterations and report the cycle

## Examples
```json
{
  "action": "validate",
  "code": "internal/engine/planner.go",
  "auto_fix": true
}
```
