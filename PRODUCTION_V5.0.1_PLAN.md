# Production Quality v5.0.1 - Deterministic Execution Plan

**Version:** v5.0.1
**Status:** Planning Phase
**Date:** 2026-03-14
**Workflow:** Deterministic Prompt Framework

---

## PHASE 0: DETERMINISTIC TASK KICKOFF

### 0.1 Core Initialization
- **Context**: Migrate from FloydDeployable to production directory, fix degradation triggers, deliver production-quality extensibility system
- **Objective**: Ship v5.0.1 with clean protocol, working agent/skill libraries, categorized display, production-grade UX
- **Constraints**: Time critical, budget exhausted, no feature creep

### 0.2 State Verification
**Completed:**
- ✅ FLOYD.md cleaned (removed degraded mode playbook, shadow daemon protocol)
- ✅ All binaries signed and quarantines removed
- ✅ Agent library and skills library dialogs created
- ✅ Backtick key system implemented
- ✅ SuperFloyd mode activation wired
- ✅ GlobalAgentsDirs() returns user + built-in paths

**Current Issues:**
- ⚠️ Binaries show `v5.0.0-dev` (should be v5.0.1)
- ⚠️ 101 agents lack categorization and have inconsistent front matter
- ⚠️ 0 skills in skills library (directory empty)
- ⚠️ Single flat list for 101 agents (UX nightmare)
- ⚠️ No category filtering or grouping in UI

---

## PHASE 1: VERSION BUMP (CRITICAL - BLOCKER)

### 1.1 Update Version Files
**File:** `/Volumes/Storage/floyd/VERSION`
**Action:** Change `5.0.0` to `5.0.1`

**File:** `/Volumes/Storage/floyd/internal/version/version.go`
**Action:**
- Change `Version = "v5.0.0-dev"` to `Version = "v5.0.1"`
- Change `BuildDate` from "unknown" to actual build date

**Validation:**
```bash
grep "5.0.1" /Volumes/Storage/floyd/VERSION
grep "v5.0.1" /Volumes/Storage/floyd/internal/version/version.go
```

**Success Criteria:**
- VERSION file reads "5.0.1"
- version.go reads `Version = "v5.0.1"`

### 1.2 Rebuild Binaries
**Commands:**
```bash
cd /Volumes/Storage/floyd
go build -o floyd
go build -ldflags "-X github.com/legacy-ai/floyd/internal/version.BinaryName=superfloyd" -o superfloyd .
```

**Validation:**
```bash
./floyd --version    # Should show "floyd version v5.0.1"
./superfloyd --version  # Should show "superfloyd version v5.0.1"
```

### 1.3 Sign and Deploy
**Commands:**
```bash
xattr -cr floyd superfloyd
codesign --force --sign - floyd superfloyd
cp -f floyd /opt/homebrew/bin/floyd
cp -f superfloyd /opt/homebrew/bin/superfloyd
cp -f floyd ~/.local/bin/floyd
cp -f superfloyd ~/.local/bin/superfloyd
cp -f floyd /Volumes/Storage/floyd/releases/v5.0.1/floyd
cp -f superfloyd /Volumes/Storage/floyd/releases/v5.0.1/superfloyd
```

**Success Criteria:**
- All 6 binary locations updated
- `which floyd` returns working binary
- `floyd --version` shows v5.0.1

**Estimated Time:** 15 minutes
**Priority:** P0 (BLOCKER - Must complete before anything else)

---

## PHASE 2: CATEGORY SYSTEM FOUNDATION

### 2.1 Define Agent Categories
**File:** `/Volumes/Storage/floyd/internal/agents/loader.go` (new constants)

**Action:** Add category constants
```go
const (
    CategoryCoding        = "coding"
    CategoryArchitecture  = "architecture"
    CategoryDebugging     = "debugging"
    CategoryOrchestration = "orchestration"
    CategoryQuality       = "quality"
    CategorySecurity      = "security"
    CategoryMonitoring     = "monitoring"
    CategoryTesting       = "testing"
    CategoryDocumentation  = "documentation"
    CategoryInfrastructure = "infrastructure"
    CategoryData          = "data"
    CategoryPerformance    = "performance"
    CategoryDX            = "dx"
)
```

