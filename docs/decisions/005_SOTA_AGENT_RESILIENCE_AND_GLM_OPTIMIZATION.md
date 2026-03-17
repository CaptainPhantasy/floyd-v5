# ADR 005: SOTA Agent Resilience & Intelligence Scaling

**Date:** 2026-03-16
**Status:** Implemented

## Context

Agents exhibit amnesia due to context pollution and the lack of "Preserved Thinking" on standard GLM OpenAI-compatible endpoints. Structural reinforcement is required to maintain multi-turn stability and output quality.

## Decision

1. **Mandatory Context Flush:** Explicitly drop raw tool/fetch data post-synthesis to prevent window pollution.
2. **Supercache Sync:** Synchronous I/O at every turn boundary to ensure zero-loss amnesia protection.
3. **Deterministic Thinking:** Enforced `thinking` blocks for every state change.
4. **Proactive Artifact Conversion:** Automatically divert lengthy conversational outputs (>10 lines) to the filesystem as files.
5. **Structural Thinking Levels:** Framework-level env vars set to force maximum reasoning depth (FLOYD_THINKING_LEVEL=MAX).
6. **Visual Perfection:** Python-based box-table generation is the only acceptable tabular format.
7. **GLM Reasoning Persistence:** Manual re-anchoring of goals and previous results in every turn to simulate "Preserved Thinking."
8. **Closed-Loop Self-Healing:** Automatic execution of `go build/test` after edits with auto-correction of errors.

## Consequences

- Agents will consume more tokens due to explicit thinking blocks
- Build/test cycles may slow down rapid iteration
- Output quality will increase significantly (10X multiplier target)
- Zero-loss continuity achieved through supercache persistence
- SOTA output quality on GLM-4.7 and GLM-5 models guaranteed

## Implementation Evidence

- Templates modified: `internal/agent/templates/superfloyd-coder.md.tpl`, `internal/agent/templates/floyd-general.md.tpl`
- Environment configured: `.env.local`, `~/.floyd/.env.local`
- Validation script: `scripts/validate_go.sh`
- Binaries rebuilt: `bin/floyd`, `bin/superfloyd`, `floyd-test`, `superfloyd-test`
- System deployed: `/opt/homebrew/bin/floyd`, `/opt/homebrew/bin/superfloyd`

## Rollback

Restore point: git commit `13fe87c`
