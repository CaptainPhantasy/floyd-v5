---
name: mit-analysis
description: Multiple Instance Theory analysis for decomposing complex problems into ranked hypotheses with evidence weighting
category: core
version: "2.0.0"
---

# MIT Analysis

> Multiple Instance Theory analysis for decomposing complex problems into ranked hypotheses with evidence weighting.

## When to Use
- WHEN `mode=DEBUG` and a bug has multiple possible causes that need systematic evaluation
- WHEN `mode=EXPLORE` and a complex architectural decision has many competing approaches
- WHEN `mode=BUILD` and a design problem requires evaluating trade-offs across multiple dimensions

## Actions
`'hypothesize' | 'evaluate' | 'rank' | 'select'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| problem | string | yes | The problem statement or bug description |
| instances | string[] | no | Specific instances or examples of the problem occurring |
| constraints | string[] | no | Constraints on the solution (performance, compatibility, complexity) |
| dimensions | string[] | no | Evaluation dimensions: correctness, performance, complexity, maintainability (default: all four) |

## Execution Pipeline

### Step 1: Hypothesize
Decompose the `problem` into a set of mutually exclusive hypotheses. For each hypothesis, identify:
- **Mechanism**: how it would cause the observed behavior
- **Predictions**: what else would be true if this hypothesis is correct
- **Disprovers**: what evidence would refute it

Use `grep` to search the codebase for evidence supporting or refuting each hypothesis. Use `view` to inspect relevant code sections.

### Step 2: Evaluate
For each hypothesis, collect evidence using `grep` and `view`. Score each hypothesis on each `dimension` from 0 to 1. Weight the scores by the constraints (e.g., if performance is a hard constraint, weight it 3×).

### Step 3: Rank
Compute a weighted composite score for each hypothesis. Sort by descending score. Identify the top candidate and the closest competitor. If the top two scores are within 15%, flag as ambiguous.

### Step 4: Select
Return the top-ranked hypothesis with its evidence chain. If ambiguous, return both candidates and recommend an experiment to discriminate between them.

## Output Shape
```json
{
  "problem": "string — the problem statement",
  "hypotheses": [
    {
      "rank": 1,
      "description": "string — the hypothesis",
      "mechanism": "string — how it causes the problem",
      "evidence": ["string — supporting evidence"],
      "disprovers": ["string — what would refute it"],
      "composite_score": 0.87,
      "dimension_scores": {"correctness": 0.9, "performance": 0.8, "complexity": 0.85, "maintainability": 0.9}
    }
  ],
  "ambiguous": false,
  "recommended_experiment": "string — if ambiguous, what test would discriminate"
}
```

## Failure Modes
- IF no hypotheses can be formulated: the problem is underspecified — ask for more concrete instances or error output
- IF all hypotheses score below 0.3: the root cause is likely outside the analyzed scope — recommend expanding the search

## Examples
```json
{
  "action": "evaluate",
  "problem": "Memory usage grows linearly during long-running processing",
  "instances": ["after 10k items: 200MB", "after 100k items: 1.8GB"],
  "constraints": ["performance"],
  "dimensions": ["correctness", "performance", "maintainability"]
}
```
