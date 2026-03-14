# Floyd vs SuperFloyd: Architectural Analysis

> **Analysis Date:** 2026-03-13
> **Analyst:** Code review simulation (no external documentation consulted)
> **Purpose:** Document the actual relationship between Floyd and SuperFloyd, multi-agent capabilities, and implementation path for differentiated agents

---

## Executive Summary

**Key Finding:** Floyd and SuperFloyd are the **same binary** with **runtime-detected differences** only. There is no multi-agent spawning capability currently active. The `beast`, `balanced`, `safe`, and `sf` commands merely change parallelism limits and UI theming.

---

## 1. Single Binary Architecture

### 1.1 Binary Detection

**Source:** `main.go:10-12`
```go
binName := "floyd"
if len(os.Args) > 0 {
    binName = filepath.Base(os.Args[0])
    cmd.SetRootUse(binName)
    cmd.SetupSuperFloydMode(binName)
}
```

**Finding:** The binary name is detected via `os.Args[0]` and passed to `SetupSuperFloydMode()`. No separate compilation exists.

### 1.2 SuperFloyd Mode Activation

**Source:** `internal/cmd/superfloyd_resilience.go:61-77`
```go
func isSuperFloydBinary() bool {
    name := strings.ToLower(strings.TrimSpace(filepathBase(os.Args[0])))
    return strings.Contains(name, "superfloyd") ||
        name == "beast" ||
        name == "balanced" ||
        name == "balance" ||
        name == "safe" ||
        name == "sf"
}

func SetupSuperFloydMode(binName string) {
    name := strings.ToLower(binName)
    if !isSuperFloydBinary() {
        return
    }
    // ... enables safety systems
}
```

**Finding:** All these names (`superfloyd`, `beast`, `balanced`, `safe`, `sf`) trigger the same "SuperFloyd mode" - there is no differentiation between them except parallelism.

### 1.3 Parallelism Configuration

**Source:** `internal/cmd/superfloyd_resilience.go:88-105`
```go
currParallel := os.Getenv("SUPERFLOYD_MAX_PARALLEL")

switch {
case name == "beast":
    if currParallel == "" {
        _ = os.Setenv("SUPERFLOYD_MAX_PARALLEL", "16")
    }
case name == "balanced" || name == "balance" || name == "sf" || name == "superfloyd":
    if currParallel == "" {
        _ = os.Setenv("SUPERFLOYD_MAX_PARALSLEL", "12")
    }
case name == "safe":
    if currParallel == "" {
        _ = os.Setenv("SUPERFLOYD_MAX_PARALLEL", "6")
    }
}
```

**Finding:** The ONLY difference between modes is the `SUPERFLOYD_MAX_PARALLEL` environment variable value.

---

## 2. Version Number Conflicts

### 2.1 Code Version

**Source:** `internal/version/version.go:5`
```go
var Version = "v1.8"
```

**Finding:** The canonical version in code is `v1.8`.

### 2.2 Prompt Template Version

**Source:** `internal/agent/templates/coder.md.tpl:5`
```markdown
- YOU ARE NOT CLAUDE. You are FLOYD v4.6.1.
```

**Source:** `FLOYD.md:9`
```markdown
- Version: v4.6.1
```

**Finding:** Prompts hardcode `v4.6.1`, which is neither the code version (`v1.8`) nor the intended Floyd version (`v5.0`). This is stale and should be synchronized.

---

## 3. UI/Theme Differences

### 3.1 Logo Detection

**Source:** `internal/ui/logo/logo.go:148-162`
```go
func isSuperFloyd() bool {
    name := strings.ToLower(filepath.Base(os.Args[0]))
    return strings.Contains(name, "superfloyd")
}

// In Render():
if isSuperFloyd() {
    floydASCII = `_____/\\\\\\\\\\\____/\\\__...` // Elaborate art
} else {
    floydASCII = `    __/\\\\\\\\\\\\\\\___/\\\__...` // Simple art
}
```

**Finding:** Different ASCII art is rendered based on binary name.

### 3.2 Theme Application

**Source:** `internal/ui/styles/themes.go:191-209`
```go
func IsSuperFloydBinary() bool {
    name := strings.ToLower(filepath.Base(os.Args[0]))
    return strings.Contains(name, "superfloyd")
}

func (s *Styles) AutoApplySuperFloydTheme() bool {
    if !IsSuperFloydBinary() {
        return false
    }
    for _, preset := range ThemePresets() {
        if preset.Name == ThemeSuperFloyd {
            s.ApplyTheme(preset)
            return true
        }
    }
    return false
}
```