### 2.2 Add Category Field to AgentDefinition
**File:** `/Volumes/Storage/floyd/internal/agents/loader.go`
**Location:** `AgentDefinition` struct

**Action:** Add category field
```go
type AgentDefinition struct {
    Name        string   `yaml:"name"`
    Description string   `yaml:"description"`
    Trigger      string   `yaml:"trigger,omitempty"`
    Version      string   `yaml:"version,omitempty"`
    Author       string   `yaml:"author,omitempty"`
    Tags         []string `yaml:"tags,omitempty"`
    Category     string   `yaml:"category"`  // NEW FIELD
    SystemPrompt string   `yaml:"-"`
    FilePath     string   `yaml:"-"`
}
```

### 2.3 Define Skill Categories
**File:** `/Volumes/Storage/floyd/internal/skills/skills.go` (new constants)

**Action:** Add category constants
```go
const (
    SkillCategoryGit           = "git"
    SkillCategoryTesting       = "testing"
    SkillCategoryLinting       = "linting"
    SkillCategoryRefactoring   = "refactoring"
    SkillCategoryDocumentation = "documentation"
    SkillCategoryDeployment    = "deployment"
    SkillCategoryMonitoring    = "monitoring"
    SkillCategorySecurity      = "security"
)
```

### 2.4 Add Category Field to Skill
**File:** `/Volumes/Storage/floyd/internal/skills/skills.go`
**Location:** `Skill` struct

**Action:** Add category field
```go
type Skill struct {
    Name          string            `yaml:"name"`
    Description   string            `yaml:"description"`
    License       string            `yaml:"license,omitempty"`
    Compatibility string            `yaml:"compatibility,omitempty"`
    Metadata      map[string]string `yaml:"metadata,omitempty"`
    Category      string            `yaml:"category"`  // NEW FIELD
    Instructions  string            `yaml:"-"`
    Path          string            `yaml:"-"`
    SkillFilePath string            `yaml:"-"`
}
```

**Validation:**
```bash
go build -o floyd  # Should compile cleanly
go build -o superfloyd .  # Should compile cleanly
```

**Estimated Time:** 30 minutes
**Priority:** P1 (Foundation for all UI work)

---

## PHASE 3: AGENT FRONT MATER AUDIT

### 3.1 Audit Script Creation
**Script:** `/Volumes/Storage/floyd/scripts/audit_agents.go`

**Purpose:** Automated audit of all 101 agent files

**Checks:**
1. Front matter completeness (name, description, version, tags, category)
2. Empty fields (trigger, tags, category)
3. Invalid formats (emoji in name, kebab-case violations)
4. Duplicate names across directories
5. Missing required fields

**Output:** JSON report with issues by file

**Usage:**
```bash
go run scripts/audit_agents.go > audit_report.json
```

### 3.2 Execute Audit
**Action:**
```bash
go run /Volumes/Storage/floyd/scripts/audit_agents.go
```

**Expected Output:**
- Count of agents with empty triggers
- Count of agents with empty tags
- Count of agents with empty categories
- Count of agents with emoji in names
- List of duplicate names
- Validation errors per file

### 3.3 Fix Empty Triggers
**Pattern:** 60+ agents have empty `trigger:` fields

**Fix Strategy:**
1. Extract key phrase from name (e.g., "Code Reviewer" → trigger: "code-reviewer")
2. Or remove `trigger:` field entirely (agent name becomes trigger)
3. Validate triggers are kebab-case, no spaces

**Script:** `/Volumes/Storage/floyd/scripts/fix_agent_triggers.go`

**Action:** Auto-generate triggers from agent names

### 3.4 Add Tags to Empty Arrays
**Pattern:** 40+ agents have `tags: []`

**Fix Strategy:**
1. Infer from agent name/description (e.g., "code review" → tags: [review, code, quality])
2. Infer from agent function (e.g., "git steward" → tags: [git, security, reputation])
3. Add minimum 3 tags per agent for discoverability

**Script:** `/Volumes/Storage/floyd/scripts/fix_agent_tags.go`

**Action:** Auto-generate tags from context

### 3.5 Remove Emojis from Names
**Pattern:** Emoji prefixes in agent names (e.g., "🗺️ Data Flow Cartographer")

