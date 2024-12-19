# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `SOCIAL_AUTH_GOOGLE_OAUTH2_CALLBACK_URL`: Provide this URL as the callback URL for your application as part of your registration process. Refer to the documentation for more detail. (string)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_KEY`: The OAuth2 key from your web application. (string)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET`: The OAuth2 secret from your web application. (string)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS`: Update this setting to restrict the domains who are allowed to login using Google OAuth2. (list)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS`: Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail. (nested object)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:



* `SOCIAL_AUTH_GOOGLE_OAUTH2_KEY`: The OAuth2 key from your web application. (string, default=`""`)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_SECRET`: The OAuth2 secret from your web application. (string, default=`""`)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_WHITELISTED_DOMAINS`: Update this setting to restrict the domains who are allowed to login using Google OAuth2. (list, default=`[]`)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_AUTH_EXTRA_ARGUMENTS`: Extra arguments for Google OAuth2 login. You can restrict it to only allow a single domain to authenticate, even if the user is logged in with multple Google accounts. Refer to the documentation for more detail. (nested object, default=`{}`)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object, default=`None`)
* `SOCIAL_AUTH_GOOGLE_OAUTH2_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object, default=`None`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.