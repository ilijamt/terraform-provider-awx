---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_credential Resource - awx"
subcategory: ""
description: |-
  
---

# awx_credential (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `credential_type` (Number) Specify the type of credential you want to create. Refer to the documentation for details on each type.
- `name` (String) Name of this credential.

### Optional

- `description` (String) Optional description of this credential.
- `inputs` (String) Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.
- `organization` (Number) Inherit permissions from organization roles. If provided on creation, do not give either user or team.

### Read-Only

- `cloud` (Boolean)
- `id` (Number) Database ID for this credential.
- `kind` (String)
- `kubernetes` (Boolean)
- `managed` (Boolean)
