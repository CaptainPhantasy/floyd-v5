---
name: patch-oracle
description: Predicts the 'blast radius' of a change by identifying architectural fault lines using the Fiedler vector.
---

# Patch Oracle

> Predicts the 'blast radius' of a change by identifying architectural fault lines using the Fiedler vector.

**Category**: Ghost Algorithm

## When to Use
- When system conditions change or require adaptation

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| dep_graph | object | yes | No description |
| target_files | array | yes | No description |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `patch-oracle`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.