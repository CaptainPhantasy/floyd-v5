---
name: viterbi-resolver
description: Resolves ambiguous or 'fuzzy' agent output by calculating the maximum-likelihood correct token sequence.
---

# Viterbi Resolver

> Resolves ambiguous or 'fuzzy' agent output by calculating the maximum-likelihood correct token sequence.

**Category**: Ghost Algorithm

## When to Use
- When coordinating multiple agents or systems

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| ambiguous_code | string | yes | No description |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `viterbi-resolver`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.