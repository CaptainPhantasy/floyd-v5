# Investigation Report: 70,000+ Token Startup Bloat & Database Migration Error

## 1. The 70,000+ Token Initialization Bloat

### The Symptoms
When starting a new session, the Floyd/SuperFloyd agent consumes roughly 70,000 to 90,000 tokens on the very first turn (initialization), even for a simple "Say hi" prompt. 

### What We Found
The token bloat is **not** coming from the system prompt text itself. We successfully truncated the XML tool descriptions inside the system prompt (`mcp.ToPromptXML()`), but the token usage remained massively high.

The true source of the token bloat is the **JSON Tool Schemas** being sent to the LLM API. 
When Floyd initializes, it connects to every MCP server defined in the configuration file (`~/.config/floyd/floyd.json` and local `floyd.json`). If servers like `pattern-crystallizer-v2`, `hivemind-v2`, `omega-v2`, and others are enabled, Floyd fetches the JSON schema for every single tool (over 200 tools total) and attaches them to the `tools` array of the API request.

The API provider (GLM-5 / Z.AI / Anthropic) charges tokens for the entire JSON schema of every tool provided in the request. Passing 200+ complex tool schemas on every single turn costs ~70,000 tokens.

### Why Previous Fixes Failed
1. We modified `mcp.ToPromptXML()` to shrink the text prompt, but the underlying `tools` array sent to the API was untouched.
2. We attempted to disable the heavy MCP servers by setting `"disabled": true` in `~/.config/floyd/floyd.json`, but Floyd might still be loading them if the config parser doesn't properly respect the `"disabled"` flag for MCP servers, or if it's reading a different config file.

### Recommended Next Steps for Token Bloat
1. **Remove Heavy MCP Servers from Config:** Instead of relying on a `"disabled"` flag, physically remove the heavy cognitive MCP servers from the `mcp` block in `~/.config/floyd/floyd.json`. 
2. **Rely on Skills:** The 73 "cognitive" tools have already been downloaded as Markdown skills to `~/.config/floyd/skills/`. Floyd can read these skills on-demand without paying the upfront token cost of loading them as API tools.
3. **Keep Only Core Execution Servers:** Only keep `floyd-lab`, `floyd-runner`, `floyd-git`, `floyd-supercache`, and `floyd-explorer` in the MCP config.

---

## 2. The Database Migration Crash (`cache_read_tokens`)

### The Symptoms
Running `floyd -y` or `superfloyd -y` in a new directory with no existing database results in a crash:
`Failed to apply migrations: ERROR 20260314200000_add_cache_read_tokens.sql: failed to run SQL migration: failed to execute SQL query "ALTER TABLE sessions ADD COLUMN cache_read_tokens INTEGER NOT NULL DEFAULT 0;": SQL logic error: duplicate column name: cache_read_tokens (1).`

### What We Found
There is a race condition/conflict between Floyd's manual schema generation and the `goose` migration engine.
1. When Floyd boots with a missing database, it runs `ensureColumns()` in `internal/db/connect.go`. We removed `cache_read_tokens` from `backfills` in `ensureColumns()`.
2. However, the initial database schema in `internal/db/migrations/20250424200609_initial.sql` might not have `cache_read_tokens`.
3. When `goose` runs `20260314200000_add_cache_read_tokens.sql`, it crashes with "duplicate column". This implies the column *was* somehow created before the migration ran, or there are multiple migration files trying to create it, or `ensureColumns()` was not properly purged from the compiled binary due to Go build caching.

### Why Previous Fixes Failed
1. We edited `connect.go` to remove the backfill, and we manipulated the SQL migration files. 
2. We ran `go generate ./internal/db/...` to refresh the `//go:embed` cache.
3. Despite this, the compiled binaries in `/opt/homebrew/bin/` are still experiencing the collision. This means either the Go embed cache wasn't truly cleared, or another piece of code (like a hardcoded struct or another migration file) is pre-creating the `cache_read_tokens` column before Goose reaches the specific migration file.

### Recommended Next Steps for DB Crash
1. **Audit Initial Schema:** Check exactly what `20250424200609_initial.sql` contains in the final source tree.
2. **Check SQLite State Pre-Migration:** The next model should run the app in debug mode and inspect the SQLite schema *exactly* before `goose.Up()` is called in `connect.go`. 
3. **Consolidate Migrations:** To permanently fix this, delete the `20260314...` migration file entirely, and just add the `cache_read_tokens` column directly into the `CREATE TABLE sessions` block of `20250424200609_initial.sql`. Then clear the `goose_db_version` table so it rebuilds cleanly from scratch.