#!/bin/bash
# Floyd Build Script
# Usage: ./scripts/build.sh [--clean]

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
FLOYD_ROOT="$(dirname "$SCRIPT_DIR")"
cd "$FLOYD_ROOT"

# Read version from VERSION file
VERSION=$(cat VERSION 2>/dev/null || echo "0.0.0-dev")
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")

LDFLAGS="-X 'internal/version.Version=v${VERSION}' \
         -X 'internal/version.BuildDate=${BUILD_DATE}' \
         -X 'internal/version.GitCommit=${GIT_COMMIT}'"

echo "=========================================="
echo "Building Floyd v${VERSION}"
echo "Commit: ${GIT_COMMIT}"
echo "Date: ${BUILD_DATE}"
echo "=========================================="

# Clean if requested
if [ "$1" == "--clean" ]; then
    echo "Cleaning previous builds..."
    rm -f floyd superfloyd
fi

# Build floyd (general enterprise)
echo ""
echo "Building floyd (general enterprise)..."
go build -ldflags "${LDFLAGS} -X 'internal/version.BinaryName=floyd'" -o floyd .

# Build superfloyd (coding-only)
echo "Building superfloyd (coding-only)..."
go build -ldflags "${LDFLAGS} -X 'internal/version.BinaryName=superfloyd'" -o superfloyd .

echo ""
echo "=========================================="
echo "Build complete!"
echo "  - floyd       (general enterprise)"
echo "  - superfloyd  (coding-only)"
echo "=========================================="

# Show versions
echo ""
echo "Verifying versions:"
./floyd --version 2>/dev/null || echo "  floyd: (version check failed)"
./superfloyd --version 2>/dev/null || echo "  superfloyd: (version check failed)"
