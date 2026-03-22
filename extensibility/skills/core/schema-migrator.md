---
name: schema-migrator
description: Generate database schema migrations with validation, rollback plans, and SQL generation for safe schema evolution.
---

# Schema Migrator

> Generate database schema migrations with validation, rollback plans, and SQL generation for safe schema evolution.

**Category**: General

## When to Use
- When executing this skill's core functionality
- When you need the capabilities this skill provides

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| source_schema | object | yes | source_schema parameter |
| target_schema | object | yes | target_schema parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `schema-migrator`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.