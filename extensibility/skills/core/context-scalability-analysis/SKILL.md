---
name: context-scalability-analysis
description: Analyze context handling scaling characteristics with increasing complexity
category: core
version: "2.0.0"
---

# Context Scalability Analysis

> Analyze context handling scaling characteristics with increasing complexity

## When to Use
- WHEN mode=BUILD and the task requires context scalability analysis
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`predict`: perform predict operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'analyze' | 'predict' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| current_context | object | yes | Current context |
| data_volume | number | yes | Data volume |
| complexity | number | yes | Complexity |
| processing_time | number | yes | Processing time |
| resource_usage | Record<string, number> | yes | Resource usage |

## Invocation
```
floyd skill:context-scalability-analysis --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
