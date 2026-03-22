---
name: context-compression
description: Compress context while preserving semantic meaning and critical information
category: core
version: "2.0.0"
---

# Context Compression

> Compress context while preserving semantic meaning and critical information

## When to Use
- WHEN mode=BUILD and the task requires context compression
- WHEN action=`compress`: perform compress operation
- WHEN action=`decompress`: perform decompress operation
- WHEN action=`analyze`: inspect and report without modification

## Actions
`'compress' | 'decompress' | 'analyze'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| context_data | object | yes | Context data |
| content | string | object | yes | Content |
| size_kb | number | yes | Size kb |
| type | enum | yes | Type |

## Invocation
```
floyd skill:context-compression --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
