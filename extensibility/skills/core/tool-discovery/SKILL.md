---
name: tool-discovery
description: Discover and recommend relevant tools based on task context and agent capabilities
category: core
version: "2.0.0"
---

# Tool Discovery

> Discover and recommend relevant tools based on task context and agent capabilities

## When to Use
- WHEN mode=BUILD and the task requires tool discovery
- WHEN action=`find`: locate or enumerate matching items
- WHEN action=`list`: locate or enumerate matching items
- WHEN action=`recommend`: locate or enumerate matching items

## Actions
`'find' | 'list' | 'recommend'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| query | string | no | Query |
| category | string | no | Category |
| task_description | string | no | Task description |
| server_type | enum | no | Server type |

## Invocation
```
floyd skill:tool-discovery --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
