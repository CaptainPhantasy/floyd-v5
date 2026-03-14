#!/bin/bash
# Floyd Release Validation Script
# Usage: ./scripts/validate-release.sh [version]

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
FLOYD_ROOT="$(dirname "$SCRIPT_DIR")"
cd "$FLOYD_ROOT"

VERSION=${1:-$(cat VERSION 2>/dev/null || echo "unknown")}
RELEASE_DIR="releases/v${VERSION}"

echo "=========================================="
echo "Validating Release v${VERSION}"
echo "=========================================="
echo ""

ERRORS=0
WARNINGS=0

# Check 1: VERSION file exists
echo "1. Checking VERSION file..."
if [ -f "VERSION" ]; then
    VFILE=$(cat VERSION)
    if [ "$VFILE" == "$VERSION" ]; then
        echo "   ✓ VERSION file matches: $VERSION"
    else
        echo "   ✗ VERSION file mismatch: $VFILE != $VERSION"
        ((ERRORS++))
    fi
else
    echo "   ✗ VERSION file not found"
    ((ERRORS++))
fi

# Check 2: Release directory exists
echo "2. Checking release directory..."
if [ -d "$RELEASE_DIR" ]; then
    echo "   ✓ Release directory exists: $RELEASE_DIR"
else
    echo "   ✗ Release directory not found: $RELEASE_DIR"
    ((ERRORS++))
fi

# Check 3: Binaries exist
echo "3. Checking binaries..."
for binary in floyd superfloyd; do
    if [ -f "$RELEASE_DIR/$binary" ]; then
        SIZE=$(ls -lh "$RELEASE_DIR/$binary" | awk '{print $5}')
        echo "   ✓ $binary exists ($SIZE)"
    else
        echo "   ✗ $binary not found"
        ((ERRORS++))
    fi
done

# Check 4: Binaries are executable
echo "4. Checking binaries are executable..."
for binary in floyd superfloyd; do
    if [ -x "$RELEASE_DIR/$binary" ]; then
        echo "   ✓ $binary is executable"
    else
        echo "   ✗ $binary is not executable"
        ((ERRORS++))
    fi
done

# Check 5: Checksums exist
echo "5. Checking checksums..."
if [ -f "$RELEASE_DIR/checksums.txt" ]; then
    echo "   ✓ checksums.txt exists"
    # Verify checksums
    cd "$RELEASE_DIR"
    if sha256sum -c checksums.txt &>/dev/null; then
        echo "   ✓ Checksums verified"
    elif shasum -a 256 -c checksums.txt &>/dev/null; then
        echo "   ✓ Checksums verified"
    else
        echo "   ⚠ Checksum verification failed"
        ((WARNINGS++))
    fi
    cd "$FLOYD_ROOT"
else
    echo "   ⚠ checksums.txt not found"
    ((WARNINGS++))
fi

# Check 6: Manifest exists
echo "6. Checking manifest..."
if [ -f "$RELEASE_DIR/manifest.json" ]; then
    echo "   ✓ manifest.json exists"
    # Validate JSON
    if command -v jq &>/dev/null; then
        if jq empty "$RELEASE_DIR/manifest.json" 2>/dev/null; then
            echo "   ✓ manifest.json is valid JSON"
            MANIFEST_VERSION=$(jq -r '.version' "$RELEASE_DIR/manifest.json")
            if [ "$MANIFEST_VERSION" == "$VERSION" ]; then
                echo "   ✓ Manifest version matches: $VERSION"
            else
                echo "   ✗ Manifest version mismatch: $MANIFEST_VERSION != $VERSION"
                ((ERRORS++))
            fi
        else
            echo "   ✗ manifest.json is invalid JSON"
            ((ERRORS++))
        fi
    else
        echo "   ⚠ jq not available, skipping JSON validation"
        ((WARNINGS++))
    fi
else
    echo "   ⚠ manifest.json not found"
    ((WARNINGS++))
fi

# Check 7: Release notes exist
echo "7. Checking release notes..."
RELEASE_NOTES="docs/releases/v${VERSION}.md"
if [ -f "$RELEASE_NOTES" ]; then
    echo "   ✓ Release notes exist: $RELEASE_NOTES"
    # Check for required sections
    if grep -q "## Summary" "$RELEASE_NOTES" 2>/dev/null; then
        echo "   ✓ Summary section found"
    else
        echo "   ⚠ Summary section not found"
        ((WARNINGS++))
    fi
    if grep -q "## Changes" "$RELEASE_NOTES" 2>/dev/null; then
        echo "   ✓ Changes section found"
    else
        echo "   ⚠ Changes section not found"
        ((WARNINGS++))
    fi
else
    echo "   ⚠ Release notes not found: $RELEASE_NOTES"
    ((WARNINGS++))
fi

# Check 8: Run binary version checks
echo "8. Checking binary versions..."
for binary in floyd superfloyd; do
    if [ -x "$RELEASE_DIR/$binary" ]; then
        BIN_VERSION=$("$RELEASE_DIR/$binary" --version 2>/dev/null | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' | head -1)
        if [ -n "$BIN_VERSION" ]; then
            EXPECTED="v${VERSION}"
            if [ "$BIN_VERSION" == "$EXPECTED" ]; then
                echo "   ✓ $binary version correct: $BIN_VERSION"
            else
                echo "   ✗ $binary version mismatch: $BIN_VERSION != $EXPECTED"
                ((ERRORS++))
            fi
        else
            echo "   ⚠ Could not determine $binary version"
            ((WARNINGS++))
        fi
    fi
done

# Check 9: Go tests pass (if go is available)
echo "9. Running tests..."
if command -v go &>/dev/null; then
    if go test ./... -short &>/dev/null; then
        echo "   ✓ All tests passed"
    else
        echo "   ⚠ Some tests failed (run 'go test ./...' for details)"
        ((WARNINGS++))
    fi
else
    echo "   ⚠ Go not available, skipping tests"
    ((WARNINGS++))
fi

# Summary
echo ""
echo "=========================================="
echo "Validation Summary"
echo "=========================================="
echo ""
echo "Version: v${VERSION}"
echo "Release: $RELEASE_DIR"
echo ""
if [ $ERRORS -eq 0 ] && [ $WARNINGS -eq 0 ]; then
    echo "✓ All checks passed!"
    echo ""
    echo "Ready to release. Next steps:"
    echo "  git tag -a v${VERSION} -m 'Release v${VERSION}'"
    echo "  git push origin v${VERSION}"
    exit 0
elif [ $ERRORS -eq 0 ]; then
    echo "✓ All critical checks passed"
    echo "⚠ $WARNINGS warning(s) found"
    echo ""
    echo "Review warnings before releasing."
    exit 0
else
    echo "✗ $ERRORS error(s) found"
    echo "⚠ $WARNINGS warning(s) found"
    echo ""
    echo "Fix errors before releasing."
    exit 1
fi
