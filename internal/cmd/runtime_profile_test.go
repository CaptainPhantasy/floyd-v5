package cmd

import (
	"os"
	"testing"

	"github.com/legacy-ai/floyd/internal/config"
)

func TestProfileFromBinaryName(t *testing.T) {
	tests := map[string]config.RuntimeProfile{
		"floyd":      config.RuntimeProfileFloyd,
		"superfloyd": config.RuntimeProfileSuperFloyd,
		"beast":      config.RuntimeProfileSuperFloyd,
		"balanced":   config.RuntimeProfileSuperFloyd,
		"safe":       config.RuntimeProfileSuperFloyd,
		"sf":         config.RuntimeProfileSuperFloyd,
	}

	for bin, want := range tests {
		got := profileFromBinaryName(bin)
		if got != want {
			t.Fatalf("profileFromBinaryName(%q) = %q, want %q", bin, got, want)
		}
	}
}

func TestResolveRuntimeProfilePrecedence(t *testing.T) {
	t.Run("config wins over env and binary", func(t *testing.T) {
		t.Setenv("FLOYD_RUNTIME_PROFILE", "floyd")
		cfg := &config.Config{Options: &config.Options{RuntimeProfile: "superfloyd"}}
		got := resolveRuntimeProfile(cfg, "floyd")
		if got != config.RuntimeProfileSuperFloyd {
			t.Fatalf("got %q, want %q", got, config.RuntimeProfileSuperFloyd)
		}
	})

	t.Run("env wins when config empty", func(t *testing.T) {
		t.Setenv("FLOYD_RUNTIME_PROFILE", "superfloyd")
		cfg := &config.Config{Options: &config.Options{}}
		got := resolveRuntimeProfile(cfg, "floyd")
		if got != config.RuntimeProfileSuperFloyd {
			t.Fatalf("got %q, want %q", got, config.RuntimeProfileSuperFloyd)
		}
	})

	t.Run("binary fallback when config and env absent", func(t *testing.T) {
		_ = os.Unsetenv("FLOYD_RUNTIME_PROFILE")
		cfg := &config.Config{Options: &config.Options{}}
		got := resolveRuntimeProfile(cfg, "beast")
		if got != config.RuntimeProfileSuperFloyd {
			t.Fatalf("got %q, want %q", got, config.RuntimeProfileSuperFloyd)
		}
	})
}
