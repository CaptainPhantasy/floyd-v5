---
name: benchmark-runner
description: Execute code benchmarks with statistical analysis regression detection and performance recommendations
category: core
version: "2.0.0"
---

# Benchmark Runner

> Execute code benchmarks with statistical analysis regression detection and performance recommendations

## When to Use
- WHEN mode=BUILD and the task requires benchmark runner
- WHEN action=`run`: execute benchmark
- WHEN action=`compare`: diff two inputs and report differences
- WHEN action=`baseline`: establish baseline metrics
- WHEN action=`report`: generate report

## Actions
`'run' | 'compare' | 'baseline' | 'report'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| benchmark_id | string | yes | Benchmark id |
| code_snippet | string | no | Code snippet |
| iterations | number | no | Iterations |
| warmup_runs | number | no | Warmup runs |
| compare_to | string | no | Compare to |
| threshold_percent | number | no | Threshold percent |

## Invocation
```
floyd skill:benchmark-runner --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
