---
id: tooling-review
title: Integration readiness review
---

# Integration readiness review

This matrix distinguishes observed facts from planned work. **Presence checks are not integrations.**

| System | Available integration entry point | Current Factory support | Main risk |
| --- | --- | --- | --- |
| BMAD v6 | Project-local `_bmad/` and installed skill files | Detect paths only | Version-specific workflows and customizations |
| BMAD v7 preview | Skills CLI + setup in coding tool | No adapter | Preview changes can break compatibility |
| OpenCode | `opencode run` non-interactive prompt | Executable detection only | Sessions, permissions, subprocess lifetime and model routing |
| GitHub | `gh` CLI or approved API | Executable detection only | Authentication, permissions, rate limits |
| Jira | Organization-approved CLI | Optional executable detection via `FACTORY_ISSUE_CLI` | No webhook, output contract unknown, governance |
| OpenSpec | Existing project configuration and skills | Local path detection only | Custom schemas and migrations |
| Models | Native coding-runtime configuration | Not implemented | Availability, costs, thinking effort, fallback rules |
| Processes | `.factory/PROCESS.md` | Name discovery only | No validated schema, conditions or state machine |

## Required engineering gates before autonomous execution

1. Define a typed capability contract for each approved integration, including timeouts, structured output and exit codes.
2. Validate workflow graphs and permissions before running anything.
3. Record execution snapshots (effective runtime, model, process revision and tool versions).
4. Isolate working directories and credentials per project/customer.
5. Use a durable event store with retries, idempotency, leasing and recovery.
6. Resolve explicit human approval through structured work-item references.
7. Verify a real, company-approved Jira CLI and real BMAD/OpenCode integration in a safe test project.

## What the current diagnostics do

`factory tools` calls OS executable lookup for `git`, `gh`, `opencode`, `claude`, `bmad`, `openspec` and a configurable enterprise issue CLI. It **does not run** any third-party command, inspect authentication or validate semantics. Missing executables are informational; none is mandatory for the local scanner.

The variable `FACTORY_ISSUE_CLI` only indicates the executable path/name to inspect; it is **not** a configured authenticated Jira adapter.

Do not confuse illustrative process examples with a functional end-to-end integration.
