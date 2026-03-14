# Floyd Extensibility Guide

> **Skills • Agents • Plugins** — Extend Floyd with custom capabilities

This guide covers three powerful extensibility mechanisms in Floyd that let you customize behavior, add specialized knowledge, and create reusable workflow bundles.

---

## Table of Contents

1. [Quick Comparison](#quick-comparison)
2. [Skills](#skills)
   - [What are Skills?](#what-are-skills)
   - [Creating a Skill](#creating-a-skill)
   - [Skill Directory Structure](#skill-directory-structure)
   - [Skill Discovery](#skill-discovery)
3. [Agents](#agents)
   - [What are Agents?](#what-are-agents)
   - [Built-in Agents](#built-in-agents)
   - [Agent Configuration](#agent-configuration)
4. [Plugins](#plugins)
   - [What are Plugins?](#what-are-plugins)
   - [Plugin Components](#plugin-components)
   - [Creating a Plugin](#creating-a-plugin)
   - [Plugin Examples](#plugin-examples)
5. [Configuration Reference](#configuration-reference)
6. [Best Practices](#best-practices)
7. [Troubleshooting](#troubleshooting)

---

## Quick Comparison

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                    FLOYD EXTENSIBILITY LAYERS                               │
├─────────────────┬─────────────────┬─────────────────────────────────────────┤
│                 │                 │                                         │
│    SKILLS       │    AGENTS       │              PLUGINS                    │
│    ───────      │    ──────       │              ───────                    │
│                 │                 │                                         │
│  Single-purpose │  Pre-configured │  MCP Server Integrations                │
│  instruction    │  persona and    │                                         │
│  sets           │  tool access    │  • Model Context Protocol               │
│                 │                 │  • External Tool Access                 │
│  "How to do X"  │  "What role to  │  • Resource Bridging                    │
│                 │   embody"       │  • Real-time Status                     │
│                 │                 │                                         │
├─────────────────┼─────────────────┼─────────────────────────────────────────┤
│  SKILL.md       │  Agent Library  │  MCP Registry                           │
│  (YAML + MD)    │  (.md files)    │  (via floyd.json)                       │
├─────────────────┼─────────────────┼─────────────────────────────────────────┤
│  20+ Core       │  90+ Domain     │  Active/Available                       │
│  Library        │  Specialists    │  Monitoring                             │
└─────────────────┴─────────────────┴─────────────────────────────────────────┘
```

---

## The Unified Library System

In v5.0.1, Floyd introduces a unified category-based UI for browsing all extensibility components. Use `ctrl+p` (Command Palette) or the dedicated library dialogs to browse by domain.

### Navigation Shortcuts

| Key | Action |
|-----|--------|
| `Tab` | Cycle Categories |
| `Alt+1..6` | Direct Jump to Domain (Commands, Agents, Skills, etc.) |
| `0..9` | Jump to specific Category |
| `e` | Expand/Collapse Description Preview |
| `/` | Focus Search Filter |

---

## Skills

### What are Skills?

Skills are reusable instruction sets that teach Floyd how to handle specific tasks or domains. They follow the [Agent Skills Open Standard](https://agentskills.io). Floyd v5.0.1 ships with **20 core skills** pre-installed.

### Categories (Alt+5)

1. **Git**: `git-commit`, `git-diff-expert`
2. **Testing**: `test-write-go`, `test-coverage-fix`
3. **Linting**: `lint-fix-go`, `dependency-unused-cleanup`
4. **Refactoring**: `refactor-extract`, `dry-logic-unifier`
5. **Documentation**: `doc-readme-gen`, `changelog-autogen`
6. **Deployment**: `dockerfile-pro`, `github-action-ci`
7. **Security**: `security-audit-checklist`
8. **Data**: `migration-create-sql`
9. **Debugging**: `debug-breakpoint-inject`, `log-trace-analysis`
10. **DX**: `env-validate`

---

## Agents

### What are Agents?

Agents are pre-configured personas with specific tool access levels and model settings. Floyd v5.0.1 includes **90 production-ready agents** categorized for frictionless discovery.

### Domain Specialization (Alt+4)

- **Architecture**: System Behavior Mappers, Monorepo Cartographers.
- **Infrastructure**: Git Stewards, Deployment Orchestrators.
- **Orchestration**: Swarm Coordinators, Type-Error Orchestrators.
- **Coding**: SuperFloyd SOTA Coder, Force Multiplier Architect.
- **Security**: Compliance Enforcers, Legal Shield Agents.
- **Quality**: Code Reviewers, Usability Inspectors.
- **Monitoring**: Incident Analysts, Postmortem Synthesizers.

---

## Plugins (MCP Servers)

### What are Plugins?

Plugins in Floyd are **MCP (Model Context Protocol) Servers**. They bridge the gap between Floyd's reasoning and your external tools (GitHub, Slack, Databases, etc.).

### Categories (Alt+6)

- **Connected**: Actively running and providing tools to the current session.
- **Configured**: Defined in your `floyd.json` but not yet active.
- **Available**: Suggested community plugins for instant integration.

### Persistent Terminal

The Terminal is now a first-class feature available at the top of your **System Commands (Alt+1)**. It is persistent, multiplexed, and bypasses agent restrictions for direct manual control.

```
┌──────────────────────────────────────────────────────────────────────────┐
│                           SKILL STRUCTURE                                │
├──────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│   ┌─────────────────────────────────────────────────────────────────┐   │
│   │  YAML FRONTMATTER (metadata)                                    │   │
│   │  ───────────────────────────                                    │   │
│   │  name: code-review                                              │   │
│   │  description: Perform thorough code reviews                     │   │
│   │  license: MIT                                                   │   │
│   │  compatibility: floyd >= 4.0                                    │   │
│   └─────────────────────────────────────────────────────────────────┘   │
│                                                                          │
│   ┌─────────────────────────────────────────────────────────────────┐   │
│   │  MARKDOWN BODY (instructions)                                   │   │
│   │  ──────────────────────────                                     │   │
│   │  ## Code Review Checklist                                       │   │
│   │                                                                 │   │
│   │  1. Check for security vulnerabilities                         │   │
│   │  2. Verify error handling                                       │   │
│   │  3. Review test coverage                                        │   │
│   │  ...                                                            │   │
│   └─────────────────────────────────────────────────────────────────┘   │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘
```

### Creating a Skill

**Step 1:** Create a directory with your skill name

```bash
mkdir -p ~/.config/floyd/skills/code-review
```

**Step 2:** Create a `SKILL.md` file

```markdown
---
name: code-review
description: Perform thorough code reviews focusing on security, performance, and maintainability
license: MIT
compatibility: floyd >= 4.0
metadata:
  author: your-name
  version: 1.0.0
---

## Code Review Guidelines

When reviewing code, systematically check:

### 1. Security
- SQL injection vulnerabilities
- XSS vulnerabilities
- Authentication/authorization issues
- Sensitive data exposure

### 2. Performance
- N+1 query problems
- Memory leaks
- Inefficient algorithms
- Unnecessary re-renders (for frontend)

### 3. Maintainability
- Code readability
- Naming conventions
- Documentation
- Test coverage

### 4. Best Practices
- SOLID principles
- DRY principle
- Error handling
- Logging

## Review Format

Structure your review as:

1. **Summary**: Brief overview of changes
2. **Critical Issues**: Must-fix problems
3. **Suggestions**: Nice-to-have improvements
4. **Questions**: Clarifications needed
```

**Step 3:** Floyd automatically discovers the skill

Skills are auto-loaded from default directories. No configuration needed!

### Skill Directory Structure

```
~/.config/floyd/skills/
├── code-review/
│   └── SKILL.md              # Required: skill definition
├── api-design/
│   └── SKILL.md
├── database-migrations/
│   └── SKILL.md
└── testing-strategies/
    └── SKILL.md

~/.config/agents/skills/      # Alternative location
├── deployment/
│   └── SKILL.md
└── monitoring/
    └── SKILL.md
```

### Skill Discovery

Floyd discovers skills from these directories (in order):

| Priority | Path | Environment Override |
|----------|------|---------------------|
| 1 | `$FLOYD_SKILLS_DIR` | `FLOYD_SKILLS_DIR` env var |
| 2 | `~/.config/floyd/skills/` | — |
| 3 | `~/.config/agents/skills/` | — |
| 4 | Custom paths in config | `options.skills_paths` |

**Configuration:**

```json
{
  "options": {
    "skills_paths": [
      "~/.config/floyd/skills",
      "./project-skills"
    ]
  }
}
```

### How Skills Work

When Floyd starts, it:

1. **Discovers** all `SKILL.md` files in configured paths
2. **Validates** each skill against the spec
3. **Injects** skill metadata into the system prompt:

```xml
<available_skills>
  <skill>
    <name>code-review</name>
    <description>Perform thorough code reviews...</description>
    <location>/Users/you/.config/floyd/skills/code-review/SKILL.md</location>
  </skill>
</available_skills>
```

4. **On-demand loading**: When you ask about code review, Floyd reads the full instructions from the skill file

---

## Agents

### What are Agents?

Agents are pre-configured personas with specific tool access levels and model settings. They define what Floyd CAN do in a given context.

```
┌──────────────────────────────────────────────────────────────────────────┐
│                         AGENT CONFIGURATION                              │
├──────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│   ┌─────────────────┐   ┌─────────────────┐   ┌─────────────────────┐   │
│   │     CODER       │   │      TASK       │   │    CUSTOM AGENT     │   │
│   │     ──────      │   │      ────       │   │    ────────────     │   │
│   │                 │   │                 │   │                     │   │
│   │  All tools      │   │  Read-only      │   │  Configurable       │   │
│   │  Full MCP       │   │  Limited MCP    │   │  Custom tools       │   │
│   │  Large model    │   │  Large model    │   │  Specific model     │   │
│   │                 │   │                 │   │                     │   │
│   │  FOR:           │   │  FOR:           │   │  FOR:               │   │
│   │  Development    │   │  Research       │   │  Specialized        │   │
│   │  Editing        │   │  Analysis       │   │  Workflows          │   │
│   │  Refactoring    │   │  Search         │   │                     │   │
│   └─────────────────┘   └─────────────────┘   └─────────────────────┘   │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘
```

### Built-in Agents

Floyd includes two built-in agents:

#### 1. Coder Agent
**Purpose:** Full-featured development work

```json
{
  "coder": {
    "id": "coder",
    "name": "Coder",
    "description": "An agent that helps with executing coding tasks.",
    "model": "large",
    "allowed_tools": ["bash", "edit", "write", "view", "grep", "glob", "ls", "..."],
    "allowed_mcp": null
  }
}
```

**When to use:** Writing code, refactoring, file operations, running commands

#### 2. Task Agent
**Purpose:** Read-only research and analysis

```json
{
  "task": {
    "id": "task",
    "name": "Task",
    "description": "An agent that helps with searching for context and finding implementation details.",
    "model": "large",
    "allowed_tools": ["glob", "grep", "ls", "sourcegraph", "view"],
    "allowed_mcp": null
  }
}
```

**When to use:** Code exploration, searching, documentation lookup

### Agent Configuration

Configure agent behavior in `floyd.json`:

```json
{
  "agents": {
    "coder": {
      "allowed_mcp": {
        "floyd-supercache-server": null,
        "web-search-prime": null
      }
    }
  }
}
```

**MCP Access Control:**

| Setting | Meaning |
|---------|---------|
| `null` | All tools from this MCP are allowed |
| `["tool1", "tool2"]` | Only listed tools are allowed |
| `[]` | No tools from this MCP |

### Agent Library

Create custom agents in the Agent Library:

**Location:** `~/.config/floyd/agents/` or project `.floyd/agents/`

**Format:**

```markdown
---
id: security-auditor
name: Security Auditor
description: Specialized agent for security audits
model: large
tools:
  - view
  - grep
  - glob
  - ls
mcps:
  - web-search-prime
---

## Security Audit Protocol

As a Security Auditor agent, you focus exclusively on:

1. **Vulnerability Scanning**
   - Check for known CVEs in dependencies
   - Scan for hardcoded secrets
   - Review authentication flows

2. **Code Analysis**
   - Identify injection vulnerabilities
   - Review access control implementations
   - Check cryptographic usage

3. **Reporting**
   - Document findings with severity levels
   - Provide remediation steps
   - Link to relevant security advisories
```

---

## Plugins

### What are Plugins?

Plugins are comprehensive capability packages that bundle multiple extensibility components together. They're similar to Claude's plugin architecture.

```
┌──────────────────────────────────────────────────────────────────────────┐
│                           PLUGIN ANATOMY                                 │
├──────────────────────────────────────────────────────────────────────────┤
│                                                                          │
│   PLUGIN.md                                                              │
│   ─────────                                                              │
│   ┌────────────────────────────────────────────────────────────────┐    │
│   │  METADATA (YAML)                                                │    │
│   │  name: devops-toolkit                                           │    │
│   │  version: 1.0.0                                                 │    │
│   │  category: development                                          │    │
│   └────────────────────────────────────────────────────────────────┘    │
│                                                                          │
│   ┌────────────────────────────────────────────────────────────────┐    │
│   │  COMPONENTS                                                     │    │
│   │  ──────────                                                     │    │
│   │                                                                 │    │
│   │  [Skills]        → Domain knowledge/instructions                │    │
│   │  [Slash Cmds]    → /create-pr, /deploy, /audit                  │    │
│   │  [Sub-Agents]    → code-reviewer, doc-writer                    │    │
│   │  [Connectors]    → MCP references (github, slack)               │    │
│   │                                                                 │    │
│   └────────────────────────────────────────────────────────────────┘    │
│                                                                          │
│   ┌────────────────────────────────────────────────────────────────┐    │
│   │  INSTRUCTIONS (Markdown body)                                   │    │
│   │  Detailed guidance loaded into the system prompt                │    │
│   └────────────────────────────────────────────────────────────────┘    │
│                                                                          │
└──────────────────────────────────────────────────────────────────────────┘
```

### Plugin Components

#### 1. Skills (Instructions)

The markdown body of `PLUGIN.md` serves as skill instructions that get injected into the system prompt.

#### 2. Slash Commands

Custom shortcuts that trigger automated actions:

```yaml
slash_commands:
  - name: Create Pull Request
    trigger: /create-pr
    description: Creates a PR with the current changes
    template: |
      Analyze the current changes and create a pull request with:
      - Clear title describing the changes
      - Detailed description of what and why
      - Test plan for verification
      - Any breaking changes noted
    auto_execute: false
```

#### 3. Sub-Agents

Specialized mini-instances for complex sub-tasks:

```yaml
sub_agents:
  - name: code-reviewer
    description: Specialized agent for thorough code review
    model_type: large
    tools:
      - view
      - grep
      - glob
    mcps:
      - web-search-prime
```

#### 4. Connectors

MCP connector references:

```yaml
connectors:
  - name: github
    description: GitHub API for repository operations
    type: required
  - name: slack
    description: Slack notifications
    type: optional
```

### Creating a Plugin

**Step 1:** Create the plugin directory

```bash
mkdir -p ~/.config/floyd/plugins/devops-toolkit
```

**Step 2:** Create `PLUGIN.md`

```markdown
---
name: devops-toolkit
version: 1.0.0
description: Complete DevOps toolkit with CI/CD, deployment, and monitoring workflows
license: MIT
author: Your Name
category: devops
tags:
  - ci-cd
  - deployment
  - monitoring
  - automation

# Slash Commands
slash_commands:
  - name: Deploy
    trigger: /deploy
    description: Deploy to the specified environment
    template: |
      Deploy the current codebase to {{environment}}:
      1. Run pre-deployment checks
      2. Build and tag containers
      3. Push to registry
      4. Update deployment manifests
      5. Verify deployment health

  - name: Pipeline Status
    trigger: /pipeline
    description: Check CI/CD pipeline status
    template: |
      Check the status of CI/CD pipelines:
      - List recent pipeline runs
      - Show failed stages
      - Display test results summary

# Sub-Agents
sub_agents:
  - name: deployment-coordinator
    description: Coordinates multi-service deployments
    model_type: large
  - name: log-analyzer
    description: Analyzes logs for issues
    model_type: small

# Connectors
connectors:
  - name: github
    description: GitHub API for PR and workflow management
    type: required
  - name: docker
    description: Docker daemon for container operations
    type: required

---

## DevOps Toolkit Instructions

This plugin provides comprehensive DevOps capabilities.

### Deployment Workflow

1. **Pre-Flight Checks**
   - Verify all tests pass
   - Check for security vulnerabilities
   - Validate configuration

2. **Build Phase**
   - Build containers with proper tagging
   - Run container security scans
   - Push to container registry

3. **Deploy Phase**
   - Update Kubernetes manifests
   - Apply changes with rolling update
   - Monitor deployment health

4. **Post-Deploy**
   - Run smoke tests
   - Update monitoring dashboards
   - Notify stakeholders

### Monitoring Guidelines

- Set up alerts for critical metrics
- Use structured logging
- Implement distributed tracing
- Monitor SLOs, not just availability
```

### Plugin Examples

#### Example 1: Code Review Plugin

```markdown
---
name: code-review
version: 1.0.0
description: Comprehensive code review with security and performance focus
category: development

slash_commands:
  - name: Review
    trigger: /review
    description: Perform a comprehensive code review
    template: "Review the current changes for: correctness, security, performance, maintainability."

sub_agents:
  - name: security-scanner
    description: Focuses on security vulnerabilities
  - name: perf-analyzer
    description: Analyzes performance implications

connectors:
  - name: github
    description: GitHub API for PR comments
    type: required
---

## Code Review Checklist

[Detailed review guidelines...]
```

#### Example 2: Documentation Plugin

```markdown
---
name: documentation
version: 1.0.0
description: Documentation generation and maintenance
category: productivity

slash_commands:
  - name: Generate API Docs
    trigger: /api-docs
    description: Generate API documentation from code
  - name: Update README
    trigger: /readme
    description: Update README with current project state

sub_agents:
  - name: doc-writer
    description: Specialized in writing clear documentation
    model_type: small
---

## Documentation Standards

[Documentation guidelines...]
```

#### Example 3: Financial Analysis Plugin

```markdown
---
name: financial-analysis
version: 2.0.0
description: Financial modeling and analysis capabilities
category: finance

slash_commands:
  - name: DCF Model
    trigger: /dcf
    description: Build a discounted cash flow model
  - name: Valuation
    trigger: /valuation
    description: Perform company valuation analysis

connectors:
  - name: bloomberg
    description: Financial data API
    type: optional
---

## Financial Analysis Framework

[Analysis methodologies...]
```

---

## Configuration Reference

### Complete Configuration Schema

```json
{
  "$schema": "./floyd-schema.json",
  
  "options": {
    "skills_paths": [
      "~/.config/floyd/skills",
      "~/.config/agents/skills",
      "./project-skills"
    ],
    "plugins_paths": [
      "~/.config/floyd/plugins",
      "~/.config/agents/plugins",
      "./project-plugins"
    ],
    "context_paths": [
      "FLOYD.md",
      "AGENTS.md"
    ]
  },
  
  "agents": {
    "coder": {
      "allowed_mcp": {
        "floyd-supercache-server": null,
        "web-search-prime": null
      }
    },
    "task": {
      "allowed_mcp": {
        "web-search-prime": ["search"]
      }
    }
  }
}
```

### Environment Variables

| Variable | Purpose | Default |
|----------|---------|---------|
| `FLOYD_SKILLS_DIR` | Override skills directory | `~/.config/floyd/skills` |
| `FLOYD_PLUGINS_DIR` | Override plugins directory | `~/.config/floyd/plugins` |
| `XDG_CONFIG_HOME` | Base config directory | `~/.config` |

### Directory Precedence

Floyd searches directories in this order:

```
1. Environment variable override (FLOYD_SKILLS_DIR, FLOYD_PLUGINS_DIR)
2. XDG_CONFIG_HOME/floyd/{skills,plugins}/
3. ~/.config/floyd/{skills,plugins}/
4. ~/.config/agents/{skills,plugins}/
5. Config file options.skills_paths / options.plugins_paths
```

---

## Best Practices

### Skills Best Practices

✅ **DO:**
- Keep skills focused on a single domain
- Write clear, actionable instructions
- Include examples in your skill body
- Version your skills with metadata
- Use descriptive names (kebab-case)

❌ **DON'T:**
- Create overly broad skills ("do everything")
- Duplicate instructions across skills
- Include sensitive information

### Agents Best Practices

✅ **DO:**
- Use the Task agent for read-only operations
- Configure MCP access explicitly for sensitive environments
- Create custom agents for specialized workflows

❌ **DON'T:**
- Grant unnecessary tool access
- Mix concerns in a single agent

### Plugins Best Practices

✅ **DO:**
- Group related functionality into plugins
- Use slash commands for common workflows
- Document required connectors
- Include sub-agents for complex tasks

❌ **DON'T:**
- Create monolithic plugins
- Duplicate skills across plugins
- Require connectors that aren't universally available

---

## Troubleshooting

### Skills Not Loading

**Symptom:** Skills don't appear in Floyd

**Solutions:**

1. Check file location:
   ```bash
   ls ~/.config/floyd/skills/*/SKILL.md
   ```

2. Validate YAML frontmatter:
   ```bash
   head -20 ~/.config/floyd/skills/my-skill/SKILL.md
   ```

3. Check debug logs:
   ```bash
   FLOYD_DEBUG=1 floyd
   ```

### Plugins Not Discovered

**Symptom:** Plugin slash commands don't work

**Solutions:**

1. Verify directory structure:
   ```bash
   ls ~/.config/floyd/plugins/*/PLUGIN.md
   ```

2. Check for validation errors in logs

3. Ensure plugin name matches directory name

### Agent MCP Access Issues

**Symptom:** MCP tools not available

**Solutions:**

1. Check `floyd.json` configuration:
   ```json
   {
     "agents": {
       "coder": {
         "allowed_mcp": {
           "your-mcp-server": null
         }
       }
     }
   }
   ```

2. Verify MCP server is configured:
   ```json
   {
     "mcp": {
       "your-mcp-server": {
         "type": "stdio",
         "command": "node",
         "args": ["path/to/server.js"]
       }
     }
   }
   ```

---

## File Format Reference

### SKILL.md Format

```yaml
---
# Required
name: skill-name                    # kebab-case, max 64 chars
description: What this skill does   # max 1024 chars

# Optional
license: MIT
compatibility: floyd >= 4.0
metadata:
  author: Your Name
  version: 1.0.0
---

## Instructions

Your skill instructions in Markdown...
```

### PLUGIN.md Format

```yaml
---
# Required
name: plugin-name                   # kebab-case, max 64 chars
description: What this plugin does  # max 2048 chars

# Optional
version: 1.0.0                      # max 32 chars
license: MIT
author: Your Name
category: development               # development, finance, productivity, etc.
tags:
  - tag1
  - tag2
metadata:
  key: value

# Components
slash_commands:
  - name: Command Name
    trigger: /trigger
    description: What it does

sub_agents:
  - name: agent-name
    description: What it does

connectors:
  - name: connector-name
    description: What it connects to
    type: required | optional
---

## Plugin Instructions

Your plugin instructions in Markdown...
```

---

## Summary

| Extensibility | File | Purpose | Best For |
|--------------|------|---------|----------|
| **Skills** | `SKILL.md` | Teach specialized knowledge | Domain expertise |
| **Agents** | `.md` library | Configure tool access | Role-based restrictions |
| **Plugins** | `PLUGIN.md` | Bundle capabilities | Complete workflows |

```
┌─────────────────────────────────────────────────────────────────────┐
│                    CHOOSE THE RIGHT TOOL                            │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│   "I want Floyd to know how to..." ───────────────► SKILL           │
│                                                                     │
│   "I want to restrict what Floyd can do..." ──────► AGENT           │
│                                                                     │
│   "I want a complete workflow package..." ─────────► PLUGIN         │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

*For more information, see the [Floyd Documentation](./docs/) or visit the [Agent Skills Specification](https://agentskills.io).*
