---
name: long-term-visioning
description: Develop long-term visions with scenario planning and milestone decomposition
category: core
version: "2.0.0"
---

# Long Term Visioning

> Develop long-term visions with scenario planning and milestone decomposition

## When to Use
- WHEN mode=EXPLORE and the task requires long term visioning
- WHEN action=`envision`: perform envision operation
- WHEN action=`evaluate`: perform evaluate operation
- WHEN action=`refine`: perform refine operation

## Actions
`'envision' | 'evaluate' | 'refine'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| vision_context | object | yes | Vision context |
| time_horizon | number | yes | Time horizon |
| domain_scope | string[] | yes | Domain scope |
| influencing_factors | object | yes | Influencing factors |
| factor | string | yes | Factor |
| impact_direction | enum | yes | Impact direction |
| probability | number | yes | Probability |

## Invocation
```
floyd skill:long-term-visioning --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
