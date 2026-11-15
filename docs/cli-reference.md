---
id: cli-reference
title: CLI reference
sidebar_position: 3
---

# CLI reference

The installed binary is called **`factory`**, even though the GitHub repository is `ai-factory-cli`.

| Command | Current behavior |
| --- | --- |
| `factory`, `factory help` | Print help text |
| `factory version` | Print the build version |
| `factory tools` | Read-only CLI executable discovery |
| `factory pilot --repo X/Y --issue N` | Prepare a read-only issue-analysis pilot (dry run) |
| `factory pilot ... --execute` | Call approved local `gh` and `opencode` for issue analysis; review local permissions first |
| `factory demo --scenario running` | Live-refreshing terminal dashboard with **synthetic** telemetry |
| `factory scan [path]` | Read-only inspection of known local paths |
| `factory processes [path]` | Find named `PROCESS.md` files, including under `.factory/` |
| `factory doctor [path]` | Minimal presence check for `.factory/PROCESS.md` |
| `factory status [path]` | Explicitly reports that the supervisor is not implemented |

### Planned, not implemented

`factory onboard`, `factory run`, `factory serve`, `factory tui`, `factory logs`, `factory update`, `factory report` and `factory offboard` are **future commands**. Invoking an unimplemented command currently returns an error.

### Output and failure expectations

`factory scan .` only checks existence, never installs software or changes project files. Missing paths are reported as `missing`. `factory doctor` returns an error if the process file is absent. Other runtime and provider validation remains future work.

### Development

~~~bash
go test ./...
go vet ./...
~~~

See [Binary installation](./guides/releases.md), [WSL pilot](./guides/wsl-pilot.md), and [Terminal demo](./guides/terminal-demo.md).
