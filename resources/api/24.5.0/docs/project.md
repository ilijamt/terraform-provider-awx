# List Projects:

Make a GET request to this resource to retrieve the list of
projects.

The resulting data structure contains:

    {
        "count": 99,
        "next": null,
        "previous": null,
        "results": [
            ...
        ]
    }

The `count` field indicates the total number of projects
found for the given query.  The `next` and `previous` fields provides links to
additional results if there are more than will fit on a single page.  The
`results` list contains zero or more project records.  

## Results

Each project data structure includes the following fields:

* `id`: Database ID for this project. (integer)
* `type`: Data type for this project. (choice)
* `url`: URL for this project. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this project was created. (datetime)
* `modified`: Timestamp when this project was last modified. (datetime)
* `name`: Name of this project. (string)
* `description`: Optional description of this project. (string)
* `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string)
* `scm_type`: Specifies the source control system used to store the project. (choice)
    - `""`: Manual
    - `git`: Git
    - `svn`: Subversion
    - `insights`: Red Hat Insights
    - `archive`: Remote Archive
* `scm_url`: The location where the project is stored. (string)
* `scm_branch`: Specific branch, tag or commit to checkout. (string)
* `scm_refspec`: For git projects, an additional refspec to fetch. (string)
* `scm_clean`: Discard any local changes before syncing the project. (boolean)
* `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean)
* `scm_delete_on_update`: Delete the project before syncing. (boolean)
* `credential`:  (id)
* `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer)
* `scm_revision`: The last revision fetched by a project update (string)
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
    - `ok`: OK
    - `missing`: Missing
* `organization`: The organization used to determine access to this template. (id)
* `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean)
* `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer)
* `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean)
* `custom_virtualenv`: Local absolute file path containing a custom Python virtualenv to use (string)
* `default_environment`: The default execution environment for jobs run using this project. (id)
* `signature_validation_credential`: An optional credential used for validating files in the project against unexpected changes. (id)
* `last_update_failed`:  (boolean)
* `last_updated`:  (datetime)



## Sorting

To specify that projects are returned in a particular
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




# Create a Project:

Make a POST request to this resource with the following project
fields to create a new project:









* `name`: Name of this project. (string, required)
* `description`: Optional description of this project. (string, default=`""`)
* `local_path`: Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project. (string, default=`""`)
* `scm_type`: Specifies the source control system used to store the project. (choice)
    - `""`: Manual (default)
    - `git`: Git
    - `svn`: Subversion
    - `insights`: Red Hat Insights
    - `archive`: Remote Archive
* `scm_url`: The location where the project is stored. (string, default=`""`)
* `scm_branch`: Specific branch, tag or commit to checkout. (string, default=`""`)
* `scm_refspec`: For git projects, an additional refspec to fetch. (string, default=`""`)
* `scm_clean`: Discard any local changes before syncing the project. (boolean, default=`False`)
* `scm_track_submodules`: Track submodules latest commits on defined branch. (boolean, default=`False`)
* `scm_delete_on_update`: Delete the project before syncing. (boolean, default=`False`)
* `credential`:  (id, default=``)
* `timeout`: The amount of time (in seconds) to run before the task is canceled. (integer, default=`0`)





* `organization`: The organization used to determine access to this template. (id, default=``)
* `scm_update_on_launch`: Update the project when a job is launched that uses the project. (boolean, default=`False`)
* `scm_update_cache_timeout`: The number of seconds after the last project update ran that a new project update will be launched as a job dependency. (integer, default=`0`)
* `allow_override`: Allow changing the SCM branch or revision in a job template that uses this project. (boolean, default=`False`)

* `default_environment`: The default execution environment for jobs run using this project. (id, default=``)
* `signature_validation_credential`: An optional credential used for validating files in the project against unexpected changes. (id, default=``)