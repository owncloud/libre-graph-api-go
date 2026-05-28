# Libre Graph API Go Client

<!-- OSPO-managed README | Generated: 2026-04-16 | v2 -->

[![License](https://img.shields.io/badge/License-Apache--2.0-blue.svg)](LICENSE) [![ownCloud OSPO](https://img.shields.io/badge/OSPO-ownCloud-blue)](https://kiteworks.com/opensource) [![Docker Hub](https://img.shields.io/docker/pulls/owncloud)](https://hub.docker.com/r/owncloud/ocis)

Auto-generated Go client library for the Libre Graph API, produced by the OpenAPI Generator. It provides idiomatic Go types, API method wrappers, and configurable HTTP transport for accessing all Libre Graph endpoints -- including drives, users, groups, education resources, shares, tags, and application role assignments -- used internally by ownCloud Infinite Scale.

## Getting Started

Follow the steps below to install and use the Go client.

### Installation

```bash
go get github.com/owncloud/libre-graph-api-go
```

### Usage

```go
import libregraph "github.com/owncloud/libre-graph-api-go"

cfg := libregraph.NewConfiguration()
client := libregraph.NewAPIClient(cfg)
```

## Documentation

- [Libre Graph API Documentation](https://owncloud.dev/libre-graph-api/)
- [Go Package Documentation](https://pkg.go.dev/github.com/owncloud/libre-graph-api-go)

## Part of ownCloud Infinite Scale

This is the Go SDK used internally by [oCIS](https://github.com/owncloud/ocis) to interact with Libre Graph endpoints. It is generated from the [libre-graph-api](https://github.com/owncloud/libre-graph-api) OpenAPI spec.

> **Note:** This library contains generated code. Do not edit the generated files directly; instead modify the OpenAPI spec and regenerate.

This component is part of the [oCIS Docker image](https://hub.docker.com/r/owncloud/ocis).

## Community & Support

**[Star](https://github.com/owncloud/libre-graph-api-go)** this repo and **Watch** for release notifications!

- [ownCloud Website](https://owncloud.com)
- [Community Discussions](https://github.com/orgs/owncloud/discussions)
- [Matrix Chat](https://app.element.io/#/room/#owncloud:matrix.org)
- [Documentation](https://doc.owncloud.com)
- [Enterprise Support](https://owncloud.com/contact-us/)
- [OSPO Home](https://kiteworks.com/opensource)

## Contributing

We welcome contributions! Please read the [Contributing Guidelines](CONTRIBUTING.md)
and our [Code of Conduct](CODE_OF_CONDUCT.md) before getting started.

### Workflow

- **Rebase Early, Rebase Often!** We use a rebase workflow. Always rebase on the target branch before submitting a PR.
- **Dependabot**: Automated dependency updates are managed via Dependabot. Review and merge dependency PRs promptly.
- **Signed Commits**: All commits **must** be PGP/GPG signed. See [GitHub's signing guide](https://docs.github.com/en/authentication/managing-commit-signature-verification).
- **DCO Sign-off**: Every commit must carry a `Signed-off-by` line:
  ```
  git commit -s -S -m "your commit message"
  ```
- **GitHub Actions Policy**: Workflows may only use actions that are (a) owned by `owncloud`, (b) created by GitHub (`actions/*`), or (c) verified in the GitHub Marketplace.

## Security

**Do not open a public GitHub issue for security vulnerabilities.**

Report vulnerabilities at **<https://security.owncloud.com>** -- see [SECURITY.md](SECURITY.md).

Bug bounty: [YesWeHack ownCloud Program](https://yeswehack.com/programs/owncloud-bug-bounty-program)

## License

This project is licensed under the [Apache-2.0](LICENSE).

## About the ownCloud OSPO

The [Kiteworks Open Source Program Office](https://kiteworks.com/opensource), operating under
the [ownCloud](https://owncloud.com) brand, launched on May 5, 2026, to steward the open source
ecosystem around ownCloud's products. The OSPO ensures transparent governance, license compliance,
community health, and sustainable collaboration between the open source community and
[Kiteworks](https://www.kiteworks.com), which acquired ownCloud in 2023.

- **OSPO Home**: <https://kiteworks.com/opensource>
- **GitHub**: <https://github.com/owncloud>
- **ownCloud**: <https://owncloud.com>

For questions about the OSPO or licensing, contact ospo@kiteworks.com.

> **License status:** This repository is already licensed under Apache-2.0 -- the OSPO target license.
> No migration is required.
