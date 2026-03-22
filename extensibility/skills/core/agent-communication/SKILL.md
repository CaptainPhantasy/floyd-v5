---
name: agent-communication
description: Facilitate structured communication between agents with message routing and protocol handling
category: core
version: "2.0.0"
---

# Agent Communication

> Facilitate structured communication between agents with message routing and protocol handling

## When to Use
- WHEN mode=BUILD and the task requires agent communication
- WHEN action=`establish`: perform establish operation
- WHEN action=`mediate`: perform mediate operation
- WHEN action=`optimize`: optimize for efficiency

## Actions
`'establish' | 'mediate' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| agents | object | yes | Agents |
| id | string | yes | Id |
| role | string | yes | Role |
| communication_preferences | object | yes | Communication preferences |
| current_state | object | yes | Current state |

## Invocation
```
floyd skill:agent-communication --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
