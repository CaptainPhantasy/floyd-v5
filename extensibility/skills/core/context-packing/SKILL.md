---
name: context-packing
description: Efficiently pack relevant context into limited token windows with semantic compression
category: core
version: "2.0.0"
---

# Context Packing

> Efficiently pack relevant context into limited token windows with semantic compression

## When to Use
- WHEN mode=BUILD and the task requires context packing
- WHEN action=`pack`: reduce size or improve quality of input
- WHEN action=`unpack`: decompress packed context
- WHEN action=`optimize`: reduce size or improve quality of input

## Actions
`'pack' | 'unpack' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| data | object | yes | Data |
| text_content | string | no | Text content |
| structured_data | object | no | Structured data |
| metadata | object | no | Metadata |
| priorities | Record<string, number> | no | Priorities |

## Invocation
```
floyd skill:context-packing --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
