---
name: dependency-hologram
description: Generate interactive 2D/3D dependency visualizations with architectural insights
category: core
version: "2.0.0"
---

# Dependency Hologram

> Generate interactive 2D/3D dependency visualizations with architectural insights

## When to Use
- WHEN mode=BUILD and the task requires dependency hologram
- WHEN action=`generate`: produce new artifacts from inputs
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`compare`: diff two inputs and report differences

## Actions
`'generate' | 'analyze' | 'compare'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| project_path | string | yes | Project path |
| visualization_type | enum | no | Visualization type |
| include_dev_deps | boolean | no | Include dev deps |
| depth | number | no | Depth |

## Invocation
```
floyd skill:dependency-hologram --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
