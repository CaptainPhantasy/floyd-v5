---
name: trace-replay-debugger
description: Record and replay execution traces for debugging with anomaly detection
category: core
version: "2.0.0"
---

# Trace Replay Debugger

> Record and replay execution traces for debugging with anomaly detection

## When to Use
- WHEN mode=DEBUG and the task requires trace replay debugger
- WHEN action=`replay`: re-execute trace data
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`compare`: diff two inputs and report differences

## Actions
`'replay' | 'analyze' | 'compare'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| trace_data | object | yes | Trace data |
| function_calls | object | yes | Function calls |
| function_name | string | yes | Function name |
| parameters | object | yes | Parameters |
| return_value | any | yes | Return value |
| timestamp | number | yes | Timestamp |
| duration | number | yes | Duration |
| call_stack | string[] | yes | Call stack |

## Invocation
```
floyd skill:trace-replay-debugger --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
