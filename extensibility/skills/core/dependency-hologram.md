---
name: dependency-hologram
description: Generate interactive 2D/3D dependency visualizations for complex project architectures.
---

# Dependency Hologram

> Generate interactive 2D/3D dependency visualizations for complex project architectures.

**Category**: General

## When to Use
- When analyzing code architecture or dependencies

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| project_path | string | yes | project_path parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `dependency-hologram`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.