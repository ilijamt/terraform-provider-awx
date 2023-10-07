# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `SESSION_COOKIE_AGE`: Number of seconds that a user is inactive before they will need to login again. (integer)
* `SESSIONS_PER_USER`: Maximum number of simultaneous logged in sessions a user may have. To disable enter -1. (integer)
* `DISABLE_LOCAL_AUTH`: Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration. (boolean)
* `AUTH_BASIC_ENABLED`: Enable HTTP Basic Auth for the API Browser. (boolean)
* `OAUTH2_PROVIDER`: Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds. (nested object)
* `ALLOW_OAUTH2_FOR_EXTERNAL_USERS`: For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off. (boolean)
* `LOGIN_REDIRECT_OVERRIDE`: URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page. (string)
* `AUTHENTICATION_BACKENDS`: List of authentication backends that are enabled based on license features and other authentication settings. (list)
* `SOCIAL_AUTH_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object)
* `SOCIAL_AUTH_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object)
* `SOCIAL_AUTH_USER_FIELDS`: When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login. (list)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `SESSION_COOKIE_AGE`: Number of seconds that a user is inactive before they will need to login again. (integer, required)
* `SESSIONS_PER_USER`: Maximum number of simultaneous logged in sessions a user may have. To disable enter -1. (integer, required)
* `DISABLE_LOCAL_AUTH`: Controls whether users are prevented from using the built-in authentication system. You probably want to do this if you are using an LDAP or SAML integration. (boolean, required)
* `AUTH_BASIC_ENABLED`: Enable HTTP Basic Auth for the API Browser. (boolean, required)
* `OAUTH2_PROVIDER`: Dictionary for customizing OAuth 2 timeouts, available items are `ACCESS_TOKEN_EXPIRE_SECONDS`, the duration of access tokens in the number of seconds, `AUTHORIZATION_CODE_EXPIRE_SECONDS`, the duration of authorization codes in the number of seconds, and `REFRESH_TOKEN_EXPIRE_SECONDS`, the duration of refresh tokens, after expired access tokens, in the number of seconds. (nested object, default=`{&#x27;ACCESS_TOKEN_EXPIRE_SECONDS&#x27;: 31536000000, &#x27;AUTHORIZATION_CODE_EXPIRE_SECONDS&#x27;: 600, &#x27;REFRESH_TOKEN_EXPIRE_SECONDS&#x27;: 2628000}`)
* `ALLOW_OAUTH2_FOR_EXTERNAL_USERS`: For security reasons, users from external auth providers (LDAP, SAML, SSO, Radius, and others) are not allowed to create OAuth2 tokens. To change this behavior, enable this setting. Existing tokens will not be deleted when this setting is toggled off. (boolean, default=`False`)
* `LOGIN_REDIRECT_OVERRIDE`: URL to which unauthorized users will be redirected to log in.  If blank, users will be sent to the login page. (string, default=`""`)

* `SOCIAL_AUTH_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object, default=`None`)
* `SOCIAL_AUTH_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object, default=`None`)
* `SOCIAL_AUTH_USER_FIELDS`: When set to an empty list `[]`, this setting prevents new user accounts from being created. Only users who have previously logged in using social auth or have a user account with a matching email address will be able to login. (list, default=`None`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.