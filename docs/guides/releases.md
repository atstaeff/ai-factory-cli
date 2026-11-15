---
id: releases
title: Installing Factory binaries
---

# macOS, Windows and Linux binaries

A release workflow packages a standalone **`factory`** CLI, never an `ai-factory-cli` binary. It builds six target combinations with CGO disabled:

| Platform | Archive | Binary |
| --- | --- | --- |
| Linux x86_64 | `factory_vX.Y.Z_linux_amd64.tar.gz` | `factory` |
| Linux ARM64 | `factory_vX.Y.Z_linux_arm64.tar.gz` | `factory` |
| Windows x64 | `factory_vX.Y.Z_windows_amd64.zip` | `factory.exe` |
| Windows ARM64 | `factory_vX.Y.Z_windows_arm64.zip` | `factory.exe` |
| macOS Intel | `factory_vX.Y.Z_darwin_amd64.tar.gz` | `factory` |
| macOS Apple Silicon | `factory_vX.Y.Z_darwin_arm64.tar.gz` | `factory` |

The workflow also generates `SHA256SUMS`. No Go runtime is required on the target machine. BMAD and OpenCode are **not bundled** and require separately approved installation.

## No-Go installation via GitHub Releases

Once the first GitHub prerelease has **actually been published**, download only the archive for your OS/CPU from [Releases](https://github.com/atstaeff/ai-factory-cli/releases). On macOS, WSL/Linux and Windows, the included installation scripts verify checksums and install only the native `factory` executable to your user directory:

~~~bash
# macOS / WSL / Linux (after reviewing the script from the repository)
curl -fsSLO https://raw.githubusercontent.com/atstaeff/ai-factory-cli/main/scripts/install.sh
bash install.sh v0.1.0-rc.2
export PATH="$HOME/.local/bin:$PATH"
factory processes examples
factory scan examples/software-engineering
~~~

~~~powershell
# Windows PowerShell (after reviewing the script from the repository)
Invoke-WebRequest https://raw.githubusercontent.com/atstaeff/ai-factory-cli/main/scripts/install.ps1 -OutFile install.ps1
.\install.ps1 -Version v0.1.0-rc.2
factory version
factory scan .
~~~

**Important:** These commands work only after release `v0.1.0-rc.2` exists. They do not install OpenCode or BMAD. On managed devices, use approved download channels and review the scripts before execution. macOS builds are currently cross-compiled on Linux CI; test on the actual target machines before enterprise use. Unsigned/unnotarized macOS binaries may be blocked by Gatekeeper. To run the `examples` commands, run them inside a checkout of this repository or supply another existing project path. No Go toolchain is needed for either native binary.

## Download

When the first tagged release is actually published, visit the [Releases page](https://github.com/atstaeff/ai-factory-cli/releases). Select the archive matching your CPU and OS, verify its checksum, and extract it.

### Linux (example)

~~~bash
tar -xzf factory_vX.Y.Z_linux_amd64.tar.gz
chmod +x factory
./factory version
./factory scan .
./factory tools
~~~

### Windows PowerShell (example)

~~~powershell
Expand-Archive .\factory_vX.Y.Z_windows_amd64.zip -DestinationPath .\factory-dist
.\factory-dist\factory.exe version
.\factory-dist\factory.exe scan .
.\factory-dist\factory.exe tools
~~~

Replace `vX.Y.Z` with the exact published release version. For convenience, put the extracted executable on your `PATH` according to your enterprise workstation policy.

## Publish the first release without local Git or Go

After the manual-release workflow is merged into `main`, open [Release CLI binaries on GitHub Actions](https://github.com/atstaeff/ai-factory-cli/actions/workflows/release.yml), select **Run workflow**, choose **main**, enter `v0.1.0-rc.2`, and start it. The workflow runs Go tests on GitHub's runner, cross-compiles all six archives, verifies files, creates the tag on the selected main commit, and publishes a prerelease with SHA256SUMS. Check the run's publish job for success; this is not automatic when merging a PR.

Only use a version that does not already exist. If the workflow fails, inspect its logs before attempting a new release. The action is restricted to the reviewed `main` branch.

## Maintainer release procedure (Git CLI alternative)

First merge the release workflow to `main` and ensure it passes CI. Then create a reviewed annotated SemVer tag on the exact approved commit:

~~~bash
git checkout main
git pull --ff-only
git tag -a v0.1.0-rc.2 -m "Factory v0.1.0-rc.2"
git push origin v0.1.0-rc.2
~~~

The tag triggers the workflow, which runs tests, builds six native binaries, and publishes a **prerelease** after successful verification. It refuses to overwrite an existing release and checks that the tag already exists.

:::caution Pre-release scope
v0.1.0-rc.2 is the proposed next prerelease; it is available only after the release workflow's publish job succeeds. The current CLI supports discovery and diagnostics. It cannot autonomously execute BMAD, OpenCode, GitHub or Jira processes.
:::

### Security and distribution

Publishing a GitHub release exposes downloadable binaries according to the repository's visibility and the applicable GitHub permissions. Keep internal adapters and customer settings out of public assets. Prefer testing on a non-production machine. Consider artifact signing and SBOM generation before enterprise distribution.


## Verify publication

After the `Release CLI binaries` workflow runs, open [GitHub Releases](https://github.com/atstaeff/ai-factory-cli/releases). A Git tag alone is not a release. The publishing job must validate **six platform archives plus SHA256SUMS**, change the draft to a published prerelease, and confirm its state.

Expected assets: `linux_amd64`, `linux_arm64`, `windows_amd64`, `windows_arm64`, `darwin_amd64`, `darwin_arm64`, plus `SHA256SUMS`. If a workflow fails, review the failing `verify-and-build` or `publish` job rather than assuming that the presence of a tag means success.

For troubleshooting with an authenticated GitHub CLI:

~~~bash
gh release view v0.1.0-rc.2 --repo atstaeff/ai-factory-cli --json isDraft,isPrerelease,assets
~~~
