---
name: semantic-understanding
description: Deep semantic analysis extracting meaning intent and relationships from text
category: core
version: "2.0.0"
---

# Semantic Understanding

> Deep semantic analysis extracting meaning intent and relationships from text

## When to Use
- WHEN mode=EXPLORE and the task requires semantic understanding
- WHEN action=`understand`: perform understand operation
- WHEN action=`extract`: perform extract operation
- WHEN action=`relate`: find patterns across multiple inputs

## Actions
`'understand' | 'extract' | 'relate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| content | object | yes | Content |
| text | string | yes | Text |
| domain | string | no | Domain |
| context | object | no | Context |

## Invocation
```
floyd skill:semantic-understanding --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
