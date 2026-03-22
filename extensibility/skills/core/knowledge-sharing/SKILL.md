---
name: knowledge-sharing
description: Facilitate knowledge sharing between agents with versioning and conflict resolution
category: core
version: "2.0.0"
---

# Knowledge Sharing

> Facilitate knowledge sharing between agents with versioning and conflict resolution

## When to Use
- WHEN mode=BUILD and the task requires knowledge sharing
- WHEN action=`facilitate`: perform facilitate operation
- WHEN action=`curate`: perform curate operation
- WHEN action=`disseminate`: perform disseminate operation

## Actions
`'facilitate' | 'curate' | 'disseminate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| knowledge_sources | object | yes | Knowledge sources |
| source_id | string | yes | Source id |
| knowledge_type | enum | yes | Knowledge type |
| content | any | yes | Content |
| quality_score | number | yes | Quality score |
| accessibility | number | yes | Accessibility |

## Invocation
```
floyd skill:knowledge-sharing --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