**Source:** `internal/ui/styles/themes.go:144-162`
```go
{
    // SuperFloyd - Red, White, Blue patriotic theme
    Name:        ThemeSuperFloyd,
    Primary:     mustHex("#DC143C"), // Crimson Red
    Secondary:   mustHex("#FFFFFF"), // White
    Tertiary:    mustHex("#4169E1"), // Royal Blue
    BorderFocus: mustHex("#DC143C"), // Red border
    ...
}
```

**Finding:** SuperFloyd automatically applies a red/white/blue theme.

### 3.3 Persistent Bar

**Source:** `internal/ui/logo/logo.go:242-278`
```go
func PersistentBar(s *styles.Styles, width int) string {
    if !isSuperFloyd() {
        return ""  // Only shows for SuperFloyd
    }
    // SuperFloyd persistent ASCII art (85 chars wide)
    artLines := []string{
        "███████╗██╗   ██╗██████╗ ███████╗██████╗ ███████╗██╗      ██████╗ ██╗   ██╗██████╗ ",
        ...
    }
}
```

**Finding:** A 6-line persistent ASCII bar is rendered above the chat input ONLY for SuperFloyd.

---

## 4. Safety Systems

### 4.1 Default Values

**Source:** `internal/cmd/superfloyd_resilience.go:27-50`
```go
func qualityGatesEnabled() bool {
    v := strings.TrimSpace(strings.ToLower(os.Getenv("SUPERFLOYD_QUALITY_GATES")))
    if v == "" {
        return true  // Default ON
    }
    return v != "0" && v != "false" && v != "off"
}

func degradationControlsEnabled() bool {
    // Same pattern - defaults to true
}

func consistencyLockEnabled() bool {
    // Same pattern - defaults to true
}

func autoStabilizeEnabled() bool {
    // Same pattern - defaults to true
}
```

**Finding:** All safety systems default to ON, but can be disabled via environment variables.

### 4.2 Safety System Activation

**Source:** `internal/cmd/superfloyd_resilience.go:83-93`
```go
// Default safety systems ON for SuperFloyd family
if os.Getenv("SUPERFLOYD_QUALITY_GATES") == "" {
    _ = os.Setenv("SUPERFLOYD_QUALITY_GATES", "1")
}
if os.Getenv("SUPERFLOYD_DEGRADATION_CONTROLS") == "" {
    _ = os.Setenv("SUPERFLOYD_DEGRADATION_CONTROLS", "1")
}
// ... etc
```

**Finding:** For SuperFloyd modes, safety systems are forced ON unless explicitly set before launch.

---

## 5. Parallel Execution

### 5.1 Parallel Bash Tool

**Source:** `internal/agent/tools/parallel_bash.go:1-25`
```go
const (
    ParallelBashToolName         = "parallel_bash"
    DefaultMaxParallelCommands   = 4
    SuperFloydMaxParallel        = 12
    ParallelDefaultTimeout       = 60 * time.Second
    ParallelMaxOutputPerJob      = 10000
    DefaultParallelConcurrency   = 4
    SuperFloydMaxParallelCeiling = 32
)
```

### 5.2 Dynamic Limit Selection

**Source:** `internal/agent/tools/parallel_bash.go:232-250`
```go
func maxParallelCommandsForLane() int {
    max := DefaultMaxParallelCommands  // 4
    bin := strings.ToLower(strings.TrimSpace(filepathBase(os.Args[0])))
    if strings.Contains(bin, "superfloyd") {
        max = SuperFloydMaxParallel  // 12
    }
    if v := strings.TrimSpace(os.Getenv("SUPERFLOYD_MAX_PARALLEL")); v != "" {
        if n, err := strconv.Atoi(v); err == nil && n > 0 {
            if n > SuperFloydMaxParallelCeiling {
                n = SuperFloydMaxParallelCeiling
            }
            max = n
        }
    }
    if max < 1 {
        max = 1
    }
    return max
}
```

**Finding:** The parallelism limit is determined by:
1. Default: 4 commands
2. SuperFloyd binary: 12 commands
3. Environment override: Uses `SUPERFLOYD_MAX_PARALLEL` (up to 32 ceiling)

**Important:** This only affects the `parallel_bash` tool. The agent must explicitly choose to use this tool. Regular `bash` tool calls are always single-threaded.

---

## 6. Multi-Agent Status

### 6.1 MCP Server Configuration

