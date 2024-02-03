# List Inventory Sources:

Make a GET request to this resource to retrieve the list of
inventory sources.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of inventory sources
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more inventory source records.  

## Results

Each inventory source data structure includes the following fields:

* `id`: Database ID for this inventory source. (integer)
* `type`: Data type for this inventory source. (choice)
* `url`: URL for this inventory source. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this inventory source was created. (datetime)
* `modified`: Timestamp when this inventory source was last modified. (datetime)
* `name`: Name of this inventory source. (string)
* `description`: Optional description of this inventory source. (string)
* `source`:  (choice)
    - `file`: File, Directory or Script
    - `constructed`: Template additional groups and hostvars at runtime
    - `scm`: Sourced from a Project
    - `ec2`: Amazon EC2
    - `gce`: Google Compute Engine
    - `azure_rm`: Microsoft Azure Resource Manager
    - `vmware`: VMware vCenter
    - `satellite6`: Red Hat Satellite 6
    - `openstack`: OpenStack
    - `rhv`: Red Hat Virtualization
    - `controller`: Red Hat Ansible Automation Platform
    - `insights`: Red Hat Insights
* `source_path`:  (string)
* `source_vars`: Inventory source variables in YAML or JSON format. (string)
* `scm_branch`: Inventory source SCM branch. Project default used if blank. Only allowed if project allow_override field is set to true. (string)
* `credential`: Cloud credential to use for inventory updates. (integer)
* `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string)
* `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2020-08-04T18:13:04+00:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string)
* `host_filter`: This field is deprecated and will be removed in a future release. Regex where only matching hosts will be imported. (string)
* `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean)
* `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean)
* `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string)
* `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer)
* `verbosity`:  (choice)
    - `0`: 0 (WARNING)
    - `1`: 1 (INFO)
    - `2`: 2 (DEBUG)
* `limit`: Enter host, group or pattern match (string)
* `last_job_run`:  (datetime)
* `last_job_failed`:  (boolean)
* `next_job_run`:  (datetime)
* `status`:  (choice)
    - `new`: New
    - `pending`: Pending
    - `waiting`: Waiting
    - `running`: Running
    - `successful`: Successful
    - `failed`: Failed
    - `error`: Error
    - `canceled`: Canceled
    - `never updated`: Never Updated
    - `none`: No External Source
* `execution_environment`: The container image to be used for execution. (id)
* `inventory`:  (id)
* `update_on_launch`:  (boolean)
* `update_cache_timeout`:  (integer)
* `source_project`: Project containing inventory file used as source. (id)
* `last_update_failed`:  (boolean)
* `last_updated`:  (datetime)



## Sorting

To specify that inventory sources are returned in a particular
order, use the `order_by` query string parameter on the GET request.

    ?order_by=name

Prefix the field name with a dash `-` to sort in reverse:

    ?order_by=-name

Multiple sorting fields may be specified by separating the field names with a
comma `,`:

    ?order_by=name,some_other_field

## Pagination

Use the `page_size` query string parameter to change the number of results
returned for each request.  Use the `page` query string parameter to retrieve
a particular page of results.

    ?page_size=100&page=2

The `previous` and `next` links returned with the results will set these query
string parameters automatically.

## Searching

Use the `search` query string parameter to perform a case-insensitive search
within all designated text fields of a model.

    ?search=findme

(_Added in Ansible Tower 3.1.0_) Search across related fields:

    ?related__search=findme

Note: If you want to provide more than one search term, multiple
search fields with the same key, like `?related__search=foo&related__search=bar`,
will be ORed together. Terms separated by commas, like `?related__search=foo,bar`
will be ANDed together.

## Filtering

Any additional query string parameters may be used to filter the list of
results returned to those matching a given value.  Only fields and relations
that exist in the database may be used for filtering.  Any special characters
in the specified value should be url-encoded. For example:

    ?field=value%20xyz

Fields may also span relations, only for fields and relationships defined in
the database:

    ?other__field=value

To exclude results matching certain criteria, prefix the field parameter with
`not__`:

    ?not__field=value

By default, all query string filters are AND'ed together, so
only the results matching *all* filters will be returned.  To combine results
matching *any* one of multiple criteria, prefix each query string parameter
with `or__`:

    ?or__field=value&or__field=othervalue
    ?or__not__field=value&or__field=othervalue

(_Added in Ansible Tower 1.4.5_) The default AND filtering applies all filters
simultaneously to each related object being filtered across database
relationships.  The chain filter instead applies filters separately for each
related object. To use, prefix the query string parameter with `chain__`:

    ?chain__related__field=value&chain__related__field2=othervalue
    ?chain__not__related__field=value&chain__related__field2=othervalue

