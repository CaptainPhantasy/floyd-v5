---
name: semantic-diff-validation
description: Validate that code changes preserve semantic behavior while allowing structural improvements, with quantified risk scoring
category: core
version: "2.0.0"
---

# Semantic Diff Validation

> Validate that code changes preserve semantic behavior while allowing structural improvements, with quantified risk scoring.

## When to Use
- WHEN `mode=BUILD` before merging a refactor to confirm no behavioral regression
- WHEN `mode=DEBUG` to determine if a recent change introduced a subtle logic bug
- WHEN reviewing a PR that restructures code without changing its tests

## Actions
`'validate' | 'diff' | 'score'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| diff | string | no | Raw unified diff content (if empty, reads from git) |
| base_commit | string | no | Base commit SHA for comparison (default: HEAD~1) |
| scope | string[] | no | Restrict validation to specific files (default: all changed files) |

## Execution Pipeline

### Step 1: Collect Diff
If `diff` is not provided, use `mcp_floyd-git_git_diff` to get the working tree diff, or `bash` with `git diff base_commit..HEAD` for commit ranges. Filter to `scope` if provided.

### Step 2: Analyze Semantic Impact
For each changed hunk, use `view` to read the surrounding function context (20 lines before/after). Use `list_symbols` to identify if function signatures changed. Check for:
- **Breaking changes**: modified public API signatures, removed exports, changed return types
- **Behavioral changes**: altered control flow, modified error handling, changed loop conditions
- **Structural changes**: renames, moves, reformats (low risk)

### Step 3: Score Risk
Assign a risk score (0–100) per file based on the categories above. Aggregate into an overall change risk. Flag any file scoring above 60 for manual review.

## Output Shape
```json
{
  "total_files_changed": 8,
  "risk_score": 35,
  "risk_level": "LOW | MEDIUM | HIGH | CRITICAL",
  "findings": [
    {
      "file": "string — file path",
      "risk_score": 72,
      "category": "BREAKING | BEHAVIORAL | STRUCTURAL",
      "description": "string — what changed and why it matters",
      "line_range": "120-145"
    }
  ],
  "recommendation": "string — merge/hold/request-review"
}
```

## Failure Modes
- IF the diff is empty: report that no changes were detected
- IF a file cannot be parsed for semantic analysis: flag it as UNKNOWN risk and recommend manual review

## Examples
```json
{
  "action": "validate",
  "base_commit": "abc1234",
  "scope": ["internal/engine/planner.go", "internal/engine/executor.go"]
}
```
