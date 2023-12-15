# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `LOG_AGGREGATOR_HOST`: Hostname/IP where external logs will be sent to. (string)
* `LOG_AGGREGATOR_PORT`: Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator). (integer)
* `LOG_AGGREGATOR_TYPE`: Format messages for the chosen log aggregator. (choice)
    - `None`: ---------
    - `logstash`
    - `splunk`
    - `loggly`
    - `sumologic`
    - `other`
* `LOG_AGGREGATOR_USERNAME`: Username for external log aggregator (if required; HTTP/s only). (string)
* `LOG_AGGREGATOR_PASSWORD`: Password or authentication token for external log aggregator (if required; HTTP/s only). (string)
* `LOG_AGGREGATOR_LOGGERS`: List of loggers that will send HTTP logs to the collector, these can include any or all of: 
awx - service logs
activity_stream - activity stream records
job_events - callback data from Ansible job events
system_tracking - facts gathered from scan jobs
broadcast_websocket - errors pertaining to websockets broadcast metrics
 (list)
* `LOG_AGGREGATOR_INDIVIDUAL_FACTS`: If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing. (boolean)
* `LOG_AGGREGATOR_ENABLED`: Enable sending logs to external log aggregator. (boolean)
* `LOG_AGGREGATOR_TOWER_UUID`: Useful to uniquely identify instances. (string)
* `LOG_AGGREGATOR_PROTOCOL`: Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname. (choice)
    - `https`: HTTPS/HTTP
    - `tcp`: TCP
    - `udp`: UDP
* `LOG_AGGREGATOR_TCP_TIMEOUT`: Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols. (integer)
* `LOG_AGGREGATOR_VERIFY_CERT`: Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is &quot;https&quot;. If enabled, the log handler will verify certificate sent by external log aggregator before establishing connection. (boolean)
* `LOG_AGGREGATOR_LEVEL`: Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting) (choice)
    - `DEBUG`
    - `INFO`
    - `WARNING`
    - `ERROR`
    - `CRITICAL`
