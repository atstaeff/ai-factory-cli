# AI Factory CLI

**Factory** is an early-stage, headless, process-first orchestration CLI. Its long-term goal is to connect existing methods (BMAD, OpenSpec), agent runtimes, internal tools, issue systems, quality gates and human approvals without replacing them.

Repository: `atstaeff/ai-factory-cli`  
CLI command: **`factory`**

> Status: **early prototype**. No scheduling, issue processing, model routing or autonomous execution yet.

## Development

Requires Go 1.23+.

```sh
go test ./...
go run ./cmd/factory help
go run ./cmd/factory scan .
go run ./cmd/factory processes .
go run ./cmd/factory doctor .
```

To build an executable named `factory`:

```sh
go build -o factory ./cmd/factory
./factory
```

## Commands available

| Command | Behavior |
| --- | --- |
| `factory` / `factory help` | Show help |
| `factory scan [path]` | Read-only detection of local project files |
| `factory processes [path]` | Discover PROCESS.md files |
| `factory doctor [path]` | Minimal presence check; not full compatibility validation |
| `factory status [path]` | Explicitly reports that the supervisor is not implemented |
| `factory version` | CLI version |

Other commands (run, onboard, update, serve, tui, logs) are **planned, not implemented**. They return a clear error rather than claiming execution.

## Project process format

See `.factory/PROCESS.md`. It is a human-readable draft. A validated, versioned execution schema is future work. No arbitrary frontmatter is currently executed.

## Optional BMAD setup for contributors

See [BMAD collaboration guide](docs/BMAD.md). BMAD is not preinstalled or auto-downloaded by Factory.

## Architecture direction

- **Go core:** lifecycle inspection, process registry, reliable execution state and policy evaluation
- **Adapters:** issue sources (GitHub, enterprise CLI), methodology (BMAD, OpenSpec), agent runtimes, and quality systems
- **Governance:** externally enforced permissions, approval gates, minimal credentials, auditable actions
- **Operating model:** issue ownership remains human; automated steps report progress and request human attention only when necessary

The immediate goal is to validate brownfield discovery and the process contract before autonomous execution.

## Documentation website and examples

Docusaurus source content lives under [`docs/`](docs/), with a designed landing page in `src/pages/`. Examples are in [`examples/`](examples/).

```bash
npm install
npm run start
npm run build
```

Development URL: `http://localhost:3000/ai-factory-cli/`.

Once GitHub Pages is enabled (Settings → Pages → Build and deployment → **GitHub Actions**) and a docs change is merged to `main`, the Pages workflow attempts to publish the compiled `build/` directory to `https://atstaeff.github.io/ai-factory-cli/`. The Pages URL is **not verified live** yet; the repository is private and availability/visibility depend on plan and settings.

**Important:** GitHub Pages does not directly serve the Docusaurus Markdown in `docs/`; Docusaurus first creates a static `build/` artifact. The documentation is public-facing: do not include internal customer details, credentials or proprietary CLI snippets.

The examples show illustrative processes only. They are not yet executable. See the [roadmap](docs/roadmap.md) for implementation status.

Documentation dependencies are version-pinned in `package.json`, but a `package-lock.json` still needs to be generated and committed before strict, reproducible npm CI using `npm ci` is enabled.