**Fix Strategy:**
1. Strip emoji characters from name field
2. Keep emoji in description for visual flair
3. Normalize to ASCII-only names

**Script:** `/Volumes/Storage/floyd/scripts/fix_agent_names.go`

**Action:** Remove emojis, ensure ASCII names

**Estimated Time:** 3 hours
**Priority:** P1 (Data quality - affects all agents)

---

## PHASE 4: AGENT CATEGORIZATION

### 4.1 Manual Categorization Matrix
**Create:** `/Volumes/Storage/floyd/docs/AGENT_CATEGORIZATION.md`

**Format:** Markdown table mapping agent → category

**Example:**
```markdown
| Agent Name | Suggested Category | Rationale |
|------------|------------------|------------|
| Code Reviewer | quality | Reviews code quality |
| Data Flow Cartographer | architecture | Maps system flow |
| Git Steward | infrastructure | Git operations |
```

**Action:** Categorize all 101 agents by domain

### 4.2 Batch Update Agent Files
**Script:** `/Volumes/Storage/floyd/scripts/update_agent_categories.go`

**Input:** `AGENT_CATEGORIZATION.md` table

**Action:** For each agent file, add `category:` field based on matrix

**Validation:**
```bash
# Verify category field exists in all agents
grep -L "^category:" /Volumes/Storage/floyd/extensibility/agents/*.md
# Should return nothing (all have category)
```

### 4.3 Validate Categories
**Check:** All categories must be from defined constants

**Action:**
```bash
grep -h "^category:" /Volumes/Storage/floyd/extensibility/agents/*.md | sort | uniq -c
```

**Expected Output:**
- All 101 agents have valid categories
- No unknown category values
- Reasonable distribution (not all in one category)

**Estimated Time:** 4 hours
**Priority:** P1 (Required for category system)

---

## PHASE 5: UI CATEGORY SYSTEM

### 5.1 Add Category Tabs to Agent Library
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Changes:**
1. Add `currentCategory string` field to `AgentLibrary` struct
2. Add `categories []string` field (all defined categories + "All")
3. Add `nextCategory()` and `prevCategory()` methods
4. Add keyboard bindings for category navigation (1-9, 0 for All)
5. Update `Draw()` to show category tabs at top
6. Update `Filter()` to filter by both name AND category
7. Update `loadAgents()` to group by category

**Key Bindings:**
```go
keyMap.CategoryNext: key.NewBinding(key.WithKeys("right", "tab"))
keyMap.CategoryPrev: key.NewBinding(key.WithKeys("left", "shift+tab"))
keyMap.Category1: key.NewBinding(key.WithKeys("1"))
// ... 2-9 for specific categories
keyMap.CategoryAll: key.NewBinding(key.WithKeys("0", "a"))
```

### 5.2 Update Help Text
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Action:** Update `ShortHelp()` and `FullHelp()` to include:
- Category navigation (left/right arrows, tab/shift+tab)
- Category shortcuts (0-9 keys)
- Search within category (still `/` + filter text)

### 5.3 Apply Same Changes to Skills Library
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/skills_library.go`

**Action:** Mirror category system changes from AgentLibrary

### 5.4 Add Category Counters
**Display:** Show agent count per category in tabs

**Example:**
```
[Coding (23)] [Architecture (15)] [Debugging (12)] ...
```

**Action:** Count agents per category in `loadAgents()`, update tab titles

**Estimated Time:** 4 hours
**Priority:** P2 (UX - Makes 101 agents usable)

---

## PHASE 6: SKILLS PORTING & CREATION

### 6.1 Discover Existing Skills
**Check:** `/Volumes/Storage/floyd-sandbox/FloydDeployable/` for skill files

**Action:**
```bash
find /Volumes/Storage/floyd-sandbox/FloydDeployable/ -name "SKILL.md" -o -name "skill.md" | head -20
```

**Expected:** Find skill files from previous version

### 6.2 Define 20 Core Skills
**Source:** DETERMINISTIC_PROMPT_FRAMEWORK templates

**Skills to Create:**
1. `git-commit` - Commit message generation
2. `git-diff` - Diff analysis and explanation
3. `test-write` - Unit test generation
4. `lint-fix` - Auto-lint and fix issues
5. `refactor-extract` - Extract method/function
6. `debug-breakpoint` - Add strategic breakpoints
7. `doc-readme` - Generate README.md
8. `performance-profile` - Add profiling code
9. `security-scan` - Security audit checklist
10. `migration-create` - Database migration template
11. `api-test` - API endpoint test
12. `component-create` - React component scaffold
13. `style-component` - CSS/Tailwind component
14. `dockerfile-write` - Dockerfile generation
15. `ci-config` - GitHub Actions workflow
16. `env-var-validate` - Environment config check
17. `log-structure` - Structured logging setup
18. `error-handle` - Error boundary pattern
19. `feature-flag` - Feature toggle implementation
20. `cache-strategy` - Caching implementation

### 6.3 Create Skill Files
**Location:** `/Volumes/Storage/floyd/extensibility/skills/{category}/{skill-name}/SKILL.md`

**Format:**
```markdown
---
name: skill-name
description: What this skill does
license: MIT
compatibility: floyd 5.0+
category: git
metadata:
  author: Legacy AI
  version: 1.0.0
