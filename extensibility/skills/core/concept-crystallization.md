---
name: concept-crystallization
description: Transform vague concepts into precise, actionable definitions with clear boundaries and relationships.
---

# Concept Crystallization

> Transform vague concepts into precise, actionable definitions with clear boundaries and relationships.

**Category**: General

## When to Use
- When executing this skill's core functionality
- When you need the capabilities this skill provides

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| concept | string | yes | concept parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `concept-crystallization`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.