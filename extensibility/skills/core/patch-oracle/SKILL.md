---
name: patch-oracle
description: Predicts the blast radius of a change by analyzing dependency graphs and identifying files at risk of cascading failure
category: core
version: "2.0.0"
---

# Patch Oracle

> Predicts the blast radius of a change by analyzing dependency graphs and identifying files at risk of cascading failure.

## When to Use
- WHEN `mode=BUILD` before applying a large refactor or API change to assess impact
- WHEN `mode=DEBUG` to understand why a change in one module broke an apparently unrelated module
- WHEN reviewing a PR to identify files that should be tested but are not in the diff

## Actions
`'analyze' | 'simulate' | 'report'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| target_files | string[] | yes | Files being modified by the change |
| project_path | string | no | Root of the project (default: working directory) |
| include_tests | boolean | no | Whether to include test files in blast radius (default: true) |

## Execution Pipeline

### Step 1: Analyze Dependency Graph
Use `mcp_floyd-devtools_dependency_analyzer` with action `analyze` on the `project_path`. This builds the import dependency graph using language-aware parsing. If the MCP tool is unavailable, fall back to `grep` to extract import/require statements from all source files.

### Step 2: Simulate Blast Radius
For each file in `target_files`, trace all direct and transitive dependents in the graph. Use `mcp_floyd-devtools_dependency_analyzer` with action `visualize` to see the full dependency tree. Rank affected files by distance (direct > indirect) and by number of incoming edges (high-coupling nodes are highest risk).

### Step 3: Report
Cross-reference the blast radius against the test files in the diff. Flag any high-risk files that lack test coverage. Output a structured report with risk tiers.

## Output Shape
```json
{
  "target_files": ["string — files in the change"],
  "blast_radius": {
    "direct": ["string — files that directly import the targets"],
    "transitive": ["string — files that indirectly depend on the targets"],
    "total_count": 42
  },
  "risk_tiers": {
    "critical": ["string — high-coupling files with no test coverage"],
    "moderate": ["string — files with indirect dependency only"],
    "low": ["string — test files or generated code"]
  },
  "untested_high_risk": ["string — files that should have tests but don't"]
}
```

## Failure Modes
- IF the dependency analyzer fails on a language: fall back to `grep` for import statements and build a manual graph
- IF the project has no import structure (e.g., shell scripts): report that blast radius analysis is not applicable and list all files that reference the target by name

## Examples
```json
{
  "action": "analyze",
  "target_files": ["internal/engine/planner.go", "internal/engine/executor.go"],
  "project_path": "/workspace/myproject",
  "include_tests": true
}
```
