---
name: adaptive-behavior
description: Enable agents to adapt based on environment, feedback, and learning signals
category: core
version: "2.0.0"
---

# Adaptive Behavior

> Enable agents to adapt based on environment, feedback, and learning signals

## When to Use
- WHEN mode=BUILD and the task requires adaptive behavior
- WHEN action=`adapt`: perform adapt operation
- WHEN action=`learn`: perform learn operation
- WHEN action=`predict`: perform predict operation

## Actions
`'adapt' | 'learn' | 'predict'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| agents | object | yes | Agents |
| id | string | yes | Id |
| behavior_patterns | object | yes | Behavior patterns |
| performance_history | Record<string, number> | yes | Performance history |
| adaptation_capabilities | object | yes | Adaptation capabilities |

## Invocation
```
floyd skill:adaptive-behavior --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
