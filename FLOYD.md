# Floyd — General Purpose Operational Agent

> v5.3.0 | Go | MCP | TUI

## Quick Reference

| Item | Detail |
|---|---|
| Language | Go 1.23+ |
| Entry | `main.go` → `internal/cmd/root.go` |
| TUI | Bubble Tea (`charmbracelet/bubbletea`) |
| MCP Servers | `floyd-supercache`, `floyd-terminal` |
| Config | `floyd.json`, `floyd-schema.json` |
| DB | SQLite via `modernc.org/ncruces` |
| Build | `task build` or `scripts/build.sh` |
| Test | `go test ./...` |

## Architecture

```
internal/
├── agent/       → Core agent loop, tool dispatch, Ralph Loop
├── ai/          → LLM client abstraction
├── app/         → Application bootstrap, LSP bridge
├── cmd/         → CLI commands (cobra)
├── config/      → Config loading, providers, resolution
├── db/          → SQLite queries (sqlc generated)
├── skills/      → Runtime skill loader
├── ui/          → TUI views (chat, diff, completions)
└── ...          → Shell, telemetry, permissions, etc.
```

## Conventions

- All edits MUST pass `go build` + `go test ./...`
- Use `gofumpt` formatting
- Error types in `internal/errors/`
- Events via `internal/event/` pub/sub
- Config resolved through `internal/config/resolve.go`

## Recent Work (v5.3.0)

- Production hardening & protocol kernel
- Ralph Loop (self-correction cycle)
- Skill restructure (`extensibility/skills/`)
- Agent tool dispatch rewrite (`agent_tool.go`)
- Test harness modernization (VCR → layered)
