---
name: context-orchestration
description: Orchestrate multiple context sources with prioritization and conflict resolution
category: core
version: "2.0.0"
---

# Context Orchestration

> Orchestrate multiple context sources with prioritization and conflict resolution

## When to Use
- WHEN mode=BUILD and the task requires context orchestration
- WHEN action=`orchestrate`: perform orchestrate operation
- WHEN action=`optimize`: reduce size or improve quality of input
- WHEN action=`balance`: perform balance operation

## Actions
`'orchestrate' | 'optimize' | 'balance'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| contexts | object | yes | Contexts |
| id | string | yes | Id |
| content | any | yes | Content |
| priority | number | yes | Priority |
| dependencies | string[] | no | Dependencies |
| constraints | object | no | Constraints |

## Invocation
```
floyd skill:context-orchestration --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
