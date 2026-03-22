---
name: distributed-task-management
description: Manage tasks across distributed agents with load balancing and fault tolerance
category: core
version: "2.0.0"
---

# Distributed Task Management

> Manage tasks across distributed agents with load balancing and fault tolerance

## When to Use
- WHEN mode=BUILD and the task requires distributed task management
- WHEN action=`create`: produce new artifacts from inputs
- WHEN action=`assign`: perform assign operation
- WHEN action=`coordinate`: perform coordinate operation
- WHEN action=`monitor`: perform monitor operation

## Actions
`'create' | 'assign' | 'coordinate' | 'monitor'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| tasks | object | yes | Tasks |
| id | string | yes | Id |
| description | string | yes | Description |
| requirements | object | yes | Requirements |
| skills | string[] | yes | Skills |
| resources | Record<string, number> | yes | Resources |
| dependencies | string[] | yes | Dependencies |

## Invocation
```
floyd skill:distributed-task-management --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
