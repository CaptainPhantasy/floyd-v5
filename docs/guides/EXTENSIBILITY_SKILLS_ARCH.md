# Floyd LLM-Agnostic Skills Architecture

## System Overview
As of v5.2.x/v5.3.0, the Floyd ecosystem shifted from a highly monolithic Node.js MCP server architecture (12+ separate, heavy local servers) to a modern hybrid **LLM-Agnostic Skills** framework. 

Rather than executing proprietary algorithms via JSON-RPC calls over localhost to a Node.js daemon (which consumes V8 memory, causes bloat, and makes NPM package distribution immensely difficult), we extract the algorithmic logic, execution steps, and context data into plain-text Markdown documents called **Skills**.

## Why Markdown Skills?
LLMs are exceptionally powerful at reading instructions and executing standard logic (like mathematical graphing, tarjan circular dependency detection, or syntax AST validation) if properly prompted. 
- **Zero Latency:** No RPC calls required. The logic is loaded directly into the active prompt context.
- **Portability:** These skills can be distributed inside the Git repository naturally.
- **Maintainability:** You just edit text files instead of rebuilding TypeScript servers.

## How Agents Use Skills
Inside `extensibility/skills/`, there are distinct sub-directories by category. Each folder represents a unique operation (e.g., `extensibility/skills/linting/lint-fix-go/SKILL.md`).

1. **Discovery:** The Agent observes the overarching tool/skill catalog in `extensibility/SKILLS_INDEX.md`.
2. **Access:** The Agent reads the required `SKILL.md` via the standard `view` or `read_file` system tool.
3. **Execution:** The document instructs the agent exactly how to orchestrate the internal tools (like running `glob`, executing AST checks via CLI, or mutating code via `multiedit`) to achieve the result.

## Integrating the Cloud & Premium Skills
Out of the box, the Floyd NPM distribution ships natively with the **Top 20 absolute core skills** (such as `tarjan-circular-dependency-detection`, `code-review-workflow`, `semantic-diff-validation`, etc.) bundled directly in the `extensibility/skills/core/` directory.

However, some skills require extreme computational overhead, multi-agent swarming, or API access logic that a local `.md` file cannot handle. For these, Floyd leverages the **Remote Hub Interface**.

The Floyd agent can connect to **floydslabs.com** directly, which exposes **50+ additional high-velocity AI skills** over a single HTTP SSE JSON-RPC tunnel, preventing your host machine from having to natively host complex API orchestration loops.

### Getting an API Key
If you downloaded Floyd via NPM and want access to the full 73+ skill library:
1. Visit [https://floydslabs.com/connect](https://floydslabs.com/connect)
2. Submit a request for API Key access.
3. Configure your `floyd.json` with the remote JSON-RPC URL and your new `floyd_...` key.

## Migration Reference
To see which older generation MCP servers were migrated, refer to the root `COMPLETE_SKILLS_MIGRATION_SUMMARY.md`. 
*Note:* The only MCP servers retained locally are those requiring stateful disk operations (`SuperCache`, `Terminal`, `Lab`).