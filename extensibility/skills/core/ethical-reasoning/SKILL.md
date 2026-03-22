---
name: ethical-reasoning
description: Apply ethical frameworks to decision making with impact analysis and audit trails
category: core
version: "2.0.0"
---

# Ethical Reasoning

> Apply ethical frameworks to decision making with impact analysis and audit trails

## When to Use
- WHEN mode=EXPLORE and the task requires ethical reasoning
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`evaluate`: perform evaluate operation
- WHEN action=`recommend`: locate or enumerate matching items

## Actions
`'analyze' | 'evaluate' | 'recommend'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| ethical_scenario | object | yes | Ethical scenario |
| situation | string | yes | Situation |
| stakeholders | object | yes | Stakeholders |
| role | string | yes | Role |
| interests | string[] | yes | Interests |
| rights | string[] | yes | Rights |

## Invocation
```
floyd skill:ethical-reasoning --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