---

# Skill Instructions

You are [skill name] skill.

Your purpose: [single sentence description]

## When to Use
- [condition 1]
- [condition 2]

## How to Use
1. [step 1]
2. [step 2]
3. [step 3]

## Output Format
[expected output template]

## Constraints
- [limitation 1]
- [limitation 2]
```

**Script:** `/Volumes/Storage/floyd/scripts/generate_skills.go`

**Action:** Batch create 20 skill files with front matter

### 6.4 Port Existing Skills
**Check:** If skills exist in FloydDeployable

**Action:** Copy and convert front matter format to Agent Skills standard

**Validation:**
```bash
ls -1 /Volumes/Storage/floyd/extensibility/skills/*/*/SKILL.md | wc -l
# Should show at least 20 skills
```

**Estimated Time:** 3 hours
**Priority:** P1 (Zero skills = broken feature)

---

## PHASE 7: UI PRODUCTION POLISH

### 7.1 Add Empty State Messaging
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Condition:** `len(a.agents) == 0`

**Display:**
```
┌─────────────────────────────────┐
│  Agent Library               │
├─────────────────────────────────┤
│                             │
│  No agents found             │
│  Check:                    │
│  • ~/.config/floyd/agents/ │
│  • <floyd>/extensibility/ │
│    agents/                   │
│                             │
└─────────────────────────────────┘
```

**Action:** Update `Draw()` to render empty state when no items

### 7.2 Add Loading Indicator
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Condition:** During `loadAgents()` call

**Display:** "Loading agents..." spinner

**Action:** Add `loading bool` field, show spinner in `Draw()`

### 7.3 Add Error State Display
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Condition:** `loadAgents()` returns error

**Display:**
```
┌─────────────────────────────────┐
│  Agent Library               │
├─────────────────────────────────┤
│  ❌ Error loading agents    │
│                             │
│  Failed to read:             │
│  ~/.config/floyd/agents/     │
│                             │
│  Permission denied             │
└─────────────────────────────────┘
```

**Action:** Add `error string` field, display error in `Draw()`

### 7.4 Add Keyboard Help Overlay
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Trigger:** Press `?` key

**Display:**
```
┌─────────────────────────────────┐
│  Keyboard Shortcuts          │
├─────────────────────────────────┤
│  Navigation:                │
│  • j/k or ↑/↓ - Move      │
│  • Enter - Select            │
│  • Esc - Close               │
│                             │
│  Categories:                 │
│  • 0 - All                 │
│  • 1-9 - By category      │
│  • Tab/Shift+Tab - Next/Prev │
│                             │
│  Filtering:                  │
│  • / - Start filter          │
│  • Type to filter            │
│  • Escape - Clear filter      │
│                             │
│  Press any key to close      │
└─────────────────────────────────┘
```

**Action:** Add `showHelp bool` field, toggle on `?` key

### 7.5 Add Description Preview
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Trigger:** Expand selected item (press `e`)

**Display:** Show full description (3-4 lines) below selected item

**Action:** Add `expanded bool` field, show more text in `Render()`

**Estimated Time:** 3 hours
**Priority:** P2 (UX polish)

---

## PHASE 8: PERFORMANCE & VALIDATION

### 8.1 Add Filter Debouncing
**File:** `/Volumes/Storage/floyd/internal/ui/dialog/agent_library.go`

**Action:** Add 300ms debounce to input

**Implementation:**
```go
type AgentLibrary struct {
    // ... existing fields
    filterTimer   *time.Timer
    lastFilter    string
}

