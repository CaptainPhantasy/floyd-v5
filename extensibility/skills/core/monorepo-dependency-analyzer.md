---
name: monorepo-dependency-analyzer
description: Analyze dependencies across monorepo packages, detect circular dependencies, and visualize the dependency graph.
---

# Monorepo Dependency Analyzer

> Analyze dependencies across monorepo packages, detect circular dependencies, and visualize the dependency graph.

**Category**: General

## When to Use
- When managing or optimizing context
- When analyzing code architecture or dependencies

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| repo_path | string | yes | repo_path parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `monorepo-dependency-analyzer`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.