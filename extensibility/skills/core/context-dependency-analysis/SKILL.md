---
name: context-dependency-analysis
description: Analyze dependencies between context elements to identify critical information chains
category: core
version: "2.0.0"
---

# Context Dependency Analysis

> Analyze dependencies between context elements to identify critical information chains

## When to Use
- WHEN mode=BUILD and the task requires context dependency analysis
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`map`: perform map operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'analyze' | 'map' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| contexts | object | yes | Contexts |
| id | string | yes | Id |
| content | any | yes | Content |
| dependencies | object | no | Dependencies |
| context_id | string | yes | Context id |
| dependency_type | enum | yes | Dependency type |
| strength | number | yes | Strength |

## Invocation
```
floyd skill:context-dependency-analysis --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
