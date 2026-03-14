#!/bin/bash
# extract_from_db.sh
# Extracts MoE training data from Floyd's SQLite session database
# This provides the actual conversation history with tool calls

set -e

DB_FILE="${1:-/Volumes/Storage/.floyd/floyd.db}"
OUTPUT_DIR="${2:-./moe_datasets}"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)
OUTPUT_FILE="${OUTPUT_DIR}/glm5_moe_full_${TIMESTAMP}.jsonl"

mkdir -p "$OUTPUT_DIR"

echo "=== Floyd v3.5 Database Extractor ==="
echo "Database: $DB_FILE"
echo "Output: $OUTPUT_FILE"

# Check if database exists
if [[ ! -f "$DB_FILE" ]]; then
    echo "ERROR: Database not found: $DB_FILE"
    exit 1
fi

# Check if sqlite3 is available
if ! command -v sqlite3 &> /dev/null; then
    echo "ERROR: sqlite3 is required"
    exit 1
fi

# Get table schema
echo ""
echo "=== Database Schema ==="
sqlite3 "$DB_FILE" ".tables" 2>/dev/null || echo "Could not list tables"

# Extract session data
# Note: Adjust table names based on actual Floyd schema
echo ""
echo "=== Extracting Sessions ==="

# Try to extract messages table
sqlite3 -json "$DB_FILE" "
  SELECT
    s.id as session_id,
    s.created_at,
    m.role,
    m.content,
    m.tool_calls,
    m.tool_call_id
  FROM sessions s
  LEFT JOIN messages m ON s.id = m.session_id
  ORDER BY s.created_at DESC, m.created_at ASC
  LIMIT 1000
" 2>/dev/null > "${OUTPUT_FILE}.sessions" || {
  echo "Note: Standard schema not found, trying alternate..."
  # Try alternate schema
  sqlite3 -json "$DB_FILE" ".dump" 2>/dev/null | head -100 > "${OUTPUT_FILE}.schema"
}

# Count sessions
SESSION_COUNT=$(wc -l < "${OUTPUT_FILE}.sessions" 2>/dev/null || echo "0")
echo "Session entries extracted: $SESSION_COUNT"

# Format for GLM-5 training
if [[ -s "${OUTPUT_FILE}.sessions" ]]; then
  jq -c '
    # Group by session
    group_by(.session_id) |
    # Format each session as a training example
    map({
      session_id: .[0].session_id,
      messages: [
        {
          role: "system",
          content: "You are a senior production engineer operating with persistent continuity via SUPERCACHE..."
        },
        (.[] | select(.role == "user") | {role: "user", content: .content}),
        (.[] | select(.role == "assistant") | {role: "assistant", content: .content}),
        (.[] | select(.role == "tool") | {role: "tool", tool_call_id: .tool_call_id, content: .content})
      ] | map(select(.content != null and .content != ""))
    })[]
  ' "${OUTPUT_FILE}.sessions" > "$OUTPUT_FILE" 2>/dev/null || {
    echo "Could not format sessions, outputting raw data..."
    cp "${OUTPUT_FILE}.sessions" "$OUTPUT_FILE"
  }
fi

# Cleanup
rm -f "${OUTPUT_FILE}.sessions" "${OUTPUT_FILE}.schema"

echo ""
echo "=== Extraction Complete ==="
echo "Output: $OUTPUT_FILE"
echo "Lines: $(wc -l < "$OUTPUT_FILE" 2>/dev/null || echo '0')"

# Show sample
echo ""
echo "=== Sample Output ==="
head -2 "$OUTPUT_FILE" 2>/dev/null | jq -C '.' 2>/dev/null || head -2 "$OUTPUT_FILE"
