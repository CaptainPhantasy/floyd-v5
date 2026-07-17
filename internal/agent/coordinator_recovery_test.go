package agent

import (
	"testing"

	"charm.land/fantasy/providers/anthropic"
	"github.com/legacy-ai/floyd/internal/config"
)

func TestAnthropicReasoningEffortMapsToThinkingBudget(t *testing.T) {
	tests := []struct {
		name     string
		effort   string
		override any
		want     int64
	}{
		{name: "low", effort: "low", want: 2000},
		{name: "medium", effort: "medium", want: 6000},
		{name: "medium alias", effort: "MED", want: 6000},
		{name: "high", effort: "high", want: 16000},
		{name: "high alias", effort: "MAX", want: 16000},
		{name: "explicit override", effort: "high", override: 4096, want: 4096},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			modelConfig := config.SelectedModel{ReasoningEffort: test.effort}
			if test.override != nil {
				modelConfig.ProviderOptions = map[string]any{"budget_tokens": test.override}
			}
			options := getProviderOptions(Model{
				ModelCfg: modelConfig,
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
