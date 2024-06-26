---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_credential Data Source - awx"
subcategory: ""
description: |-
  
---

# awx_credential (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (Number) Database ID for this credential.
- `name` (String) Name of this credential.
- `team` (Number) Write-only field used to add team to owner role. If provided, do not give either user or organization. Only valid for creation.
- `user` (Number) Write-only field used to add user to owner role. If provided, do not give either team or organization. Only valid for creation.

### Read-Only

- `cloud` (Boolean) Cloud
- `credential_type` (Number) Specify the type of credential you want to create. Refer to the documentation for details on each type.
- `description` (String) Optional description of this credential.
- `inputs` (String) Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.
- `kind` (String) Kind
- `kubernetes` (Boolean) Kubernetes
- `managed` (Boolean) Managed
- `organization` (Number) Inherit permissions from organization roles. If provided on creation, do not give either user or team.