If the first query above were written as
`?related__field=value&related__field2=othervalue`, it would return only the
primary objects where the *same* related object satisfied both conditions.  As
written using the chain filter, it would return the intersection of primary
objects matching each condition.

Field lookups may also be used for more advanced queries, by appending the
lookup to the field name:

    ?field__lookup=value

The following field lookups are supported:

* `exact`: Exact match (default lookup if not specified).
* `iexact`: Case-insensitive version of `exact`.
* `contains`: Field contains value.
* `icontains`: Case-insensitive version of `contains`.
* `startswith`: Field starts with value.
* `istartswith`: Case-insensitive version of `startswith`.
* `endswith`: Field ends with value.
* `iendswith`: Case-insensitive version of `endswith`.
* `regex`: Field matches the given regular expression.
* `iregex`: Case-insensitive version of `regex`.
* `gt`: Greater than comparison.
* `gte`: Greater than or equal to comparison.
* `lt`: Less than comparison.
* `lte`: Less than or equal to comparison.
* `isnull`: Check whether the given field or related object is null; expects a
  boolean value.
* `in`: Check whether the given field's value is present in the list provided;
  expects a list of items.

Boolean values may be specified as `True` or `1` for true, `False` or `0` for
false (both case-insensitive).

Null values may be specified as `None` or `Null` (both case-insensitive),
though it is preferred to use the `isnull` lookup to explicitly check for null
values.

Lists (for the `in` lookup) may be specified as a comma-separated list of
values.

(_Added in Ansible Tower 3.1.0_) Filtering based on the requesting user's
level of access by query string parameter.

* `role_level`: Level of role to filter on, such as `admin_role`




# Create an Inventory Source:

Make a POST request to this resource with the following inventory source
fields to create a new inventory source:









* `name`: Name of this inventory source. (string, required)
* `description`: Optional description of this inventory source. (string, default=`""`)
* `source`:  (choice)
    - `file`: File, Directory or Script
    - `constructed`: Template additional groups and hostvars at runtime
    - `scm`: Sourced from a Project
    - `ec2`: Amazon EC2
    - `gce`: Google Compute Engine
    - `azure_rm`: Microsoft Azure Resource Manager
    - `vmware`: VMware vCenter
    - `satellite6`: Red Hat Satellite 6
    - `openstack`: OpenStack
    - `rhv`: Red Hat Virtualization
    - `controller`: Red Hat Ansible Automation Platform
    - `insights`: Red Hat Insights
* `source_path`:  (string, default=`""`)
* `source_vars`: Inventory source variables in YAML or JSON format. (string, default=`""`)
* `scm_branch`: Inventory source SCM branch. Project default used if blank. Only allowed if project allow_override field is set to true. (string, default=`""`)
* `credential`: Cloud credential to use for inventory updates. (integer, default=`None`)
* `enabled_var`: Retrieve the enabled state from the given dict of host variables. The enabled variable may be specified as &quot;foo.bar&quot;, in which case the lookup will traverse into nested dicts, equivalent to: from_dict.get(&quot;foo&quot;, {}).get(&quot;bar&quot;, default) (string, default=`""`)
* `enabled_value`: Only used when enabled_var is set. Value when the host is considered enabled. For example if enabled_var=&quot;status.power_state&quot;and enabled_value=&quot;powered_on&quot; with host variables:{   &quot;status&quot;: {     &quot;power_state&quot;: &quot;powered_on&quot;,     &quot;created&quot;: &quot;2020-08-04T18:13:04+00:00&quot;,     &quot;healthy&quot;: true    },    &quot;name&quot;: &quot;foobar&quot;,    &quot;ip_address&quot;: &quot;192.168.2.1&quot;}The host would be marked enabled. If power_state where any value other than powered_on then the host would be disabled when imported. If the key is not found then the host will be enabled (string, default=`""`)
* `host_filter`: This field is deprecated and will be removed in a future release. Regex where only matching hosts will be imported. (string, default=`""`)
* `overwrite`: Overwrite local groups and hosts from remote inventory source. (boolean, default=`False`)
* `overwrite_vars`: Overwrite local variables from remote inventory source. (boolean, default=`False`)

* `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`)
* `verbosity`:  (choice)
    - `0`: 0 (WARNING)
    - `1`: 1 (INFO) (default)
    - `2`: 2 (DEBUG)
* `limit`: Enter host, group or pattern match (string, default=`""`)




* `execution_environment`: The container image to be used for execution. (id, default=``)
* `inventory`:  (id, required)
* `update_on_launch`:  (boolean, default=`False`)
* `update_cache_timeout`:  (integer, default=`0`)
* `source_project`: Project containing inventory file used as source. (id, default=``)