---
name: api-contract-validation
description: Systematically validate API contracts for compatibility, consistency, and compliance with design principles. Detects breaking changes, enforces conventions, and provides client impact analysis with de
---

# Api Contract Validation

> Systematically validate API contracts for compatibility, consistency, and compliance with design principles. Detects breaking changes, enforces conventions, and provides client impact analysis with de

**Category**: General

## When to Use
- When system conditions change or require adaptation
- When performing code review or quality checks

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| current_contract | object | yes | No description |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `api-contract-validation`
- **Args**: `{key inputs as JSON object}`

## Output
Returns success, metadata, migration_plan and related metadata.