---
name: concept-crystallization
description: Transform vague concepts into precise actionable definitions with clear boundaries
category: core
version: "2.0.0"
---

# Concept Crystallization

> Transform vague concepts into precise actionable definitions with clear boundaries

## When to Use
- WHEN mode=EXPLORE and the task requires concept crystallization
- WHEN action=`crystallize`: perform crystallize operation
- WHEN action=`refine`: perform refine operation
- WHEN action=`expand`: perform expand operation

## Actions
`'crystallize' | 'refine' | 'expand'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| raw_concepts | object | yes | Raw concepts |
| name | string | yes | Name |
| description | string | yes | Description |
| examples | any[] | yes | Examples |
| context | object | yes | Context |

## Invocation
```
floyd skill:concept-crystallization --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