**Source:** `internal/config/load.go:167-209`
```go
defaultMCPs := map[string]MCPConfig{
    "context-singularity-v2": {
        Type:     MCPStdio,
        Command:  "node",
        Args:     []string{"/Volumes/Storage/MCP/context-singularity-v2/dist/index.js"},
        Disabled: true,  // ← DISABLED
    },
    "hivemind-v2": {
        Type:     MCPStdio,
        Command:  "node",
        Args:     []string{"/Volumes/Storage/MCP/hivemind-v2/dist/index.js"},
        Disabled: true,  // ← DISABLED
    },
    "omega-v2": {
        Type:     MCPStdio,
        Command:  "node",
        Args:     []string{"/Volumes/Storage/MCP/omega-v2/dist/index.js"},
        Disabled: true,  // ← DISABLED
    },
    "pattern-crystallizer-v2": {
        Type:     MCPStdio,
        Command:  "node",
        Args:     []string{"/Volumes/Storage/MCP/pattern-crystallizer-v2/dist/index.js"},
        Disabled: true,  // ← DISABLED
    },
    // ... others are Disabled: false
}
```

**Finding:** All multi-agent MCP servers are **DISABLED by default**:
- `hivemind-v2` - Multi-agent coordination
- `omega-v2` - Meta-cognition
- `context-singularity-v2` - Context compression
- `pattern-crystallizer-v2` - Pattern extraction

### 6.2 Hivemind Capabilities (When Enabled)

**Source:** `/Volumes/Storage/MCP/hivemind-v2/src/index.ts:1-50`
```typescript
/**
 * Hivemind Orchestrator MCP Server V2
 *
 * SuperTool #3 - Multi-agent coordination via shared distributed_task_board
 *
 * Power Level: ⚡⚡⚡⚡⚡⚡⚡ (Integration-focused)
 *
 * This server shares storage with novel-concepts-server and uses:
 * - distributed_task_board for task coordination (not custom queue)
 * - SUPERCACHE reasoning tier for file locks and state
 * - concept_web_weaver for capability tracking
 */

interface Agent {
  id: string;
  name: string;
  type: string;
  capabilities: string[];
  status: 'idle' | 'busy' | 'offline';
  currentTasks: string[];
  completedTasks: number;
  averageScore: number;
  lastSeen: number;
}

interface Task {
  id: string;
  description: string;
  priority: number;
  estimatedEffort: number;
  requiredCapabilities: string[];
  state: 'pending' | 'ready' | 'claimed' | 'in_progress' | 'completed' | 'failed';
  dependencies: string[];
  claimedBy?: string;
  // ...
}
```

**Finding:** Hivemind provides:
- Agent registration and capability tracking
- Distributed task board
- Task claiming and completion tracking
- Inter-agent collaboration

**But:** It requires multiple Floyd instances to coordinate. A single instance gains no benefit.

### 6.3 Agent Library (Persona System)

**Source:** `internal/agents/loader.go:20-50`
```go
type AgentDefinition struct {
    Name        string   `yaml:"name"`        // Unique identifier
    Description string   `yaml:"description"` // Human-readable description
    Trigger     string   `yaml:"trigger,omitempty"` // Keyword to invoke
    Version     string   `yaml:"version,omitempty"`
    Author      string   `yaml:"author,omitempty"`
    Tags        []string `yaml:"tags,omitempty"`
    SystemPrompt string   `yaml:"-"` // Markdown body
    FilePath    string   `yaml:"-"`
}
```

**Source:** `internal/agents/` directory:
```
internal/agents/
├── _template.md           # Template for new agents
├── code-reviewer.md       # Code review persona
├── release-auditor.md     # Release audit persona
└── skill-and-agent-packager.md  # Packaging tool
```

**Finding:** Agent Library exists but provides **personas only** - not separate agent instances. These are prompt variations, not spawned sub-agents.

### 6.4 Agent Selector (Classification Only)

**Source:** `internal/agent/selector.go:1-50`
```go
type AgentSelector struct {
    agents         []agents.AgentDefinition
    agentsByTrigger map[string]*agents.AgentDefinition
    agentsByTag     map[string][]*agents.AgentDefinition
}

func (as *AgentSelector) ClassifyTask(prompt string) TaskClassification {
    lower := strings.ToLower(prompt)
    
    // Check for explicit trigger keywords first
    for trigger, agent := range as.agentsByTrigger {
        if strings.Contains(lower, trigger) {
            return TaskClassification{
                Type:           trigger,
                Confidence:     0.9,
                SuggestedAgent: agent.Name,
            }
        }
    }
    // ... pattern matching for task types
}
```

