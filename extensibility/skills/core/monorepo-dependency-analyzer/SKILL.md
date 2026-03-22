---
name: monorepo-dependency-analyzer
description: Analyze dependencies across monorepo packages with cycle detection
category: core
version: "2.0.0"
---

# Monorepo Dependency Analyzer

> Analyze dependencies across monorepo packages with cycle detection

## When to Use
- WHEN mode=BUILD and the task requires monorepo dependency analyzer
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`visualize`: generate visualization
- WHEN action=`find_cycles`: detect circular dependencies

## Actions
`'analyze' | 'visualize' | 'find_cycles'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| repo_path | string | yes | Repo path |
| packages_path | string | no | Packages path |
| include_dev_deps | boolean | no | Include dev deps |
| max_depth | number | no | Max depth |

## Invocation
```
floyd skill:monorepo-dependency-analyzer --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
