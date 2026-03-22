---
name: git-bisect
description: Automated git bisect to identify commits that introduced bugs
category: core
version: "2.0.0"
---

# Git Bisect

> Automated git bisect to identify commits that introduced bugs

## When to Use
- WHEN mode=DEBUG and the task requires git bisect
- WHEN action=`start`: begin the operation
- WHEN action=`next`: advance to next step
- WHEN action=`reset`: reset state
- WHEN action=`log`: retrieve operation log

## Actions
`'start' | 'next' | 'reset' | 'log'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | enum | yes | Action |
| repo_path | string | yes | Repo path |
| start_commit | string | yes | Start commit |
| end_commit | string | yes | End commit |
| test_command | string | yes | Test command |
| max_steps | number | no | Max steps |

## Invocation
```
floyd skill:git-bisect --action <action> [--input <json>]
```

## Output
Returns structured JSON with results, metadata, and execution status.
