---
name: knowledge-synthesis
description: Synthesize knowledge from multiple sources into structured actionable understanding
category: core
version: "2.0.0"
---

# Knowledge Synthesis

> Synthesize knowledge from multiple sources into structured actionable understanding

## When to Use
- WHEN mode=EXPLORE and the task requires knowledge synthesis
- WHEN action=`synthesize`: perform synthesize operation
- WHEN action=`integrate`: perform integrate operation
- WHEN action=`generate`: produce new artifacts from inputs

## Actions
`'synthesize' | 'integrate' | 'generate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| knowledge_sources | object | yes | Knowledge sources |
| type | enum | yes | Type |
| content | any | yes | Content |
| confidence | number | yes | Confidence |
| relevance | number | yes | Relevance |

## Invocation
```
floyd skill:knowledge-synthesis --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
