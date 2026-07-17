package agent

import (
	"testing"

	"charm.land/fantasy/providers/anthropic"
	"github.com/legacy-ai/floyd/internal/config"
)

func TestAnthropicReasoningEffortMapsToThinkingBudget(t *testing.T) {
	tests := []struct {
		effort string
		want   int64
	}{
		{effort: "low", want: 2000},
		{effort: "medium", want: 6000},
		{effort: "high", want: 16000},
	}

	for _, test := range tests {
		t.Run(test.effort, func(t *testing.T) {
			options := getProviderOptions(Model{
				ModelCfg: config.SelectedModel{ReasoningEffort: test.effort},
			}, config.ProviderConfig{Type: anthropic.Name})

			got, ok := options[anthropic.Name].(*anthropic.ProviderOptions)
			if !ok || got.Thinking == nil {
				t.Fatalf("anthropic thinking options missing: %#v", options[anthropic.Name])
			}
			if got.Thinking.BudgetTokens != test.want {
				t.Fatalf("budget = %d, want %d", got.Thinking.BudgetTokens, test.want)
			}
		})
	}
}
