---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_constructed_inventories Resource - awx"
subcategory: ""
description: |-
  
---

# awx_constructed_inventories (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of this inventory.
- `organization` (Number) Organization containing this inventory.

### Optional

- `description` (String) Optional description of this inventory.
- `limit` (String) The limit to restrict the returned hosts for the related auto-created inventory source, special to constructed inventory.
- `prevent_instance_group_fallback` (Boolean) If enabled, the inventory will prevent adding any organization instance groups to the list of preferred instances groups to run associated job templates on.If this setting is enabled and you provided an empty list, the global instance groups will be applied.
- `source_vars` (String) The source_vars for the related auto-created inventory source, special to constructed inventory.
- `update_cache_timeout` (Number) The cache timeout for the related auto-created inventory source, special to constructed inventory
- `variables` (String) Inventory variables in JSON or YAML format.
- `verbosity` (Number) The verbosity level for the related auto-created inventory source, special to constructed inventory

### Read-Only

- `has_active_failures` (Boolean) This field is deprecated and will be removed in a future release. Flag indicating whether any hosts in this inventory have failed.
- `has_inventory_sources` (Boolean) This field is deprecated and will be removed in a future release. Flag indicating whether this inventory has any external inventory sources.
- `hosts_with_active_failures` (Number) This field is deprecated and will be removed in a future release. Number of hosts in this inventory with active failures.
- `id` (Number) Database ID for this inventory.
- `inventory_sources_with_failures` (Number) Number of external inventory sources in this inventory with failures.
- `kind` (String) Kind of inventory being represented.
- `pending_deletion` (Boolean) Flag indicating the inventory is being deleted.
- `total_groups` (Number) This field is deprecated and will be removed in a future release. Total number of groups in this inventory.
- `total_hosts` (Number) This field is deprecated and will be removed in a future release. Total number of hosts in this inventory.
- `total_inventory_sources` (Number) Total number of external inventory sources configured within this inventory.