---
name: linear-cli
description: Lightweight CLI for working with Linear issues, teams, projects, and labels. Use when managing Linear issues, creating/updating tasks, querying project status, or working with Linear programmatically. Provides alternative to Linear MCP with minimal dependencies.
---

# Linear CLI Skill

A Floyd skill providing a lightweight CLI for working with Linear issues. Written in JavaScript with minimal dependencies (Linear SDK and dotenv).

## What's Included

This skill bundles:
- `scripts/linear` – Standalone CLI tool
- Minimal dependencies (only @linear/sdk and dotenv)
- GitHub CLI-style conventions

## Setup

### 1. Get API Key

1. Go to https://linear.app/settings/api
2. Navigate to Settings > API > Personal API keys
3. Click "Create key"

### 2. Configure API Key

**Option 1: Environment variable**
```bash
export LINEAR_API_KEY="your-api-key"
```

**Option 2: .env file (recommended)**
```bash
# In the skill directory
echo 'LINEAR_API_KEY=your-api-key' > scripts/.env
```

## Usage

```bash
./scripts/linear <resource> <action> [arguments] [options]
```

## Core Commands

### Teams

```bash
# List all teams
./scripts/linear team list

# Get team details
./scripts/linear team view TEAM_KEY
```

### Issues

```bash
# List issues
./scripts/linear issue list [options]
  --limit N          Show N issues (default: 10)
  --team TEAM_KEY    Filter by team
  --assignee EMAIL   Filter by assignee
  --state STATE      Filter by state (backlog, unstarted, started, completed, canceled)
  --label LABEL      Filter by label

# View issue details
./scripts/linear issue view ISSUE_KEY

# Create issue
./scripts/linear issue create [options]
  --title "Title"    Issue title (required)
  --description "..."  Issue description
  --team TEAM_KEY    Team key (required)
  --assignee EMAIL   Assign to user
  --priority N       Priority (0-4: none, urgent, high, medium, low)
  --label LABEL      Add label
  --state STATE      Set state

# Update issue
./scripts/linear issue update ISSUE_KEY [options]
  --title "..."
  --description "..."
  --assignee EMAIL
  --priority N
  --state STATE
  --label LABEL      Add label
  --remove-label LABEL  Remove label

# Add comment
./scripts/linear issue comment ISSUE_KEY --body "Comment text"
```

### Projects

```bash
# List projects
./scripts/linear project list [options]
  --team TEAM_KEY    Filter by team

# View project
./scripts/linear project view PROJECT_KEY

# Create project
./scripts/linear project create [options]
  --name "Name"      Project name (required)
  --description "..."
  --team TEAM_KEY    Team key (required)
  --state STATE      Project state
```

### Labels

```bash
# List labels
./scripts/linear label list

# Create label
./scripts/linear label create [options]
  --name "Name"      Label name (required)
  --description "..."
  --color "#FF0000"  Hex color
```

## Common Workflows

### Create Task with Assignment

```bash
./scripts/linear issue create \
  --title "Fix authentication bug" \
  --description "Users can't log in with SSO" \
  --team ENG \
  --assignee john@company.com \
  --priority 1 \
  --label bug
```

### List Team's Active Issues

```bash
./scripts/linear issue list \
  --team ENG \
  --state started \
  --limit 20
```

### Update Issue Status

```bash
./scripts/linear issue update ENG-123 \
  --state completed \
  --label "deployed"
```

### Create Project and Link Issue

```bash
# Create project
./scripts/linear project create \
  --name "Q1 Migration" \
  --team ENG \
  --description "Database migration project"

# Update issue to link to project
./scripts/linear issue update ENG-123 \
  --project "Q1 Migration"
```

## Output Format

All commands output JSON by default for easy parsing:

```json
{
  "id": "abc123",
  "identifier": "ENG-123",
  "title": "Fix bug",
  "state": {
    "name": "In Progress"
  },
  "assignee": {
    "name": "John Doe",
    "email": "john@company.com"
  }
}
```

## Error Handling

Common errors:

**Invalid API key:**
```
Error: Invalid API key. Check LINEAR_API_KEY environment variable or .env file.
```

**Team not found:**
```
Error: Team 'XYZ' not found. Use 'linear team list' to see available teams.
```

**Missing required field:**
```
Error: --title is required for issue creation
```

## Field Reference

### Priority Values
- `0` - None
- `1` - Urgent
- `2` - High
- `3` - Medium
- `4` - Low

### State Values
- `backlog` - In backlog
- `unstarted` - Not started
- `started` - In progress
- `completed` - Done
- `canceled` - Canceled

### Common Team Keys
Use `linear team list` to find your team keys. Common patterns:
- `ENG` - Engineering
- `PROD` - Product
- `DES` - Design

## Troubleshooting

**Command not found:**
```bash
# Ensure linear script is executable
chmod +x scripts/linear
```

**Dependencies missing:**
```bash
# Install dependencies manually
cd scripts/
npm install
```

**API rate limiting:**
Linear API has rate limits. If you hit them, wait a few minutes before retrying.

## Integration with Other Skills

Works well with:
- `project-management` - For tracking work
- `plan-implementer` - For creating implementation tasks
- `systematic-debugging` - For bug tracking workflow
