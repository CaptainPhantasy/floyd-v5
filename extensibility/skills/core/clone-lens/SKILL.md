---
name: clone-lens
description: Identify structural code duplicates across the codebase for refactoring candidates and maintenance reduction
category: core
version: "2.0.0"
---

# Clone Lens

> Identify structural code duplicates across the codebase for refactoring candidates and maintenance reduction.

## When to Use
- WHEN `mode=EXPLORE` to audit a codebase for code duplication before planning a refactor
- WHEN `mode=BUILD` to verify that a new function doesn't duplicate existing logic
- WHEN `mode=DEBUG` when the same bug appears in multiple locations (indicating cloned code)

## Actions
`'scan' | 'report' | 'suggest'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| scope | string[] | no | Directories or files to scan (default: entire project) |
| min_lines | integer | no | Minimum clone length in lines to report (default: 6) |
| min_copies | integer | no | Minimum number of copies to flag (default: 2) |
| language | string | no | Language filter: go, typescript, python, javascript (default: all) |

## Execution Pipeline

### Step 1: Scan
Use `glob` to collect source files within `scope`. For each file, use `view` to read the content. Extract function bodies and structural blocks. Compare blocks using normalized form (stripped whitespace, collapsed variable names) to find duplicates. Use `grep` with common patterns to identify near-clones (same logic, different variable names).

### Step 2: Cluster
Group detected clones into families. Within each family, identify the canonical (earliest or most complete) version and all copies. Calculate the duplication metric: lines duplicated / total lines scanned.

### Step 3: Suggest
For each clone family, recommend a refactoring strategy: extract shared function, use generics/templates, or consolidate into a utility module. Estimate the line savings.

## Output Shape
```json
{
  "total_lines_scanned": 12400,
  "duplication_rate": 0.08,
  "clone_families": [
    {
      "canonical": "string — file:line of the original",
      "copies": ["string — file:line locations of duplicates"],
      "lines_per_clone": 12,
      "similarity": 0.95,
      "recommendation": "string — suggested refactoring action"
    }
  ],
  "total_clones_found": 7
}
```

## Failure Modes
- IF the codebase is too large to scan in one pass (>10k files): split into directory-based batches and report per-directory
- IF normalization strips too much meaning (false positives): increase `min_lines` and `min_copies` thresholds

## Examples
```json
{
  "action": "scan",
  "scope": ["internal/engine/", "internal/parser/"],
  "min_lines": 8,
  "language": "go"
}
```
