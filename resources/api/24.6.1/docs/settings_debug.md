# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `AWX_CLEANUP_PATHS`: Enable or Disable TMP Dir cleanup (boolean)
* `AWX_REQUEST_PROFILE`: Debug web request python timing (boolean)
* `RECEPTOR_RELEASE_WORK`: Release receptor work (boolean)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `AWX_CLEANUP_PATHS`: Enable or Disable TMP Dir cleanup (boolean, default=`True`)
* `AWX_REQUEST_PROFILE`: Debug web request python timing (boolean, default=`False`)
* `RECEPTOR_RELEASE_WORK`: Release receptor work (boolean, default=`True`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.