---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_host Resource - awx"
subcategory: ""
description: |-
  
---

# awx_host (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `inventory` (Number) Inventory
- `name` (String) Name of this host.

### Optional

- `description` (String) Optional description of this host.
- `enabled` (Boolean) Is this host online and available for running jobs?
- `instance_id` (String) The value used by the remote inventory source to uniquely identify the host
- `variables` (String) Host variables in JSON or YAML format.

### Read-Only

- `id` (Number) Database ID for this host.
- `last_job` (Number)
- `last_job_host_summary` (Number)
