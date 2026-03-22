---
name: test-generation-patterns
description: Automated test generation using structural analysis, edge case inference, and framework-specific patterns
category: core
version: "2.0.0"
---

# Test Generation Patterns

> Automated test generation using structural analysis, edge case inference, and framework-specific patterns.

## When to Use
- WHEN `mode=BUILD` and a new function or module has no test coverage
- WHEN `mode=DEBUG` and a bug was not caught by existing tests — generate missing test cases
- WHEN `mode=BUILD` and you need to quickly scaffold tests for an existing codebase

## Actions
`'generate' | 'suggest-edges' | 'mock'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| source | string | yes | File path or function signature to generate tests for |
| framework | string | no | Test framework: jest, vitest, pytest, go (default: auto-detect from project config) |
| test_type | string | no | Test type: unit, integration, property (default: unit) |
| include_edges | boolean | no | Whether to include edge case tests (default: true) |

## Execution Pipeline

### Step 1: Analyze Source
Use `view` to read the source code. Use `list_symbols` to extract function signatures, parameter types, and return types. Identify the function's behavior: pure computation, I/O, state mutation, or error-prone.

### Step 2: Generate Test Cases
Use `mcp_floyd-devtools_test_generator` with action `generate`. Pass the source code, framework, and test type. This produces a complete test file with test cases for:
- **Happy path**: typical valid inputs producing expected outputs
- **Boundary cases**: empty inputs, max values, zero values
- **Error cases**: invalid inputs, nil/null, type mismatches

### Step 3: Suggest Edge Cases
Use `mcp_floyd-devtools_test_generator` with action `suggest_edge_cases` to identify additional edge cases specific to the function's domain (e.g., concurrent access, large inputs, unicode strings).

### Step 4: Generate Mocks
If the function has external dependencies (database, API, file system), use `mcp_floyd-devtools_test_generator` with action `generate_mocks` to create mock implementations.

### Step 5: Write and Verify
Use `write` to create the test file in the appropriate location. Use `bash` to run the tests and verify they pass.

## Output Shape
```json
{
  "source": "string — the analyzed source",
  "test_file": "string — path to the generated test file",
  "framework": "string — framework used",
  "test_cases": [
    {
      "name": "string — test case name",
      "type": "HAPPY | BOUNDARY | ERROR | EDGE",
      "input": "string — test input description",
      "expected": "string — expected output description"
    }
  ],
  "total_cases": 8,
  "mocks_generated": ["string — mock dependency names"],
  "test_run_result": "PASS | FAIL"
}
```

## Failure Modes
- IF the test framework cannot be auto-detected: prompt for explicit `framework` specification
- IF the source has side effects that cannot be safely tested: flag the test as integration-level and recommend sandbox execution

## Examples
```json
{
  "action": "generate",
  "source": "internal/engine/planner.go",
  "framework": "go",
  "test_type": "unit",
  "include_edges": true
}
```
