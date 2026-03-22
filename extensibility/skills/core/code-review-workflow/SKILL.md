---
name: code-review-workflow
description: Systematic code review combining automated analysis with structured checklists for quality, security, and maintainability
category: core
version: "2.0.0"
---

# Code Review Workflow

> Systematic code review combining automated analysis with structured checklists for quality, security, and maintainability.

## When to Use
- WHEN `mode=BUILD` before merging a PR to validate code quality
- WHEN `mode=DEBUG` to audit code for potential bugs after a fix
- WHEN `mode=EXPLORE` to assess the health of a codebase or module

## Actions
`'review' | 'checklist' | 'security' | 'summary'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| target | string | yes | PR number, commit SHA, branch name, or file path to review |
| focus | string[] | no | Review focus areas: correctness, performance, security, style, testing (default: all) |
| project_path | string | no | Root directory (default: working directory) |

## Execution Pipeline

### Step 1: Collect Changes
If `target` is a PR number, use `bash` with `gh pr diff <number>` to get the diff. If a commit SHA, use `bash` with `git show <sha>`. If a branch, use `bash` with `git diff main..<branch>`. Use `mcp_floyd-git_git_log` to understand recent commit history for context.

### Step 2: Automated Analysis
Run automated checks in parallel:
- **Lint**: Use `mcp_floyd-runner_lint` or `bash` with the project's linter
- **Type check**: Use `mcp_floyd-devtools_typescript_semantic_analyzer` with action `find_type_mismatches` for TypeScript projects, or `bash go vet` for Go
- **Security**: Scan for hardcoded secrets using `grep` with patterns like `password`, `api_key`, `token`
- **Test coverage**: Use `bash` to run the test suite and capture pass/fail

### Step 3: Structured Review
For each focus area, walk through the changed files using `view`. Apply the review checklist:
- **Correctness**: Does the logic match the intent? Are edge cases handled?
- **Performance**: Are there unnecessary allocations, N+1 queries, or synchronous I/O in hot paths?
- **Security**: Are inputs validated? Are there injection risks? Are secrets exposed?
- **Style**: Does the code match the project's formatting conventions?
- **Testing**: Are new functions covered? Do tests assert meaningful behavior?

### Step 4: Summary
Aggregate findings into a structured report. Classify each finding as BLOCKER, WARNING, or SUGGESTION.

## Output Shape
```json
{
  "target": "string — PR/commit/branch reviewed",
  "files_reviewed": 8,
  "findings": [
    {
      "file": "string — file path",
      "line": 42,
      "severity": "BLOCKER | WARNING | SUGGESTION",
      "category": "correctness | performance | security | style | testing",
      "description": "string — what the issue is",
      "recommendation": "string — how to fix it"
    }
  ],
  "automated_checks": {
    "lint": "PASS | FAIL",
    "type_check": "PASS | FAIL",
    "tests": "42/42 PASS",
    "secrets_found": 0
  },
  "verdict": "APPROVE | REQUEST_CHANGES | NEEDS_DISCUSSION"
}
```

## Failure Modes
- IF the target cannot be resolved (invalid PR, missing commit): use `bash` with `gh pr list` or `git log` to find the correct identifier
- IF automated checks timeout: report partial results and note which checks could not complete

## Examples
```json
{
  "action": "review",
  "target": "123",
  "focus": ["correctness", "security", "performance"]
}
```
