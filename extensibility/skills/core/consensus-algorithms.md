---
name: consensus-algorithms
description: Systematic approaches to achieving agreement among multiple agents, entities, or decision-makers through structured protocols that handle conflicts, preferences, and requirements for fair and reliable
---

# Consensus Algorithms

> Systematic approaches to achieving agreement among multiple agents, entities, or decision-makers through structured protocols that handle conflicts, preferences, and requirements for fair and reliable

**Category**: General

## When to Use
- When coordinating multiple agents or systems
- When making collective decisions or reaching consensus
- When analyzing code architecture or dependencies

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| issue | object | yes | The decision problem definition |
| preferences | object | yes | Participant preferences keyed by participant ID |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `consensus-algorithms`
- **Args**: `{key inputs as JSON object}`

## Output
Returns success, metadata, audit_trail and related metadata.