---
name: tarjan-circular-dependency-detection
description: Detect circular dependencies in codebases using Tarjan's SCC algorithm, identifying import cycles that cause initialization problems
category: core
version: "2.0.0"
---

# Tarjan Circular Dependency Detection

> Detect circular dependencies in codebases using Tarjan's SCC algorithm, identifying import cycles that cause initialization problems.

## When to Use
- WHEN `mode=EXPLORE` to audit a codebase for architectural dependency problems
- WHEN `mode=DEBUG` and an initialization error or import loop is suspected
- WHEN `mode=BUILD` before adding a new import to verify it won't create a cycle

## Actions
`'analyze' | 'visualize' | 'find_cycles' | 'suggest_fixes'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| project_path | string | no | Root directory of the project (default: working directory) |
| language | string | no | Language to analyze: typescript, javascript, python, go, auto (default: auto) |
| entry_point | string | no | Specific file to use as analysis root (default: full project) |

## Execution Pipeline

### Step 1: Detect Cycles
Use `mcp_floyd-devtools_dependency_analyzer` with action `find_cycles` on the `project_path`. This runs Tarjan's SCC algorithm on the import graph. If the MCP tool is unavailable, use `grep` to extract all import/require statements and build a manual adjacency list.

### Step 2: Visualize
Use `mcp_floyd-devtools_dependency_analyzer` with action `visualize` to produce an ASCII dependency graph. Identify which nodes participate in cycles.

### Step 3: Suggest Fixes
For each cycle found, analyze the import graph to determine the weakest edge (the import that is least depended upon by other modules). Recommend breaking the cycle by introducing an interface or moving the shared dependency to a common package.

## Output Shape
```json
{
  "cycles_found": 3,
  "cycles": [
    {
      "nodes": ["string — module/file in the cycle"],
      "edges": [["string", "string"] — directed import pairs],
      "severity": "HIGH | MEDIUM | LOW",
      "suggested_break": {
        "from": "string — module to remove import from",
        "to": "string — module no longer imported",
        "strategy": "string — extract-interface | move-to-common | invert-dependency"
      }
    }
  ],
  "total_modules_analyzed": 142,
  "graph_density": 0.34
}
```

## Failure Modes
- IF the language is not supported: fall back to regex-based import extraction using `grep`
- IF no cycles are found: report a clean dependency graph with the total module count

## Examples
```json
{
  "action": "find_cycles",
  "project_path": "/workspace/myproject",
  "language": "go"
}
```
