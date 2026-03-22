---
name: context-compression
description: Compress context while preserving semantic meaning using abstractive and extractive techniques.
---

# Context Compression

> Compress context while preserving semantic meaning using abstractive and extractive techniques.

**Category**: General

## When to Use
- When managing or optimizing context

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| context | string | yes | context parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `context-compression`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.