---
name: goal-alignment
description: Ensure alignment between agent goals and system objectives with drift detection
category: core
version: "2.0.0"
---

# Goal Alignment

> Ensure alignment between agent goals and system objectives with drift detection

## When to Use
- WHEN mode=BUILD and the task requires goal alignment
- WHEN action=`align`: perform align operation
- WHEN action=`assess`: perform assess operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'align' | 'assess' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| entities | object | yes | Entities |
| id | string | yes | Id |
| goals | object | yes | Goals |
| goal | string | yes | Goal |
| priority | number | yes | Priority |
| timeline | string | yes | Timeline |
| success_criteria | string[] | yes | Success criteria |

## Invocation
```
floyd skill:goal-alignment --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
