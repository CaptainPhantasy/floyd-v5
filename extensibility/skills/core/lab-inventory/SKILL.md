---
name: lab-inventory
description: Track and manage Floyd Lab tools with versioning usage analytics and health monitoring
category: core
version: "2.0.0"
---

# Lab Inventory

> Track and manage Floyd Lab tools with versioning usage analytics and health monitoring

## When to Use
- WHEN mode=BUILD and the task requires lab inventory
- WHEN action=`list`: locate or enumerate matching items
- WHEN action=`search`: locate or enumerate matching items
- WHEN action=`get_status`: perform get_status operation

## Actions
`'list' | 'search' | 'get_status'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| server_type | enum | no | Server type |
| detailed | boolean | no | Detailed |

## Invocation
```
floyd skill:lab-inventory --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
