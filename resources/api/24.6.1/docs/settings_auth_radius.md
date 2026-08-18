# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `RADIUS_SERVER`: Hostname/IP of RADIUS server. RADIUS authentication is disabled if this setting is empty. (string)
* `RADIUS_PORT`: Port of RADIUS server. (integer)
* `RADIUS_SECRET`: Shared secret for authenticating to RADIUS server. (string)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `RADIUS_SERVER`: Hostname/IP of RADIUS server. RADIUS authentication is disabled if this setting is empty. (string, default=`""`)
* `RADIUS_PORT`: Port of RADIUS server. (integer, default=`1812`)
* `RADIUS_SECRET`: Shared secret for authenticating to RADIUS server. (string, default=`""`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.