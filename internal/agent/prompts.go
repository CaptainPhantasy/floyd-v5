package agent

import (
	"context"
	_ "embed"

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
	// Select template based on binary name (floyd vs superfloyd)
	var tmpl []byte
	if version.BinaryName == "superfloyd" {
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
