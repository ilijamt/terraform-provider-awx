# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `BULK_JOB_MAX_LAUNCH`: Max jobs to allow bulk jobs to launch (integer)
* `BULK_HOST_MAX_CREATE`: Max number of hosts to allow to be created in a single bulk action (integer)
* `BULK_HOST_MAX_DELETE`: Max number of hosts to allow to be deleted in a single bulk action (integer)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `BULK_JOB_MAX_LAUNCH`: Max jobs to allow bulk jobs to launch (integer, default=`100`)
* `BULK_HOST_MAX_CREATE`: Max number of hosts to allow to be created in a single bulk action (integer, default=`100`)
* `BULK_HOST_MAX_DELETE`: Max number of hosts to allow to be deleted in a single bulk action (integer, default=`250`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.