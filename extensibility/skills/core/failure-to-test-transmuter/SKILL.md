---
name: failure-to-test-transmuter
description: Automatically generate regression test cases from production failures and crash reports
category: core
version: "2.0.0"
---

# Failure To Test Transmuter

> Automatically generate regression test cases from production failures and crash reports

## When to Use
- WHEN mode=DEBUG and the task requires failure to test transmuter
- WHEN action=`transmute`: produce new artifacts from inputs
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`generate`: produce new artifacts from inputs

## Actions
`'transmute' | 'analyze' | 'generate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| error_report | object | yes | Error report |
| error_message | string | yes | Error message |
| stack_trace | string | yes | Stack trace |
| reproduction_steps | string[] | yes | Reproduction steps |
| environment | object | yes | Environment |

## Invocation
```
floyd skill:failure-to-test-transmuter --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
