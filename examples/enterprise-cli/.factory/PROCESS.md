---
name: enterprise-governed-work
version: 1
status: example-only
work:
  adapter: company-issue-cli
  events: polling
defaults:
  execution:
    runtime: approved-coding-cli
    provider: internal-llm
    model: approved-model-id
    thinking: medium
steps:
  - id: intake
    uses: company-issues/read
    next: analysis
  - id: analysis
    uses: team/analysis
    next: quality
  - id: quality
    uses: company-ci/verify
    next:
      success: human
      failure: analysis
  - id: human
    uses: company-issues/request-decision
---

# Example: restricted enterprise integration

**Not executable yet.** Placeholder adapters represent enterprise-approved commands rather than public Jira API access.

## Rules
Use only approved CLIs and accounts. The future worker must pass explicit command arguments, handle structured output and avoid shell interpolation.

## No webhooks
A future poller should use durable cursors and backoff to notice an authorized human decision and resume only the corresponding job.
