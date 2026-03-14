# Migration Guide: Floyd v4 to v5.0.1

This document outlines the major architectural changes and required steps for migrating from legacy Floyd/superfloyd to the v5.0.1 Production Release.

---

## 1. Major Changes

### Twin-Engine Binary Architecture
Previously, `superfloyd` was often a separate compilation or a simple alias. In v5.0.1, the engine uses **Deterministic Binary Detection**:
*   The same binary acts differently based on its filename (`floyd` vs `superfloyd`).
*   `superfloyd` automatically enables **SOTA Architect** mode, optimized for complex coding.
*   `floyd` remains the high-velocity multipurpose operational agent.

### Multiplexed Persistent Terminal
The terminal has been moved from a hidden utility to a core workspace feature:
*   **Persistence**: Sessions are no longer killed when the UI panel is closed.
*   **Discovery**: Selecting "Open Terminal" now automatically transitions the UI to Chat mode.

### Unified Extensibility
Agents and Skills are now categorized and searchable in a unified Command Palette.

---

## 2. Migration Steps

### Step 1: Binary Overwrite
Ensure you are using the new ad-hoc signed binaries to avoid macOS Code Integrity Guard (CIG) issues.
```bash
cp -f floyd /opt/homebrew/bin/floyd
cp -f superfloyd /opt/homebrew/bin/superfloyd
```

### Step 2: Agent Category Update
If you have custom agents in `~/.config/floyd/agents/`, you must add a `category` field to their front matter for them to appear in the new UI.
**Valid Categories:** `architecture`, `infrastructure`, `orchestration`, `coding`, `security`, `quality`, `testing`, `monitoring`, `documentation`, `data`, `performance`, `dx`, `debugging`.

### Step 3: Skill Porting
Move any legacy instruction files to the new folder structure:
`~/.config/floyd/skills/{category}/{skill-name}/SKILL.md`

---

## 3. Configuration Updates (`floyd.json`)

The `mcp` configuration is now live-synced with the UI. Ensure your `floyd.json` matches the latest schema to see tool counts in the Plugins Library.

```json
"mcp": {
  "my-plugin": {
    "type": "stdio",
    "command": "node",
    "args": ["path/to/index.js"]
  }
}
```

---

## 4. Troubleshooting

*   **Initialization Hangs**: v5.0.1 fixes the recursive filesystem walk. If you still see delays, ensure you aren't running in a directory with >10,000 files.
*   **Command Palette Size**: If the UI looks cramped, ensure your terminal supports at least 100 characters of width.
