---
name: synchronization
description: Synchronize state and actions across distributed agents with consistency guarantees
category: core
version: "2.0.0"
---

# Synchronization

> Synchronize state and actions across distributed agents with consistency guarantees

## When to Use
- WHEN mode=BUILD and the task requires synchronization
- WHEN action=`synchronize`: perform synchronize operation
- WHEN action=`coordinate`: perform coordinate operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'synchronize' | 'coordinate' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| agents | object | yes | Agents |
| id | string | yes | Id |
| current_state | object | yes | Current state |
| synchronization_needs | object | yes | Synchronization needs |
| communication_channels | string[] | yes | Communication channels |

## Invocation
```
floyd skill:synchronization --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
