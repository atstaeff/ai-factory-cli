---
id: security
title: Security and governance
---

# Security and governance

AI Factory is intended for regulated and enterprise environments, but the **current prototype has no hardened sandbox or enterprise security certification**.

## Design requirements

- Honor the company's approved CLIs, internal repositories and model endpoints.
- Enforce restrictions via OS/container permissions, credentials and network controls, not prompt instructions alone.
- Keep each customer/workspace isolated and scope credentials to required operations.
- Regard issue comments, project files, model output and skill contents as untrusted inputs.
- Audit invoked tool, command arguments (redacted), model, policy decisions and result.
- Restrict merges, deployments and other sensitive side effects through explicit, testable approval gates.
- Provide a way to stop jobs and revoke permissions.

## Published documentation

GitHub Pages publishes a static website. **Do not place internal system topology, private model endpoints, company CLI commands, secrets or customer data in these docs or examples.** Pages content may be publicly accessible even when the source repository is private, depending on plan and policy.

For enterprise deployment, assess license, supply-chain and governance requirements independently before installing any optional agent runtime.