**Finding:** AgentSelector classifies tasks but does NOT spawn agents. It only suggests which persona's system prompt to use.

---

## 7. Prompt Template Analysis

### 7.1 Template Loading

**Source:** `internal/agent/coordinator.go:126-131`
```go
func NewCoordinator(...) (Coordinator, error) {
    // ...
    promptTemplate, err := coderPrompt(prompt.WithWorkingDir(c.cfg.WorkingDir()))
    if err != nil {
        return nil, err
    }
    // ...
}
```

**Source:** `internal/agent/coordinator.go:528-538`
```go
func coderPrompt(opts ...prompt.Option) (*prompt.Prompt, error) {
    embedded, err := templates.FS.ReadFile("coder.md.tpl")
    if err != nil {
        return nil, fmt.Errorf("reading template: %w", err)
    }
    return prompt.NewPrompt("coder", string(embedded), opts...)
}
```

**Finding:** ALL modes use `coder.md.tpl`. There is no mode-specific template selection.

### 7.2 Template Structure

**Source:** `internal/agent/templates/coder.md.tpl:1-16`
```markdown
You are **FLOYD** (File-Logged Orchestrator Yielding Deliverables), a production engineer agent.

## CRITICAL IDENTITY ANCHOR
- YOU ARE NOT CLAUDE. You are FLOYD v4.6.1.
- Protocol: FLOYD.md governs behavior. This template mirrors the deterministic edition.

## 0) POLICY PRECEDENCE (Highest → Lowest)
1. Tool/Hook Safety STOP
2. Bans (e.g., agentic_fetch)
3. Debug Hard-Gates (Hypothesis Gate, Two-Failure Reset, Prediction Rule, Circuit Breaker)
4. Rate Limits & Retry Budgets
5. SUPERCACHE Access Rules
6. Bias-for-Action
```

**Finding:** Template is coding-focused. No differentiation for research/documentation use cases.

### 7.3 FLOYD.md Protocol

**Source:** `FLOYD.md:1-50`
```markdown
FLOYD Persistent Agent Protocol v4.0 — Deterministic Edition (SUPERCACHE-First)

## 0. PRIME DIRECTIVE – IDENTITY ANCHOR
You are FLOYD (File-Logged Orchestrator Yielding Deliverables), not Claude, not Anthropic.

YOUR IDENTITY
- Name: FLOYD
- Version: v4.6.1
- Role: Production engineer agent for the FloydDeployable project
- Creator: douglastalley
- Home: /Volumes/Storage/floyd-sandbox/FloydDeployable/
```

**Finding:** FLOYD.md is project-specific and hardcoded to a single path. No mode awareness.

---

## 8. Data Directory Isolation

### 8.1 SuperFloyd Data Path

**Source:** `main.go:20-27`
```go
if strings.EqualFold(binName, "superfloyd") {
    if os.Getenv("FLOYD_GLOBAL_DATA") == "" {
        if homeDir, err := os.UserHomeDir(); err == nil {
            _ = os.Setenv("FLOYD_GLOBAL_DATA", filepath.Join(homeDir, ".superfloyd", "data"))
        }
    }
}
```

**Source:** `internal/cmd/superfloyd_resilience.go:78-80`
```go
if os.Getenv("FLOYD_GLOBAL_DATA") == "" {
    _ = os.Setenv("FLOYD_GLOBAL_DATA", "/Volumes/Storage/.floyd")
}
```

**Finding:** SuperFloyd can use a different data directory, but the second assignment overwrites the first, resulting in `/Volumes/Storage/.floyd` for all SuperFloyd modes.

---

## 9. Feature Flags Summary

| Variable | Default (Floyd) | Default (SuperFloyd) | Effect |
|----------|-----------------|----------------------|--------|
| `SUPERFLOYD_QUALITY_GATES` | User sets | `1` (forced) | Prompt validation |
| `SUPERFLOYD_DEGRADATION_CONTROLS` | User sets | `1` (forced) | Circuit breakers |
| `SUPERFLOYD_CONSISTENCY_LOCK` | User sets | `1` (forced) | Env fingerprinting |
| `SUPERFLOYD_AUTOSTABILIZE` | User sets | `1` (forced) | Auto-recovery |
| `SUPERFLOYD_MAX_PARALLEL` | (unset) | `6`/`12`/`16` | Parallel bash limit |
| `FLOYD_GLOBAL_DATA` | `~/.floyd/` | `/Volumes/Storage/.floyd` | Data directory |

