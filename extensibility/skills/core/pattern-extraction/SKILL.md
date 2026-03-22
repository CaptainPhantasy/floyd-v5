---
name: pattern-extraction
description: Extract reusable code patterns and architectural patterns from existing codebases
category: core
version: "2.0.0"
---

# Pattern Extraction

> Extract reusable code patterns and architectural patterns from existing codebases

## When to Use
- WHEN mode=BUILD and the task requires pattern extraction
- WHEN action=`extract`: perform extract operation
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`synthesize`: perform synthesize operation

## Actions
`'extract' | 'analyze' | 'synthesize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| source_data | object | yes | Source data |
| code_examples | string[] | no | Code examples |
| documentation | string | no | Documentation |
| logs | string[] | no | Logs |
| conversation_history | any[] | no | Conversation history |

## Invocation
```
floyd skill:pattern-extraction --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
