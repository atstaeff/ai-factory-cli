---
id: bmad
title: BMAD for teams
---

# BMAD for teams

BMAD is an **optional project methodology**, not the Factory core.

For the current team setup instructions, see [docs/BMAD.md in the source repository](https://github.com/atstaeff/ai-factory-cli/blob/main/docs/BMAD.md). Use the documented stable version only after reviewing the installation process and company policies.

## Team workflow

1. Install BMAD using its approved installer and select the coding tools that teammates actually use.
2. Review generated shared files, skill definitions and repository configuration.
3. Commit non-secret, portable process definitions so teammates use the same methodology.
4. Pin the expected BMAD version and run a quick team compatibility check.
5. Keep customizations project-local, rather than recreating BMAD in Factory.

:::warning
BMAD has **not** been installed by the Factory prototype. The current CLI does not invoke BMAD skills, choose the next phase or execute AI work. Those behaviors require a future method adapter.
:::

A custom Factory step such as `uses: bmad/code-review` is conceptual until the installed BMAD version is discovered and its real entrypoint is validated.