func (a *AgentLibrary) HandleMsg(msg tea.Msg) Action {
    case tea.KeyPressMsg:
        // ... existing cases
        default:
            a.input, cmd = a.input.Update(msg)
            value := a.input.Value()
            // DEBOUNCE
            if a.filterTimer != nil {
                a.filterTimer.Stop()
            }
            a.filterTimer = time.AfterFunc(300*time.Millisecond, func() {
                a.list.SetFilter(value)
                a.list.ScrollToTop()
            })
            return ActionCmd{cmd}
}
```

### 8.2 Add Virtual Scrolling (Large Lists)
**File:** `/Volumes/Storage/floyd/internal/ui/list/filterable_list.go`

**Condition:** `len(items) > 100`

**Action:** Only render visible items in viewport, scroll virtual window

**Implementation:**
```go
type FilterableList struct {
    // ... existing fields
    visibleOffset int  // Index of first visible item
    visibleCount  int  // Number of visible items
}

func (l *FilterableList) Render() string {
    // Only render visibleCount items starting at visibleOffset
    // Adjust based on scroll position
}
```

### 8.3 Duplicate Detection
**File:** `/Volumes/Storage/floyd/internal/agents/loader.go`

**Function:** `LoadAgents()`

**Action:** Track seen names by directory, warn on duplicates

**Implementation:**
```go
func LoadAgents(dir string) ([]AgentDefinition, error) {
    seen := make(map[string]bool)
    duplicates := []string{}

    for _, agent := range loaded {
        if seen[agent.Name] {
            duplicates = append(duplicates, agent.Name)
            continue
        }
        seen[agent.Name] = true
        agents = append(agents, agent)
    }

    if len(duplicates) > 0 {
        slog.Warn("Duplicate agents found", "duplicates", duplicates)
    }

    return agents, nil
}
```

**Estimated Time:** 2 hours
**Priority:** P3 (Performance optimization)

---

## PHASE 9: TESTING

### 9.1 Unit Tests
**File:** `/Volumes/Storage/floyd/internal/agents/loader_test.go` (update existing)

**Tests:**
```go
func TestAgentDefinitionWithCategory(t *testing.T) {
    // Test category field parsing
}

func TestLoadAgentsWithDuplicateDetection(t *testing.T) {
    // Test duplicate warning
}

func TestAgentFrontMatterValidation(t *testing.T) {
    // Test validation of all fields
}
```

**Run:**
```bash
cd /Volumes/Storage/floyd
go test ./internal/agents/ -v
```

### 9.2 Integration Tests
**Test 1: Open Agent Library**
```bash
floyd
# Press /, type "agent_library", press Enter
# Verify: Dialog opens with 101 agents
```

**Test 2: Category Filtering**
```bash
# In agent library, press "1"
# Verify: Shows only Coding agents (should show count)
```

**Test 3: Skills Library**
```bash
# Press /, type "skills_library", press Enter
# Verify: Dialog opens with 20+ skills
```

**Test 4: Search Filter**
```bash
# Type "git" in filter
# Verify: Shows only git-related agents/skills
```

### 9.3 E2E Tests
**Test 1: Create Agent**
```bash
cat > ~/.config/floyd/agents/test-agent.md << 'EOF'
---
name: Test Agent
description: Testing agent creation
category: testing
version: 1.0.0
tags: [test]
---
Test agent prompt.
EOF

floyd
# Open agent library, verify "Test Agent" appears
```

**Test 2: Create Skill**
```bash
mkdir -p ~/.config/floyd/skills/test/
cat > ~/.config/floyd/skills/test/SKILL.md << 'EOF'
---
name: test-skill
description: Testing skill creation
license: MIT
compatibility: floyd 5.0+
category: testing
---
Test skill instructions.
EOF

