# Software engineering example

A conceptual BMAD-backed process for a brownfield application.

**Input:** GitHub or Jira engineering issue.

**Method:** A real installation of BMAD at the project-selected version; Factory should resolve the actual installed capabilities, not assume the example names are valid.

**Output:** Changed files, tests, review and a pull request after future adapter support is implemented.

**Human gate:** Approval before merge, plus escalations for unclear requirements.

To inspect the file:

~~~bash
go run ./cmd/factory scan examples/software-engineering
go run ./cmd/factory processes examples/software-engineering
~~~

The current `factory` binary only discovers and displays files. It does not call any BMAD command.
