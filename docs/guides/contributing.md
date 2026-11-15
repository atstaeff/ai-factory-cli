---
id: contributing
title: Contributing and documentation
---

# Contributing and documentation

## Engineering

~~~bash
go test ./...
go vet ./...
go build -o factory ./cmd/factory
~~~

Keep changes focused, make scans read-only, and describe unsupported commands truthfully. Follow [AGENTS.md](https://github.com/atstaeff/ai-factory-cli/blob/main/AGENTS.md).

## Documentation

Docs live under `docs/`, powered by Docusaurus configured in the repo root.

~~~bash
npm install
npm run start
npm run build
~~~

The homepage is `src/pages/index.js`. Navigation uses `sidebars.js`. The shared theme is `src/css/custom.css`.

The PR workflow checks that the docs build; only pushes to `main` attempt deployment to GitHub Pages.

## GitHub Pages setup

1. Repository **Settings → Pages**.
2. Under **Build and deployment**, select **GitHub Actions** as the source.
3. Ensure Actions are enabled and allow Pages deployment.
4. Merge a reviewed PR into `main`; the deploy workflow publishes the generated `build/` directory.

Docusaurus source files under `docs/` are **not themselves a deployable GitHub Pages website**. GitHub Pages serves the compiled static output. The intended public URL is `https://atstaeff.github.io/ai-factory-cli/` **only after a successful deployment**.

:::caution Private repository
The source repository is private. GitHub Pages availability and visibility depend on plan, settings and organizational policies. Do not deploy sensitive documentation; configure Pages visibility before publishing.
:::
