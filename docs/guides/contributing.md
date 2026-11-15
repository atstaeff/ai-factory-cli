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
just build-docs
~~~

The homepage is `src/pages/index.js`. Navigation uses `sidebars.js`. The shared theme is `src/css/custom.css`.

The PR workflow checks that the docs build; only pushes to `main` attempt deployment to GitHub Pages.

## GitHub Pages setup

1. Repository **Settings → Pages**.
2. Under **Build and deployment**, select **GitHub Actions** as the source.
3. Ensure Actions are enabled and allow Pages deployment.
4. Merge a reviewed PR into `main`; the deploy workflow publishes the generated `docs/site/` directory.

Docusaurus source files under `docs/` are **not themselves a deployable GitHub Pages website**. GitHub Pages serves the compiled static output. The intended public URL is `https://atstaeff.github.io/ai-factory-cli/` **only after a successful deployment**.

:::caution Private repository
The source repository is private. GitHub Pages availability and visibility depend on plan, settings and organizational policies. Do not deploy sensitive documentation; configure Pages visibility before publishing.
:::


## Static output location

`just build-docs` builds Docusaurus into `docs/site/` while keeping Markdown under `docs/`. The generated directory is Git-ignored; GitHub Actions uploads it to Pages. Do **not** set GitHub Pages to deploy from the `main` branch's `/docs` folder: that contains Markdown sources, not the built website. Use GitHub Actions as the Pages source.

## Site deployment and 404 troubleshooting

The `just build-docs` recipe produces `docs/site/index.html`; the `Publish Docs to GitHub Pages` workflow uploads that folder as its root artifact. It attempts to enable GitHub Pages and switch its `build_type` to `workflow`, then checks the public URL for an HTTP success and AI Factory content. If GitHub rejects automatic activation, verify **Settings → Pages → Build and deployment → Source: GitHub Actions** and inspect the failing Actions job. A green Docs CI does **not** mean Pages is deployed.

The canonical URL is `https://atstaeff.github.io/ai-factory-cli/` (Docusaurus `baseUrl: '/ai-factory-cli/'`). Never point the Pages branch source at the Markdown-only `docs/` folder.


### One-time required GitHub Pages setting

On this repository GitHub reported `build_type: legacy`, and the workflow's token received HTTP 403 when trying to switch it to `workflow`. This permission cannot be overcome by changing build scripts. A repository administrator must visit [Settings → Pages](https://github.com/atstaeff/ai-factory-cli/settings/pages), choose **GitHub Actions** under **Build and deployment → Source**, and rerun [Publish Docs to GitHub Pages](https://github.com/atstaeff/ai-factory-cli/actions/workflows/pages.yml) on `main`. The workflow then uploads and verifies the generated site.


## Branch-based GitHub Pages setup (current)

In **Repository Settings → Pages → Build and deployment**, set **Source: Deploy from a branch**, **Branch: main**, **Folder: /docs**, then Save. The public site is built by GitHub Pages/Jekyll from `docs/index.md` and `docs/_layouts/default.html` (the entry point becomes `index.html`). No separate Pages Action is needed; this intentionally replaces the former Actions artifact deployment. All source Markdown continues to live in `docs/`.

The earlier GitHub Pages publishing instructions above that refer to the Actions source are *superseded* by this section. `just build-docs` still performs a separate Docusaurus compatibility build; it is not what branch-based GitHub Pages publishes. Public site: https://atstaeff.github.io/ai-factory-cli/.
