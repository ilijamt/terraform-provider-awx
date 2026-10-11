---
page_title: "awx Provider"
description: |-
  Manage AWX resources with Terraform. Generated from the AWX {{ .ApiVersion }} API.
---

# awx Provider

The awx provider manages [AWX](https://github.com/ansible/awx) objects such as organizations, inventories, credentials, projects and job templates. Its resources and data sources are generated from the `/api/v2/` API of AWX {{ .ApiVersion }}, so some fields may differ on other AWX versions.

Each release tag encodes that AWX version and a build counter in its patch number, so `v24.6.103` is AWX 24.6.1, build 3. Pin a range such as `~> 24.6.103` to stay on one AWX release.

## Example Usage

{{ "{{ tffile .ExampleFile }}" }}

## Authentication

Set `hostname` and log in with either:

- a `token`, or
- a `username` and `password`.

Setting both a token and a username or password is an error, whether they come from the configuration or the environment.

### Environment variables

The provider reads any argument you leave out of the configuration from the environment. If both variables for an argument are set, the `TOWER_*` one wins.

| Argument | Environment variables |
|---|---|
| `hostname` | `TOWER_HOST`, `AWX_HOST` |
| `token` | `TOWER_AUTH_TOKEN`, `AWX_AUTH_TOKEN` |
| `username` | `TOWER_USERNAME`, `AWX_USERNAME` |
| `password` | `TOWER_PASSWORD`, `AWX_PASSWORD` |
| `verify_ssl` | `TOWER_VERIFY_SSL`, `AWX_VERIFY_SSL` |

{{ "{{ .SchemaMarkdown | trimspace }}" }}
