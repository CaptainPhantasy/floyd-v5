---
name: meta-reasoning
description: Reason about reasoning processes to optimize thinking strategies
category: core
version: "2.0.0"
---

# Meta Reasoning

> Reason about reasoning processes to optimize thinking strategies

## When to Use
- WHEN mode=EXPLORE and the task requires meta reasoning
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`optimize`: optimize for efficiency
- WHEN action=`evaluate`: perform evaluate operation

## Actions
`'analyze' | 'optimize' | 'evaluate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| reasoning_process | object | yes | Reasoning process |
| problem | string | yes | Problem |
| approach | string | yes | Approach |
| steps | object | yes | Steps |
| step | number | yes | Step |
| reasoning | string | yes | Reasoning |
| evidence | string[] | yes | Evidence |
| confidence | number | yes | Confidence |

## Invocation
```
floyd skill:meta-reasoning --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
