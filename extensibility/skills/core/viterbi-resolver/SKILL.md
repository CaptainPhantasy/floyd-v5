---
name: viterbi-resolver
description: Resolves ambiguous code by computing the most probable correct interpretation using structural context and type constraints
category: core
version: "2.0.0"
---

# Viterbi Resolver

> Resolves ambiguous code by computing the most probable correct interpretation using structural context and type constraints.

## When to Use
- WHEN `mode=DEBUG` and an LLM-generated code patch contains ambiguous type usage or unclear function calls
- WHEN `mode=BUILD` and a code completion or suggestion needs disambiguation between overloaded symbols
- WHEN a merge conflict has multiple valid resolutions and the best one must be selected

## Actions
`'resolve' | 'rank' | 'apply'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| ambiguous_code | string | yes | The code snippet containing ambiguity |
| candidates | string[] | no | Possible resolutions to evaluate (if known) |
| context_file | string | no | File path providing surrounding type/function context |
| language | string | no | Programming language for type-aware analysis (default: auto-detect) |

## Execution Pipeline

### Step 1: Extract Context
Use `view` to read `context_file` if provided. Use `grep` to find the types and function signatures referenced in `ambiguous_code`. Use `list_symbols` on relevant files to enumerate available overloads or matching names.

### Step 2: Rank Candidates
For each candidate (or each discovered overload), score it against:
1. **Type compatibility** — do the argument types match the signature?
2. **Surrounding usage** — how are the same symbols used in nearby code?
3. **Import presence** — is the required module imported?

If `candidates` is empty, generate candidates from the symbol table discovered in Step 1.

### Step 3: Apply
Return the highest-scoring resolution. If the scores are within 10% of each other, flag as ambiguous and return both options with explanations.

## Output Shape
```json
{
  "ambiguous_code": "string — the original input",
  "resolution": "string — the most probable correct code",
  "confidence": 0.92,
  "candidates_evaluated": [
    {
      "code": "string — candidate resolution",
      "score": 0.92,
      "reason": "string — why this scored highest"
    }
  ],
  "flags": ["string — warnings if confidence is low"]
}
```

## Failure Modes
- IF no candidates can be discovered: return the original code with a note that context is insufficient
- IF multiple candidates tie within 10%: return all tied candidates and recommend human review

## Examples
```json
{
  "action": "resolve",
  "ambiguous_code": "result := process(data)",
  "context_file": "internal/engine/pipeline.go",
  "language": "go"
}
```
