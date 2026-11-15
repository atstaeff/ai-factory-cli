---
id: operations
title: Operations and monitoring
---

# Operations and monitoring

The operational goal is a durable background supervisor with a terminal monitor, without requiring a web dashboard.

## Target operator experience

~~~bash
factory run project-a          # planned
factory serve                  # planned
factory tui                    # planned
factory logs ISSUE-142 --follow # planned
~~~

These are **roadmap commands**, not currently implemented.

### Proposed states

`ready` → `running` → `waiting-human` / `failed` / `completed`

A future monitor should show the project, work item, execution step, runtime, model, last heartbeat, retries and decision requests. A task remaining silent beyond a timeout should be reported as potentially stalled rather than silently considered healthy.

### Human handoffs

The issue owner remains assigned. Automation should use an explicit `ai:working` label and post concise progress updates where allowed. A structured `human:required` request would pause only the relevant job, and a permitted reply would resume it via polling or webhooks.

### Recovery

Retry limits, idempotency and persistent checkpoints are mandatory before unattended operation. Logs alone must not act as the job database.
