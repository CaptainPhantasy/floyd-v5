---
name: collective-learning
description: Coordinate learning across agents with knowledge aggregation and distribution
category: core
version: "2.0.0"
---

# Collective Learning

> Coordinate learning across agents with knowledge aggregation and distribution

## When to Use
- WHEN mode=BUILD and the task requires collective learning
- WHEN action=`facilitate`: perform facilitate operation
- WHEN action=`aggregate`: perform aggregate operation
- WHEN action=`distribute`: perform distribute operation

## Actions
`'facilitate' | 'aggregate' | 'distribute'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| learning_experiences | object | yes | Learning experiences |
| agent_id | string | yes | Agent id |
| experience | object | yes | Experience |
| outcomes | object | yes | Outcomes |
| generalizability | number | yes | Generalizability |

## Invocation
```
floyd skill:collective-learning --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
