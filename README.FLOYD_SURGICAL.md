# Floyd SURGICAL mode

`SURGICAL <task>` asks Floyd to use a smallest-safe-diff workflow for a high-risk change.

## Behavior

When SURGICAL mode is active, Floyd:

1. Records the observable pre-change behavior.
2. Limits edits to the stated symptom boundary.
3. Runs focused verification after the patch.
4. Reports the changed files, remaining risk, and a normal Git revert command.
5. Stops for renewed scope approval if the required fix expands beyond the stated boundary.

The mode changes execution discipline, not provider sampling parameters. Model temperature, reasoning effort, and provider options continue to come from the selected Floyd configuration.

## Examples

```text
SURGICAL fix the nil-pointer in internal/agent/agent.go:274
SURGICAL isolate the pprof server from the default HTTP mux
```

## Implementation

The system-prompt overlay is defined in `internal/agent/templates/floyd-general.md.tpl`. Provider options are resolved independently in `internal/agent/coordinator.go`.

## Verification

Run focused tests for the affected package, followed by the project build and static analysis:

```bash
go test . -run TestProfileHandlerIsIsolated -count=1
go test ./internal/agent -run 'TestAnthropicReasoningEffortMapsToThinkingBudget|TestHardening' -count=1
go vet ./...
go build ./...
```

To roll back a published SURGICAL-mode change, revert its commit with `git revert <commit>` and publish that new revert commit through the normal governed workflow.
