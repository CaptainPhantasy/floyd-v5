package mcp

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"

	"charm.land/fantasy"
)

// --- 1. SPAWN LAB TOOL ---

type SpawnLabParams struct {
	Image string `json:"image" description:"The Docker image to use (default: golang:1.24-alpine)"`
}

func NewSpawnLabTool(hostWd string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		"spawn_lab",
		"Spawns a persistent, privileged Docker lab and mirrors the current host workspace into it for safe, high-speed iteration.",
		func(ctx context.Context, params SpawnLabParams, call fantasy.ToolCall) (fantasy.ToolResponse, error) {
			sessionID := GetSessionFromContext(ctx)
			if sessionID == "" {
				return fantasy.NewTextErrorResponse("session_id required"), nil
			}

			image := params.Image
			if image == "" {
				image = "golang:1.24-alpine"
			}

			lab, err := ProvisionLab(ctx, sessionID, hostWd, image)
			if err != nil {
				return fantasy.NewTextErrorResponse(err.Error()), nil
			}

			return fantasy.NewTextResponse(fmt.Sprintf("Persistent Lab [%s] Online. Host workspace mirrored to /workspace.", lab.Container)), nil
		},
	)
}

// --- 2. EXECUTE IN LAB TOOL ---

type ExecuteInLabParams struct {
	Command string `json:"command" description:"The shell command to execute inside the persistent lab"`
}

func NewExecuteInLabTool() fantasy.AgentTool {
	return fantasy.NewAgentTool(
		"execute_in_lab",
		"Executes a command inside the active persistent lab and returns the output.",
		func(ctx context.Context, params ExecuteInLabParams, call fantasy.ToolCall) (fantasy.ToolResponse, error) {
			sessionID := GetSessionFromContext(ctx)
			lab := GetLab(sessionID)
			if lab == nil {
				return fantasy.NewTextErrorResponse("No active lab found for this session. Run spawn_lab first."), nil
			}

			output, err := lab.Execute(ctx, params.Command)
			if err != nil {
				return fantasy.NewTextResponse(fmt.Sprintf("Execution Error:\n%s\n%v", output, err)), nil
			}

			return fantasy.NewTextResponse(output), nil
		},
	)
}

// --- 3. MIGRATE TO HOST TOOL ---

type MigrateToHostParams struct {
	Path string `json:"path" description:"The path inside the lab (e.g. /workspace/src/main.go) to copy back to the host"`
}

func NewMigrateToHostTool(hostWd string) fantasy.AgentTool {
	return fantasy.NewAgentTool(
		"migrate_to_host",
		"Surgically copies a file or directory from the persistent lab back to the host macOS environment.",
		func(ctx context.Context, params MigrateToHostParams, call fantasy.ToolCall) (fantasy.ToolResponse, error) {
			sessionID := GetSessionFromContext(ctx)
			lab := GetLab(sessionID)
			if lab == nil {
				return fantasy.NewTextErrorResponse("No active lab found."), nil
			}

			if params.Path == "" {
				return fantasy.NewTextErrorResponse("path is required"), nil
			}

			// docker cp <container>:<container_path> <host_path>
			// We ensure we copy to the correct host location relative to hostWd
			relPath, _ := filepath.Rel("/workspace", params.Path)
			destPath := filepath.Join(hostWd, relPath)

			args := []string{"cp", lab.Container + ":" + params.Path, destPath}
			if err := exec.CommandContext(ctx, "docker", args...).Run(); err != nil {
				return fantasy.NewTextErrorResponse(fmt.Sprintf("Migration failed: %v", err)), nil
			}

			return fantasy.NewTextResponse(fmt.Sprintf("Successfully migrated %s to host: %s", params.Path, destPath)), nil
		},
	)
}

// Helper to get session ID from context (consistent with agent.go)
func GetSessionFromContext(ctx context.Context) string {
	if val, ok := ctx.Value("session_id").(string); ok {
		return val
	}
	return ""
}
