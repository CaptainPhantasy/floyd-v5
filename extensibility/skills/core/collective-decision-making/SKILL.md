---
name: collective-decision-making
description: Facilitate decision making across multiple agents with weighted voting protocols
category: core
version: "2.0.0"
---

# Collective Decision Making

> Facilitate decision making across multiple agents with weighted voting protocols

## When to Use
- WHEN mode=BUILD and the task requires collective decision making
- WHEN action=`facilitate`: perform facilitate operation
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`implement`: perform implement operation

## Actions
`'facilitate' | 'analyze' | 'implement'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| decision_context | object | yes | Decision context |
| problem | string | yes | Problem |
| constraints | object | yes | Constraints |
| stakeholders | object | yes | Stakeholders |
| id | string | yes | Id |
| influence | number | yes | Influence |
| preferences | object | yes | Preferences |

## Invocation
```
floyd skill:collective-decision-making --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
