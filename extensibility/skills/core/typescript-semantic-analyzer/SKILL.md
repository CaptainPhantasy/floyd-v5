---
name: typescript-semantic-analyzer
description: Deep semantic analysis of TypeScript code with type extraction and reference finding
category: core
version: "2.0.0"
---

# Typescript Semantic Analyzer

> Deep semantic analysis of TypeScript code with type extraction and reference finding

## When to Use
- WHEN mode=BUILD and the task requires typescript semantic analyzer
- WHEN action=`analyze`: inspect and report without modification
- WHEN action=`extract_types`: extract type definitions
- WHEN action=`find_references`: find symbol references

## Actions
`'analyze' | 'extract_types' | 'find_references'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| file_path | string | yes | File path |
| project_path | string | no | Project path |
| include_types | boolean | no | Include types |
| depth | number | no | Depth |

## Invocation
```
floyd skill:typescript-semantic-analyzer --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
