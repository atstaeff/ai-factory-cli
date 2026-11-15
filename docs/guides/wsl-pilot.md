---
id: wsl-pilot
title: Test Factory in WSL
---

# First Brownfield Pilot: WSL + OpenCode + GitHub Issues + BMAD

This guide tests the **real local tooling on your own WSL machine**, without claiming that Factory is already an autonomous orchestrator.

## 1. Build or download

The release workflow is still in [PR #3](https://github.com/atstaeff/ai-factory-cli/pull/3). **A binary download is available only once that workflow has been merged and an approved release tag published.** Check [Releases](https://github.com/atstaeff/ai-factory-cli/releases) before downloading.

To test this pilot branch immediately with Go 1.23+:

~~~bash
git clone https://github.com/atstaeff/ai-factory-cli.git
cd ai-factory-cli
git fetch origin feat/wsl-brownfield-pilot
git checkout feat/wsl-brownfield-pilot
go test ./...
go build -o factory ./cmd/factory
./factory help
~~~

The executable is called **factory**.

## 2. Inspect your existing WSL environment

~~~bash
command -v opencode
opencode --version
opencode run --help
command -v gh
gh auth status
git --version
~~~

These commands run on your own machine. Factory cannot see your existing OpenCode, Agents, company configuration or local repositories remotely.

## 3. Choose one disposable or clean brownfield repository

Do not start with all repositories. Keep the original repository unchanged and preferably use a **separate checkout** without production credentials.

~~~bash
PROJECT="$HOME/projects/my-example-repo"
cd "$PROJECT"
git status --short
ls -la AGENTS.md .factory _bmad .agents/skills .opencode/agents 2>/dev/null || true
~~~

BMAD installations and custom skills vary by version. An existing \`_bmad/\` folder is evidence of files, **not** proof that an agent can invoke them. Read [official BMAD guidance](https://docs.bmad-method.org/start/install-bmad/) for the project's installed version; do not update it automatically.

## 4. Dry run with an existing issue

~~~bash
/path/to/ai-factory-cli/factory pilot \
  --project "$PROJECT" \
  --repo OWNER/REPOSITORY \
  --issue 123
~~~

This checks inputs and project markers, **without launching \`gh\` or OpenCode**.

## 5. First real integration test

Make sure your installed OpenCode configuration permits the chosen agent and model. The pilot's default OpenCode agent is \`plan\` but **a prompt is not a security boundary**; enforce read-only permissions using your local approved environment.

~~~bash
/path/to/ai-factory-cli/factory pilot \
  --project "$PROJECT" \
  --repo OWNER/REPOSITORY \
  --issue 123 \
  --agent plan \
  --execute
~~~

This genuinely invokes **\`gh issue view\`** to read the issue and **\`opencode run --agent plan\`** to analyze it in the project's directory. It uses project \`AGENTS.md\` and installed methodology files through the agent's normal discovery; it does **not** guarantee BMAD activation. It does not intentionally write an issue comment, create a PR or implement the issue.

Check afterwards:

~~~bash
cd "$PROJECT"
git status --short
~~~

If anything changed unexpectedly, stop and inspect it. Do not grant \`--auto\` tool approvals. On a corporate system, obtain approval for the local model provider/network and ensure the necessary issue access is already authorized.

## 6. Multiple projects

Run the same pilot separately for each project with its own \`--project\`, \`--repo\` and \`--issue\`. There is no background scheduler, issue scanning or multi-project parallel execution yet.

## What success looks like

- Factory reads an existing GitHub issue through the installed, authenticated GitHub CLI.
- OpenCode starts with the correct repository workspace and uses the configured \`plan\` agent.
- It identifies local BMAD/OpenSpec/custom instructions when present, explains the relevant next step and any missing capability.
- The working tree remains unchanged.
- Factory reports failures instead of falsely claiming a completed workflow.

**Not yet implemented:** automatic workflow graph selection, process execution, Jira CLI adapters, status synchronization, persistence, quality gates, or automatic PR creation.

## Troubleshooting

If \`gh auth status\` fails, authenticate through the company's approved GitHub mechanism. If \`opencode run --help\` lacks the expected flags, check your installed OpenCode version before proceeding. If BMAD is not detected, inspect the actual installed skill path instead of installing a different version blindly.
