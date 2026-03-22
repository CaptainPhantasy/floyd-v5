---
name: attention-allocation
description: Optimize attention allocation across multiple tasks contexts and priorities
category: core
version: "2.0.0"
---

# Attention Allocation

> Optimize attention allocation across multiple tasks contexts and priorities

## When to Use
- WHEN mode=BUILD and the task requires attention allocation
- WHEN action=`allocate`: perform allocate operation
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`optimize`: optimize for efficiency

## Actions
`'allocate' | 'analyze' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| information_space | object | yes | Information space |
| items | object | yes | Items |
| id | string | yes | Id |
| content | string | yes | Content |
| importance | number | yes | Importance |
| urgency | number | yes | Urgency |
| complexity | number | yes | Complexity |

## Invocation
```
floyd skill:attention-allocation --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
