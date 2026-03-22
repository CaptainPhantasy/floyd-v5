---
name: quality-scoring
description: 140-point algorithm for comprehensive artifact evaluation across code docs and architecture
category: core
version: "2.0.0"
---

# Quality Scoring

> 140-point algorithm for comprehensive artifact evaluation across code docs and architecture

## When to Use
- WHEN mode=BUILD and the task requires quality scoring
- WHEN action=`score`: evaluate and grade against criteria
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`improve`: reduce size or improve quality of input

## Actions
`'score' | 'analyze' | 'improve'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| target | object | yes | Target |
| code | string | no | Code |
| pattern | object | no | Pattern |
| design | object | no | Design |

## Invocation
```
floyd skill:quality-scoring --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
