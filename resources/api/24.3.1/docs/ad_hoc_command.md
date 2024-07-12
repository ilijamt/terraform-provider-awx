# List Ad Hoc Commands:

Make a GET request to this resource to retrieve the list of
ad hoc commands.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of ad hoc commands
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more ad hoc command records.  

## Results

Each ad hoc command data structure includes the following fields:

* `id`: Database ID for this ad hoc command. (integer)
* `type`: Data type for this ad hoc command. (choice)
* `url`: URL for this ad hoc command. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this ad hoc command was created. (datetime)
* `modified`: Timestamp when this ad hoc command was last modified. (datetime)
* `name`: Name of this ad hoc command. (string)
* `launch_type`:  (choice)
    - `manual`: Manual
    - `relaunch`: Relaunch
    - `callback`: Callback
    - `scheduled`: Scheduled
    - `dependency`: Dependency
    - `workflow`: Workflow
    - `webhook`: Webhook
    - `sync`: Sync
    - `scm`: SCM Update
* `status`:  (choice)
    - `new`: New
    - `pending`: Pending
    - `waiting`: Waiting
    - `running`: Running
    - `successful`: Successful
    - `failed`: Failed
    - `error`: Error
    - `canceled`: Canceled
* `execution_environment`: The container image to be used for execution. (id)
* `failed`:  (boolean)
* `started`: The date and time the job was queued for starting. (datetime)
* `finished`: The date and time the job finished execution. (datetime)
* `canceled_on`: The date and time when the cancel request was sent. (datetime)
* `elapsed`: Elapsed time in seconds that the job ran. (decimal)
* `job_explanation`: A status field to indicate the state of the job if it wasn&#x27;t able to run and capture stdout (string)
* `execution_node`: The node the job executed on. (string)
* `controller_node`: The instance that managed the execution environment. (string)
* `launched_by`:  (field)
* `work_unit_id`: The Receptor work unit ID associated with this job. (string)
* `job_type`:  (choice)
    - `run`: Run
    - `check`: Check
* `inventory`:  (id)
* `limit`:  (string)
* `credential`:  (id)
* `module_name`:  (choice)
    - `command`
    - `shell`
    - `yum`
    - `apt`
    - `apt_key`
    - `apt_repository`
    - `apt_rpm`
    - `service`
    - `group`
    - `user`
    - `mount`
    - `ping`
    - `selinux`
    - `setup`
    - `win_ping`
    - `win_service`
    - `win_updates`
    - `win_group`
    - `win_user`
* `module_args`:  (string)
* `forks`:  (integer)
* `verbosity`:  (choice)
    - `0`: 0 (Normal)
    - `1`: 1 (Verbose)
    - `2`: 2 (More Verbose)
    - `3`: 3 (Debug)
    - `4`: 4 (Connection Debug)
    - `5`: 5 (WinRM Debug)
* `extra_vars`:  (string)
* `become_enabled`:  (boolean)
* `diff_mode`:  (boolean)



## Sorting

To specify that ad hoc commands are returned in a particular
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




# Create an Ad Hoc Command:

Make a POST request to this resource with the following ad hoc command
fields to create a new ad hoc command:












* `execution_environment`: The container image to be used for execution. (id, default=``)










* `job_type`:  (choice)
    - `run`: Run (default)
    - `check`: Check
* `inventory`:  (id, default=``)
* `limit`:  (string, default=`""`)
* `credential`:  (id, default=``)
* `module_name`:  (choice)
    - `command` (default)
    - `shell`
    - `yum`
    - `apt`
    - `apt_key`
    - `apt_repository`
    - `apt_rpm`
    - `service`
    - `group`
    - `user`
    - `mount`
    - `ping`
    - `selinux`
    - `setup`
    - `win_ping`
    - `win_service`
    - `win_updates`
    - `win_group`
    - `win_user`
* `module_args`:  (string, default=`""`)
* `forks`:  (integer, default=`0`)
* `verbosity`:  (choice)
    - `0`: 0 (Normal) (default)
    - `1`: 1 (Verbose)
    - `2`: 2 (More Verbose)
    - `3`: 3 (Debug)
    - `4`: 4 (Connection Debug)
    - `5`: 5 (WinRM Debug)
* `extra_vars`:  (string, default=`""`)
* `become_enabled`:  (boolean, default=`False`)
* `diff_mode`:  (boolean, default=`False`)