---
id: brownfield
title: Brownfield onboarding
---

# Brownfield onboarding

The first priority is to **understand and preserve** existing projects.

## What is possible today?

~~~bash
factory scan /path/to/project
factory processes /path/to/project
factory doctor /path/to/project
~~~

`scan` detects familiar configuration paths, but does not classify their contents or install anything.

## Proposed onboarding lifecycle

1. **Discover** local project conventions (`AGENTS.md`, `_bmad`, `openspec`, scripts and process files).
2. **Inspect** installed versions and local customizations using read-only adapters.
3. **Plan** a minimal, reviewable Factory integration.
4. **Approve** any changes to existing project files.
5. **Apply** changes in a branch rather than rewriting the project in place.
6. **Verify** connections, process compatibility and security gates.

Future `factory onboard` should support `--dry-run` and leave the project fully usable after offboarding.

## Avoiding duplicate ownership

Permanent architecture, design and specifications stay in the repository; Jira or GitHub carries work status, ownership, comments and links to deliverables.

See [Project examples](../examples.md).
