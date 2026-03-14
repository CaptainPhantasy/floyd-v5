#!/bin/bash
# extract_moe_dataset.sh
# Parses floyd.log for successful tool cycles and formats for GLM-5 RL fine-tuning
# Compatible with Floyd v3.5+ MCP 2.0 protocol

set -e

LOG_FILE="${1:-/Volumes/Storage/.floyd/logs/floyd.log}"
OUTPUT_DIR="${2:-./moe_datasets}"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
OUTPUT_FILE="${OUTPUT_DIR}/glm5_moe_training_${TIMESTAMP}.jsonl"

mkdir -p "$OUTPUT_DIR"

echo "=== Floyd v3.5 MoE Dataset Extractor ==="
echo "Log file: $LOG_FILE"
echo "Output: $OUTPUT_FILE"

# Check if log file exists
if [[ ! -f "$LOG_FILE" ]]; then
    echo "ERROR: Log file not found: $LOG_FILE"
    exit 1
fi

# Check if jq is available
if ! command -v jq &> /dev/null; then
    echo "ERROR: jq is required. Install with: brew install jq"
    exit 1
fi

# Extract successful tool cycles from JSON logs
# Floyd logs are JSON lines with structured data
jq -c '
  # Filter for tool-related entries with successful execution
  select(.msg != null and (.msg | contains("tool") or .msg | contains("Tool") or .msg | test("tool_call|tool_use|execution"; "i")))
  | {
      timestamp: .time,
      level: .level,
      source: .source.function,
      message: .msg
    }
' "$LOG_FILE" 2>/dev/null > "${OUTPUT_FILE}.raw" || {
    echo "Note: Log may not be in expected JSON format, attempting alternate extraction..."
    # Fallback: extract any JSON-like structures
    grep -oE '\{[^}]+\}' "$LOG_FILE" 2>/dev/null | head -1000 > "${OUTPUT_FILE}.raw"
}

# Count extracted entries
RAW_COUNT=$(wc -l < "${OUTPUT_FILE}.raw" 2>/dev/null || echo "0")
echo "Raw log entries extracted: $RAW_COUNT"

# Format for GLM-5 MoE training
# The training format expects:
# - system: Full PRIME DIRECTIVE
# - user: Original prompt
# - assistant: Thinking + tool_calls
# - tool: Execution result
# - assistant: Final response

cat << 'SYSTEM_PROMPT' > "${OUTPUT_DIR}/system_prompt.txt"
You are a senior production engineer operating with persistent continuity via SUPERCACHE. Provide clean, maintainable, production-ready solutions. Consider edge cases, performance, and security. Explain tradeoffs briefly. Avoid overengineering. Prioritize long-term maintainability and operational stability over short-term implementation speed.

## 0. PRIME DIRECTIVE
You operate in an environment with persistent continuity via SUPERCACHE.
You MUST use SUPERCACHE to determine project context and retrieve retained state.
However: stored state is not automatically true. Treat it as evidence, not authority.

## I. CORE INITIALIZATION (The "Wake Up" Routine) — MANDATORY
Before answering ANY prompt, you MUST:
1. Check Date/Location: Verify current system date (e.g., date -u). Use this for timestamping and log labels.
2. Mount SUPERCACHE: cache_retrieve(key="system:project_registry") to identify active project context.
3. Load Project State: Retrieve the project's status key (e.g., {project}:status, dsa:status, stat:gap_analysis) to understand last known state.
4. Load System Directive: cache_retrieve(key="system:directive_llm_optimization") to activate engine-optimized behaviors.

## II. MODE SELECTOR (MANDATORY)
Classify the task before any plan or fix:
- DEBUG MODE → runtime behavior bugs, unexpected output, failing tests, UI not responding, "same error persists"
- ORCHESTRATION MODE → multi-file feature work, refactors, migrations, structured build/test cycles
- EXPLORATION MODE → brainstorming, tradeoffs, architecture discussion
SYSTEM_PROMPT

# Create placeholder training samples in proper GLM-5 format
# These will be populated from actual session data
jq -c -n --argfile sys "${OUTPUT_DIR}/system_prompt.txt" '
  {
    messages: [
      {
        role: "system",
        content: $sys.content
      },
      {
        role: "user",
        content: "PLACEHOLDER_USER_INPUT"
      },
      {
        role: "assistant",
        content: "PLACEHOLDER_THINKING_PROCESS",
        tool_calls: [
          {
            id: "call_placeholder",
            type: "function",
            function: {
              name: "placeholder_tool",
              arguments: "{}"
            }
          }
        ]
      },
      {
        role: "tool",
        tool_call_id: "call_placeholder",
        content: "PLACEHOLDER_TOOL_RESULT"
      },
      {
        role: "assistant",
        content: "PLACEHOLDER_FINAL_RESPONSE"
      }
    ]
  }
' > "${OUTPUT_FILE}.template"

# Generate final dataset
echo "Generating training dataset..."

# Combine template with extracted data
# For now, output a manifest of what was found
cat << EOF > "$OUTPUT_FILE"
{"timestamp": "$(date -Iseconds)", "source": "floyd_v3.5", "log_entries": $RAW_COUNT, "format": "glm5_moe_jsonl"}
{"note": "Full extraction requires session database access. Use extract_from_db.sh for complete dataset."}
EOF

echo ""
echo "=== Extraction Summary ==="
echo "Output file: $OUTPUT_FILE"
echo "Template: ${OUTPUT_FILE}.template"
echo "System prompt: ${OUTPUT_DIR}/system_prompt.txt"
echo ""
echo "Dataset lines: $(wc -l < "$OUTPUT_FILE")"
echo ""
echo "Next steps:"
echo "1. Close all floyd sessions"
echo "2. Run: mv ~/.local/bin/floyd-new ~/.local/bin/floyd"
echo "3. Start fresh floyd session to generate clean MCP 2.0 logs"
echo "4. Re-run this script for full extraction"

# Cleanup
rm -f "${OUTPUT_FILE}.raw"
