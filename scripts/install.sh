#!/usr/bin/env bash
set -euo pipefail
VERSION="$1"
[[ "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.-]+)?$ ]] || { echo "Usage: install.sh v0.1.0-rc.1" >&2; exit 2; }
case "$(uname -s)" in Linux) os=linux;; Darwin) os=darwin;; *) echo "This script supports macOS, Linux and WSL; use install.ps1 on Windows." >&2; exit 2;; esac
case "$(uname -m)" in x86_64|amd64) arch=amd64;; aarch64|arm64) arch=arm64;; *) echo "Unsupported CPU" >&2; exit 2;; esac
for cmd in curl tar; do command -v "$cmd" >/dev/null || { echo "Missing $cmd" >&2; exit 1; }; done
asset="factory_${VERSION}_${os}_${arch}.tar.gz"
base="https://github.com/atstaeff/ai-factory-cli/releases/download/${VERSION}"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT
curl -fsSL "$base/$asset" -o "$tmp/$asset"
curl -fsSL "$base/SHA256SUMS" -o "$tmp/SHA256SUMS"
expected="$(awk -v file="$asset" '$2 == file { print $1; exit }' "$tmp/SHA256SUMS")"
[[ "$expected" =~ ^[0-9a-fA-F]{64}$ ]] || { echo "Missing or invalid SHA256 entry" >&2; exit 1; }
if [[ "$os" == "darwin" ]]; then
  command -v shasum >/dev/null || { echo "Missing shasum" >&2; exit 1; }
  actual="$(shasum -a 256 "$tmp/$asset" | awk '{print $1}')"
else
  command -v sha256sum >/dev/null || { echo "Missing sha256sum" >&2; exit 1; }
  actual="$(sha256sum "$tmp/$asset" | awk '{print $1}')"
fi
[[ "$actual" == "$expected" ]] || { echo "Checksum verification failed" >&2; exit 1; }
tar -xzf "$tmp/$asset" -C "$tmp" factory
mkdir -p "$HOME/.local/bin"
install -m 0755 "$tmp/factory" "$HOME/.local/bin/factory"
"$HOME/.local/bin/factory" version
echo 'To use from any directory, add ~/.local/bin to PATH (e.g. export PATH="$HOME/.local/bin:$PATH").'
