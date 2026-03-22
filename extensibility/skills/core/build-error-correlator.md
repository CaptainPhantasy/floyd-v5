---
name: build-error-correlator
description: Correlate build errors across multiple builds to identify patterns, recurring issues, and suggest automated fixes.
---

# Build Error Correlator

> Correlate build errors across multiple builds to identify patterns, recurring issues, and suggest automated fixes.

**Category**: General

## When to Use
- When coordinating multiple agents or systems
- When analyzing failures or debugging issues

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| build_logs | array | yes | build_logs parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `build-error-correlator`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.