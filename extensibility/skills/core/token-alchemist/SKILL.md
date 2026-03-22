---
name: token-alchemist
description: Compress context windows to maximum information density with semantic preservation and structured summarization
category: core
version: "2.0.0"
---

# Token Alchemist

> Compress context windows to maximum information density with semantic preservation and structured summarization.

## When to Use
- WHEN `mode=BUILD` or `mode=DEBUG` and the context window is approaching capacity (>70% utilization)
- WHEN a long file or log output must be summarized without losing critical details
- WHEN switching between tasks and previous context needs to be compressed into a compact state

## Actions
`'compress' | 'analyze' | 'optimize'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| content | string | yes | The text or code to compress/analyze |
| strategy | string | no | Compression strategy: `summary`, `outline`, `key-facts`, `code-structure` (default: summary) |
| target_tokens | integer | no | Maximum token budget for the output (default: 50% of input) |
| preserve | string[] | no | Patterns or keywords that must be preserved in the output |

## Execution Pipeline

### Step 1: Analyze
Count the input token length. Identify the content type (code, prose, log, mixed). Use `grep` to extract the `preserve` patterns and mark them as immutable. If the content is a file, use `list_symbols` to extract the structural skeleton (function names, class names, imports) as a compression anchor.

### Step 2: Compress
Apply the selected `strategy`:
- **summary**: Generate a concise prose summary preserving key decisions and outcomes
- **outline**: Extract a hierarchical outline with section headers and one-line descriptions
- **key-facts**: Extract only factual statements, data points, and concrete findings
- **code-structure**: Strip implementation details, keep only signatures, types, and doc comments

Ensure all `preserve` patterns appear verbatim in the output. Trim to `target_tokens` if the initial compression exceeds the budget.

### Step 3: Optimize
Verify the compressed output retains the information needed to resume work. Check that function signatures, error messages, and file paths referenced in the original are preserved or summarized with enough context to relocate them.

## Output Shape
```json
{
  "input_tokens": 4200,
  "output_tokens": 1800,
  "compression_ratio": 0.43,
  "strategy": "string — strategy applied",
  "preserved_patterns": ["string — patterns found in output"],
  "compression_quality": "HIGH | MEDIUM | LOW",
  "compressed_content": "string — the compressed text"
}
```

## Failure Modes
- IF compression loses a `preserve` pattern: re-run with a higher `target_tokens` budget and report the near-miss
- IF the content type cannot be determined: default to `summary` strategy and note the uncertainty

## Examples
```json
{
  "action": "compress",
  "content": "[long file content or log output]",
  "strategy": "outline",
  "target_tokens": 500,
  "preserve": ["parseIdentifier", "TODO", "BUG"]
}
```
