# AI Factory developer automation
# The Docusaurus Markdown sources remain in docs/.
# Generated HTML is written to docs/site/ and excluded from Git.

default:
    @just --list

install-docs:
    npm install --no-audit --no-fund

build-docs:
    npm run build -- --out-dir docs/site
    node -e "require('fs').accessSync('docs/site/index.html')"

serve-docs:
    npm run serve -- --dir docs/site --no-open

test-go:
    go test ./...
    go vet ./...
