---
id: examples
title: Example processes
---

# Example processes

The `examples/` directory contains representative, *non-executable* process blueprints.

| Example | Purpose |
| --- | --- |
| [Software engineering](https://github.com/atstaeff/ai-factory-cli/tree/main/examples/software-engineering) | Existing BMAD engineering workflow with review gates |
| [Product management](https://github.com/atstaeff/ai-factory-cli/tree/main/examples/product-management) | Idea to validated concept with a human decision |
| [Enterprise CLI](https://github.com/atstaeff/ai-factory-cli/tree/main/examples/enterprise-cli) | Company-approved CLI and polling without webhooks |

All process step identifiers are **proposed Factory references**, not automatically supported commands. The current CLI only discovers the process name.

## Try the read-only scanner

~~~bash
go run ./cmd/factory processes examples
go run ./cmd/factory scan examples/software-engineering
~~~

See [Process-as-code](./concepts/processes.md) to understand the proposed format.
