# FLOYD v5.3.0 - Quick Start & Feature Guide

Welcome to **FLOYD v5.3.0** and **SUPERFLOYD v5.3.0**! This guide is designed to be easily ingestible by both human developers and LLM agents to rapidly understand the state-of-the-art (SOTA) enterprise upgrades.

## Core Identity Schism: Floyd vs. SuperFloyd

Version 5.3.0 introduces a hard split in reasoning architecture:
- **FLOYD**: General-purpose operational agent. Optimized for rapid project analysis, state management, and file traversal.
- **SUPERFLOYD**: SOTA coding specialist, full-stack architect, and UI expert. Handles deep, surgical edits, and complex AST-aware transformations.

## Side-by-Side: Former vs. Upgraded Capabilities

| Capability | v5.0 / Former | v5.3.0 (New!) |
| :--- | :--- | :--- |
| **Tool Scope** | Basic bash, edit, file view | 6 new highly-specialized tools mapped to roles |
| **Logic Reasoning** | Contradictory instructions on `<think>` | Mandatory `<think>` logic blocks via deterministic prompt |
| **UI Management** | Basic Chat | Enterprise Config Audit & Plugins Library via Commands Palette |
| **Prompt Bloat** | Redundant SOTA blocks per agent | Deduplicated globally via `.tpl` System Prompts |
| **Web Search** | Disconnected | Native `web_search` enabled globally |
| **Output Bounds** | Truncated stdout / lost code | LLM intercepts large text limits, pill notifications via UI |
| **Autonomic UX** | Terminal spews errors during load | *Auto-Stabilize Diagnostic Bus* natively warns UI via Bubbletea |
| **SuperFloyd Editing** | Line-based replacement only | Tree-aware `apply_patch` & AST tokens with `smart_replace` |
| **Floyd Context** | Cat entire 10k-line files | `list_symbols` & `project_map` to conserve tokens |

## SOTA Tool Usage (For LLM Intelligence)

### FLOYD (Operational Tools)
1. `project_map`: Generates a massive repository tree instantly, explicitly stripping `node_modules` and hidden cruft.
2. `manage_scratchpad`: Persists cross-session thoughts in `.floyd/scratchpad.md`.
3. `list_symbols`: Extracts `func`, `type`, `interface`, and class signatures natively via `grep` boundaries rather than reading full source files.

### SUPERFLOYD (Architect Tools)
1. `apply_patch`: Consumes a standard `a/b` unified diff string to safely migrate 10s of files simultaneously via git patch protocol.
2. `smart_replace`: A deterministic white-space resilient block-changer that outperforms raw `edit` targets.
3. `get_active_diff`: Runs `git diff HEAD` --staged to dynamically capture ongoing states without manual tracking.

## Enterprise UI Upgrades
Hit `cmd+j` / `ctrl+j` to access the Command Palette.
- New **Plugins Library**: Direct interface to Model Context Protocol parameters.
- New **Config Audit**: Evaluates connection, security bounds, and state integrity safely.

