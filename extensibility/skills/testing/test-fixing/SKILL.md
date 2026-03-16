---
name: test-fixing
description: Run tests and systematically fix all failing tests using smart error grouping. Use when user asks to fix failing tests, mentions test failures, runs test suite and failures occur, or requests to make tests pass.
---

# Test Fixing

## Overview

Systematically fix failing tests through intelligent error grouping and prioritized fixing.

Core principle: Group failures → Fix infrastructure first → Verify incrementally.

## The Process

### Step 1: Run Full Test Suite

\`\`\`bash
# Run complete test suite
make test
# Or: npm test, cargo test, pytest, go test ./...
\`\`\`

Capture all failures. Don't stop at first error.

### Step 2: Group Failures by Type

Categorize all failures into groups:

#### Group A: Infrastructure/Import Failures
- Module not found
- Import errors
- Missing dependencies
- Configuration issues
- File path errors

#### Group B: API/Interface Changes
- Function signature mismatches
- Renamed methods
- Changed return types
- Missing parameters
- Type errors

#### Group C: Logic/Behavior Failures
- Assertion failures
- Unexpected values
- Business logic errors
- Edge case handling

#### Group D: Test Code Issues
- Broken test setup/teardown
- Mock configuration errors
- Test data problems
- Fixture issues

### Step 3: Fix in Priority Order

Fix groups in this order: A → D → B → C

Why this order:
1. **Infrastructure first** - Nothing else will work until imports/config fixed
2. **Test code next** - Ensures test framework working properly
3. **API changes** - Mechanical fixes that affect many tests
4. **Logic last** - Requires understanding business requirements

### Step 4: Verify Each Group

After fixing each group:

\`\`\`bash
# Run focused test subset
pytest tests/test_module.py  # Just the fixed module
# Or
npm test -- --testPathPattern=module  # Jest example
\`\`\`

Don't proceed to next group until current group passes.

### Step 5: Final Full Run

After all groups fixed:

\`\`\`bash
# Full suite verification
make test
\`\`\`

All tests must pass.

## Example Workflow

\`\`\`
Step 1: Run full suite
  → 47 failures found

Step 2: Group failures
  Group A (Infrastructure): 12 failures
    - ImportError: cannot import 'old_module'
    - ModuleNotFoundError: No module named 'utils'
  
  Group B (API Changes): 23 failures
    - TypeError: get_user() takes 2 arguments, 3 given
    - AttributeError: 'User' has no attribute 'username'
  
  Group C (Logic): 8 failures
    - AssertionError: expected True, got False
  
  Group D (Test Code): 4 failures
    - Fixture 'mock_db' not found

Step 3: Fix Group A (Infrastructure)
  - Update imports to new module structure
  - Install missing dependencies
  Run focused tests: 12 → 0 failures ✓

Step 4: Fix Group D (Test Code)
  - Update fixture definitions
  - Fix test setup
  Run focused tests: 4 → 0 failures ✓

Step 5: Fix Group B (API Changes)
  - Update function calls to new signatures
  - Rename attributes
  Run focused tests: 23 → 0 failures ✓

Step 6: Fix Group C (Logic)
  - Update assertions for changed behavior
  - Fix edge case handling
  Run focused tests: 8 → 0 failures ✓

Step 7: Final verification
  make test: 0 failures ✓
\`\`\`

## Smart Grouping Tips

**Look for patterns:**
- Same error message → likely same root cause
- Same module → related failures
- Same test file → shared setup issues

**Don't treat each failure individually:**
- 20 ImportErrors for same module = 1 fix
- 15 TypeError for same signature = 1 fix

**Identify cascading failures:**
- Infrastructure failure can cause logic failures
- Fix infrastructure first, re-run to see true failures

## Common Mistakes

**Fixing failures one at a time**
- Problem: Misses patterns, wastes time
- Fix: Group similar failures, fix root cause once

**Skipping focused verification**
- Problem: Don't know if fix worked until full run
- Fix: Verify each group immediately after fixing

**Fixing logic before infrastructure**
- Problem: Can't test logic fixes if imports broken
- Fix: Always infrastructure → test code → API → logic

**Not re-running after each group**
- Problem: Stacking multiple changes, unclear what fixed what
- Fix: Verify each group before moving on

## Red Flags

Never:
- Fix random failures without grouping
- Skip verification between groups
- Change test assertions to match broken code
- Proceed to next group with failures in current group

Always:
- Run full suite first
- Group intelligently by root cause
- Fix infrastructure before logic
- Verify incrementally

## Integration

Pairs with:
- systematic-debugging — Use for complex failure investigation
- test-driven-development — Ensures tests stay green
