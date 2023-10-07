# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `AD_HOC_COMMANDS`: List of modules allowed to be used by ad-hoc jobs. (list)
* `ALLOW_JINJA_IN_EXTRA_VARS`: Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to &quot;template&quot; or &quot;never&quot;. (choice)
    - `always`: Always
    - `never`: Never
    - `template`: Only On Job Template Definitions
* `AWX_ISOLATION_BASE_PATH`: The directory in which the service will create new temporary directories for job execution and isolation (such as credential files). (string)
* `AWX_ISOLATION_SHOW_PATHS`: List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]].  (list)
* `AWX_TASK_ENV`: Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending. (nested object)
* `GALAXY_TASK_ENV`: Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git. (nested object)
* `PROJECT_UPDATE_VVV`: Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates. (boolean)
* `AWX_ROLES_ENABLED`: Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects. (boolean)
* `AWX_COLLECTIONS_ENABLED`: Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects. (boolean)
* `AWX_SHOW_PLAYBOOK_LINKS`: Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself. (boolean)
* `AWX_MOUNT_ISOLATED_PATHS_ON_K8S`: Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible.  (boolean)
* `GALAXY_IGNORE_CERTS`: If set to true, certificate validation will not be done when installing content from any Galaxy server. (boolean)
* `STDOUT_MAX_BYTES_DISPLAY`: Maximum Size of Standard Output in bytes to display before requiring the output be downloaded. (integer)
* `EVENT_STDOUT_MAX_BYTES_DISPLAY`: Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated. (integer)
* `MAX_WEBSOCKET_EVENT_RATE`: Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit. (integer)
* `SCHEDULE_MAX_JOBS`: Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created. (integer)
* `AWX_ANSIBLE_CALLBACK_PLUGINS`: List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line. (list)
* `DEFAULT_JOB_TIMEOUT`: Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this. (integer)
* `DEFAULT_JOB_IDLE_TIMEOUT`: If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed. (integer)
* `DEFAULT_INVENTORY_UPDATE_TIMEOUT`: Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this. (integer)
* `DEFAULT_PROJECT_UPDATE_TIMEOUT`: Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this. (integer)
* `ANSIBLE_FACT_CACHE_TIMEOUT`: Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed. (integer)
* `MAX_FORKS`: Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied. (integer)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `AD_HOC_COMMANDS`: List of modules allowed to be used by ad-hoc jobs. (list, default=`[&#x27;command&#x27;, &#x27;shell&#x27;, &#x27;yum&#x27;, &#x27;apt&#x27;, &#x27;apt_key&#x27;, &#x27;apt_repository&#x27;, &#x27;apt_rpm&#x27;, &#x27;service&#x27;, &#x27;group&#x27;, &#x27;user&#x27;, &#x27;mount&#x27;, &#x27;ping&#x27;, &#x27;selinux&#x27;, &#x27;setup&#x27;, &#x27;win_ping&#x27;, &#x27;win_service&#x27;, &#x27;win_updates&#x27;, &#x27;win_group&#x27;, &#x27;win_user&#x27;]`)
* `ALLOW_JINJA_IN_EXTRA_VARS`: Ansible allows variable substitution via the Jinja2 templating language for --extra-vars. This poses a potential security risk where users with the ability to specify extra vars at job launch time can use Jinja2 templates to run arbitrary Python.  It is recommended that this value be set to &quot;template&quot; or &quot;never&quot;. (choice, required)
    - `always`: Always
    - `never`: Never
    - `template`: Only On Job Template Definitions (default)
