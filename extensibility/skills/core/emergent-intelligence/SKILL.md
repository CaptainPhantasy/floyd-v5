---
name: emergent-intelligence
description: Detect and nurture emergent intelligent behaviors in agent systems
category: core
version: "2.0.0"
---

# Emergent Intelligence

> Detect and nurture emergent intelligent behaviors in agent systems

## When to Use
- WHEN mode=BUILD and the task requires emergent intelligence
- WHEN action=`detect`: perform detect operation
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`enhance`: perform enhance operation

## Actions
`'detect' | 'analyze' | 'enhance'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| system_state | object | yes | System state |
| agents | object | yes | Agents |
| id | string | yes | Id |
| state | object | yes | State |
| interactions | object | yes | Interactions |
| with | string | yes | With |
| interaction_type | string | yes | Interaction type |
| outcome | object | yes | Outcome |

## Invocation
```
floyd skill:emergent-intelligence --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
