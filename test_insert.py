import sys

# Simulate the template
lines = [
    "## CORE RULES\n",
    "- Rule 1\n",
    "- Rule 2\n",
    "\n",
    "---\n",
    "\n",
    "## MCP TOOLS REFERENCE\n",
    "Some content\n"
]

# Find the line "## MCP TOOLS REFERENCE"
insert_idx = -1
for i, line in enumerate(lines):
    if line.strip() == "## MCP TOOLS REFERENCE":
        insert_idx = i
        break

print(f"Found at index: {insert_idx}")

# Insert SOTA section
sota_section = """\n---\n\n## SOTA SECTION\n- Item 1\n- Item 2\n\n"""
sota_lines = sota_section.splitlines(keepends=True)
new_lines = lines[:insert_idx] + sota_lines + lines[insert_idx:]

print("Result:")
print(''.join(new_lines))
