---
name: adaptive-strategy
description: Develop strategies that adapt to changing conditions with trigger-based pivots
category: core
version: "2.0.0"
---

# Adaptive Strategy

> Develop strategies that adapt to changing conditions with trigger-based pivots

## When to Use
- WHEN mode=BUILD and the task requires adaptive strategy
- WHEN action=`develop`: perform develop operation
- WHEN action=`implement`: perform implement operation
- WHEN action=`evolve`: perform evolve operation

## Actions
`'develop' | 'implement' | 'evolve'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| strategic_context | object | yes | Strategic context |
| environment | object | yes | Environment |
| objectives | string[] | yes | Objectives |
| constraints | object | yes | Constraints |
| uncertainty_level | number | yes | Uncertainty level |

## Invocation
```
floyd skill:adaptive-strategy --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
