---
id: terminal-demo
title: Real terminal demo
---

# Real terminal demo: observability scenarios

The command below starts a **genuinely refreshing ANSI terminal display**, but its data is explicitly **synthetic**. This is a display prototype, not an operational supervisor.

~~~bash
factory demo --scenario running
factory demo --scenario approvals
factory demo --scenario incidents
factory demo --scenario multi
~~~

Use **Ctrl+C** to stop. For scripted tests or screenshots, run exactly one frame:

~~~bash
factory demo --scenario running --frames 1
~~~

The demo prints an ANSI frame on every refresh and overwrites the terminal screen. It works best in a terminal of at least 105 columns.

## Scenarios

| Name | What it demonstrates |
| --- | --- |
| running | Active work and queued jobs |
| approvals | Human decisions and paused jobs |
| incidents | Failures, retry count and missing heartbeats |
| multi | Cross-domain BMAD, OpenSpec and custom processes |

## Explicit limitations

- All jobs, event timestamps, models, runtimes and issues in demo mode are **mock data**.
- OpenCode, BMAD, OpenSpec, GitHub and your internal Jira CLI are **not invoked**.
- Counts and progress are synthetic; no metrics are collected from the OS or live processes.
- The demo is deliberately separate from the future real \`factory tui\` command.

Screenshots should always be captured from an **executed terminal session** and labeled *Live-rendered terminal demo — synthetic data*. Do not present mock telemetry as proof of an implemented integration.

## Technical implementation

Currently uses the Go standard library, ANSI clearing/color sequences, a clock-based refresh loop and scripted scenario data. For production, a TUI should subscribe to durable supervisor state, with terminal-size handling, interactive selection, accessibility, robust Windows compatibility and appropriate process lifecycle management.
