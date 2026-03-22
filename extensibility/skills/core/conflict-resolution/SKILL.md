---
name: conflict-resolution
description: Detect and resolve conflicts between agents, resources, and concurrent operations
category: core
version: "2.0.0"
---

# Conflict Resolution

> Detect and resolve conflicts between agents, resources, and concurrent operations

## When to Use
- WHEN mode=BUILD and the task requires conflict resolution
- WHEN action=`identify`: perform identify operation
- WHEN action=`mediate`: perform mediate operation
- WHEN action=`resolve`: perform resolve operation

## Actions
`'identify' | 'mediate' | 'resolve'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| conflicts | object | yes | Conflicts |
| conflict_id | string | yes | Conflict id |
| parties | string[] | yes | Parties |
| nature | enum | yes | Nature |
| intensity | number | yes | Intensity |
| description | string | yes | Description |

## Invocation
```
floyd skill:conflict-resolution --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