* `LOG_AGGREGATOR_ACTION_QUEUE_SIZE`: Defines how large the rsyslog action queue can grow in number of messages stored. This can have an impact on memory utilization. When the queue reaches 75% of this number, the queue will start writing to disk (queue.highWatermark in rsyslog). When it reaches 90%, NOTICE, INFO, and DEBUG messages will start to be discarded (queue.discardMark with queue.discardSeverity=5). (integer)
* `LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB`: Amount of data to store (in gigabytes) if an rsyslog action takes time to process an incoming message (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting on the action (e.g. omhttp). It stores files in the directory specified by LOG_AGGREGATOR_MAX_DISK_USAGE_PATH. (integer)
* `LOG_AGGREGATOR_MAX_DISK_USAGE_PATH`: Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting. (string)
* `LOG_AGGREGATOR_RSYSLOGD_DEBUG`: Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation. (boolean)
* `API_400_ERROR_LOG_FORMAT`: The format of logged messages when an API 4XX error occurs, the following variables will be substituted: 
status_code - The HTTP status code of the error
user_name - The user name attempting to use the API
url_path - The URL path to the API endpoint called
remote_addr - The remote address seen for the user
error - The error set by the api endpoint
Variables need to be in the format {&lt;variable name&gt;}. (string)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `LOG_AGGREGATOR_HOST`: Hostname/IP where external logs will be sent to. (string, default=`""`)
* `LOG_AGGREGATOR_PORT`: Port on Logging Aggregator to send logs to (if required and not provided in Logging Aggregator). (integer, default=`None`)
* `LOG_AGGREGATOR_TYPE`: Format messages for the chosen log aggregator. (choice)
    - `None`: --------- (default)
    - `logstash`
    - `splunk`
    - `loggly`
    - `sumologic`
    - `other`
* `LOG_AGGREGATOR_USERNAME`: Username for external log aggregator (if required; HTTP/s only). (string, default=`""`)
* `LOG_AGGREGATOR_PASSWORD`: Password or authentication token for external log aggregator (if required; HTTP/s only). (string, default=`""`)
* `LOG_AGGREGATOR_LOGGERS`: List of loggers that will send HTTP logs to the collector, these can include any or all of: 
awx - service logs
activity_stream - activity stream records
job_events - callback data from Ansible job events
system_tracking - facts gathered from scan jobs
broadcast_websocket - errors pertaining to websockets broadcast metrics
 (list, default=`[&#x27;awx&#x27;, &#x27;activity_stream&#x27;, &#x27;job_events&#x27;, &#x27;system_tracking&#x27;, &#x27;broadcast_websocket&#x27;]`)
* `LOG_AGGREGATOR_INDIVIDUAL_FACTS`: If set, system tracking facts will be sent for each package, service, or other item found in a scan, allowing for greater search query granularity. If unset, facts will be sent as a single dictionary, allowing for greater efficiency in fact processing. (boolean, default=`False`)
* `LOG_AGGREGATOR_ENABLED`: Enable sending logs to external log aggregator. (boolean, default=`False`)
* `LOG_AGGREGATOR_TOWER_UUID`: Useful to uniquely identify instances. (string, default=`""`)
* `LOG_AGGREGATOR_PROTOCOL`: Protocol used to communicate with log aggregator.  HTTPS/HTTP assumes HTTPS unless http:// is explicitly used in the Logging Aggregator hostname. (choice)
    - `https`: HTTPS/HTTP (default)
    - `tcp`: TCP
    - `udp`: UDP
* `LOG_AGGREGATOR_TCP_TIMEOUT`: Number of seconds for a TCP connection to external log aggregator to timeout. Applies to HTTPS and TCP log aggregator protocols. (integer, default=`5`)
* `LOG_AGGREGATOR_VERIFY_CERT`: Flag to control enable/disable of certificate verification when LOG_AGGREGATOR_PROTOCOL is &quot;https&quot;. If enabled, the log handler will verify certificate sent by external log aggregator before establishing connection. (boolean, default=`True`)
* `LOG_AGGREGATOR_LEVEL`: Level threshold used by log handler. Severities from lowest to highest are DEBUG, INFO, WARNING, ERROR, CRITICAL. Messages less severe than the threshold will be ignored by log handler. (messages under category awx.anlytics ignore this setting) (choice)
    - `DEBUG`
    - `INFO` (default)
    - `WARNING`
    - `ERROR`
    - `CRITICAL`
* `LOG_AGGREGATOR_ACTION_QUEUE_SIZE`: Defines how large the rsyslog action queue can grow in number of messages stored. This can have an impact on memory utilization. When the queue reaches 75% of this number, the queue will start writing to disk (queue.highWatermark in rsyslog). When it reaches 90%, NOTICE, INFO, and DEBUG messages will start to be discarded (queue.discardMark with queue.discardSeverity=5). (integer, default=`131072`)
* `LOG_AGGREGATOR_ACTION_MAX_DISK_USAGE_GB`: Amount of data to store (in gigabytes) if an rsyslog action takes time to process an incoming message (defaults to 1). Equivalent to the rsyslogd queue.maxdiskspace setting on the action (e.g. omhttp). It stores files in the directory specified by LOG_AGGREGATOR_MAX_DISK_USAGE_PATH. (integer, default=`1`)
* `LOG_AGGREGATOR_MAX_DISK_USAGE_PATH`: Location to persist logs that should be retried after an outage of the external log aggregator (defaults to /var/lib/awx). Equivalent to the rsyslogd queue.spoolDirectory setting. (string, default=`"/var/lib/awx"`)
* `LOG_AGGREGATOR_RSYSLOGD_DEBUG`: Enabled high verbosity debugging for rsyslogd.  Useful for debugging connection issues for external log aggregation. (boolean, default=`False`)
* `API_400_ERROR_LOG_FORMAT`: The format of logged messages when an API 4XX error occurs, the following variables will be substituted: 
status_code - The HTTP status code of the error
user_name - The user name attempting to use the API
url_path - The URL path to the API endpoint called
remote_addr - The remote address seen for the user
error - The error set by the api endpoint
Variables need to be in the format {&lt;variable name&gt;}. (string, default=`"status {status_code} received by user {user_name} attempting to access {url_path} from {remote_addr}"`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.