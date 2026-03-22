---
name: code-review-workflow
description: Systematic code review process that combines automated analysis with human review patterns to ensure code quality, maintainability, and team knowledge sharing. Provides structured checklists, focus ar
---

# Code Review Workflow

> Systematic code review process that combines automated analysis with human review patterns to ensure code quality, maintainability, and team knowledge sharing. Provides structured checklists, focus ar

**Category**: General

## When to Use
- When performing code review or quality checks
- When analyzing code architecture or dependencies

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| pull_request | object | yes | No description |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `code-review-workflow`
- **Args**: `{key inputs as JSON object}`

## Output
Returns success, findings, metadata and related metadata.