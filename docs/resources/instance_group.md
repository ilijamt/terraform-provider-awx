---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awx_instance_group Resource - terraform-provider-awx"
subcategory: ""
description: |-
  
---

# awx_instance_group (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of this instance group.

### Optional

- `credential` (Number) Credential
- `is_container_group` (Boolean) Indicates whether instances in this group are containerized.Containerized groups have a designated Openshift or Kubernetes cluster.
- `pod_spec_override` (String) Pod spec override
- `policy_instance_list` (String) List of exact-match Instances that will be assigned to this group
- `policy_instance_minimum` (Number) Static minimum number of Instances that will be automatically assign to this group when new instances come online.
- `policy_instance_percentage` (Number) Minimum percentage of all instances that will be automatically assigned to this group when new instances come online.

### Read-Only

- `capacity` (Number)
- `consumed_capacity` (Number)
- `id` (Number) Database ID for this instance group.
- `instances` (Number)
- `jobs_running` (Number) Count of jobs in the running or waiting state that are targeted for this instance group
- `jobs_total` (Number) Count of all jobs that target this instance group
- `percent_capacity_remaining` (Number)


