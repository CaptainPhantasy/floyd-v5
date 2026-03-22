# Personal Enterprise Quality Mode (Local-First)

This repository is configured for **personal environment quality** first.

## Why GitHub Actions existed
GitHub-triggered automation came from workflow files under `.github/workflows/`.
Those files can run on push/PR/schedule unless restricted.

## Current mode
All repository workflows are now **manual-only** (`workflow_dispatch`).
This prevents automatic GitHub runs and most CI email noise.

## Local strict quality gate
Run this from repo root:

```bash
task qa:personal
```

This gate performs:
1. Formatting check on changed Go files only (strict for your active delta)
2. Linting on changed Go files only (strict for your active delta)
3. Deterministic package tests for your core runtime paths
4. `go vet ./...`
5. Race tests for core runtime paths
6. Local binary build + smoke test for `floyd_56` and `superfloyd_56`

## Optional GitHub-side cleanup
To fully silence Actions noise for this repo:
1. GitHub repo → Settings → Actions → General
2. Set actions permissions as desired or disable Actions for the repo
3. GitHub Settings → Notifications → disable Actions emails for this repo

## Philosophy
Quality requirements should match your environment and goals.
Public-readiness gates should only be re-enabled when you decide to distribute publicly.
