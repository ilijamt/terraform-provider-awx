# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `TACACSPLUS_HOST`: Hostname of TACACS+ server. (string)
* `TACACSPLUS_PORT`: Port number of TACACS+ server. (integer)
* `TACACSPLUS_SECRET`: Shared secret for authenticating to TACACS+ server. (string)
* `TACACSPLUS_SESSION_TIMEOUT`: TACACS+ session timeout value in seconds, 0 disables timeout. (integer)
* `TACACSPLUS_AUTH_PROTOCOL`: Choose the authentication protocol used by TACACS+ client. (choice)
    - `ascii`
    - `pap`
* `TACACSPLUS_REM_ADDR`: Enable the client address sending by TACACS+ client. (boolean)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `TACACSPLUS_HOST`: Hostname of TACACS+ server. (string, default=`""`)
* `TACACSPLUS_PORT`: Port number of TACACS+ server. (integer, default=`49`)
* `TACACSPLUS_SECRET`: Shared secret for authenticating to TACACS+ server. (string, default=`""`)
* `TACACSPLUS_SESSION_TIMEOUT`: TACACS+ session timeout value in seconds, 0 disables timeout. (integer, default=`5`)
* `TACACSPLUS_AUTH_PROTOCOL`: Choose the authentication protocol used by TACACS+ client. (choice)
    - `ascii` (default)
    - `pap`
* `TACACSPLUS_REM_ADDR`: Enable the client address sending by TACACS+ client. (boolean, default=`False`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.