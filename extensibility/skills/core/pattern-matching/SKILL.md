---
name: pattern-matching
description: Identify structural and behavioral patterns across diverse data sources
category: core
version: "2.0.0"
---

# Pattern Matching

> Identify structural and behavioral patterns across diverse data sources

## When to Use
- WHEN mode=BUILD and the task requires pattern matching
- WHEN action=`match`: perform match operation
- WHEN action=`search`: locate or enumerate matching items
- WHEN action=`cluster`: perform cluster operation

## Actions
`'match' | 'search' | 'cluster'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| target_pattern | object | yes | Target pattern |
| structure | object | yes | Structure |
| features | string[] | yes | Features |
| constraints | object | yes | Constraints |

## Invocation
```
floyd skill:pattern-matching --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
