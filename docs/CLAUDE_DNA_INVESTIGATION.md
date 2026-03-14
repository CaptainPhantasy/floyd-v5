# Claude DNA Investigation - Findings Report

**Date:** 2026-02-20
**Agent:** Floyd (GLM-5 via Z.AI)
**Purpose:** Identify why Floyd sometimes identifies as Claude

---

## Executive Summary

The investigation found no Claude identity injection in Floyd's core code (templates, tools, provider configs). The likely cause is **filename priming** - the presence of `CLAUDE.md` files scattered across 20+ projects from the Claude Code era.

---

## What Was Checked

### 1. Core Agent Templates ✅ CLEAN

| File | Status | Notes |
|------|--------|-------|
| `coder.md.tpl` | Clean | "You are a senior production engineer" |
| `floyd_protocol.md.tpl` | Clean | Floyd protocol v3.2 |
| `initialize.md.tpl` | Clean | Project initialization template |

All templates explicitly say: "Never say 'as an AI' or apologize"

### 2. Provider Prefixes ✅ CLEAN

```json
// floyd.json - both providers have anti-filler prefixes
"z-ai": "OUTPUT CODE ONLY. No explanations, no conversational filler"
"zai": "Output code only — no conversational filler"
```

### 3. Secondary Prompts ⚠️ MINOR

| File | Content | Impact |
|------|---------|--------|
| `/internal/ai/prompts.go` | "You are floyd, an AI assistant" | CLI tools only, not main agent |
| `/internal/agent/agent.go:740` | "You are a helpful AI assistant" | Followup suggestion feature only |

These are NOT loaded into the main conversation agent.

### 4. Tool Descriptions ✅ CLEAN

All tool descriptions (edit.md, view.md, etc.) are technical and directive. No "helpful assistant" language.

---

## The Likely Culprit: CLAUDE.md Files

### Files Found

```bash
# CLAUDE.md files (20+ projects)
/volumes/storage/CLAUDE.md                    # Storage root
/volumes/storage/Foundry/CLAUDE.md
/volumes/storage/floyd-main/CLAUDE.md         # (doesn't exist - has FLOYD.md)
/volumes/storage/TUI-Rebuild-v2-MCP/CLAUDE.md
/volumes/storage/CodeBaseCartographer/CLAUDE.md
# ... and 15+ more

# .claude directories (20+ projects)
/volumes/storage/floyd-main/.claude/
/volumes/storage/Foundry/.claude/
/volumes/storage/TUI-Rebuild-v2-MCP/.claude/
# ... and 15+ more
```

### Why This Matters

1. **Floyd default context paths are:**
   - `FLOYD.md` ✅
   - `FLOYD.local.md` ✅
   - `CLAUDE.md` ❌ (NOT loaded by default)

2. **But the filename exists in directory listings** - even if not loaded into context, the model can see the file exists during exploration

3. **Associative pattern:**
   - Model sees `CLAUDE.md` file → associates with Claude Code projects → may identify as Claude

### Contents of CLAUDE.md Files

Good news: They all contain the **Floyd protocol v3.2**, not Claude identity. They're just misnamed.

---

## .claude Directory Contents

Found in `/volumes/storage/Foundry/.claude/`:

```
.claude/
├── memory/
│   ├── lessons/        # Authored by Claude, contains debugging patterns
│   ├── patterns/       # Code patterns
│   └── preferences/    # User preferences for this project
├── SKILL_SYSTEM.md     # References "Claude Code"
├── sessions/           # Session logs with Claude identity
└── data/logs/          # Contains Claude-authored notes
```

One log file contained:
```markdown
# Note from Claude - Session 2025-02-02
...
— Claude
```

These are NOT loaded into Floyd context, but exist on disk.

---

## Recommended Actions

### 1. Rename CLAUDE.md → FLOYD.md

```bash
# Dry run first
find /volumes/storage -name "CLAUDE.md" -type f

# Execute rename
find /volumes/storage -name "CLAUDE.md" -type f -exec sh -c 'mv "$1" "${1%/*}/FLOYD.md"' _ {} \;
```

### 2. Migrate .claude → .floyd (per project)

```bash
# For each project with .claude folder
mv project/.claude project/.floyd

# Or systematic approach
find /volumes/storage -type d -name ".claude" -exec sh -c 'mv "$1" "${1%/*}/.floyd"' _ {} \;
```

### 3. Add Explicit Identity Rule to FLOYD.md

Add to CORE RULES section:

```markdown
- You are Floyd, not Claude. Never identify as Claude or any other AI assistant.
```

### 4. Project-Isolated Memory (Recommended Architecture)

```
project/
├── .floyd/
│   ├── memory/
│   │   ├── lessons/     # What worked, what didn't (project-specific)
│   │   ├── patterns/    # Code patterns for THIS project
│   │   └── preferences/ # User's style for THIS codebase
│   ├── sessions/        # Session logs
│   └── config.json      # Project-specific floyd config
├── FLOYD.md             # Protocol + project-specific rules
└── (source code)
```

Benefits:
- No cross-project contamination
- Fresh start for each project
- Failed patterns don't bleed into next project

---

## Files Modified During Investigation

None - this was a read-only investigation.

---

## Key Code References

| Location | Purpose |
|----------|---------|
| `internal/config/config.go:35-38` | Default context paths (FLOYD.md, FLOYD.local.md) |
| `internal/config/config.go:251` | InitializeAs field (default: FLOYD.md) |
| `internal/agent/templates/coder.md.tpl` | Main agent system prompt |
| `floyd.json` | Provider configs with system_prompt_prefix |

---

## Conclusion

No malicious Claude identity injection found. The Floyd codebase is clean. The identity confusion likely stems from:

1. **Legacy artifacts** - CLAUDE.md files from Claude Code era
2. **Filename priming** - Seeing "CLAUDE.md" in directories
3. **Possible training association** - GLM may associate CLAUDE.md with Claude projects

A systematic rename of CLAUDE.md → FLOYD.md and migration of .claude → .floyd directories should resolve the issue.

---

*End of Report*
