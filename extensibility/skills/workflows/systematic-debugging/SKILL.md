---
name: systematic-debugging
description: Use when encountering any bug, test failure, or unexpected behavior, before proposing fixes - four-phase framework (root cause investigation, pattern analysis, hypothesis testing, implementation) that ensures understanding before attempting solutions
---

# Systematic Debugging

## Overview

Random fixes waste time and create new bugs. Quick patches mask underlying issues.

**Core principle:** Always find **root cause** before attempting fixes. Symptom fixes are failure.

Violating the letter of this process is violating the spirit of debugging.

## The Iron Law

> **No fixes without root cause investigation first.**

If you have not completed Phase 1, you cannot propose fixes.

## When to Use

Use for **any** technical issue:
- Test failures
- Bugs in production
- Unexpected behavior
- Performance problems
- Build failures
- Integration issues

Especially when:
- You are under time pressure
- "Just one quick fix" seems obvious
- You have already tried multiple fixes
- A previous fix did not work
- You do not fully understand the issue

## The Four Phases

You must complete each phase before proceeding to the next.

### Phase 1: Root Cause Investigation

Before attempting any fix:

**1. Read Error Messages Carefully**
- Do not skip past errors or warnings
- Read stack traces completely
- Note line numbers, file paths, error codes

**2. Reproduce Consistently**
- Can you trigger it reliably?
- What are the exact steps?
- Does it happen every time?
- If not reproducible → gather more data, do not guess

**3. Check Recent Changes**
- What changed that could cause this?
- Git diff, recent commits
- New dependencies, config changes
- Environmental differences

**4. Gather Evidence in Multi-Component Systems**

When the system has multiple components (CI → build → signing, API → service → database), add diagnostic instrumentation **before** proposing fixes.

For each component boundary:
- Log what data enters the component
- Log what data exits the component
- Verify environment/config propagation
- Check state at each layer

Run once to gather evidence showing **where** it breaks. Then:
- Analyze evidence to identify failing component
- Investigate that specific component

**5. Trace Data Flow**

When the error is deep in the call stack, trace backwards:
- Where does the bad value originate?
- What called this with the bad value?
- Keep tracing up until you find the source
- Fix at the source, not at the symptom

### Phase 2: Pattern Analysis

Find the pattern before fixing:

**1. Find Working Examples**
- Locate similar working code in the same codebase
- Ask: what works that is similar to what is broken?

**2. Compare Against References**
- If implementing a pattern, read the reference implementation completely
- Do not skim
- Understand the pattern fully before applying

**3. Identify Differences**
- What is different between working and broken?
- List every difference, however small

**4. Understand Dependencies**
- What other components does this need?
- What settings, config, environment?
- What assumptions does it make?

### Phase 3: Hypothesis and Testing

Apply the scientific method:

**1. Form a Single Hypothesis**
- "I think X is the root cause because Y"
- Be specific, not vague

**2. Test Minimally**
- Make the smallest possible change to test the hypothesis
- One variable at a time

**3. Verify Before Continuing**
- Did it work? Yes → Phase 4
- Did it fail? Form a new hypothesis
- Do not stack multiple speculative fixes

**4. When You Do Not Know**
- Say "I do not understand X"
- Ask for help
- Research more

### Phase 4: Implementation

Fix the root cause, not the symptom.

**1. Create a Failing Test Case**
- Simplest possible reproduction
- Automated test if possible
- One-off script if no framework

**2. Implement a Single Fix**
- Address the root cause identified
- One change at a time
- No "while I am here" refactors

**3. Verify Fix**
- New test passes
- Existing tests still pass
- The issue is actually resolved

**4. If Fix Does Not Work**
- Count how many fixes you have tried
- If < 3 → return to Phase 1 with new information
- If ≥ 3 → question the architecture instead of trying fix #4

## Red Flags — Stop and Return to Phase 1

If you catch yourself thinking:
- "Quick fix for now, investigate later"
- "Just try changing X and see if it works"
- "Add multiple changes, then run tests"
- "I will write the test after confirming the fix works"
- "It is probably X, let me fix that"
- "I do not fully understand, but this might work"

All of these mean: stop, and go back to Phase 1.

## Quick Reference

| Phase | Key Activities | Success Criteria |
|-------|----------------|------------------|
| **1. Root Cause** | Read errors, reproduce, check changes, gather evidence | Understand **what** and **why** |
| **2. Pattern** | Find working examples, compare | Differences identified |
| **3. Hypothesis** | Form theory, test minimally | Confirmed or revised theory |
| **4. Implementation** | Create test, fix, verify | Bug resolved, tests pass |

## When Process Reveals "No Root Cause"

If systematic investigation reveals the issue is truly environmental, timing-dependent, or external:

1. Document what you investigated
2. Implement appropriate handling (retry, timeout, error message)
3. Add monitoring / logging for future investigation

But in most cases, "no root cause" means investigation is incomplete.

## Integration

Pairs with:
- root-cause-tracing — for deep call-stack tracing
- test-driven-development — for creating failing tests first
- defense-in-depth — for adding guards at multiple layers
