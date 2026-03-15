package mcp

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

// EphemeralSandbox represents a sterile, single-use execution environment.
type EphemeralSandbox struct {
	ID        string
	Image     string
	cmd       *exec.Cmd
	stdin     io.WriteCloser
	stdout    io.ReadCloser
	CreatedAt time.Time
}

// SandboxConfig defines the constraints for the Superfloyd sandbox.
type SandboxConfig struct {
	Image     string
	Timeout   time.Duration
	MountPath string // e.g., mapping ./.floyd/.supercache
	MaxMemory string // e.g., "2g"
}

// Provision dynamically spins up a new isolated container and starts the MCP server inside it.
func Provision(ctx context.Context, cfg SandboxConfig) (*EphemeralSandbox, error) {
	sandboxID := fmt.Sprintf("floyd-sandbox-%s", uuid.New().String()[:8])

	args := []string{
		"run", "--rm", "-i",
		"--name", sandboxID,
		"--memory", cfg.MaxMemory,
		"--network", "bridge",
		"-v", fmt.Sprintf("%s:/workspace", cfg.MountPath),
		"-w", "/workspace",
		cfg.Image,
		"/bin/sh", // Entrypoint for passing standard commands
	}

	cmd := exec.CommandContext(ctx, "docker", args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to map sandbox stdin: %w", err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("failed to map sandbox stdout: %w", err)
	}

	cmd.Stderr = nil // Suppress raw stderr to avoid polluting the host stream

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to ignite ephemeral sandbox: %w", err)
	}

	return &EphemeralSandbox{
		ID:        sandboxID,
		Image:     cfg.Image,
		cmd:       cmd,
		stdin:     stdin,
		stdout:    stdout,
		CreatedAt: time.Now(),
	}, nil
}

// Teardown forcefully destroys the sandbox, flushing the state.
func (s *EphemeralSandbox) Teardown() error {
	_ = s.stdin.Close()
	killCmd := exec.Command("docker", "rm", "-f", s.ID)
	if err := killCmd.Run(); err != nil {
		return fmt.Errorf("failed to destroy sandbox %s: %w", s.ID, err)
	}
	return s.cmd.Wait()
}
