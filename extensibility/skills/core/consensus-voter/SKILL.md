---
name: consensus-voter
description: Mathematically identify and discard hallucinating agent outliers via statistical consensus analysis
category: core
version: "2.0.0"
---

# Consensus Voter

> Mathematically identify and discard hallucinating agent outliers via statistical consensus analysis.

## When to Use
- WHEN `mode=BUILD` and multiple agents or LLM calls produce conflicting results for the same query
- WHEN `mode=DEBUG` and a diagnostic from one tool contradicts diagnostics from others
- WHEN validating code generation by running multiple analysis passes and selecting the most reliable output

## Actions
`'vote' | 'outlier-detect' | 'aggregate'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| responses | object[] | yes | Array of responses: `{source: string, result: string, confidence: number}` |
| question | string | yes | The question or task that was posed to all sources |
| threshold | number | no | Z-score threshold for outlier detection (default: 2.0) |

## Execution Pipeline

### Step 1: Normalize Responses
Parse each response to extract a comparable output. For code: compare AST structure. For factual answers: compare semantic similarity. For classifications: compare labels. Assign each response a numerical score.

### Step 2: Detect Outliers
Compute the mean and standard deviation of scores. Flag any response with a z-score exceeding `threshold` as a potential outlier (hallucination). Use `mcp_floyd-terminal_execute_code` with Python to compute the statistics.

### Step 3: Aggregate
Weight each remaining response by its `confidence` score (if provided) or by inverse distance from the mean. Produce the consensus result. Report which sources agreed, which disagreed, and which were discarded.

## Output Shape
```json
{
  "question": "string — the original query",
  "consensus_result": "string — the agreed-upon answer",
  "agreement_rate": 0.83,
  "total_sources": 5,
  "outliers_discarded": [
    {
      "source": "string — agent or tool name",
      "z_score": 3.2,
      "reason": "string — why it was flagged"
    }
  ],
  "confidence_distribution": {
    "mean": 0.78,
    "std_dev": 0.12
  }
}
```

## Failure Modes
- IF all responses are outliers (no consensus): report a stalemate and recommend escalating to human review
- IF responses cannot be normalized to comparable scores: fall back to majority voting on raw string similarity

## Examples
```json
{
  "action": "vote",
  "question": "What is the return type of parser.Evaluate?",
  "responses": [
    {"source": "agent-1", "result": "*Plan", "confidence": 0.9},
    {"source": "agent-2", "result": "Plan", "confidence": 0.85},
    {"source": "agent-3", "result": "error", "confidence": 0.3}
  ]
}
```
