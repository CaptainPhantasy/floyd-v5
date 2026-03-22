---
name: typescript-semantic-analyzer
description: Perform deep semantic analysis of TypeScript code including symbol extraction, type inference, and reference tracking across projects.
---

# Typescript Semantic Analyzer

> Perform deep semantic analysis of TypeScript code including symbol extraction, type inference, and reference tracking across projects.

**Category**: General

## When to Use
- When executing this skill's core functionality
- When you need the capabilities this skill provides

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| file_path | string | yes | file_path parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `typescript-semantic-analyzer`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.