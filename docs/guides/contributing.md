---
id: contributing
title: Contributing and documentation
---

# Contributing and documentation

## Engineering

```bash
go test ./...
go vet ./...
go build -o factory ./cmd/factory
```

Keep changes focused, make scans read-only, and describe unsupported commands accurately. See [AGENTS.md](https://github.com/atstaeff/ai-factory-cli/blob/main/AGENTS.md).

## Documentation

All public documentation sources are in `docs/`.

**Production publishing:** GitHub Pages uses its native branch-source Jekyll renderer. It publishes directly from `main` → `/docs`, with `docs/index.md` as the homepage, `docs/_layouts/default.html` for navigation and layout, and `docs/assets/css/theme.css` for the stefanposs.com-inspired theme.

**Optional local Docusaurus compatibility build:** The repository still contains its original Docusaurus preview under `src/`, `sidebars.js` and `docusaurus.config.js`. This is separate from the Pages publishing path. `just build-docs` generates `docs/site/` for local testing; that generated directory is Git-ignored and is **not** the hosted Pages root.

```bash
npm install
just build-docs
```

## GitHub Pages setup: main /docs

Configure this once as a repository administrator:

1. Open [Repository Settings → Pages](https://github.com/atstaeff/ai-factory-cli/settings/pages).
2. Set **Build and deployment → Source: Deploy from a branch**.
3. Select **Branch: main**, **Folder: /docs**, then **Save**.
4. GitHub Pages builds Jekyll pages from the tracked Markdown and layout files and serves the generated `index.html` at [AI Factory Documentation](https://atstaeff.github.io/ai-factory-cli/).

The previous `Publish Docs to GitHub Pages` Actions workflow has been removed intentionally: it was configured for a different deployment model and GitHub rejected its attempt to modify the repository Pages setting (HTTP 403). The branch-based flow does not need that workflow.

## Troubleshooting

A green Docusaurus CI check only shows that the optional preview builds; it does not guarantee a live GitHub Pages deployment. If the public site returns 404:

- Check that the **publishing source is main /docs**, not main root or GitHub Actions.
- Confirm `docs/index.md`, `docs/_config.yml`, and `docs/_layouts/default.html` exist on `main`.
- Check the **Pages build/deployment status** under Actions and the Pages settings screen for any Jekyll compilation error.
- Use the canonical URL `https://atstaeff.github.io/ai-factory-cli/` (with trailing slash).

The repository is public. Never include sensitive company information or credentials in documentation or generated examples.
