---
name: self-organization
description: Enable agent systems to self-organize based on emergent goals and environmental feedback
category: core
version: "2.0.0"
---

# Self Organization

> Enable agent systems to self-organize based on emergent goals and environmental feedback

## When to Use
- WHEN mode=BUILD and the task requires self organization
- WHEN action=`initiate`: perform initiate operation
- WHEN action=`guide`: perform guide operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'initiate' | 'guide' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| agents | object | yes | Agents |
| id | string | yes | Id |
| capabilities | object | yes | Capabilities |
| preferences | object | yes | Preferences |
| current_role | string | no | Current role |

## Invocation
```
floyd skill:self-organization --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
