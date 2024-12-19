# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_CALLBACK_URL`: Create an organization-owned application at https://github.com/organizations/&lt;yourorg&gt;/settings/applications and obtain an OAuth2 key (Client ID) and secret (Client Secret). Provide this URL as the callback URL for your application. (string)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL`: The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details. (string)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL`: The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details. (string)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY`: The OAuth2 key (Client ID) from your GitHub Enterprise organization application. (string)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET`: The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application. (string)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID`: Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/. (string)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:



* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_URL`: The URL for your Github Enterprise instance, e.g.: http(s)://hostname/. Refer to Github Enterprise documentation for more details. (string, default=`""`)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_API_URL`: The API URL for your GitHub Enterprise instance, e.g.: http(s)://hostname/api/v3/. Refer to Github Enterprise documentation for more details. (string, default=`""`)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_KEY`: The OAuth2 key (Client ID) from your GitHub Enterprise organization application. (string, default=`""`)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_SECRET`: The OAuth2 secret (Client Secret) from your GitHub Enterprise organization application. (string, default=`""`)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ID`: Find the numeric team ID using the Github Enterprise API: http://fabian-kostadinov.github.io/2015/01/16/how-to-find-a-github-team-id/. (string, default=`""`)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object, default=`None`)
* `SOCIAL_AUTH_GITHUB_ENTERPRISE_TEAM_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object, default=`None`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.