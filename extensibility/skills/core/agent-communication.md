---
name: agent-communication
description: Facilitate structured communication between agents with message routing and protocol handling.
---

# Agent Communication

> Facilitate structured communication between agents with message routing and protocol handling.

**Category**: General

## When to Use
- When coordinating multiple agents or systems
- When analyzing code architecture or dependencies

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `agent-communication`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.