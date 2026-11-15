#!/usr/bin/env bash
set -euo pipefail
VERSION="$1"
[[ "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?$ ]] || { echo "Usage: install.sh v0.1.0-rc.1" >&2; exit 2; }
[[ "$(uname -s)" == "Linux" ]] || { echo "Use install.ps1 on Windows." >&2; exit 2; }
case "$(uname -m)" in x86_64|amd64) arch=amd64;; aarch64|arm64) arch=arm64;; *) echo "Unsupported CPU" >&2; exit 2;; esac
for cmd in curl tar sha256sum; do command -v "$cmd" >/dev/null || { echo "Missing $cmd" >&2; exit 1; }; done
asset="factory_${VERSION}_linux_${arch}.tar.gz"
base="https://github.com/atstaeff/ai-factory-cli/releases/download/${VERSION}"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT
curl -fsSL "$base/$asset" -o "$tmp/$asset"
curl -fsSL "$base/SHA256SUMS" -o "$tmp/SHA256SUMS"
(cd "$tmp" && grep -F "  $asset" SHA256SUMS | sha256sum --check --status) || { echo "Checksum verification failed" >&2; exit 1; }
tar -xzf "$tmp/$asset" -C "$tmp" factory
mkdir -p "$HOME/.local/bin"
install -m 0755 "$tmp/factory" "$HOME/.local/bin/factory"
"$HOME/.local/bin/factory" version
echo 'To use from any directory, add ~/.local/bin to PATH (e.g. export PATH="$HOME/.local/bin:$PATH").'
