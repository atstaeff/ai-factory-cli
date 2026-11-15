---
id: processes
title: Process-as-code
---

# Process-as-code

The proposed project-level entry point is `.factory/PROCESS.md`: YAML frontmatter to describe machine-verifiable metadata and Markdown for human and agent instructions.

**Today:** the prototype only discovers file names and reads the `name` frontmatter value. It does **not** validate or execute the steps below.

## Proposed contract

~~~yaml
---
name: feature-delivery
version: 1
defaults:
  execution:
    runtime: opencode
    provider: your-approved-llm-provider
    model: your-approved-model-id
    thinking: medium
steps:
  - id: analysis
    uses: bmad/help
    next: build
  - id: build
    uses: bmad/dev-story
    next: quality
  - id: quality
    uses: approved-ci/tests
    next:
      success: review
      failure: remediation
  - id: review
    uses: bmad/code-review
    next: human
  - id: human
    uses: factory/approval
---
~~~

The `uses` identifiers are proposed factory references, **not official BMAD commands**. Future adapters must map each identifier to an installed, version-compatible capability.

## Rules for an executable future schema

- Each step requires a stable ID.
- An unknown provider or capability must fail validation.
- An LLM-backed step can inherit defaults and override model/thinking settings.
- A deterministic CI or Git step should not spawn an LLM.
- Process conditions, required checks and approval policies must be evaluated by code, not inferred from prose.
- Running jobs should snapshot the process version, resolved capabilities and effective model settings.
- Changing a project's GitHub board view should not silently change its executable workflow.

See the [examples](../examples.md).
