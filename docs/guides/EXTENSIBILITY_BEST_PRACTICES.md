# Extensibility Best Practices

Follow these guidelines to ensure your custom Agents and Skills maintain the high-integrity, deterministic standards of Floyd v5.0.1.

---

## 1. Creating Custom Agents

Agents define the **persona** and **access level**.

### Front Matter Standards
Every agent file in `~/.config/floyd/agents/` must include:
*   **name**: Descriptive and ASCII-only (no emojis).
*   **category**: Must be one of the 13 defined categories (architecture, infrastructure, coding, etc.).
*   **trigger**: A kebab-case keyword for quick invocation.
*   **tools**: Only grant the tools strictly necessary for the persona's role.

### Deterministic Prompts
*   **Prime Directive**: Start with a clear "You are..." statement.
*   **Silent Reasoning**: Instruct the agent to think step-by-step before using tools.
*   **Output Consistency**: Explicitly forbid conversational filler.

---

## 2. Creating Custom Skills

Skills provide **domain expertise**.

### Design Principles
*   **Atomic Scope**: One skill should do one thing (e.g., `git-commit` not `git-everything`).
*   **Input Requirements**: Clearly define what the skill needs to function (e.g., "Provide a diff of staged changes").
*   **Output Templates**: Provide a markdown template for the response to ensure consistency.

### File Structure
Always place your skill in a subfolder: `~/.config/floyd/skills/{category}/{skill-name}/SKILL.md`.

---

## 3. General Best Practices

### Visual Integrity
*   **Tables**: Never use Markdown tables. Always use box-drawing characters in code blocks for terminal compatibility.
*   **Diagrams**: Use Mermaid syntax for workflows.

### Security
*   **Redaction**: Skills should never ask the agent to output raw secrets.
*   **Verification**: Always include a "Verification" step in your skill instructions (e.g., "Run tests to verify the fix").

### Documentation
*   **Line References**: Use `file_path:line_number` when instructing the agent to reference code.
*   **Context**: Ensure the agent understands the difference between **Facts** (logs), **Decisions** (architectural choices), and **Hypotheses** (unverified theories).
