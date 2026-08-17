# Security Policy

Security is treated as a first-class engineering concern in this repository because the curriculum includes authentication, authorization, APIs, networking, databases, and distributed systems.

## Scope

This policy covers:

- vulnerable production-style projects under `projects/`;
- insecure-by-design examples that could accidentally be mistaken for safe defaults;
- repository tooling and CI configuration;
- dependency vulnerabilities that materially affect the repository;
- accidental exposure of credentials or sensitive configuration.

## Educational examples

Some lessons intentionally demonstrate insecure or vulnerable behavior so that learners can understand the failure mode.

Such examples should be clearly labeled and must never contain:

- real credentials;
- real API keys;
- private certificates or private keys;
- production connection strings;
- personal customer data;
- live authentication tokens.

An intentionally vulnerable example is educational material, not production guidance.

## Reporting a vulnerability

Please do not publish sensitive vulnerability details in a public issue before the maintainer has had an opportunity to investigate.

Use the private security-reporting mechanism provided by the GitHub repository. Include:

- affected path or component;
- vulnerability type;
- reproduction steps;
- security impact;
- suggested mitigation, when known.

If GitHub private reporting is not enabled for a fork or local copy, contact the repository owner through a private channel before public disclosure.

## Secret handling

Before committing, inspect changes for:

```text
API_KEY=
PASSWORD=
SECRET=
PRIVATE_KEY=
TOKEN=
DATABASE_URL=
```

Use placeholders and local environment configuration instead.

Recommended local patterns include:

```text
.env.local
config.local.yaml
```

and ensuring such files are excluded by `.gitignore`.

## Dependency hygiene

For meaningful dependency changes:

```bash
go mod tidy
go mod verify
go list -m all
```

Review new dependencies before accepting them into a project. Consider maintenance health, license compatibility, security history, transitive dependencies, and whether the standard library already provides the required capability.

## Secure development principles

Examples and projects should prefer:

- input validation at trust boundaries;
- least privilege;
- explicit timeouts;
- context cancellation;
- safe error messages;
- secure password hashing;
- parameterized SQL;
- bounded resource usage;
- safe secret handling;
- observable authentication and authorization failures.

Security should be taught as a system property, not as a collection of isolated snippets.
