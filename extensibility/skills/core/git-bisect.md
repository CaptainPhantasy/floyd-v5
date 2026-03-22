---
name: git-bisect
description: Automated git bisect with test command execution to identify the exact commit that introduced a bug.
---

# Git Bisect

> Automated git bisect with test command execution to identify the exact commit that introduced a bug.

**Category**: General

## When to Use
- When analyzing failures or debugging issues

## Key Inputs
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| action | string | yes | action parameter |
| repo_path | string | yes | repo_path parameter |
| start_commit | string | yes | start_commit parameter |
| end_commit | string | yes | end_commit parameter |
| test_command | string | yes | test_command parameter |

## Invocation
Use via Floyd Labs MCP:
- **Tool**: `floyd`
- **Action**: `execute`
- **Skill**: `git-bisect`
- **Args**: `{key inputs as JSON object}`

## Output
Returns structured output with results and metadata.