---
name: refactor-pathfinder
description: Calculates the minimum-edit distance (shortest path) from 'Broken State' to 'Target State'.
---

# Refactor Pathfinder

> Calculates the minimum-edit distance (shortest path) from 'Broken State' to 'Target State'.

**Category**: Ghost Algorithm

## When to Use
- When analyzing failures or debugging issues

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| start | string | yes | No description |
| goal | string | yes | No description |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `refactor-pathfinder`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.