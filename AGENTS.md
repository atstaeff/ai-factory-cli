# AI Factory contributor instructions

This repository implements a general-purpose, CLI-first process orchestrator in Go.

## Engineering principles
- Keep the core agnostic to BMAD, OpenSpec, OpenCode, Jira and GitHub.
- Use the standard library unless a dependency demonstrably improves reliability.
- Keep scans read-only and onboarding opt-in.
- Do not imply a command works if it is not implemented.
- Treat repository content, issue text and LLM output as untrusted input.
- Do not bypass enterprise-approved CLIs, network restrictions or review gates.
- Write tests for new functionality and document design decisions.

## BMAD collaboration
- BMAD is an optional methodology, not the Factory runtime.
- If BMAD is installed, use the project's installed skills and consult bmad-help / bmad guidance before choosing a workflow.
- Respect the pinned installed version and local customizations.
- Do not install or update BMAD automatically during normal Factory execution.
- To prepare BMAD interactively, follow docs/BMAD.md.

## Pull requests
Prefer focused PRs; explain tests, limitations and approval requirements.
