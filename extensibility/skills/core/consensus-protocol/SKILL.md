---
name: consensus-protocol
description: Implement consensus protocols for multi-agent decision making with conflict resolution
category: core
version: "2.0.0"
---

# Consensus Protocol

> Implement consensus protocols for multi-agent decision making with conflict resolution

## When to Use
- WHEN mode=BUILD and the task requires consensus protocol
- WHEN action=`establish`: perform establish operation
- WHEN action=`execute`: perform operation
- WHEN action=`analyze`: inspect and report without modification

## Actions
`'establish' | 'execute' | 'analyze'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| participants | object | yes | Participants |
| id | string | yes | Id |
| voting_power | number | yes | Voting power |
| preferences | object | yes | Preferences |
| communication_reliability | number | yes | Communication reliability |

## Invocation
```
floyd skill:consensus-protocol --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
