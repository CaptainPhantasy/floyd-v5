---
name: cognitive-load-analysis
description: Analyze cognitive load of tasks and interfaces for optimization
category: core
version: "2.0.0"
---

# Cognitive Load Analysis

> Analyze cognitive load of tasks and interfaces for optimization

## When to Use
- WHEN mode=EXPLORE and the task requires cognitive load analysis
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`optimize`: optimize for efficiency
- WHEN action=`compare`: diff two inputs and report differences

## Actions
`'analyze' | 'optimize' | 'compare'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| task_definition | object | yes | Task definition |
| description | string | yes | Description |
| steps | object | yes | Steps |
| description | string | yes | Description |
| type | enum | yes | Type |
| complexity | number | yes | Complexity |

## Invocation
```
floyd skill:cognitive-load-analysis --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
