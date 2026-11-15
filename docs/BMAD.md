# BMAD collaboration setup

BMAD is **optional** and is not bundled with Factory. This repository supports contributors using OpenCode, Claude Code or other coding tools, but the BMAD installer must be run in each approved environment.

## Stable team baseline

Official documentation currently recommends BMAD v6.12 for stability while v7 is a preview. Review security policy and licensing before executing external installers.

From the repository root, for an approved environment with Node.js 20.12+, Python 3.10+ and uv:

```sh
npx bmad-method@6.12.4 install
```

Select the coding tools your team uses (for example OpenCode and Claude Code) and keep project-specific configuration under version control where appropriate. The installer is interactive; commit its generated shared BMAD files and tool-specific skill definitions only after reviewing them for secrets and machine-specific paths. Never blindly copy another user's credentials.

To work with BMAD, use its installed `bmad-help` capability. Different BMAD versions have different installation layouts and workflows; Factory will eventually discover these through a dedicated method adapter.

## Current state

**BMAD has not been installed into this branch.** This guide prepares team onboarding; it is not a claim that a real installation or workflow execution has been validated. Do not treat this repository as BMAD-enabled until the installer has been run, the files reviewed and committed, and contributors can discover its skills.

Official installation guide: https://docs.bmad-method.org/start/install-bmad/
