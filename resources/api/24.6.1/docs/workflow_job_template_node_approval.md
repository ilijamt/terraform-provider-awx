# Retrieve a Workflow Approval Template:

Make GET request to this resource to retrieve a single workflow approval template
record containing the following fields:

* `id`: Database ID for this workflow approval template. (integer)
* `type`: Data type for this workflow approval template. (choice)
* `url`: URL for this workflow approval template. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this workflow approval template was created. (datetime)
* `modified`: Timestamp when this workflow approval template was last modified. (datetime)
* `name`: Name of this workflow approval template. (string)
* `description`: Optional description of this workflow approval template. (string)
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
    - `none`: No External Source
    - `updating`: Updating
* `execution_environment`: The container image to be used for execution. (id)
* `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer)





# Update a Workflow Approval Template:

Make a PUT or PATCH request to this resource to update this
workflow approval template.  The following fields may be modified:









* `name`: Name of this workflow approval template. (string, required)
* `description`: Optional description of this workflow approval template. (string, default=`""`)




* `execution_environment`: The container image to be used for execution. (id, default=``)
* `timeout`: The amount of time (in seconds) before the approval node expires and fails. (integer, default=`0`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Workflow Approval Template:

Make a DELETE request to this resource to delete this workflow approval template.