---
name: schema-migrator
description: Generate database schema migrations with validation and rollback plans
category: core
version: "2.0.0"
---

# Schema Migrator

> Generate database schema migrations with validation and rollback plans

## When to Use
- WHEN mode=BUILD and the task requires schema migrator
- WHEN action=`migrate`: transform schema
- WHEN action=`validate`: check conformance and report violations
- WHEN action=`generate_diff`: produce diff between schemas

## Actions
`'migrate' | 'validate' | 'generate_diff'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| source_schema | object | yes | Source schema |
| target_schema | object | yes | Target schema |
| migration_path | string | no | Migration path |
| validate_only | boolean | no | Validate only |

## Invocation
```
floyd skill:schema-migrator --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
