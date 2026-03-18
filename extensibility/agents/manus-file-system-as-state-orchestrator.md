---
name: MANUS File-System-as-State Orchestrator
description: Orchestrates a Manus-style file-system-as-state workspace, enforcing persistent file-based memory, self-correction, and logging
trigger: manus-init
version: 1.0.0
tags:
    - manus
    - file-system
    - state
    - orchestration
    - memory
    - persistence
category: architecture
---


I need you to upgrade this environment to use a "File-System-as-State" (Manus-style) architecture. We are moving away from ephemeral context to persistent file-based memory.

Execute the following 3 steps immediately:

### Step 1: Initialize the .manus Directory
Create a directory named .manus in the root. Inside, create these three files with the exact content below:

File 1: .manus/master_plan.md
```markdown
# Master Plan & Objectives
## Primary Goal
[User: Enter the main goal here, e.g., "Build a Voice Agent using ElevenLabs"]

## Strategic Steps
- [ ] Phase 1: Setup & Config
- [ ] Phase 2: Core Logic Implementation
- [ ] Phase 3: Testing & Refinement

## Context & Constraints
- Must use: [Tech Stack]
- Avoid: [Constraints]
```

File 2: .manus/scratchpad.md
```markdown
# Current Thinking & Scratchpad
*Use this file to store temporary code snippets, error logs, or research notes so we don't clog the chat context.*

## Current Focus
initializing...
```

File 3: .manus/progress.md
```markdown
# Execution Log
| Timestamp | Action Taken | Result/Status | Next Step |
|-----------|--------------|---------------|-----------|
| [Date]    | Init         | Ready         | Awaiting User Input |
```

### Step 2: Create the Agent "Self-Correction" Hook
Create a file named .manus/AGENT_INSTRUCTIONS.md. This is the file you (the agent) must read at the start of every prompt.

File Content for .manus/AGENT_INSTRUCTIONS.md:
```
CRITICAL INSTRUCTION SET FOR AGENT:

1. READ FIRST: Before answering any user request, you MUST read .manus/master_plan.md to ground yourself in the big picture.
2. UPDATE OFTEN: If you complete a sub-task, mark the checkbox in master_plan.md immediately using file manipulation.
3. LOG ACTIONS: After running a shell command or writing code, append a one-line summary to .manus/progress.md.
4. NO HALLUCINATION: If you need to store a variable, API key format, or documentation snippet, write it to .manus/scratchpad.md rather than keeping it in your context window.
5. SELF-CORRECTION: If a tool fails, log the error in scratchpad.md, explicitly state "I am updating the plan to fix this," and then edit master_plan.md with the new fix approach.
```

### Step 3: The Alias
If this is a Linux/Mac environment, generate a simple shell script ./init_manus.sh that the user can run to reset these files for a new project.

GO.

---

