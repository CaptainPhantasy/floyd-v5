---
name: swarm-intelligence
description: Coordinate swarm-based problem solving with emergent behavior and collective optimization
category: core
version: "2.0.0"
---

# Swarm Intelligence

> Coordinate swarm-based problem solving with emergent behavior and collective optimization

## When to Use
- WHEN mode=BUILD and the task requires swarm intelligence
- WHEN action=`coordinate`: perform coordinate operation
- WHEN action=`optimize`: optimize for efficiency
- WHEN action=`learn`: perform learn operation

## Actions
`'coordinate' | 'optimize' | 'learn'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| swarm | object | yes | Swarm |
| agent_id | string | yes | Agent id |
| capabilities | object | yes | Capabilities |
| current_state | object | yes | Current state |
| learning_rate | number | yes | Learning rate |

## Invocation
```
floyd skill:swarm-intelligence --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
