package mcp

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"charm.land/fantasy"
)

type SpawnSandboxParams struct {
	Image   string `json:"image" description:"The Docker image to use (e.g., golang:1.24-alpine)"`
	Command string `json:"command" description:"The shell command to execute inside the sandbox"`
}

func NewSpawnSandboxTool() fantasy.AgentTool {
	return fantasy.NewAgentTool(
		"spawn_sandbox",
		"Spawns an ephemeral, sterile Docker container to execute untrusted code, run complex builds, or bypass host constraints. Returns the execution output.",
		func(ctx context.Context, params SpawnSandboxParams, call fantasy.ToolCall) (fantasy.ToolResponse, error) {
			image := params.Image
			if image == "" {
				image = "golang:1.24-alpine"
			}

			if params.Command == "" {
				return fantasy.NewTextErrorResponse("a command must be provided for the sandbox to execute"), nil
			}

			// Resolve the local cache path for mounting
			cachePath, err := filepath.Abs("./.floyd/.supercache")
			if err != nil {
				return fantasy.ToolResponse{}, fmt.Errorf("failed to resolve mount path: %w", err)
			}

			cfg := SandboxConfig{
				Image:     image,
				Timeout:   5 * time.Minute,
				MountPath: cachePath,
				MaxMemory: "2g",
			}

			sandbox, err := Provision(ctx, cfg)
			if err != nil {
				return fantasy.ToolResponse{}, fmt.Errorf("sandbox provisioning failed: %w", err)
			}
			defer sandbox.Teardown()

			// Dispatch the command to the sandbox's stdin
			_, err = io.WriteString(sandbox.stdin, params.Command+"\n")
			if err != nil {
				return fantasy.ToolResponse{}, fmt.Errorf("failed to dispatch command to sandbox: %w", err)
			}
			sandbox.stdin.Close()

			// Read the output
			output, err := io.ReadAll(sandbox.stdout)
			if err != nil {
				return fantasy.ToolResponse{}, fmt.Errorf("failed to read sandbox output: %w", err)
			}

			return fantasy.NewTextResponse(fmt.Sprintf("Sandbox [%s] Execution Complete.\nOutput:\n%s", sandbox.ID, string(output))), nil
		},
	)
}
