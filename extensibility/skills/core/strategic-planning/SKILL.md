---
name: strategic-planning
description: Develop strategic plans with goal decomposition milestone tracking and risk assessment
category: core
version: "2.0.0"
---

# Strategic Planning

> Develop strategic plans with goal decomposition milestone tracking and risk assessment

## When to Use
- WHEN mode=EXPLORE and the task requires strategic planning
- WHEN action=`plan`: perform plan operation
- WHEN action=`evaluate`: perform evaluate operation
- WHEN action=`adjust`: perform adjust operation

## Actions
`'plan' | 'evaluate' | 'adjust'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| planning_context | object | yes | Planning context |
| objectives | object | yes | Objectives |
| objective | string | yes | Objective |
| priority | number | yes | Priority |
| timeline | string | yes | Timeline |
| success_criteria | string[] | yes | Success criteria |

## Invocation
```
floyd skill:strategic-planning --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
