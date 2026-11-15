---
name: factory-development
version: 1
status: draft
---

# Factory Development Process

This process is a **draft**, not an executable workflow definition.

## Intake
Inspect the issue and existing repository. Preserve brownfield changes.

## Plan
Write a concise implementation plan. Identify security-sensitive operations and get approval when needed.

## Implement
Make small changes with tests and clear commit history.

## Verify
Run Go tests and check that documented commands match the implementation.

## Deliver
Open a pull request. Do not merge automatically.

The future engine will validate executable steps and gate definitions using a versioned schema.
