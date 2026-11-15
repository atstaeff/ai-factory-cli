---
id: integrations
title: Integrations and adapters
---

# Integrations and adapters

Factory should connect existing systems, not force teams to replace them.

## Work-source adapters

A work-source adapter will eventually normalize issue identity, title, status, assignee, comments and permitted transitions. The same process could be used with GitHub Issues or a restricted enterprise Jira CLI.

A custom CLI adapter should receive **structured arguments** rather than interpolating issue content into shell strings. It must document executable location, approved operations, output schema, timeout and exit codes.

## Method adapters

BMAD and OpenSpec may coexist with custom processes. Their installed project definitions are authoritative. Never assume a skill exists based solely on a guessed command name.

## Runtime adapters

OpenCode, Claude Code and other tools may support different model IDs and thinking modes. The runtime adapter must validate its effective model config and log the *actual* runtime/model used.

## Event delivery

Where webhooks are unavailable (for example in restricted Jira environments), use a polling adapter with backoff and durable event cursors. Human replies must be matched to an outstanding authorized decision, not treated as blanket approval.

## Capability discovery

The current CLI discovers local files only. The interfaces and adapters described here are **planned**.
