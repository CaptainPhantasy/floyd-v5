---
name: secure-hook-executor
description: Execute Git hooks in sandboxed environment with timeout and resource limits
category: core
version: "2.0.0"
---

# Secure Hook Executor

> Execute Git hooks in sandboxed environment with timeout and resource limits

## When to Use
- WHEN mode=BUILD and the task requires secure hook executor
- WHEN action=`execute`: perform the operation
- WHEN action=`validate`: check conformance and report violations
- WHEN action=`dry_run`: simulate without side effects

## Actions
`'execute' | 'validate' | 'dry_run'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| hook_type | enum | yes | Hook type |
| hook_script | string | yes | Hook script |
| context | object | no | Context |
| sandbox | boolean | no | Sandbox |
| timeout | number | no | Timeout |

## Invocation
```
floyd skill:secure-hook-executor --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
