# Terminal Keybinding Compatibility Guide

## Overview

Floyd uses specific key combinations for AI suggestions and ghost text features. This guide documents known compatibility across different terminal emulators.

## Key Bindings for Suggestions

| Action | Primary | Alternatives | Description |
|--------|---------|---------------|-------------|
| **Request Suggestion** | `Ctrl+E` | - | Requests an AI suggestion on demand |
| **Accept Suggestion** | Backtick (`` ` ``) | `Ctrl+Y`, `Ctrl+]` | Accepts the ghost text into the editor |

## Terminal Compatibility

### Confirmed Working

| Terminal | `Ctrl+E` | Backtick | `Ctrl+Y` | `Ctrl+]` |
|----------|----------|----------|----------|----------|
| **Ghostty** | ✅ | ✅ | ✅ | ✅ |
| **iTerm2** | ⚠️ | ✅ | ✅ | ✅ |
| **Terminal.app** | ⚠️ | ✅ | ⚠️ | ⚠️ |

### Notes

- **Ghostty**: Full support for all suggestion keybindings.
- **iTerm2**: `Ctrl+E` may be intercepted by shell. Go to **Preferences → Keys → Keys → Hex Code** and ensure `0x05` is not mapped.
- **Terminal.app**: `Ctrl+E` works in most applications. `Ctrl+Y` may conflict with shell yank. `Ctrl+]` is generally safe.

### Fallback Keys

If `Ctrl+E` doesn't work reliably in your terminal:

1. **iTerm2**: Remove or remap `Ctrl+E` in Preferences
2. **Terminal.app**: Use as-is (generally works)
3. **VS Code Terminal**: Should work by default

For accepting suggestions, backtick is the most reliable across all terminals.

## Testing Your Setup

1. Open Floyd
2. Type some text in the editor
3. Press `Ctrl+E` - you should see "suggest now" in the status bar
4. If a suggestion appears, press backtick to accept it

## Troubleshooting

### `Ctrl+E` does nothing
- Check terminal preferences for key interception
- Try `Ctrl+[` as an alternative (may not be bound)

### Backtick doesn't accept
- Ensure cursor is at the end of the current text
- The suggestion must be visible for acceptance to work

### Suggestion doesn't appear
- Ensure an AI session is active
- Check that `FLOYD_ENABLE_TEST_SUGGESTIONS` is not set if running tests
- Try the `Ctrl+E` request explicitly

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `FLOYD_ENABLE_TEST_SUGGESTIONS` | Enable suggestions in test mode (set to any value) |

## Related Documentation

- [Keyboard Shortcuts](../reference/keyboard-shortcuts.md)
- [Ghost Text/Suggestion UX Redesign](../proposals/SUGGESTION_UX_REDESIGN_PLAN.md)