* `AWX_ISOLATION_BASE_PATH`: The directory in which the service will create new temporary directories for job execution and isolation (such as credential files). (string, required)
* `AWX_ISOLATION_SHOW_PATHS`: List of paths that would otherwise be hidden to expose to isolated jobs. Enter one path per line. Volumes will be mounted from the execution node to the container. The supported format is HOST-DIR[:CONTAINER-DIR[:OPTIONS]].  (list, default=`[&#x27;/etc/pki/ca-trust:/etc/pki/ca-trust:O&#x27;, &#x27;/usr/share/pki:/usr/share/pki:O&#x27;]`)
* `AWX_TASK_ENV`: Additional environment variables set for playbook runs, inventory updates, project updates, and notification sending. (nested object, default=`{}`)
* `GALAXY_TASK_ENV`: Additional environment variables set for invocations of ansible-galaxy within project updates. Useful if you must use a proxy server for ansible-galaxy but not git. (nested object, required)
* `PROJECT_UPDATE_VVV`: Adds the CLI -vvv flag to ansible-playbook runs of project_update.yml used for project updates. (boolean, required)
* `AWX_ROLES_ENABLED`: Allows roles to be dynamically downloaded from a requirements.yml file for SCM projects. (boolean, default=`True`)
* `AWX_COLLECTIONS_ENABLED`: Allows collections to be dynamically downloaded from a requirements.yml file for SCM projects. (boolean, default=`True`)
* `AWX_SHOW_PLAYBOOK_LINKS`: Follow symbolic links when scanning for playbooks. Be aware that setting this to True can lead to infinite recursion if a link points to a parent directory of itself. (boolean, default=`False`)
* `AWX_MOUNT_ISOLATED_PATHS_ON_K8S`: Expose paths via hostPath for the Pods created by a Container Group. HostPath volumes present many security risks, and it is a best practice to avoid the use of HostPaths when possible.  (boolean, default=`False`)
* `GALAXY_IGNORE_CERTS`: If set to true, certificate validation will not be done when installing content from any Galaxy server. (boolean, default=`False`)
* `STDOUT_MAX_BYTES_DISPLAY`: Maximum Size of Standard Output in bytes to display before requiring the output be downloaded. (integer, required)
* `EVENT_STDOUT_MAX_BYTES_DISPLAY`: Maximum Size of Standard Output in bytes to display for a single job or ad hoc command event. `stdout` will end with `…` when truncated. (integer, required)
* `MAX_WEBSOCKET_EVENT_RATE`: Maximum number of messages to update the UI live job output with per second. Value of 0 means no limit. (integer, default=`30`)
* `SCHEDULE_MAX_JOBS`: Maximum number of the same job template that can be waiting to run when launching from a schedule before no more are created. (integer, required)
* `AWX_ANSIBLE_CALLBACK_PLUGINS`: List of paths to search for extra callback plugins to be used when running jobs. Enter one path per line. (list, default=`[]`)
* `DEFAULT_JOB_TIMEOUT`: Maximum time in seconds to allow jobs to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual job template will override this. (integer, default=`0`)
* `DEFAULT_JOB_IDLE_TIMEOUT`: If no output is detected from ansible in this number of seconds the execution will be terminated. Use value of 0 to indicate that no idle timeout should be imposed. (integer, default=`0`)
* `DEFAULT_INVENTORY_UPDATE_TIMEOUT`: Maximum time in seconds to allow inventory updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual inventory source will override this. (integer, default=`0`)
* `DEFAULT_PROJECT_UPDATE_TIMEOUT`: Maximum time in seconds to allow project updates to run. Use value of 0 to indicate that no timeout should be imposed. A timeout set on an individual project will override this. (integer, default=`0`)
* `ANSIBLE_FACT_CACHE_TIMEOUT`: Maximum time, in seconds, that stored Ansible facts are considered valid since the last time they were modified. Only valid, non-stale, facts will be accessible by a playbook. Note, this does not influence the deletion of ansible_facts from the database. Use a value of 0 to indicate that no timeout should be imposed. (integer, default=`0`)
* `MAX_FORKS`: Saving a Job Template with more than this number of forks will result in an error. When set to 0, no limit is applied. (integer, default=`200`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.