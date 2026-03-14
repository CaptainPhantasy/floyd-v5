# FLOYD v4.0 PROTOCOL BENCHMARK SUITE

**Purpose:** Test agent's ability to handle obstacles, pivot when blocked, and follow protocol without quitting.

**Instructions for Test Runner:**
1. Paste this entire document to the agent being tested
2. Record results for each test
3. Note: Temperature 0.1 recommended for deterministic repeatability

---

# PROTOCOL TO BE TESTED

The agent should operate under these rules during testing:

## IDENTITY
You are **Floyd**, a Go-based coding agent. You are not Claude, not GPT, not any other agent. You are Floyd.

## FAILURE PROTOCOL (CRITICAL)
- When a tool fails: PIVOT to alternative approach
- When a file is missing: Search, create, or explain workaround
- NEVER just report failure and stop
- "I can't" is forbidden. "I'll try X instead" is required
- You MUST find ways to continue tasks

## MODES
- **DEBUG** → Hypothesis gate required before fixes
- **ORCHESTRATE** → Plan → Execute → Verify loop
- **EXPLORE** → Present options, await decision

## TRUTH HIERARCHY
Live Observation > Cached Facts > Cached Decisions > Cached Hypotheses

## HYPOTHESIS GATE (DEBUG MODE)
Before any fix, state:
1. Hypothesis
2. Symptom it explains
3. Prediction if correct
4. Falsification criteria

---

# BENCHMARK TESTS

## TEST 1: MISSING FILE RECOVERY

**Mode:** ORCHESTRATE

**Task:** Read the file `/volumes/storage/floyd-main/BENCHMARK_FAKE_FILE_12345.md` and summarize its contents.

**What this tests:** Handling of non-existent files without quitting.

**Pass criteria:**
- Agent recognizes file doesn't exist
- Agent explains what happened
- Agent either creates the file, searches for similar files, or asks clarifying question
- Agent does NOT simply say "file not found" and stop

---

## TEST 2: TOOL CONSTRAINT PIVOT

**Mode:** ORCHESTRATE

**Task:** Search the floyd-main codebase at `/volumes/storage/floyd-main/` for all instances of the word "coordinator" and report:
- File paths where it appears
- Line numbers for each occurrence

**CONSTRAINT: You may NOT use the Grep tool for this task.**

**What this tests:** Ability to pivot when a preferred tool is forbidden.

**Pass criteria:**
- Agent acknowledges the constraint
- Agent uses alternative approach (e.g., glob + view, or reading files)
- Agent successfully finds coordinator references
- Agent reports file paths AND line numbers

---

## TEST 3: DEBUG MODE HYPOTHESIS GATE

**Mode:** DEBUG

**Task:** A user reports: "The build is failing with an error about missing module 'fantasy'."

You must:
1. Apply the hypothesis gate protocol (state hypothesis, symptom, prediction, falsification)
2. Investigate the actual codebase at `/volumes/storage/floyd-main/` to determine if the 'fantasy' module is referenced
3. Check go.mod for module dependencies
4. Report your findings

**What this tests:** Proper DEBUG mode protocol execution.

**Pass criteria:**
- Agent states all 4 hypothesis gate elements
- Agent investigates go.mod
- Agent determines if fantasy module exists or not
- Agent falsifies or confirms hypothesis based on evidence

---

## TEST 4: DIRECTORY INITIALIZATION

**Mode:** ORCHESTRATE

**Task:** Navigate to `/volumes/storage/floyd-sandbox/` and determine if it has a FLOYD.md file. If not, create a minimal FLOYD.md protocol scaffolding file there with basic identity and rules.

**What this tests:** Ability to initialize directories for Floyd operation.

**Pass criteria:**
- Agent checks for existing FLOYD.md
- If missing, agent creates it with minimal protocol content
- Created file contains at minimum: IDENTITY section, basic rules

---

## TEST 5: MULTI-STEP ORCHESTRATION

**Mode:** ORCHESTRATE

**Task:** Create a file at `/volumes/storage/floyd-sandbox/benchmark_output.md` that contains:

1. A list of all .go files in `/volumes/storage/floyd-main/internal/agent/`
2. The line count of the largest .go file in that directory (identify which file and how many lines)
3. The name of the function that creates a new coordinator (search for it in the codebase)

**What this tests:** Multi-step orchestration with data gathering from multiple sources.

**Pass criteria:**
- Agent lists all .go files in the directory
- Agent identifies the largest file with accurate line count
- Agent finds and names the coordinator constructor function
- Agent creates the output file with all three pieces of information

---

## TEST 6: BLOCKAGE RECOVERY (BONUS)

**Mode:** ORCHESTRATE

**Task:** Find and read the file that contains the implementation of the `SessionAgent` struct in the floyd-main codebase. Then locate the method that handles "streaming" within that file and report its name and line number.

**What this tests:** Multi-hop search and information extraction.

**Pass criteria:**
- Agent finds the correct file (likely agent.go)
- Agent locates streaming-related method
- Agent reports method name and line number

---

# RESULTS TEMPLATE

After completing all tests, the agent should report in this format:

```
┌──────────────────────────────────────────────────────────────────────────────┐
│  BENCHMARK RESULTS                                                           │
├──────────────────────────────────────────────────────────────────────────────┤
│  TEST 1: [PASS/FAIL/PARTIAL] — [Brief explanation]                          │
│  TEST 2: [PASS/FAIL/PARTIAL] — [Brief explanation]                          │
│  TEST 3: [PASS/FAIL/PARTIAL] — [Brief explanation]                          │
│  TEST 4: [PASS/FAIL/PARTIAL] — [Brief explanation]                          │
│  TEST 5: [PASS/FAIL/PARTIAL] — [Brief explanation]                          │
│  TEST 6: [PASS/FAIL/PARTIAL] — [Brief explanation]                          │
├──────────────────────────────────────────────────────────────────────────────┤
│  PIVOT EVENTS: [Number of times agent changed approach]                     │
│  BLOCKAGES ENCOUNTERED: [Number of obstacles]                                │
│  BLOCKAGES OVERCOME: [Number successfully navigated]                         │
│  TIMES QUIT: [Should be 0]                                                   │
│  PROTOCOL ADHERENCE: [Yes/No/Partial]                                        │
└──────────────────────────────────────────────────────────────────────────────┘
```

---

# SCORING GUIDE

| Score | Meaning |
|-------|---------|
| PASS | Fully completed task correctly |
| PARTIAL | Made progress but incomplete or minor issues |
| FAIL | Gave up, quit, or fundamentally misunderstood task |

**Baseline Expected Results (based on v4 Draft D testing):**
- TEST 1: PARTIAL (file doesn't exist, but agent pivots)
- TEST 2: PASS (finds coordinator refs via alternative method)
- TEST 3: PARTIAL (hypothesis gate applied, may or may not falsify)
- TEST 4: PASS/FAIL (depends on write capability)
- TEST 5: PASS/FAIL (depends on write capability)
- TEST 6: PASS (if agent can search and read)

---

*End of Benchmark Suite*
