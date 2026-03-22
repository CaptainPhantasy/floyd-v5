---
name: attention-allocation
description: Optimize attention allocation across multiple tasks, contexts, and priorities.
---

# Attention Allocation

> Optimize attention allocation across multiple tasks, contexts, and priorities.

**Category**: General

## When to Use
- When coordinating multiple agents or systems
- When managing or optimizing context
- When analyzing or optimizing performance

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| tasks | array | yes | tasks parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `attention-allocation`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.