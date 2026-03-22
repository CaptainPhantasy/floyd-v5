---
name: value-alignment
description: Ensure agent behavior aligns with specified values through constraint monitoring
category: core
version: "2.0.0"
---

# Value Alignment

> Ensure agent behavior aligns with specified values through constraint monitoring

## When to Use
- WHEN mode=BUILD and the task requires value alignment
- WHEN action=`assess`: perform assess operation
- WHEN action=`align`: perform align operation
- WHEN action=`monitor`: perform monitor operation

## Actions
`'assess' | 'align' | 'monitor'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| agent_values | object | yes | Agent values |
| agent_id | string | yes | Agent id |
| core_values | object | yes | Core values |
| value | string | yes | Value |
| importance | number | yes | Importance |
| manifestation | string[] | yes | Manifestation |

## Invocation
```
floyd skill:value-alignment --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
