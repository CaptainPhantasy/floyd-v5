---
name: api-contract-validation
description: Validate API contracts for compatibility, detect breaking changes, and generate migration plans with client impact analysis
category: core
version: "2.0.0"
---

# API Contract Validation

> Validate API contracts for compatibility, detect breaking changes, and generate migration plans with client impact analysis.

## When to Use
- WHEN `mode=BUILD` before releasing a new API version to ensure backward compatibility
- WHEN `mode=DEBUG` when an API client breaks after a server update
- WHEN `mode=EXPLORE` to audit an API surface for design inconsistencies

## Actions
`'validate' | 'diff' | 'migrate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| current_contract | string | yes | Current API specification (file path, URL, or inline JSON/OpenAPI) |
| proposed_contract | string | no | Proposed new specification (if diffing two versions) |
| api_type | string | no | API type: openai, anthropic, google, custom (default: auto-detect) |
| strict | boolean | no | Whether to fail on warnings in addition to errors (default: false) |

## Execution Pipeline

### Step 1: Parse Contracts
Use `view` to read the contract files. Parse the API specification to extract endpoints, request/response schemas, parameters, and authentication requirements.

### Step 2: Validate Schema
Use `mcp_floyd-devtools_api_format_verifier` with action `validate_schema` to check the contract structure. Verify all endpoints have complete request/response definitions, all parameters have types and descriptions, and all error responses are documented.

### Step 3: Diff Versions (if proposed_contract provided)
Compare `current_contract` against `proposed_contract`. Classify changes:
- **Breaking**: removed endpoints, changed parameter types, removed response fields
- **Additive**: new endpoints, new optional parameters, new response fields (safe)
- **Modified**: changed descriptions, reordered parameters (usually safe)

### Step 4: Generate Migration Plan
For each breaking change, generate a migration step: what clients must change, example before/after, and a deprecation timeline recommendation.

## Output Shape
```json
{
  "validation": {
    "valid": true,
    "errors": ["string — schema violations"],
    "warnings": ["string — non-critical issues"]
  },
  "diff": {
    "breaking_changes": [
      {
        "endpoint": "string — affected endpoint",
        "change": "string — what changed",
        "impact": "string — which clients are affected",
        "migration_step": "string — how to adapt"
      }
    ],
    "additive_changes": 5,
    "modified_changes": 3
  },
  "migration_plan": {
    "steps": ["string — ordered migration actions"],
    "deprecation_timeline": "string — recommended timeline"
  }
}
```

## Failure Modes
- IF the contract cannot be parsed: report the parse error with line number and suggest fixing the schema syntax
- IF `api_type` auto-detection fails: prompt for explicit `api_type` specification

## Examples
```json
{
  "action": "diff",
  "current_contract": "specs/api-v1.yaml",
  "proposed_contract": "specs/api-v2.yaml",
  "api_type": "custom"
}
```
