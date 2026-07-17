# Floyd STABILITY mode

`STABILITY <task>` asks Floyd to harden a concrete component for production use.

## Behavior

When STABILITY mode is active, Floyd:

1. Inspects the named code path and records its observable failure modes.
2. Applies concrete safeguards such as validation, bounded retries, timeouts, cancellation, and resource cleanup where relevant.
3. Adds focused regression tests for every mitigation.
4. Reports residual risks that remain outside the stated scope.

The mode does not override model sampling or provider settings. Those remain controlled by Floyd's selected model and provider configuration.

## Example

```text
STABILITY harden the SQLite session open path against cancellation and partial initialization
```

## Implementation

The system-prompt overlay is defined in `internal/agent/templates/floyd-general.md.tpl`. Runtime model options remain independently resolved in `internal/agent/coordinator.go`.
