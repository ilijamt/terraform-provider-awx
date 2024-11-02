# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `ACTIVITY_STREAM_ENABLED`: Enable capturing activity for the activity stream. (boolean)
* `ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC`: Enable capturing activity for the activity stream when running inventory sync. (boolean)
* `ORG_ADMINS_CAN_SEE_ALL_USERS`: Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization. (boolean)
* `MANAGE_ORGANIZATION_AUTH`: Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration. (boolean)
* `TOWER_URL_BASE`: This setting is used by services like notifications to render a valid url to the service. (string)
* `REMOTE_HOST_HEADERS`: HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as &quot;HTTP_X_FORWARDED_FOR&quot;, if behind a reverse proxy. See the &quot;Proxy Support&quot; section of the AAP Installation guide for more details. (list)
* `PROXY_IP_ALLOWED_LIST`: If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally&#x27;) (list)
* `CSRF_TRUSTED_ORIGINS`: If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values.  (list)
* `LICENSE`: The license controls which features and functionality are enabled. Use /api/v2/config/ to update or change the license. (nested object)
* `REDHAT_USERNAME`: This username is used to send data to Automation Analytics (string)
* `REDHAT_PASSWORD`: This password is used to send data to Automation Analytics (string)
* `SUBSCRIPTIONS_USERNAME`: This username is used to retrieve subscription and content information (string)
* `SUBSCRIPTIONS_PASSWORD`: This password is used to retrieve subscription and content information (string)
* `AUTOMATION_ANALYTICS_URL`: This setting is used to to configure the upload URL for data collection for Automation Analytics. (string)
* `INSTALL_UUID`:  (string)
* `DEFAULT_CONTROL_PLANE_QUEUE_NAME`:  (string)
* `DEFAULT_EXECUTION_QUEUE_NAME`:  (string)
* `DEFAULT_EXECUTION_ENVIRONMENT`: The Execution Environment to be used when one has not been configured for a job template. (field)
* `CUSTOM_VENV_PATHS`: Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line. (list)
* `INSIGHTS_TRACKING_STATE`: Enables the service to gather data on automation and send it to Automation Analytics. (boolean)
* `AUTOMATION_ANALYTICS_LAST_GATHER`:  (datetime)
* `AUTOMATION_ANALYTICS_LAST_ENTRIES`:  (string)
* `AUTOMATION_ANALYTICS_GATHER_INTERVAL`: Interval (in seconds) between data gathering. (integer)
* `IS_K8S`: Indicates whether the instance is part of a kubernetes-based deployment. (boolean)
* `UI_NEXT`: Enable preview of new user interface. (boolean)
* `SUBSCRIPTION_USAGE_MODEL`:  (choice)
    - `""`: Default model for AWX - no subscription. Deletion of host_metrics will not be considered for purposes of managed host counting
    - `unique_managed_hosts`: Usage based on unique managed nodes in a large historical time frame and delete functionality for no longer used managed nodes
* `CLEANUP_HOST_METRICS_LAST_TS`:  (datetime)
* `HOST_METRIC_SUMMARY_TASK_LAST_TS`:  (datetime)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `ACTIVITY_STREAM_ENABLED`: Enable capturing activity for the activity stream. (boolean, required)
* `ACTIVITY_STREAM_ENABLED_FOR_INVENTORY_SYNC`: Enable capturing activity for the activity stream when running inventory sync. (boolean, required)
* `ORG_ADMINS_CAN_SEE_ALL_USERS`: Controls whether any Organization Admin can view all users and teams, even those not associated with their Organization. (boolean, required)
* `MANAGE_ORGANIZATION_AUTH`: Controls whether any Organization Admin has the privileges to create and manage users and teams. You may want to disable this ability if you are using an LDAP or SAML integration. (boolean, required)
* `TOWER_URL_BASE`: This setting is used by services like notifications to render a valid url to the service. (string, required)
* `REMOTE_HOST_HEADERS`: HTTP headers and meta keys to search to determine remote host name or IP. Add additional items to this list, such as &quot;HTTP_X_FORWARDED_FOR&quot;, if behind a reverse proxy. See the &quot;Proxy Support&quot; section of the AAP Installation guide for more details. (list, required)
* `PROXY_IP_ALLOWED_LIST`: If the service is behind a reverse proxy/load balancer, use this setting to configure the proxy IP addresses from which the service should trust custom REMOTE_HOST_HEADERS header values. If this setting is an empty list (the default), the headers specified by REMOTE_HOST_HEADERS will be trusted unconditionally&#x27;) (list, default=`[]`)
* `CSRF_TRUSTED_ORIGINS`: If the service is behind a reverse proxy/load balancer, use this setting to configure the schema://addresses from which the service should trust Origin header values.  (list, default=`[]`)

* `REDHAT_USERNAME`: This username is used to send data to Automation Analytics (string, default=`""`)
* `REDHAT_PASSWORD`: This password is used to send data to Automation Analytics (string, default=`""`)
* `SUBSCRIPTIONS_USERNAME`: This username is used to retrieve subscription and content information (string, default=`""`)
* `SUBSCRIPTIONS_PASSWORD`: This password is used to retrieve subscription and content information (string, default=`""`)
* `AUTOMATION_ANALYTICS_URL`: This setting is used to to configure the upload URL for data collection for Automation Analytics. (string, default=`"https://example.com"`)



* `DEFAULT_EXECUTION_ENVIRONMENT`: The Execution Environment to be used when one has not been configured for a job template. (field, default=`None`)
* `CUSTOM_VENV_PATHS`: Paths where Tower will look for custom virtual environments (in addition to /var/lib/awx/venv/). Enter one path per line. (list, default=`[]`)
* `INSIGHTS_TRACKING_STATE`: Enables the service to gather data on automation and send it to Automation Analytics. (boolean, default=`False`)
* `AUTOMATION_ANALYTICS_LAST_GATHER`:  (datetime, default=`None`)
* `AUTOMATION_ANALYTICS_LAST_ENTRIES`:  (string, default=`""`)
* `AUTOMATION_ANALYTICS_GATHER_INTERVAL`: Interval (in seconds) between data gathering. (integer, default=`14400`)

* `UI_NEXT`: Enable preview of new user interface. (boolean, default=`True`)
* `SUBSCRIPTION_USAGE_MODEL`:  (choice)
    - `""`: Default model for AWX - no subscription. Deletion of host_metrics will not be considered for purposes of managed host counting (default)
    - `unique_managed_hosts`: Usage based on unique managed nodes in a large historical time frame and delete functionality for no longer used managed nodes
* `CLEANUP_HOST_METRICS_LAST_TS`:  (datetime, required)
* `HOST_METRIC_SUMMARY_TASK_LAST_TS`:  (datetime, required)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.