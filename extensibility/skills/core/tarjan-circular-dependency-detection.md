---
name: tarjan-circular-dependency-detection
description: Detect circular dependencies in codebases using Tarjan's Strongly Connected Components (SCC) algorithm. Identifies problematic import cycles that cause infinite loops, memory leaks, and initialization
---

# Tarjan Circular Dependency Detection

> Detect circular dependencies in codebases using Tarjan's Strongly Connected Components (SCC) algorithm. Identifies problematic import cycles that cause infinite loops, memory leaks, and initialization

**Category**: General

## When to Use
- When executing this skill's core functionality
- When you need the capabilities this skill provides

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | Operation to perform |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `tarjan-circular-dependency-detection`
- **Args**: `{key inputs as JSON object}`

## Output
Returns cycles, metrics, success and related metadata.