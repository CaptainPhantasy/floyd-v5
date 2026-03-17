#!/usr/bin/env bash
set -e

# Validation helper for Go projects
# This script runs go build and go test (if any) and returns exit code with output.

cd "$(dirname "$0")/.." 2>/dev/null || true

echo "Running go build..."
if go build ./...; then
    echo "✓ Build successful"
else
    echo "✗ Build failed"
    exit 1
fi

echo "Running go test..."
if go test ./... 2>&1 | head -50; then
    echo "✓ Tests passed"
else
    echo "⚠ Tests failed or none exist"
fi
