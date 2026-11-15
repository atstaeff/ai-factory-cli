---
name: product-discovery
version: 1
status: example-only
defaults:
  execution:
    runtime: approved-agent-cli
    provider: approved-provider
    model: approved-model-id
    thinking: medium
steps:
  - id: research
    uses: team/research
    next: concept
  - id: concept
    uses: team/product-brief
    execution:
      thinking: high
    next: review
  - id: review
    uses: factory/human-approval
    next: final
  - id: final
    uses: team/publish-document
---

# Example: product management

**Not executable yet.**

## Intake
A feature idea is submitted in the existing work tracker.

## Deliverable
A product brief stored as a versioned repository document, with a link in the work item.

## Review
A person approves scope, evidence and next steps. A comment alone must not count as authorization unless the future policy and adapter explicitly permit it.
