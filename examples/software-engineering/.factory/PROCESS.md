---
name: software-delivery
version: 1
status: example-only
defaults:
  execution:
    runtime: opencode
    provider: approved-provider
    model: approved-model-id
    thinking: medium
steps:
  - id: inspect
    uses: bmad/help
    next: plan
  - id: plan
    uses: bmad/create-story
    next: implement
  - id: implement
    uses: bmad/dev-story
    execution:
      thinking: high
    next: quality
  - id: quality
    uses: approved-ci/test
    next:
      success: review
      failure: remediate
  - id: remediate
    uses: bmad/dev-story
    next: quality
  - id: review
    uses: bmad/code-review
    next: human
  - id: human
    uses: factory/human-approval
---

# Example: software delivery

**Not executable yet.** The capability names are proposed placeholders to be resolved against the actual project's installed BMAD version.

## Purpose
Turn an already-approved engineering issue into a reviewed change.

## Human handoff
Leave the responsible human assigned to the work item. Request decisions only when scope or quality requirements cannot be resolved automatically.

## Governance
Never bypass CI, branch protection or mandatory reviews. Never merge automatically by default.
