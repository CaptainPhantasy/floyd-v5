#!/bin/bash
# Floyd Release Script
# Usage: ./scripts/release.sh [--dry-run]

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
FLOYD_ROOT="$(dirname "$SCRIPT_DIR")"
cd "$FLOYD_ROOT"

DRY_RUN=false
if [ "$1" == "--dry-run" ]; then
    DRY_RUN=true
    echo "=== DRY RUN MODE ==="
fi

# Read version
VERSION=$(cat VERSION 2>/dev/null || echo "0.0.0")
RELEASE_DIR="releases/v${VERSION}"
BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
GIT_COMMIT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")

echo "=========================================="
echo "Creating Release v${VERSION}"
echo "=========================================="
echo "Commit: ${GIT_COMMIT}"
echo "Branch: ${GIT_BRANCH}"
echo "Date: ${BUILD_DATE}"
echo "=========================================="

# Pre-flight checks
echo ""
echo "Pre-flight checks:"

# Check for uncommitted changes
if ! git diff-index --quiet HEAD -- 2>/dev/null; then
    echo "  ⚠️  WARNING: Uncommitted changes detected"
    read -p "  Continue anyway? [y/N] " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Aborted."
        exit 1
    fi
else
    echo "  ✓ No uncommitted changes"
fi

# Check if release already exists
if [ -d "$RELEASE_DIR" ]; then
    echo "  ⚠️  WARNING: Release directory already exists: $RELEASE_DIR"
    read -p "  Overwrite? [y/N] " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Aborted."
        exit 1
    fi
    rm -rf "$RELEASE_DIR"
fi

# Check if release notes exist
RELEASE_NOTES="docs/releases/v${VERSION}.md"
if [ ! -f "$RELEASE_NOTES" ]; then
    echo "  ⚠️  WARNING: Release notes not found: $RELEASE_NOTES"
    echo "  Creating template..."
    cp docs/templates/RELEASE_NOTES.tmpl.md "$RELEASE_NOTES" 2>/dev/null || true
fi

echo "  ✓ Pre-flight checks passed"

# Build
echo ""
echo "Building binaries..."
./scripts/build.sh --clean

# Create release directory
echo ""
echo "Creating release directory: $RELEASE_DIR"
mkdir -p "$RELEASE_DIR"

# Copy binaries
echo "Copying binaries..."
cp floyd "$RELEASE_DIR/"
cp superfloyd "$RELEASE_DIR/"

# Generate checksums
echo "Generating checksums..."
cd "$RELEASE_DIR"
if command -v shasum &> /dev/null; then
    shasum -a 256 floyd superfloyd > checksums.txt
elif command -v sha256sum &> /dev/null; then
    sha256sum floyd superfloyd > checksums.txt
else
    echo "  ⚠️  WARNING: No SHA256 tool found, skipping checksums"
fi
cd "$FLOYD_ROOT"

# Get file sizes
FLOYD_SIZE=$(ls -lh floyd | awk '{print $5}')
SUPERFLOYD_SIZE=$(ls -lh superfloyd | awk '{print $5}')

# Create manifest
echo "Creating manifest..."
cat > "$RELEASE_DIR/manifest.json" << EOF
{
  "version": "${VERSION}",
  "date": "${BUILD_DATE}",
  "commit": "${GIT_COMMIT}",
  "branch": "${GIT_BRANCH}",
  "binaries": {
    "floyd": {
      "file": "floyd",
      "size": "${FLOYD_SIZE}",
      "description": "General enterprise agent (research, docs, web, code)"
    },
    "superfloyd": {
      "file": "superfloyd",
      "size": "${SUPERFLOYD_SIZE}",
      "description": "Strict coding agent (code only, no web/research)"
    }
  },
  "checksums": "checksums.txt",
  "changelog": "docs/releases/v${VERSION}.md",
  "compatibility": {
    "go_version": "$(go version | awk '{print $3}')",
    "platform": "$(uname -s)-$(uname -m)"
  }
}
EOF

# Summary
echo ""
echo "=========================================="
echo "Release v${VERSION} created!"
echo "=========================================="
echo ""
echo "Directory: $RELEASE_DIR"
echo ""
ls -la "$RELEASE_DIR"
echo ""
echo "Checksums:"
cat "$RELEASE_DIR/checksums.txt"
echo ""
echo "Next steps:"
echo "  1. Review/update release notes: $RELEASE_NOTES"
echo "  2. Test binaries in $RELEASE_DIR"
echo "  3. Tag release: git tag -a v${VERSION} -m 'Release v${VERSION}'"
echo "  4. Push tag: git push origin v${VERSION}"

if [ "$DRY_RUN" = true ]; then
    echo ""
    echo "=== DRY RUN COMPLETE ==="
    echo "No permanent changes made."
fi
