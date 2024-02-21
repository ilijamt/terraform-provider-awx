# List Hosts:

Make a GET request to this resource to retrieve the list of
hosts.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of hosts
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more host records.  

## Results

Each host data structure includes the following fields:

* `id`: Database ID for this host. (integer)
* `type`: Data type for this host. (choice)
* `url`: URL for this host. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this host was created. (datetime)
* `modified`: Timestamp when this host was last modified. (datetime)
* `name`: Name of this host. (string)
* `description`: Optional description of this host. (string)
* `inventory`:  (id)
* `enabled`: Is this host online and available for running jobs? (boolean)
* `instance_id`: The value used by the remote inventory source to uniquely identify the host (string)
* `variables`: Host variables in JSON or YAML format. (json)
* `has_active_failures`:  (field)
* `has_inventory_sources`:  (field)
* `last_job`:  (id)
* `last_job_host_summary`:  (id)
* `ansible_facts_modified`: The date and time ansible_facts was last modified. (datetime)



## Sorting

To specify that hosts are returned in a particular
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




`host_filter` is available on this endpoint. The filter supports: relational queries, `and` `or` boolean logic, as well as expression grouping via `()`.

    ?host_filter=name=my_host
    ?host_filter=name="my host" or name=my_host
    ?host_filter=groups__name="my group"
    ?host_filter=name=my_host and groups__name="my group"
    ?host_filter=name=my_host and groups__name="my group"
    ?host_filter=(name=my_host and groups__name="my group") or (name=my_host2 and groups__name=my_group2)

`host_filter` can also be used to query JSON data in the related `ansible_facts`. `__` may be used to traverse JSON dictionaries. `[]` may be used to traverse JSON arrays.

    ?host_filter=ansible_facts__ansible_processor_vcpus=8
    ?host_filter=ansible_facts__ansible_processor_vcpus=8 and name="my_host" and ansible_facts__ansible_lo__ipv6[]__scope=host