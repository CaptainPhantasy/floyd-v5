---
name: consensus-algorithms
description: Systematic approaches to achieving agreement among multiple agents through structured protocols that handle conflicts and preferences
category: core
version: "2.0.0"
---

# Consensus Algorithms

> Systematic approaches to achieving agreement among multiple agents through structured protocols that handle conflicts and preferences.

## When to Use
- WHEN `mode=BUILD` and multiple sub-agents or components must agree on a plan, architecture decision, or conflict resolution
- WHEN `mode=EXPLORE` to evaluate different decision-making strategies for a multi-agent system
- WHEN `mode=DEBUG` and conflicting diagnostics from multiple tools need reconciliation

## Actions
`'propose' | 'vote' | 'resolve' | 'audit'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| issue | string | yes | The decision problem or conflict to resolve |
| participants | string[] | yes | IDs or names of the participating agents/entities |
| preferences | object | no | Map of participant ID to their stated preference or position |
| protocol | string | no | Consensus protocol: majority, unanimous, ranked-choice, quorum (default: majority) |

## Execution Pipeline

### Step 1: Propose
Each participant submits their position on the `issue`. If `preferences` is provided, use it directly. Otherwise, use `manage_scratchpad` to collect positions from each participant's analysis.

### Step 2: Vote
Apply the selected `protocol`:
- **Majority**: the option with >50% support wins
- **Unanimous**: all participants must agree; if not, enter conflict resolution
- **Ranked-choice**: participants rank options; eliminate lowest-ranked until a majority emerges
- **Quorum**: a minimum threshold of participants must agree (default: 66%)

### Step 3: Resolve
If consensus is reached, record the decision and rationale. If not, identify the blocking participants and the nature of their disagreement. Use `manage_scratchpad` to document the deadlock and suggest mediation (e.g., compromise option, escalation to human).

### Step 4: Audit
Record the full decision trace: issue, participants, preferences, protocol applied, result, and any dissenting opinions.

## Output Shape
```json
{
  "issue": "string — the decision problem",
  "protocol": "string — consensus protocol used",
  "result": "CONSENSUS | DEADLOCK | ESCALATED",
  "decision": "string — the agreed-upon resolution (if consensus)",
  "vote_tally": {
    "option_a": 3,
    "option_b": 1
  },
  "dissenters": ["string — participants who disagreed"],
  "audit_trail": [
    {
      "participant": "string — agent ID",
      "position": "string — their stated preference",
      "rationale": "string — why they chose this"
    }
  ]
}
```

## Failure Modes
- IF DEADLOCK occurs with unanimous protocol: fall back to majority or suggest the human operator decide
- IF a participant fails to submit a preference: exclude them from the quorum count and note their absence in the audit trail

## Examples
```json
{
  "action": "vote",
  "issue": "Should we use PostgreSQL or SQLite for the local cache?",
  "participants": ["planner-agent", "security-agent", "performance-agent"],
  "protocol": "majority"
}
```
