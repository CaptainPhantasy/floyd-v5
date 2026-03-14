package agent

import (
	"context"
	"testing"

	"github.com/legacy-ai/floyd/internal/config"
	"github.com/legacy-ai/floyd/internal/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPromptSelection(t *testing.T) {
	// Test floyd selection
	version.BinaryName = "floyd"
	p, err := coderPrompt()
	require.NoError(t, err)
	content, err := p.Build(context.Background(), "", "", config.Config{})
	require.NoError(t, err)
	assert.Contains(t, content, "MULTIPURPOSE OPERATIONAL AGENT")

	// Test superfloyd selection
	version.BinaryName = "superfloyd"
	p, err = coderPrompt()
	require.NoError(t, err)
	content, err = p.Build(context.Background(), "", "", config.Config{})
	require.NoError(t, err)
	assert.Contains(t, content, "SOTA CODING SPECIALIST")
}
