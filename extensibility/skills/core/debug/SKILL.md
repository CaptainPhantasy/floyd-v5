---
name: debug
description: Hypothesis-driven debugging workflow with structured fault isolation, evidence collection, and post-fix verification
category: core
version: "2.0.0"
---

# Debug

> Hypothesis-driven debugging workflow with structured fault isolation, evidence collection, and post-fix verification.

## When to Use
- WHEN `mode=DEBUG` and a build error, test failure, or runtime panic occurs
- WHEN `mode=BUILD` and a code change produces unexpected behavior or a regression
- WHEN the agent has attempted a fix twice and needs to escalate analysis before a third attempt

## Actions
`'hypothesize' | 'isolate' | 'verify' | 'postmortem'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| error | string | yes | The error message, panic trace, or failing test output |
| files | string[] | yes | File paths suspected of involvement in the fault |
| context | string | no | Additional context: recent changes, environment info, or reproducer steps |

## Execution Pipeline

### Step 1: Hypothesize
Formulate at most 3 hypotheses ranked by likelihood. Use `grep` to search for error strings across the codebase. Use `view` to inspect the suspected `files` around error line numbers. Write hypotheses to the scratchpad via `manage_scratchpad`.

- **Hypothesis template:** "The error occurs because `[mechanism]` in `[file:line]`. Evidence: `[grep/view result]`."

### Step 2: Isolate
For the top-ranked hypothesis, use `view` to read the full function or module. Trace the data flow backward from the error site. Use `list_symbols` to understand the function signatures involved. Identify the minimal change that would validate or refute the hypothesis.

### Step 3: Verify
Apply the fix using `edit` or `multiedit`. If this is a Go project, the harness will auto-run `go build` — check the `<build_check>` result. For other projects, use `bash` to run the relevant build/test command. If the fix succeeds, proceed to Step 4. If it fails, return to Step 1 with the new error information.

**Two-failure reset rule:** If two consecutive fix attempts fail, STOP. Re-read the error from scratch using `view` on the failing file. Write a root-cause analysis to the scratchpad before attempting a third fix.

### Step 4: Postmortem
After a successful fix, verify no regressions by running the full test suite via `bash`. Record the root cause, fix, and evidence chain to the scratchpad. Clear intermediate hypothesis data from context.

## Output Shape
```json
{
  "root_cause": "string — the confirmed mechanism that produced the error",
  "fix_applied": "string — description of the change made",
  "files_modified": ["string — paths of files changed"],
  "evidence_chain": ["string — ordered list of observations leading to the fix"],
  "regression_check": "PASS | FAIL — result of post-fix test run"
}
```

## Failure Modes
- IF hypothesis isolation yields no relevant code paths: expand search with `grep` using broader patterns or alternative error substrings
- IF two consecutive fixes fail: invoke the two-failure reset rule — re-analyze root cause from scratch before a third attempt
- IF the fix introduces a new error: treat the new error as input and restart from Step 1

## Examples
```json
{
  "action": "hypothesize",
  "error": "panic: runtime error: index out of range [5] with length 5",
  "files": ["internal/parser/lexer.go"],
  "context": "occurs when parsing multiline string literals"
}
```
