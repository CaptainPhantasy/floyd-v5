package mcp

import (
	"context"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/google/uuid"
)

// PersistentLab represents a long-lived, privileged execution environment.
type PersistentLab struct {
	ID         string
	Container  string
	WorkingDir string
	CreatedAt  time.Time
	mu         sync.Mutex
}

var (
	activeLabs = make(map[string]*PersistentLab)
	labsMu     sync.Mutex
)

// GetLab retrieves an existing lab for a session or returns nil.
func GetLab(sessionID string) *PersistentLab {
	labsMu.Lock()
	defer labsMu.Unlock()
	return activeLabs[sessionID]
}

// ProvisionLab ignites a privileged, persistent Docker container and mirrors the host workspace.
func ProvisionLab(ctx context.Context, sessionID string, hostWd string, image string) (*PersistentLab, error) {
	labsMu.Lock()
	if lab, ok := activeLabs[sessionID]; ok {
		labsMu.Unlock()
		return lab, nil
	}
	labsMu.Unlock()

	containerName := fmt.Sprintf("floyd-lab-%s", uuid.New().String()[:8])

	// 1. Ignite the privileged container in detached mode
	args := []string{
		"run", "-d", "--privileged",
		"--name", containerName,
		"--network", "bridge",
		"-w", "/workspace",
		image,
		"tail", "-f", "/dev/null", // Keep alive
	}

	if err := exec.CommandContext(ctx, "docker", args...).Run(); err != nil {
		return nil, fmt.Errorf("failed to ignite persistent lab: %w", err)
	}

	// 2. Mirror host workspace to container (Isolated Clone)
	// We copy the contents of the current host directory into the container's /workspace
	copyArgs := []string{"cp", hostWd + "/.", containerName + ":/workspace/"}
	if err := exec.CommandContext(ctx, "docker", copyArgs...).Run(); err != nil {
		// Cleanup on failure
		_ = exec.Command("docker", "rm", "-f", containerName).Run()
		return nil, fmt.Errorf("failed to mirror workspace to lab: %w", err)
	}

	lab := &PersistentLab{
		ID:         sessionID,
		Container:  containerName,
		WorkingDir: "/workspace",
		CreatedAt:  time.Now(),
	}

	labsMu.Lock()
	activeLabs[sessionID] = lab
	labsMu.Unlock()

	return lab, nil
}

// Execute runs a command inside the persistent lab and returns output.
func (l *PersistentLab) Execute(ctx context.Context, command string) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	args := []string{"exec", l.Container, "/bin/sh", "-c", command}
	cmd := exec.CommandContext(ctx, "docker", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("lab execution failed: %w", err)
	}
	return string(out), nil
}

// TeardownLab destroys the specific container.
func TeardownLab(sessionID string) error {
	labsMu.Lock()
	lab, ok := activeLabs[sessionID]
	if !ok {
		labsMu.Unlock()
		return nil
	}
	delete(activeLabs, sessionID)
	labsMu.Unlock()

	return exec.Command("docker", "rm", "-f", lab.Container).Run()
}
