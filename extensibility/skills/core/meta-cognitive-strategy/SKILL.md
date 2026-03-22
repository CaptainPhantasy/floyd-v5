---
name: meta-cognitive-strategy
description: Develop meta-cognitive strategies for improved reasoning and self-awareness
category: core
version: "2.0.0"
---

# Meta Cognitive Strategy

> Develop meta-cognitive strategies for improved reasoning and self-awareness

## When to Use
- WHEN mode=EXPLORE and the task requires meta cognitive strategy
- WHEN action=`develop`: perform develop operation
- WHEN action=`apply`: perform apply operation
- WHEN action=`evaluate`: perform evaluate operation

## Actions
`'develop' | 'apply' | 'evaluate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| task_context | object | yes | Task context |
| domain | string | yes | Domain |
| complexity | number | yes | Complexity |
| objectives | string[] | yes | Objectives |
| constraints | object | no | Constraints |

## Invocation
```
floyd skill:meta-cognitive-strategy --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
