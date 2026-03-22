---
name: knowledge-graph-building
description: Build knowledge graphs from unstructured data with entity extraction and relationship inference
category: core
version: "2.0.0"
---

# Knowledge Graph Building

> Build knowledge graphs from unstructured data with entity extraction and relationship inference

## When to Use
- WHEN mode=EXPLORE and the task requires knowledge graph building
- WHEN action=`build`: perform build operation
- WHEN action=`extend`: perform extend operation
- WHEN action=`validate`: check conformance and report violations

## Actions
`'build' | 'extend' | 'validate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| entities | object | yes | Entities |
| id | string | yes | Id |
| type | string | yes | Type |
| properties | Record<string, any> | yes | Properties |

## Invocation
```
floyd skill:knowledge-graph-building --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
