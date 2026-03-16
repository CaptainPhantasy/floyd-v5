---
name: legacy-agent-orchestrator
description: Orchestrate multi-agent AI workflows for CASPER and Legacy Prime architectures with MindPal and Cursor IDE integration
---

# Legacy AI Agent Orchestrator

This skill enables sophisticated multi-agent AI workflows following the LegacyAI methodology for CASPER (Cognitive Agent for Synthesis, Planning, Execution & Refinement) and Legacy Prime architectures.

## Core Agent Chain Pattern

### Standard Three-Agent Flow
1. **Strategist Agent**: Initial analysis and planning
2. **Critique Agent**: Review and refinement
3. **Overwatch Agent**: Final validation and optimization

## MindPal Integration Format

When creating agent chains, output in MindPal-ready JSON:
```json
{
  "workflow": "agent_chain_name",
  "agents": [
    {
      "role": "Strategist",
      "prompt": "Analyze [TASK] and create comprehensive strategy",
      "outputs": ["strategy_doc", "implementation_plan"]
    },
    {
      "role": "Critique", 
      "prompt": "Review strategy from Agent 1 and identify gaps",
      "inputs": ["strategy_doc"],
      "outputs": ["critique_report", "improvements"]
    },
    {
      "role": "Overwatch",
      "prompt": "Synthesize all inputs into final deliverable",
      "inputs": ["strategy_doc", "critique_report"],
      "outputs": ["final_output"]
    }
  ]
}
```

## Cursor IDE Integration

Generate Cursor-compatible agent instructions with:
- Clear role definitions
- Step-by-step task breakdowns
- Context preservation between agents
- Error handling patterns

## CASPER Framework Implementation

Follow the CASPER methodology:
- **C**ognitive: Deep analysis and understanding
- **A**gent: Autonomous execution capabilities  
- **S**ynthesis: Combining multiple data sources
- **P**lanning: Strategic roadmap creation
- **E**xecution: Implementation with monitoring
- **R**efinement: Continuous improvement loops

## Best Practices

1. Always maintain context between agent handoffs
2. Include validation checkpoints
3. Design for iterative refinement
4. Output production-ready code/documents (no placeholders)
5. Optimize token usage while maintaining quality
