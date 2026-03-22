---
name: trace-replay-debugger
description: Record and replay execution traces for time-travel debugging and root cause analysis.
---

# Trace Replay Debugger

> Record and replay execution traces for time-travel debugging and root cause analysis.

**Category**: General

## When to Use
- When analyzing failures or debugging issues

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `trace-replay-debugger`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.