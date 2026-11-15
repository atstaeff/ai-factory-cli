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
| `factory tools` | Read-only executable availability |
| `factory pilot --repo X/Y --issue N [--execute]` | Optional GitHub issue/OpenCode analysis pilot |
| `factory demo --scenario running` | Live terminal redraw with simulated data |

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
just build-docs
```

Development URL: `http://localhost:3000/ai-factory-cli/`.

Once GitHub Pages is enabled (Settings → Pages → Build and deployment → **GitHub Actions**) and a docs change is merged to `main`, the Pages workflow attempts to publish the compiled `build/` directory to `https://atstaeff.github.io/ai-factory-cli/`. The Pages URL is **not verified live** yet; the repository is public; publishing requires a successful Pages Actions workflow.

**Important:** GitHub Pages does not directly serve the Docusaurus Markdown in `docs/`; Docusaurus first generates `docs/site/` via `just build-docs`; GitHub Actions uploads that artifact. The documentation is public-facing: do not include internal customer details, credentials or proprietary CLI snippets.

The examples show illustrative processes only. They are not yet executable. See the [roadmap](docs/roadmap.md) for implementation status.

Documentation dependencies are version-pinned in `package.json`, but a `package-lock.json` still needs to be generated and committed before strict, reproducible npm CI using `npm ci` is enabled.


### Reproducible documentation build

Install the `just` command runner, then:

```bash
npm install
just build-docs
just serve-docs
```

The Markdown sources remain in `docs/`; generated static HTML is placed in `docs/site/` and deliberately Git-ignored. The PR Docs CI and GitHub Pages workflow both call **`just build-docs`**. GitHub Pages is configured to use **GitHub Actions** (rather than selecting `/docs` as a branch folder); its artifact is `docs/site/`.

## Native binaries (no Go required)

After a tagged prerelease is published, get a prebuilt macOS/Linux/Windows binary from [GitHub Releases](https://github.com/atstaeff/ai-factory-cli/releases). Install instructions for WSL and PowerShell: [release guide](docs/guides/releases.md). Native binary users can run `factory scan .`, `factory processes .` and `factory tools` without a Go toolchain. The `examples/` directory is not embedded in the binary.

## Functional validation

Start with [the WSL pilot](docs/guides/wsl-pilot.md) on a clean repository and an existing test issue. The opt-in `--execute` path invokes installed `gh` and `opencode` locally; BMAD execution is **not** implemented. [Terminal demo](docs/guides/terminal-demo.md) uses synthetic data only.


## Documentation brand alignment

The Docusaurus UI follows the design tokens in `stefanposs/sp-website-blog` (`data/theme.json`): primary navy `#122E53`, white and light-grey surfaces, Heebo body text, and Signika headings. The style source is intentionally kept separate from website code; the documentation remains a standalone static Docusaurus build. After merging a theme change, GitHub Actions must successfully rebuild and deploy `docs/site/` for it to appear on GitHub Pages.
