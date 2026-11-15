---
id: architecture
title: Architecture
---

# Architecture

Factory is intended to be a **headless process supervisor**, not an alternative to a work tracker, coding agent or methodology.

## Responsibilities

| Boundary | Examples | Owns |
| --- | --- | --- |
| Work source | GitHub Issues, Jira via approved internal CLI | Work identity, assignee, human collaboration |
| Process method | BMAD, OpenSpec, custom | Method-specific instructions and artifacts |
| Runtime | OpenCode, Claude Code, Copilot CLI | AI execution and model invocation |
| Capability | Skill, script, CI job | One specialized task |
| Factory core | Go process engine | State, transitions, policies, scheduling, recovery |
| Output | Git repo, issue comment, PR, document | Durable deliverables |

No runtime or provider is required by the initial scanner. The above boundaries are **architectural goals**, not shipped integrations.

## High-level flow

~~~text
Issue / work source
        |
     process
        |
   step -> capability + runtime + model
        |
  verified result / gate
        |
next step OR human decision
~~~

## Engineering constraints

- Keep the source project authoritative for its own process definitions.
- Preserve brownfield modifications during adoption.
- Make process versions immutable for already-running jobs.
- Never treat model output or untrusted issue text as a governance permission.
- Require approved tools; a prompt-only allowlist is insufficient.
- Do not duplicate project documentation into issue bodies.
- Keep the supervisor independent of the TUI lifecycle.

See [Process concepts](./processes.md) and [Integrations](./integrations.md).
