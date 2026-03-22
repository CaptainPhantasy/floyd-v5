---
name: api-format-verifier
description: Verify API request and response payloads against provider-specific schema requirements and token limits
category: core
version: "2.0.0"
---

# API Format Verifier

> Verify API request and response payloads against provider-specific schema requirements and token limits.

## When to Use
- WHEN `mode=BUILD` and constructing an API integration to validate payloads before sending
- WHEN `mode=DEBUG` when an API call returns a schema validation error from the provider
- WHEN `mode=BUILD` and switching between LLM providers (OpenAI, Anthropic, Google) to verify format compatibility

## Actions
`'verify-request' | 'verify-response' | 'check-compatibility'`

## Parameters
| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| payload | object | yes | The API request or response payload to verify |
| api_type | string | yes | Target API: openai, anthropic, google, custom |
| direction | string | yes | Payload direction: request, response |
| check_token_limits | boolean | no | Whether to verify token count against model limits (default: true) |

## Execution Pipeline

### Step 1: Validate Structure
Use `mcp_floyd-devtools_api_format_verifier` with action `verify_request` or `verify_response`. This checks the payload against the provider's JSON schema, ensuring all required fields are present, types are correct, and enums contain valid values.

### Step 2: Check Token Limits
If `check_token_limits=true`, use `mcp_floyd-devtools_api_format_verifier` with `estimate_cost=true` to compute the token count. Compare against the model's context window and output limits. Flag payloads that exceed 80% of the limit as risky.

### Step 3: Compatibility Check
If switching providers, use `mcp_floyd-devtools_api_format_verifier` with action `check_compatibility` to identify format differences between the source and target API. Report field mappings, unsupported features, and required transformations.

## Output Shape
```json
{
  "valid": true,
  "errors": [
    {
      "field": "string — field path (e.g., messages.0.role)",
      "message": "string — validation error description",
      "severity": "ERROR | WARNING"
    }
  ],
  "token_estimate": {
    "input_tokens": 1500,
    "output_tokens": 0,
    "context_window_usage": 0.03,
    "within_limits": true
  },
  "compatibility_notes": ["string — provider-specific warnings"]
}
```

## Failure Modes
- IF the payload is invalid JSON: report the parse error with line and column
- IF the `api_type` is not recognized: fall back to a generic JSON schema validation and report unsupported provider

## Examples
```json
{
  "action": "verify-request",
  "payload": {"model": "gpt-4o", "messages": [{"role": "user", "content": "Hello"}]},
  "api_type": "openai",
  "direction": "request",
  "check_token_limits": true
}
```
