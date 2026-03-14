#!/bin/bash
# Floyd Version Bump Script
# Usage: ./scripts/version-bump.sh [major|minor|patch] [--dry-run]

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
FLOYD_ROOT="$(dirname "$SCRIPT_DIR")"
cd "$FLOYD_ROOT"

DRY_RUN=false
PART=""

# Parse arguments
for arg in "$@"; do
    case $arg in
        --dry-run) DRY_RUN=true ;;
        major|minor|patch) PART="$arg" ;;
        *) echo "Unknown argument: $arg"; exit 1 ;;
    esac
done

if [ -z "$PART" ]; then
    echo "Usage: $0 [major|minor|patch] [--dry-run]"
    echo ""
    echo "Examples:"
    echo "  $0 patch    # 5.0.0 -> 5.0.1"
    echo "  $0 minor    # 5.0.0 -> 5.1.0"
    echo "  $0 major    # 5.0.0 -> 6.0.0"
    exit 1
fi

# Read current version
CURRENT=$(cat VERSION 2>/dev/null || echo "0.0.0")
echo "Current version: $CURRENT"

# Parse version components
IFS='.' read -r MAJOR MINOR PATCH <<< "$CURRENT"

# Validate
if [ -z "$MAJOR" ] || [ -z "$MINOR" ] || [ -z "$PATCH" ]; then
    echo "Error: Invalid version format: $CURRENT"
    echo "Expected: X.Y.Z"
    exit 1
fi

# Calculate new version
case "$PART" in
    major)
        NEW_MAJOR=$((MAJOR + 1))
        NEW_MINOR=0
        NEW_PATCH=0
        ;;
    minor)
        NEW_MAJOR=$MAJOR
        NEW_MINOR=$((MINOR + 1))
        NEW_PATCH=0
        ;;
    patch)
        NEW_MAJOR=$MAJOR
        NEW_MINOR=$MINOR
        NEW_PATCH=$((PATCH + 1))
        ;;
esac

NEW_VERSION="${NEW_MAJOR}.${NEW_MINOR}.${NEW_PATCH}"

echo ""
echo "=========================================="
echo "Version Bump: $CURRENT -> $NEW_VERSION"
echo "=========================================="
echo ""
echo "Type: $PART"
echo "Changes:"
case "$PART" in
    major) echo "  - Breaking changes" ;;
    minor) echo "  - New features" ;;
    patch) echo "  - Bug fixes" ;;
esac

if [ "$DRY_RUN" = true ]; then
    echo ""
    echo "=== DRY RUN MODE ==="
    echo "Would update VERSION file to: $NEW_VERSION"
    exit 0
fi

# Confirm
read -p "Proceed with version bump? [y/N] " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 1
fi

# Update VERSION file
echo "$NEW_VERSION" > VERSION
echo "Updated VERSION file"

# Create release notes template if not exists
RELEASE_NOTES="docs/releases/v${NEW_VERSION}.md"
if [ ! -f "$RELEASE_NOTES" ]; then
    if [ -f "docs/templates/RELEASE_NOTES.tmpl.md" ]; then
        cp docs/templates/RELEASE_NOTES.tmpl.md "$RELEASE_NOTES"
        # Replace placeholders
        sed -i '' "s/{VERSION}/${NEW_VERSION}/g" "$RELEASE_NOTES" 2>/dev/null || \
        sed -i "s/{VERSION}/${NEW_VERSION}/g" "$RELEASE_NOTES"
        sed -i '' "s/{DATE}/$(date -u +"%Y-%m-%d")/g" "$RELEASE_NOTES" 2>/dev/null || \
        sed -i "s/{DATE}/$(date -u +"%Y-%m-%d")/g" "$RELEASE_NOTES"
        sed -i '' "s/{PREVIOUS_VERSION}/${CURRENT}/g" "$RELEASE_NOTES" 2>/dev/null || \
        sed -i "s/{PREVIOUS_VERSION}/${CURRENT}/g" "$RELEASE_NOTES"
        echo "Created release notes template: $RELEASE_NOTES"
    else
        echo "  ⚠️  WARNING: Release notes template not found"
    fi
else
    echo "Release notes already exist: $RELEASE_NOTES"
fi

# Update changelog.go if it exists
CHANGELOG_GO="internal/version/changelog.go"
if [ -f "$CHANGELOG_GO" ]; then
    # Try to update CurrentVersion
    if grep -q "CurrentVersion:" "$CHANGELOG_GO"; then
        sed -i '' "s/CurrentVersion: \"v[^\"]*\"/CurrentVersion: \"v${NEW_VERSION}\"/" "$CHANGELOG_GO" 2>/dev/null || \
        sed -i "s/CurrentVersion: \"v[^\"]*\"/CurrentVersion: \"v${NEW_VERSION}\"/" "$CHANGELOG_GO"
        echo "Updated CurrentVersion in $CHANGELOG_GO"
    fi
fi

echo ""
echo "=========================================="
echo "Version bump complete!"
echo "=========================================="
echo ""
echo "Updated files:"
echo "  - VERSION -> $NEW_VERSION"
echo "  - $RELEASE_NOTES (created if new)"
echo ""
echo "Next steps:"
echo "  1. Update release notes: $RELEASE_NOTES"
echo "  2. Commit changes: git add -A && git commit -m 'chore: bump version to v${NEW_VERSION}'"
echo "  3. Build and test: ./scripts/build.sh"
echo "  4. Create release: ./scripts/release.sh"
