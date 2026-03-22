---
name: pattern-synthesis
description: Synthesize new patterns through combination and abstraction of existing patterns
category: core
version: "2.0.0"
---

# Pattern Synthesis

> Synthesize new patterns through combination and abstraction of existing patterns

## When to Use
- WHEN mode=BUILD and the task requires pattern synthesis
- WHEN action=`synthesize`: perform synthesize operation
- WHEN action=`combine`: perform combine operation
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'synthesize' | 'combine' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| base_patterns | object | yes | Base patterns |
| name | string | yes | Name |
| structure | object | yes | Structure |
| usage_context | string | yes | Usage context |
| effectiveness | number | yes | Effectiveness |

## Invocation
```
floyd skill:pattern-synthesis --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
