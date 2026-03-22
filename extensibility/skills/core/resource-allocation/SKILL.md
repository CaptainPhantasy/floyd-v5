---
name: resource-allocation
description: Optimize resource allocation with constraint satisfaction and priority scheduling
category: core
version: "2.0.0"
---

# Resource Allocation

> Optimize resource allocation with constraint satisfaction and priority scheduling

## When to Use
- WHEN mode=BUILD and the task requires resource allocation
- WHEN action=`allocate`: perform allocate operation
- WHEN action=`optimize`: reduce size or improve quality of input
- WHEN action=`rebalance`: perform rebalance operation

## Actions
`'allocate' | 'optimize' | 'rebalance'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| resources | object | yes | Resources |
| compute | number | yes | Compute |
| memory | number | yes | Memory |
| storage | number | yes | Storage |
| bandwidth | number | yes | Bandwidth |

## Invocation
```
floyd skill:resource-allocation --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
