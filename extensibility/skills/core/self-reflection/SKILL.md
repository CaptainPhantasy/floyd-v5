---
name: self-reflection
description: Enable agents to reflect on performance identify weaknesses and adapt strategies
category: core
version: "2.0.0"
---

# Self Reflection

> Enable agents to reflect on performance identify weaknesses and adapt strategies

## When to Use
- WHEN mode=BUILD and the task requires self reflection
- WHEN action=`reflect`: perform reflect operation
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`improve`: reduce size or improve quality of input

## Actions
`'reflect' | 'analyze' | 'improve'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| subject | object | yes | Subject |
| entity | string | yes | Entity |
| context | object | yes | Context |
| recent_actions | object | yes | Recent actions |
| action | string | yes | Action |
| outcome | object | yes | Outcome |
| timestamp | string | yes | Timestamp |

## Invocation
```
floyd skill:self-reflection --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
