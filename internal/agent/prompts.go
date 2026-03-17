package agent

import (
	"context"
	_ "embed"
	"os"

	"github.com/legacy-ai/floyd/internal/agent/prompt"
	"github.com/legacy-ai/floyd/internal/config"
	"github.com/legacy-ai/floyd/internal/version"
)

//go:embed templates/floyd-general.md.tpl
var floydGeneralPromptTmpl []byte

//go:embed templates/superfloyd-coder.md.tpl
var superfloydCoderPromptTmpl []byte

//go:embed templates/task.md.tpl
var taskPromptTmpl []byte

//go:embed templates/initialize.md.tpl
var initializePromptTmpl []byte

//go:embed templates/floyd_protocol.md.tpl
var floydProtocolTmpl []byte

func coderPrompt(opts ...prompt.Option) (*prompt.Prompt, error) {
	// Select template based on runtime profile, with binary-name fallback for compatibility.
	var tmpl []byte
	profile := config.NormalizeRuntimeProfile(os.Getenv("FLOYD_RUNTIME_PROFILE"))
	if profile == config.RuntimeProfileFloyd && version.BinaryName == "superfloyd" {
		profile = config.RuntimeProfileSuperFloyd
	}

	if profile == config.RuntimeProfileSuperFloyd {
		tmpl = superfloydCoderPromptTmpl
	} else {
		tmpl = floydGeneralPromptTmpl
	}

	systemPrompt, err := prompt.NewPrompt("coder", string(tmpl), opts...)
	if err != nil {
		return nil, err
	}
	return systemPrompt, nil
}

func taskPrompt(opts ...prompt.Option) (*prompt.Prompt, error) {
	systemPrompt, err := prompt.NewPrompt("task", string(taskPromptTmpl), opts...)
	if err != nil {
		return nil, err
	}
	return systemPrompt, nil
}

func InitializePrompt(cfg config.Config) (string, error) {
	systemPrompt, err := prompt.NewPrompt("initialize", string(initializePromptTmpl))
	if err != nil {
		return "", err
	}
	return systemPrompt.Build(context.Background(), "", "", cfg)
}

// FloydProtocol returns the standard FLOYD.md protocol content.
func FloydProtocol() string {
	return string(floydProtocolTmpl)
}
