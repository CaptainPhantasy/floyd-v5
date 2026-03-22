---
name: performance-profiling-patterns
description: Systematically identify performance bottlenecks through CPU, memory, I/O, and concurrency profiling with actionable optimization insights
category: core
version: "2.0.0"
---

# Performance Profiling Patterns

> Systematically identify performance bottlenecks through CPU, memory, I/O, and concurrency profiling with actionable optimization insights.

## When to Use
- WHEN `mode=DEBUG` and a user reports slowness, high memory usage, or latency spikes
- WHEN `mode=BUILD` and you need to establish a performance baseline before optimizing
- WHEN `mode=EXPLORE` and you want to understand the performance characteristics of a codebase

## Actions
`'profile' | 'baseline' | 'compare' | 'report'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| target | string | yes | File path, function name, or binary to profile |
| kind | string | yes | Profile type: cpu, memory, io, goroutine, http |
| duration | integer | no | Collection duration in seconds (default: 30) |
| baseline_id | string | no | Benchmark ID to compare against (for compare action) |

## Execution Pipeline

### Step 1: Profile
Use `mcp_floyd-devtools_benchmark_runner` with action `run` to execute a benchmark of the `target` code snippet. For Go projects, use `bash` to run `go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof` on the relevant package. Use `bash` to run `go tool pprof` for analysis.

### Step 2: Analyze
Parse the profiling output. Identify the top 5 hotspots by cumulative time or allocation. Use `view` to inspect the source code at each hotspot. Use `list_symbols` to understand call chains.

### Step 3: Report
Generate a structured report with bottleneck locations, severity ranking, and specific optimization recommendations (e.g., "reduce allocations in loop at `file.go:142` by pre-allocating the slice").

## Output Shape
```json
{
  "profile_type": "string — cpu | memory | io | goroutine | http",
  "duration_seconds": 30,
  "hotspots": [
    {
      "location": "string — file:line or function name",
      "metric": "string — e.g., 45.2% CPU, 12.3 MB alloc",
      "severity": "HIGH | MEDIUM | LOW",
      "recommendation": "string — specific actionable fix"
    }
  ],
  "total_samples": 15000,
  "comparison": {
    "baseline_id": "string — if compared",
    "delta_percent": "+12.3% slower"
  }
}
```

## Failure Modes
- IF the target cannot be benchmarked (not testable in isolation): fall back to static analysis using `view` and `list_symbols` to identify likely bottlenecks (nested loops, unchecked allocations, synchronous I/O in hot paths)
- IF the profiler times out: reduce `duration` and sample a smaller window, or profile a subset of the target

## Examples
```json
{
  "action": "profile",
  "target": "internal/engine/planner.go:Evaluate",
  "kind": "cpu",
  "duration": 60
}
```
