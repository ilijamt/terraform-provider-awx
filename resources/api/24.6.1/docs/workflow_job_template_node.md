# List Workflow Job Template Nodes:

Make a GET request to this resource to retrieve the list of
workflow job template nodes.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of workflow job template nodes
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more workflow job template node records.  

## Results

Each workflow job template node data structure includes the following fields:

* `id`: Database ID for this workflow job template node. (integer)
* `type`: Data type for this workflow job template node. (choice)
* `url`: URL for this workflow job template node. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow job template node was created. (datetime)
* `modified`: Timestamp when this workflow job template node was last modified. (datetime)
* `extra_data`:  (json)
* `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id)
* `scm_branch`:  (string)
* `job_type`:  (choice)
    - `None`: ---------
    - `""`: ---------
    - `run`: Run
    - `check`: Check
* `job_tags`:  (string)
* `skip_tags`:  (string)
* `limit`:  (string)
* `diff_mode`:  (boolean)
* `verbosity`:  (choice)
    - `None`: ---------
    - `0`: 0 (Normal)
    - `1`: 1 (Verbose)
    - `2`: 2 (More Verbose)
    - `3`: 3 (Debug)
    - `4`: 4 (Connection Debug)
    - `5`: 5 (WinRM Debug)
* `execution_environment`: The container image to be used for execution. (id)
* `forks`:  (integer)
* `job_slice_count`:  (integer)
* `timeout`:  (integer)
* `workflow_job_template`:  (id)
* `unified_job_template`:  (id)
* `success_nodes`:  (field)
* `failure_nodes`:  (field)
* `always_nodes`:  (field)
* `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean)
* `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string)



## Sorting

To specify that workflow job template nodes are returned in a particular
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




# Create a Workflow Job Template Node:

Make a POST request to this resource with the following workflow job template node
fields to create a new workflow job template node:









* `extra_data`:  (json, default=`{}`)
* `inventory`: Inventory applied as a prompt, assuming job template prompts for inventory (id, default=``)
* `scm_branch`:  (string, default=`""`)
* `job_type`:  (choice)
    - `None`: --------- (default)
    - `""`: ---------
    - `run`: Run
    - `check`: Check
* `job_tags`:  (string, default=`""`)
* `skip_tags`:  (string, default=`""`)
* `limit`:  (string, default=`""`)
* `diff_mode`:  (boolean, default=`None`)
* `verbosity`:  (choice)
    - `None`: --------- (default)
    - `0`: 0 (Normal)
    - `1`: 1 (Verbose)
    - `2`: 2 (More Verbose)
    - `3`: 3 (Debug)
    - `4`: 4 (Connection Debug)
    - `5`: 5 (WinRM Debug)
* `execution_environment`: The container image to be used for execution. (id, default=``)
* `forks`:  (integer, default=`None`)
* `job_slice_count`:  (integer, default=`None`)
* `timeout`:  (integer, default=`None`)
* `workflow_job_template`:  (id, required)
* `unified_job_template`:  (id, default=``)



* `all_parents_must_converge`: If enabled then the node will only run if all of the parent nodes have met the criteria to reach this node (boolean, default=`False`)
* `identifier`: An identifier for this node that is unique within its workflow. It is copied to workflow job nodes corresponding to this node. (string, default=`"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`)