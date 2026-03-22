---
name: context-persistence
description: Manage persistent storage and retrieval of context across sessions and restarts
category: core
version: "2.0.0"
---

# Context Persistence

> Manage persistent storage and retrieval of context across sessions and restarts

## When to Use
- WHEN mode=BUILD and the task requires context persistence
- WHEN action=`persist`: perform persist operation
- WHEN action=`retrieve`: perform retrieve operation
- WHEN action=`manage`: perform manage operation

## Actions
`'persist' | 'retrieve' | 'manage'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| context_data | object | yes | Context data |
| content | any | yes | Content |
| metadata | object | yes | Metadata |
| created_at | string | yes | Created at |
| expires_at | string | no | Expires at |
| access_count | number | yes | Access count |
| importance | number | yes | Importance |

## Invocation
```
floyd skill:context-persistence --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
