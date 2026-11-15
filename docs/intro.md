---
id: intro
title: Welcome to AI Factory
sidebar_position: 1
description: A headless process-first tool for AI-assisted work.
---

# AI Factory

**AI Factory** aims to connect your existing workflows, methods, skills, agents, models, humans and approved tooling into reliable work processes.

It is deliberately a **CLI**, not a new issue tracker, methodology or AI model provider.

:::caution Prototype status
The current Go prototype supports basic local `scan`, `processes`, `doctor`, `status`, `help` and `version`. **It does not yet run autonomous processes, invoke BMAD/OpenSpec, connect to Jira/GitHub, schedule work, or deploy a TUI.**
:::

## Choose your starting point

- [Getting started](./getting-started.md): build and run the currently implemented CLI.
- [CLI reference](./cli-reference.md): commands, expected output and limitations.
- [Architecture](./concepts/architecture.md): work providers, process steps, runtimes and gates.
- [Brownfield onboarding](./guides/brownfield.md): the desired safe adoption model.
- [BMAD collaboration](./guides/bmad.md): optional project setup for teams.
- [Examples](./examples.md): realistic illustrative process definitions.
- [Roadmap](./roadmap.md): what's implemented and what comes next.

## Philosophy

**The process belongs to the project.** The factory should only coordinate it, keeping permanent documentation in Git and human accountability in the project's existing work management system.
