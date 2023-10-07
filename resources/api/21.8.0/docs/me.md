# Retrieve Information about the current User

Make a GET request to retrieve user information about the current user.

One result should be returned containing the following fields:

* `id`: Database ID for this user. (integer)
* `type`: Data type for this user. (choice)
* `url`: URL for this user. (string)
* `related`: Data structure with URLs of related resources. (object)
* `summary_fields`: Data structure with name/description for related resources.  The output for some objects may be limited for performance reasons. (object)
* `created`: Timestamp when this user was created. (datetime)
* `modified`: Timestamp when this user was last modified. (datetime)
* `username`: Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. (string)
* `first_name`:  (string)
* `last_name`:  (string)
* `email`:  (string)
* `is_superuser`: Designates that this user has all permissions without explicitly assigning them. (boolean)
* `is_system_auditor`:  (boolean)

* `ldap_dn`:  (string)
* `last_login`:  (datetime)
* `external_account`: Set if the account is managed by an external service (field)



Use the primary URL for the user (/api/v2/users/N/) to modify the user.