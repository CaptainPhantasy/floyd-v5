---
name: concept-lattice
description: Map hidden architectural relationships and abstractions into navigable lattice structures for knowledge discovery
category: core
version: "2.0.0"
---

# Concept Lattice

> Map hidden architectural relationships and abstractions into navigable lattice structures for knowledge discovery.

## When to Use
- WHEN `mode=EXPLORE` to understand the implicit architecture and abstraction hierarchy of a codebase
- WHEN `mode=BUILD` to plan a refactoring by identifying shared abstractions that should be extracted
- WHEN onboarding to a large codebase and needing a conceptual map of modules and their relationships

## Actions
`'build' | 'query' | 'visualize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| scope | string[] | no | Directories or modules to include (default: entire project) |
| depth | integer | no | Maximum lattice depth to traverse (default: 5) |
| focus_concept | string | no | Starting concept to build the lattice around (default: auto-detect root concepts) |

## Execution Pipeline

### Step 1: Extract Concepts
Use `list_symbols` on all files in `scope` to extract type names, interface names, function names, and module names. Use `grep` to find implementation relationships (e.g., `implements`, `extends`, `interface{}`) and usage relationships (function calls, type references).

### Step 2: Build Lattice
Construct a formal concept lattice where:
- **Nodes** are concepts (types, interfaces, modules)
- **Edges** represent "is-a", "has-a", "uses", or "implements" relationships
- **Levels** represent abstraction depth (interfaces at top, concrete types at bottom)

Use `manage_scratchpad` to store the lattice as an adjacency map.

### Step 3: Query and Visualize
If `focus_concept` is provided, extract the sub-lattice rooted at that concept. Produce an ASCII tree visualization showing the hierarchy. Identify concept clusters (groups of tightly related types) and isolation (types with few relationships).

## Output Shape
```json
{
  "total_concepts": 87,
  "total_relationships": 234,
  "lattice_depth": 5,
  "root_concepts": ["string — top-level abstractions"],
  "clusters": [
    {
      "name": "string — cluster label",
      "members": ["string — concept names"],
      "coupling_score": 0.85
    }
  ],
  "isolated_concepts": ["string — concepts with fewer than 2 relationships"],
  "visualization": "string — ASCII lattice diagram"
}
```

## Failure Modes
- IF the scope contains no extractable symbols: report that the files are non-structural (e.g., configs, data files) and lattice analysis is not applicable
- IF the lattice is too deep (>10 levels): truncate at `depth` and report the truncation

## Examples
```json
{
  "action": "build",
  "scope": ["internal/engine/", "internal/types/"],
  "focus_concept": "Planner",
  "depth": 4
}
```
