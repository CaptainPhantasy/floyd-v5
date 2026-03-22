---
name: build-error-correlator
description: Correlate build errors across builds to identify patterns and suggest automated fixes
category: core
version: "2.0.0"
---

# Build Error Correlator

> Correlate build errors across builds to identify patterns and suggest automated fixes

## When to Use
- WHEN mode=DEBUG and the task requires build error correlator
- WHEN action=`correlate`: find patterns across inputs
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`suggest_fixes`: propose fixes

## Actions
`'correlate' | 'analyze' | 'suggest_fixes'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| build_logs | object | yes | Build logs |
| file | string | yes | File |
| errors | object | yes | Errors |
| message | string | yes | Message |
| line | number | yes | Line |
| column | number | yes | Column |
| code | string | yes | Code |

## Invocation
```
floyd skill:build-error-correlator --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
