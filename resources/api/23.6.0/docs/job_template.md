# List Job Templates:

Make a GET request to this resource to retrieve the list of
job templates.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of job templates
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more job template records.  

## Results

Each job template data structure includes the following fields:

* `id`: Database ID for this job template. (integer)
* `type`: Data type for this job template. (choice)
* `url`: URL for this job template. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this job template was created. (datetime)
* `modified`: Timestamp when this job template was last modified. (datetime)
* `name`: Name of this job template. (string)
* `description`: Optional description of this job template. (string)
* `job_type`:  (choice)
    - `run`: Run
    - `check`: Check
* `inventory`:  (id)
* `project`:  (id)
* `playbook`:  (string)
* `scm_branch`: Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true. (string)
* `forks`:  (integer)
* `limit`:  (string)
* `verbosity`:  (choice)
    - `0`: 0 (Normal)
    - `1`: 1 (Verbose)
    - `2`: 2 (More Verbose)
    - `3`: 3 (Debug)
    - `4`: 4 (Connection Debug)
    - `5`: 5 (WinRM Debug)
* `extra_vars`:  (json)
* `job_tags`:  (string)
* `force_handlers`:  (boolean)
* `skip_tags`:  (string)
* `start_at_task`:  (string)
* `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer)
* `use_fact_cache`: If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible. (boolean)
* `organization`: The organization used to determine access to this template. (id)
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
* `execution_environment`: The container image to be used for execution. (id)
* `host_config_key`:  (string)
* `ask_scm_branch_on_launch`:  (boolean)
* `ask_diff_mode_on_launch`:  (boolean)
* `ask_variables_on_launch`:  (boolean)
* `ask_limit_on_launch`:  (boolean)
* `ask_tags_on_launch`:  (boolean)
* `ask_skip_tags_on_launch`:  (boolean)
* `ask_job_type_on_launch`:  (boolean)
* `ask_verbosity_on_launch`:  (boolean)
* `ask_inventory_on_launch`:  (boolean)
* `ask_credential_on_launch`:  (boolean)
* `ask_execution_environment_on_launch`:  (boolean)
* `ask_labels_on_launch`:  (boolean)
* `ask_forks_on_launch`:  (boolean)
* `ask_job_slice_count_on_launch`:  (boolean)
* `ask_timeout_on_launch`:  (boolean)
* `ask_instance_groups_on_launch`:  (boolean)
* `survey_enabled`:  (boolean)
* `become_enabled`:  (boolean)
* `diff_mode`: If enabled, textual changes made to any templated files on the host are shown in the standard output (boolean)
* `allow_simultaneous`:  (boolean)
* `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string)
* `job_slice_count`: The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1. (integer)
* `webhook_service`: Service that webhook requests will be accepted from (choice)
    - `""`: ---------
    - `github`: GitHub
    - `gitlab`: GitLab
* `webhook_credential`: Personal Access Token for posting back the status to the service API (id)
* `prevent_instance_group_fallback`: If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied. (boolean)



## Sorting

To specify that job templates are returned in a particular
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




# Create a Job Template:

Make a POST request to this resource with the following job template
fields to create a new job template:









* `name`: Name of this job template. (string, required)
* `description`: Optional description of this job template. (string, default=`""`)
* `job_type`:  (choice)
    - `run`: Run (default)
    - `check`: Check
* `inventory`:  (id, default=``)
* `project`:  (id, default=``)
* `playbook`:  (string, default=`""`)
* `scm_branch`: Branch to use in job run. Project default used if blank. Only allowed if project allow_override field is set to true. (string, default=`""`)
* `forks`:  (integer, default=`0`)
* `limit`:  (string, default=`""`)
* `verbosity`:  (choice)
    - `0`: 0 (Normal) (default)
    - `1`: 1 (Verbose)
    - `2`: 2 (More Verbose)
    - `3`: 3 (Debug)
    - `4`: 4 (Connection Debug)
    - `5`: 5 (WinRM Debug)
* `extra_vars`:  (json, default=``)
* `job_tags`:  (string, default=`""`)
* `force_handlers`:  (boolean, default=`False`)
* `skip_tags`:  (string, default=`""`)
* `start_at_task`:  (string, default=`""`)
* `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`)
* `use_fact_cache`: If enabled, the service will act as an Ansible Fact Cache Plugin; persisting facts at the end of a playbook run to the database and caching facts for use by Ansible. (boolean, default=`False`)





* `execution_environment`: The container image to be used for execution. (id, default=``)
* `host_config_key`:  (string, default=`""`)
* `ask_scm_branch_on_launch`:  (boolean, default=`False`)
* `ask_diff_mode_on_launch`:  (boolean, default=`False`)
* `ask_variables_on_launch`:  (boolean, default=`False`)
* `ask_limit_on_launch`:  (boolean, default=`False`)
* `ask_tags_on_launch`:  (boolean, default=`False`)
* `ask_skip_tags_on_launch`:  (boolean, default=`False`)
* `ask_job_type_on_launch`:  (boolean, default=`False`)
* `ask_verbosity_on_launch`:  (boolean, default=`False`)
* `ask_inventory_on_launch`:  (boolean, default=`False`)
* `ask_credential_on_launch`:  (boolean, default=`False`)
* `ask_execution_environment_on_launch`:  (boolean, default=`False`)
* `ask_labels_on_launch`:  (boolean, default=`False`)
* `ask_forks_on_launch`:  (boolean, default=`False`)
* `ask_job_slice_count_on_launch`:  (boolean, default=`False`)
* `ask_timeout_on_launch`:  (boolean, default=`False`)
* `ask_instance_groups_on_launch`:  (boolean, default=`False`)
* `survey_enabled`:  (boolean, default=`False`)
* `become_enabled`:  (boolean, default=`False`)
* `diff_mode`: If enabled, textual changes made to any templated files on the host are shown in the standard output (boolean, default=`False`)
* `allow_simultaneous`:  (boolean, default=`False`)

* `job_slice_count`: The number of jobs to slice into at runtime. Will cause the Job Template to launch a workflow if value is greater than 1. (integer, default=`1`)
* `webhook_service`: Service that webhook requests will be accepted from (choice)
    - `""`: ---------
    - `github`: GitHub
    - `gitlab`: GitLab
* `webhook_credential`: Personal Access Token for posting back the status to the service API (id, default=``)
* `prevent_instance_group_fallback`: If enabled, the job template will prevent adding any inventory or organization instance groups to the list of preferred instances groups to run on.If this setting is enabled and you provided an empty list, the global instance groups will be applied. (boolean, default=`False`)