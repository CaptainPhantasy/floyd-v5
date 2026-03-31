# TASK: Full Polish of 76 Core Skills to Production Quality

## MODE: BUILD — Autonomous Batch Rewrite

## OBJECTIVE
Rewrite every SKILL.md file in `./extensibility/skills/core/*/SKILL.md` to production quality. Each skill must be a self-contained runbook that an LLM agent can execute without external documentation.

## SOURCE OF TRUTH
Read `./COMPLETE_BUILDER_DELIVERY_PACKAGE_PATH` below for TypeScript interfaces. Extract the `Input` and `Output` interfaces for each skill. If a skill has no interface in the delivery package, infer appropriate actions and params from its description and category.

**Delivery package location:** `/Volumes/Storage/MCP/COMPLETE_BUILDER_DELIVERY_PACKAGE.md`

## GOLD STANDARD
Read `./extensibility/skills/core/ouroboros-self-evolution/SKILL.md` first. This is the quality bar. Every rewritten skill must match this level of specificity — concrete tool names, exact commands, step-by-step execution pipelines, and clear trigger conditions.

## REQUIRED FORMAT FOR EVERY SKILL

```markdown
---
name: {kebab-case-name}
description: {one-line description from delivery package or website}
category: core
version: "2.0.0"
---

# {Title Case Name}

> {description}

## When to Use
- WHEN mode={DEBUG|BUILD|EXPLORE} and {specific trigger condition}
- WHEN {second concrete trigger — not generic}
{one line per action explaining exactly when that action fires}

## Actions
`'{action1}' | '{action2}' | '{action3}'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
{real typed params with meaningful descriptions, not "No description"}

## Execution Pipeline

### Step 1: {verb phrase}
{exact instructions — what tool to call, what input to pass, what to check}

### Step 2: {verb phrase}
{next step — reference specific tool names or commands}

### Step N: {verb phrase}
{final step — what constitutes success, what output to return}

## Output Shape
{describe the structure of the returned JSON — key fields, types, what they mean}

## Failure Modes
- IF {condition}: {recovery action}
- IF {condition}: {escalation path}

## Examples
```json
{minimal example invocation with expected output skeleton}
```
```

## EXECUTION PLAN

Work in 4 passes. Do NOT attempt all 76 in a single pass.

### Pass 1: Fix the broken (12 skills)
These need immediate structural repair before polish:
- `debug` — 7-line empty stub. Rewrite as the canonical DEBUG MODE skill (hypothesis gate, post-fix protocol, two-failure reset). This is the most important skill in the system.
- `refactor` — 7-line empty stub. Rewrite as the canonical BUILD MODE refactoring skill.
- `patch-oracle` — old MCP invocation, "No description" params
- `viterbi-resolver` — old MCP invocation
- `refactor-pathfinder` — old MCP invocation
- `performance-profiling-patterns` — old MCP invocation
- `semantic-diff-validation` — old MCP invocation
- `tarjan-circular-dependency-detection` — old MCP invocation
- `statistical-performance-benchmarking` — old MCP invocation
- `code-review-workflow` — old MCP invocation
- `api-contract-validation` — old MCP invocation
- `consensus-algorithms` — old MCP invocation

After each file: verify frontmatter parses (has name, description, `---` delimiters). Do NOT proceed to Pass 2 until Pass 1 is verified.

### Pass 2: Enrich the hollow (11 skills)
These have correct structure but only a generic `'execute'` action and placeholder params:
- `bloom-sentinel`
- `token-alchemist`
- `grammar-gate`
- `clone-lens`
- `merge-engine`
- `consensus-voter`
- `concept-lattice`
- `mit-analysis`
- `analogy-synthesis`
- `test-generation-patterns`
- `api-format-verifier`

For each: infer real actions from the skill name and description. A `bloom-sentinel` should have `'check' | 'register' | 'reset'`. A `token-alchemist` should have `'compress' | 'analyze' | 'optimize'`. Write real params and an execution pipeline.

### Pass 3: Add execution pipelines to the 52 schema-backed skills
These already have real actions and params. Add:
- Execution Pipeline (step-by-step with tool names)
- Output Shape (from the delivery package `Output` interfaces)
- Failure Modes
- Example invocation

Read the delivery package to get the `Output` interface for each. The pipeline should reference Floyd harness tools where applicable (`view`, `edit`, `bash`, `grep`, `glob`, `write`).

### Pass 4: Polish ouroboros-self-evolution
The gold standard needs its frontmatter fixed (missing `category: core` and `version`). Add the standard sections it's missing (Actions, Parameters table, Output Shape, Failure Modes) while preserving its excellent execution pipeline.

## CONSTRAINTS
- NEVER write a "When to Use" that says "When you need this skill's capabilities" — that is circular and useless.
- NEVER write a param description that just restates the param name — "file_path: File path" is forbidden.
- EVERY execution pipeline step must name a specific tool or command.
- EVERY output shape must list at least 3 concrete fields the agent can expect.
- WHEN reading the delivery package: you are extracting `Output` interfaces, not re-reading `Input` (params are already done).
- Write findings to `.floyd/.supercache` after each pass completes.

## VERIFICATION
After all 4 passes:
1. Run: `find ./extensibility/skills/core -name "SKILL.md" | wc -l` — must be 76
2. Run: `find ./extensibility/skills/core -name "SKILL.md" -exec grep -L "## Execution Pipeline" {} \;` — must return 0 files
3. Run: `find ./extensibility/skills/core -name "SKILL.md" -exec grep -L "## Output Shape" {} \;` — must return 0 files
4. Run: `find ./extensibility/skills/core -name "SKILL.md" -exec grep -L "## Failure Modes" {} \;` — must return 0 files
5. Confirm no skill contains "No description" or "Use via Floyd Labs MCP"

Report the verification results. If any check fails, fix the failures before reporting complete.
