---
name: benchmark-runner
description: Execute code benchmarks with statistical analysis, regression detection, and performance recommendations.
---

# Benchmark Runner

> Execute code benchmarks with statistical analysis, regression detection, and performance recommendations.

**Category**: General

## When to Use
- When analyzing or optimizing performance

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| benchmark_id | string | yes | benchmark_id parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `benchmark-runner`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.