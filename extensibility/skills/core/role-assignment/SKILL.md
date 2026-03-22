---
name: role-assignment
description: Dynamically assign roles to agents based on capabilities load and task requirements
category: core
version: "2.0.0"
---

# Role Assignment

> Dynamically assign roles to agents based on capabilities load and task requirements

## When to Use
- WHEN mode=BUILD and the task requires role assignment
- WHEN action=`assign`: perform assign operation
- WHEN action=`reassign`: perform reassign operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'assign' | 'reassign' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| agents | object | yes | Agents |
| id | string | yes | Id |
| capabilities | Record<string, number> | yes | Capabilities |
| preferences | object | yes | Preferences |
| current_role | string | no | Current role |
| performance_history | Record<string, number> | yes | Performance history |

## Invocation
```
floyd skill:role-assignment --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
