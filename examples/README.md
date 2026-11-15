# Process examples

Examples demonstrate the proposed **Process-as-Code** format. They are deliberately not wired to external systems and will not run autonomous agents.

| Example | Focus |
| --- | --- |
| `software-engineering/` | BMAD-inspired design, implementation, quality and review |
| `product-management/` | Multi-step concept work with human signoff |
| `enterprise-cli/` | Governed internal CLI in an environment without webhooks |

Inspect names with:

~~~bash
go run ./cmd/factory processes examples
~~~

The current parser extracts only the `name:` frontmatter field; all step definitions are illustrative until a validated execution schema and adapters are shipped.
