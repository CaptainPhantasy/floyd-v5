// Package home provides utilities for dealing with the user's home directory.
package home

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	homedir    string
	homedirErr error
	homedirOnce sync.Once
)

// getHomeDir lazily initializes the home directory on first access.
func getHomeDir() (string, error) {
	homedirOnce.Do(func() {
		homedir, homedirErr = os.UserHomeDir()
		if homedirErr != nil {
			slog.Error("Failed to get user home directory", "error", homedirErr)
		}
	})
	return homedir, homedirErr
}

// Dir returns the user home directory.
// Deprecated: Use DirE for proper error handling.
func Dir() string {
	dir, _ := getHomeDir()
	return dir
}

// DirE returns the user home directory with error handling.
func DirE() (string, error) {
	return getHomeDir()
}

// Short replaces the actual home path from [Dir] with `~`.
func Short(p string) string {
	dir, err := getHomeDir()
	if err != nil || dir == "" || !strings.HasPrefix(p, dir) {
		return p
	}
	return filepath.Join("~", strings.TrimPrefix(p, dir))
}

// Long replaces the `~` with actual home path from [Dir].
func Long(p string) string {
	dir, err := getHomeDir()
	if err != nil || dir == "" || !strings.HasPrefix(p, "~") {
		return p
	}
	return strings.Replace(p, "~", dir, 1)
}