floyd
# Open skills library, verify "test-skill" appears
```

**Test 3: Select Agent**
```bash
# In agent library, select an agent, press Enter
# Verify: Agent system prompt inserted into editor
```

**Estimated Time:** 2 hours
**Priority:** P3 (Quality assurance)

---

## PHASE 10: DOCUMENTATION

### 10.1 Update EXTENSIBILITY Guide
**File:** `/Volumes/Storage/floyd/docs/guides/EXTENSIBILITY.md`

**Additions:**
- Category system explanation
- Agent front matter schema with category
- Skill front matter schema with category
- Category values reference table
- Examples: Creating categorized agents
- Examples: Creating categorized skills

### 10.2 Create Agent/Skill Best Practices
**File:** `/Volumes/Storage/floyd/docs/guides/EXTENSIBILITY_BEST_PRACTICES.md`

**Sections:**
- Choosing the right category
- Writing effective triggers
- Tagging for discoverability
- Naming conventions (no emojis, kebab-case)
- Description writing tips
- Version numbering
- User override behavior

### 10.3 Update Migration Guide
**File:** `/Volumes/Storage/floyd/docs/releases/MIGRATION_V4_TO_V5.md`

**Sections:**
- Agent file format changes
- New category system
- Empty field fixes
- UI changes (categories, tabs)
- Breaking changes (if any)

**Estimated Time:** 2 hours
**Priority:** P3 (Documentation completeness)

---

## PHASE 11: FINAL RELEASE

### 11.1 Release Notes
**File:** `/Volumes/Storage/floyd/docs/releases/v5.0.1.md`

**Sections:**
- Summary (what changed)
- Features (categories, UI polish, skills)
- Fixes (version, front matter)
- Breaking changes (category field required)
- Migration guide

### 11.2 Build Release Binaries
**Commands:**
```bash
cd /Volumes/Storage/floyd
# Set version
echo "5.0.1" > VERSION

