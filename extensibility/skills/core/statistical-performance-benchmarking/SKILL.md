---
name: statistical-performance-benchmarking
description: Perform rigorous statistical performance analysis with high-resolution timing, confidence intervals, and outlier detection
category: core
version: "2.0.0"
---

# Statistical Performance Benchmarking

> Perform rigorous statistical performance analysis with high-resolution timing, confidence intervals, and outlier detection.

## When to Use
- WHEN `mode=BUILD` to measure the performance impact of an optimization
- WHEN `mode=EXPLORE` to establish a performance baseline for a codebase
- WHEN `mode=DEBUG` to quantify the overhead of a suspected bottleneck

## Actions
`'run' | 'compare' | 'baseline' | 'report'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| benchmark_id | string | yes | Unique identifier for this benchmark |
| code_snippet | string | yes | The code to benchmark (must be self-contained) |
| iterations | integer | no | Number of iterations (default: 100) |
| warmup_runs | integer | no | Number of warmup iterations excluded from results (default: 10) |
| compare_to | string | no | Baseline benchmark ID to compare against |

## Execution Pipeline

### Step 1: Execute Benchmark
Use `mcp_floyd-devtools_benchmark_runner` with action `run`. Pass the `benchmark_id`, `code_snippet`, `iterations`, and `warmup_runs`. The runner executes the code in a sandboxed environment with high-resolution timers.

### Step 2: Statistical Analysis
The runner automatically computes: mean, median, min, max, standard deviation, P95, P99, and throughput (ops/sec). Review the returned statistics for outlier presence (if P99 > 3× median, flag potential environmental interference).

### Step 3: Compare (Optional)
If `compare_to` is provided, use `mcp_floyd-devtools_benchmark_runner` with action `compare`. This computes the percent change between current and baseline results.

### Step 4: Report
Set the result as a baseline using action `baseline` if this is the first run. Output the full statistical report.

## Output Shape
```json
{
  "benchmark_id": "string — identifier",
  "iterations": 100,
  "warmup_runs": 10,
  "statistics": {
    "mean_ns": 1234.5,
    "median_ns": 1200.0,
    "min_ns": 1100.0,
    "max_ns": 5600.0,
    "std_dev_ns": 89.2,
    "p95_ns": 1400.0,
    "p99_ns": 3200.0,
    "throughput_ops_per_sec": 809716
  },
  "comparison": {
    "baseline_id": "string — if compared",
    "delta_percent": -15.3,
    "regression": false
  }
}
```

## Failure Modes
- IF the code snippet has a syntax error: the runner will fail — use `view` to verify the snippet before retrying
- IF outliers dominate (>20% of samples are >3× median): report environmental interference and suggest re-running with fewer concurrent processes

## Examples
```json
{
  "action": "run",
  "benchmark_id": "json-serialize-v2",
  "code_snippet": "const data = Array.from({length: 1000}, () => Math.random()); JSON.stringify(data);",
  "iterations": 200,
  "warmup_runs": 20
}
```
