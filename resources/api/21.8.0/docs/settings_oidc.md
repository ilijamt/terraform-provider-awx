# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `SOCIAL_AUTH_OIDC_KEY`: The OIDC key (Client ID) from your IDP. (string)
* `SOCIAL_AUTH_OIDC_SECRET`: The OIDC secret (Client Secret) from your IDP. (string)
* `SOCIAL_AUTH_OIDC_OIDC_ENDPOINT`: The URL for your OIDC provider including the path up to /.well-known/openid-configuration (string)
* `SOCIAL_AUTH_OIDC_VERIFY_SSL`: Verify the OIDV provider ssl certificate. (boolean)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `SOCIAL_AUTH_OIDC_KEY`: The OIDC key (Client ID) from your IDP. (string, default=`""`)
* `SOCIAL_AUTH_OIDC_SECRET`: The OIDC secret (Client Secret) from your IDP. (string, default=`""`)
* `SOCIAL_AUTH_OIDC_OIDC_ENDPOINT`: The URL for your OIDC provider including the path up to /.well-known/openid-configuration (string, default=`""`)
* `SOCIAL_AUTH_OIDC_VERIFY_SSL`: Verify the OIDV provider ssl certificate. (boolean, default=`True`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.