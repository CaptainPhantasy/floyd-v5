---
name: concept-web-analysis
description: Analyze interconnected concepts as a web/graph with relationship scoring
category: core
version: "2.0.0"
---

# Concept Web Analysis

> Analyze interconnected concepts as a web/graph with relationship scoring

## When to Use
- WHEN mode=EXPLORE and the task requires concept web analysis
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`visualize`: generate visual representation
- WHEN action=`navigate`: perform navigate operation

## Actions
`'analyze' | 'visualize' | 'navigate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| concepts | object | yes | Concepts |
| name | string | yes | Name |
| definition | string | yes | Definition |
| attributes | Record<string, any> | yes | Attributes |
| relationships | object | yes | Relationships |
| target | string | yes | Target |
| type | string | yes | Type |
| strength | number | yes | Strength |

## Invocation
```
floyd skill:concept-web-analysis --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