---

## 10. Current Mode Comparison Matrix

| Feature | floyd | safe | balanced/sf | beast |
|---------|-------|------|-------------|-------|
| **Binary** | Same | Same | Same | Same |
| **Template** | coder.md.tpl | coder.md.tpl | coder.md.tpl | coder.md.tpl |
| **Protocol** | FLOYD.md | FLOYD.md | FLOYD.md | FLOYD.md |
| **Theme** | Default | SuperFloyd | SuperFloyd | SuperFloyd |
| **ASCII Art** | Simple | Elaborate | Elaborate | Elaborate |
| **Persistent Bar** | No | Yes | Yes | Yes |
| **Quality Gates** | Optional | ON | ON | ON |
| **Degradation Controls** | Optional | ON | ON | ON |
| **Consistency Lock** | Optional | ON | ON | ON |
| **Auto-Stabilize** | Optional | ON | ON | ON |
| **Max Parallel** | 4 | 6 | 12 | 16 |
| **Paranoia Checks** | Optional | ON | ON | ON |
| **Multi-Agent** | ❌ | ❌ | ❌ | ❌ |

---

## 11. Recommendations

### 11.1 Immediate: Separate Templates

Create mode-specific templates:

```
internal/agent/templates/
├── floyd-general.md.tpl      # Research/docs/web versatile
├── superfloyd-coder.md.tpl   # Strict coding agent
└── coder.md.tpl              # (keep as fallback)
```

Implementation in `coordinator.go`:
```go
func coderPrompt(opts ...prompt.Option) (*prompt.Prompt, error) {
    templateName := "coder.md.tpl"
    if isSuperFloydBinary() {
        templateName = "superfloyd-coder.md.tpl"
    } else {
        templateName = "floyd-general.md.tpl"
    }
    embedded, err := templates.FS.ReadFile(templateName)
    // ...
}
```

### 11.2 Medium: Tool Whitelisting

Add to `config.go`:
```go
type AgentMode string
const (
    AgentModeGeneral AgentMode = "general"
    AgentModeCoder   AgentMode = "coder"
)

var ModeConfigs = map[AgentMode]AgentConfig{
    AgentModeGeneral: {
        EnableWebResearch: true,
        EnableDocCreation: true,
        StrictDebugGate:   false,
    },
    AgentModeCoder: {
        DisabledTools:     []string{"web_search", "web_fetch"},
        EnableWebResearch: false,
        StrictDebugGate:   true,
    },
}
```

### 11.3 Future: Enable Hivemind

In `floyd.json`:
```json
{
  "mcp": {
    "hivemind-v2": {
      "disabled": false
    }
  }
}
```

This enables multi-agent coordination but requires multiple Floyd instances.

---

## 12. Version Synchronization Fix

**Current state:**
- Code: `v1.8` (`internal/version/version.go`)
- Prompts: `v4.6.1` (hardcoded in templates)

**Recommended fix:**
```go
// internal/version/version.go
var (
    Version   = "v1.8"
    AgentName = "FLOYD"
)

// Build with:
// go build -ldflags "-X internal/version.Version=v5.0 -X internal/version.AgentName=FLOYD" -o floyd .
// go build -ldflags "-X internal/version.Version=v1.8 -X internal/version.AgentName=SUPERFLOYD" -o superfloyd .
```

Then in templates:
```markdown
You are **{{.AgentName}}** v{{.Version}}
```

---

## Appendix: File References

| File | Lines | Purpose |
|------|-------|---------|
| `main.go` | 10-27 | Binary name detection, mode setup |
| `internal/cmd/superfloyd_resilience.go` | 27-105 | Safety systems, parallelism config |
| `internal/version/version.go` | 5 | Version constant |
| `internal/agent/templates/coder.md.tpl` | 1-100 | System prompt template |
| `FLOYD.md` | 1-100 | Persistent protocol file |
| `internal/agent/tools/parallel_bash.go` | 1-250 | Parallel execution tool |
| `internal/config/load.go` | 167-209 | MCP server defaults |
| `internal/ui/logo/logo.go` | 148-278 | ASCII art selection |
| `internal/ui/styles/themes.go` | 144-209 | Theme definitions |
| `internal/agents/loader.go` | 20-50 | Agent definition parsing |
| `internal/agent/selector.go` | 1-50 | Task classification |
| `/Volumes/Storage/MCP/hivemind-v2/src/index.ts` | 1-50 | Multi-agent coordination (disabled) |

---

*End of Analysis — 2026-03-13*
