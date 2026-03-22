---
name: analogy-synthesis
description: Generate cross-domain analogies to explain complex code concepts, architectures, or debugging patterns
category: core
version: "2.0.0"
---

# Analogy Synthesis

> Generate cross-domain analogies to explain complex code concepts, architectures, or debugging patterns.

## When to Use
- WHEN `mode=EXPLORE` to understand an unfamiliar architecture by mapping it to a known domain
- WHEN `mode=DEBUG` to reframe a bug in terms of a well-understood pattern from another domain
- WHEN `mode=BUILD` to communicate architectural decisions to non-technical stakeholders

## Actions
`'generate' | 'map' | 'explain'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| source | string | yes | The code concept, architecture, or pattern to explain |
| target_domain | string | no | Preferred analogy domain (e.g., cooking, architecture, biology, traffic) — default: auto-select |
| depth | string | no | Explanation depth: `brief`, `detailed`, `comprehensive` (default: brief) |

## Execution Pipeline

### Step 1: Analyze Source
Use `view` and `list_symbols` to understand the structure of the `source` concept. Identify its key components, relationships, and behaviors. Map these to abstract properties: flow, transformation, state, coordination, constraint.

### Step 2: Select Target Domain
If `target_domain` is specified, use it directly. Otherwise, select a domain whose natural structures match the abstract properties of the source (e.g., a pipeline maps well to an assembly line; a state machine maps well to a traffic light).

### Step 3: Generate Analogy
Map each component of the source to an element in the target domain. Ensure the mapping is structurally faithful — the relationships in the analogy must mirror the relationships in the code. Validate that the analogy does not break down for the `depth` level requested.

### Step 4: Explain
Produce the explanation with the mapping table and a narrative connecting the analogy back to the code.

## Output Shape
```json
{
  "source": "string — the original concept",
  "target_domain": "string — the analogy domain used",
  "mapping": [
    {
      "code_element": "string — component in the code",
      "analogy_element": "string — corresponding element in the domain",
      "explanation": "string — why this mapping works"
    }
  ],
  "narrative": "string — the analogy explanation",
  "limitations": ["string — where the analogy breaks down"]
}
```

## Failure Modes
- IF the source concept is too abstract to map: break it into smaller sub-concepts and generate analogies for each
- IF no suitable domain is found: use a general systems-theory analogy (input → process → output)

## Examples
```json
{
  "action": "generate",
  "source": "Kubernetes pod scheduling with affinity constraints",
  "target_domain": "restaurant seating with dietary requirements",
  "depth": "detailed"
}
```
