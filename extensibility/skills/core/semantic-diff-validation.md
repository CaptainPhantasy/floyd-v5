---
name: semantic-diff-validation
description: Validate that code changes preserve semantic behavior while allowing structural improvements. Detects breaking changes, behavioral modifications, and potential bugs with quantified risk scoring and ac
---

# Semantic Diff Validation

> Validate that code changes preserve semantic behavior while allowing structural improvements. Detects breaking changes, behavioral modifications, and potential bugs with quantified risk scoring and ac

**Category**: General

## When to Use
- When system conditions change or require adaptation
- When performing code review or quality checks

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| diff | unknown | yes | No description |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `semantic-diff-validation`
- **Args**: `{key inputs as JSON object}`

## Output
Returns success, findings, metadata and related metadata.