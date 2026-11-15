---
id: roadmap
title: Roadmap
---

# Roadmap and current status

| Capability | Status |
| --- | --- |
| Go `factory` executable, help and version | Implemented |
| Read-only local scan | Implemented |
| PROCESS.md discovery | Implemented (name only) |
| Minimal doctor and status | Implemented |
| Docusaurus documentation | Implemented; verify deployment separately |
| Binary release workflow | Implemented in this PR; publication requires a tagged release |
| Tool executable diagnostics | Implemented |
| OpenCode / GitHub issue analysis pilot | Implemented; real WSL acceptance testing pending |
| Synthetic terminal demo | Implemented (not real telemetry) |
| Validated process schema | Planned |
| Brownfield onboarding, updates and offboarding | Planned |
| BMAD/OpenSpec runtime integration | Planned |
| GitHub, Jira and custom enterprise CLI adapters | Planned |
| Scheduler, durable checkpoints and recovery | Planned |
| Human-on-exception and resumption | Planned |
| Terminal monitoring and model routing | Planned |

## Planned vertical slice

Read a mock work item through an approved CLI, validate a single process, record execution state and produce a structured result. Full automation and privileged actions require further engineering and review.

The project intentionally avoids claiming that a static example already executes.