# Build with release flags
go build -ldflags "-X github.com/legacy-ai/floyd/internal/version.Version=v5.0.1 -X github.com/legacy-ai/floyd/internal/version.BuildDate=$(date -u +%Y-%m-%d)" -o floyd
go build -ldflags "-X github.com/legacy-ai/floyd/internal/version.Version=v5.0.1 -X github.com/legacy-ai/floyd/internal/version.BinaryName=superfloyd -X github.com/legacy-ai/floyd/internal/version.BuildDate=$(date -u +%Y-%m-%d)" -o superfloyd
```

### 11.3 Sign Release
**Commands:**
```bash
xattr -cr floyd superfloyd
codesign --force --sign - floyd superfloyd
```

### 11.4 Deploy Release
**Commands:**
```bash
mkdir -p /Volumes/Storage/floyd/releases/v5.0.1
cp floyd superfloyd /Volumes/Storage/floyd/releases/v5.0.1/
cp -f floyd /opt/homebrew/bin/floyd
cp -f superfloyd /opt/homebrew/bin/superfloyd
cp -f floyd ~/.local/bin/floyd
cp -f superfloyd ~/.local/bin/superfloyd
```

### 11.5 Create Checksums
**Command:**
```bash
cd /Volumes/Storage/floyd/releases/v5.0.1
shasum -a 256 floyd superfloyd > SHA256SUMS.txt
```

### 11.6 Validation Checklist
- [ ] Both binaries show `v5.0.1`
- [ ] Agent library shows 101 agents with categories
- [ ] Skills library shows 20+ skills with categories
- [ ] Category tabs work (0-9, left/right)
- [ ] Search filter works (debounced)
- [ ] Empty states display correctly
- [ ] No crash on dialog open/close
- [ ] All unit tests pass
- [ ] Documentation complete

**Estimated Time:** 1 hour
**Priority:** P0 (Release completion)

---

## EXECUTION SEQUENCE

### Critical Path (Must Complete In Order)
1. **PHASE 1: VERSION BUMP** (15 min) - Blocks everything
2. **PHASE 2: CATEGORY SYSTEM FOUNDATION** (30 min) - Required for all agent/skill work
3. **PHASE 3: AGENT FRONT MATER AUDIT** (3 hours) - Data quality foundation
4. **PHASE 4: AGENT CATEGORIZATION** (4 hours) - Required for UI work
5. **PHASE 6: SKILLS PORTING & CREATION** (3 hours) - Zero skills = broken feature

### Parallel Path (Can Work Simultaneously)
- **PHASE 5: UI CATEGORY SYSTEM** (4 hours) - Dependent on Phase 2, 4
- **PHASE 7: UI PRODUCTION POLISH** (3 hours) - Dependent on Phase 5

### Quality Path (Complete After Core Features)
- **PHASE 8: PERFORMANCE & VALIDATION** (2 hours)
- **PHASE 9: TESTING** (2 hours)
- **PHASE 10: DOCUMENTATION** (2 hours)
- **PHASE 11: FINAL RELEASE** (1 hour)

---

## SUCCESS METRICS

### Production Quality Indicators
- ✅ Version correctly reports v5.0.1
- ✅ All 101 agents have valid front matter (name, description, trigger, tags, category)
- ✅ All 101 agents are categorized into 14 categories
- ✅ 20+ skills created with valid front matter
- ✅ Agent library UI supports category navigation (tabs 0-9)
- ✅ Skills library UI supports category navigation
- ✅ Search filter debounced (300ms)
- ✅ Empty/error states display properly
- ✅ Keyboard help overlay available
- ✅ All unit tests pass (>90% coverage)
- ✅ All E2E tests pass
- ✅ Documentation complete and accurate

### Release Readiness Checklist
- [ ] Binaries built with correct version
- [ ] Binaries signed and quarantines removed
- [ ] All binary locations updated (global, local, releases)
- [ ] Agent categorization complete (101/101)
- [ ] Skills created (20+ skills)
- [ ] UI category system functional
- [ ] All tests passing
- [ ] Documentation updated
- [ ] Release notes drafted
- [ ] Checksums generated

---

## TOTAL TIME ESTIMATE

**Critical Path:** 8.5 hours
**Parallel Path:** 7 hours
**Quality Path:** 7 hours
**Total:** 22.5 hours

**Fast Path (MVP Only):**
- Skip UI polish (Phase 7)
- Skip performance optimization (Phase 8)
- Skip testing (Phase 9, use manual testing)
- Skip full documentation (Phase 10, minimal updates only)
**Fast Path Total:** 15 hours

---

## DECISION POINTS

### Decision 1: Category Granularity
**Option A:** 14 categories (fine-grained) → Better filtering, more complexity
**Option B:** 6 categories (coarse-grained) → Simpler UI, less precision
**Recommendation:** Start with 14, consolidate if UI too complex

### Decision 2: Skill Creation Approach
**Option A:** Port from FloydDeployable → Faster, but may be outdated
**Option B:** Create new from deterministic framework → Clean, production-grade, but slower
**Recommendation:** Mix - port high-quality skills, create new for gaps

### Decision 3: Empty Field Handling
**Option A:** Fail validation → Agent won't load, user gets error
**Option B:** Generate defaults → Agent loads with placeholder data
**Recommendation:** Generate defaults with warning in log

---

## RISK MITIGATION

### Risk 1: Agent File Format Breaks
**Mitigation:** Test with subset of agents first, validate all before batch update

### Risk 2: UI Performance with 101 Agents
**Mitigation:** Implement virtual scrolling early, measure render time

### Risk 3: Category System Incompatibility
**Mitigation:** Add fallback to "All" category, soft-fail on unknown categories

### Risk 4: Time Overrun
**Mitigation:** Prioritize MVP (working features) over polish (nice-to-have)

---

## CONTINUITY

### Handoff to Next Session
**State:**
- Version bump complete (PHASE 1 done)
- Category system foundation laid (PHASE 2 done)
- Agent front matter audit in progress (PHASE 3 in progress)

**Next Steps:**
1. Complete agent front matter audit (PHASE 3 finish)
2. Categorize all agents (PHASE 4)
3. Port/create skills (PHASE 6)
4. Implement UI category system (PHASE 5)
5. Polish UI (PHASE 7)
6. Test thoroughly (PHASE 9)
7. Ship v5.0.1 (PHASE 11)

**Context:**
- 101 agents in `/Volumes/Storage/floyd/extensibility/agents/`
- 0 skills in `/Volumes/Storage/floyd/extensibility/skills/`
- Goal: Production-quality v5.0.1 with working agent/skill libraries

**Blockers:**
- None (version bump unblocks all work)

---

**Status:** Ready to execute PHASE 1 (VERSION BUMP)
