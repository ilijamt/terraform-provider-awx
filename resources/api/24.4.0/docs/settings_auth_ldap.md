# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `AUTH_LDAP_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string)
* `AUTH_LDAP_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string)
* `AUTH_LDAP_BIND_PASSWORD`: Password used to bind LDAP user account. (string)
* `AUTH_LDAP_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean)
* `AUTH_LDAP_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object)
* `AUTH_LDAP_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list)
* `AUTH_LDAP_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string)
* `AUTH_LDAP_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object)
* `AUTH_LDAP_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list)
* `AUTH_LDAP_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType`
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object)
* `AUTH_LDAP_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string)
* `AUTH_LDAP_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string)
* `AUTH_LDAP_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object)
* `AUTH_LDAP_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_1_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string)
* `AUTH_LDAP_1_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string)
* `AUTH_LDAP_1_BIND_PASSWORD`: Password used to bind LDAP user account. (string)
* `AUTH_LDAP_1_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean)
* `AUTH_LDAP_1_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object)
* `AUTH_LDAP_1_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list)
* `AUTH_LDAP_1_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string)
* `AUTH_LDAP_1_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object)
* `AUTH_LDAP_1_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list)
* `AUTH_LDAP_1_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType`
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_1_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object)
* `AUTH_LDAP_1_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string)
* `AUTH_LDAP_1_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string)
* `AUTH_LDAP_1_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object)
* `AUTH_LDAP_1_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_1_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_2_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string)
* `AUTH_LDAP_2_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string)
* `AUTH_LDAP_2_BIND_PASSWORD`: Password used to bind LDAP user account. (string)
* `AUTH_LDAP_2_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean)
* `AUTH_LDAP_2_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object)
* `AUTH_LDAP_2_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list)
* `AUTH_LDAP_2_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string)
* `AUTH_LDAP_2_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object)
* `AUTH_LDAP_2_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list)
* `AUTH_LDAP_2_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType`
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_2_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object)
* `AUTH_LDAP_2_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string)
* `AUTH_LDAP_2_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string)
* `AUTH_LDAP_2_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object)
* `AUTH_LDAP_2_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_2_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_3_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string)
* `AUTH_LDAP_3_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string)
* `AUTH_LDAP_3_BIND_PASSWORD`: Password used to bind LDAP user account. (string)
* `AUTH_LDAP_3_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean)
* `AUTH_LDAP_3_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object)
* `AUTH_LDAP_3_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list)
* `AUTH_LDAP_3_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string)
* `AUTH_LDAP_3_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object)
* `AUTH_LDAP_3_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list)
* `AUTH_LDAP_3_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType`
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_3_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object)
* `AUTH_LDAP_3_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string)
* `AUTH_LDAP_3_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string)
* `AUTH_LDAP_3_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object)
* `AUTH_LDAP_3_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_3_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_4_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string)
* `AUTH_LDAP_4_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string)
* `AUTH_LDAP_4_BIND_PASSWORD`: Password used to bind LDAP user account. (string)
* `AUTH_LDAP_4_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean)
* `AUTH_LDAP_4_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object)
* `AUTH_LDAP_4_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list)
* `AUTH_LDAP_4_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string)
* `AUTH_LDAP_4_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object)
* `AUTH_LDAP_4_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list)
* `AUTH_LDAP_4_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType`
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_4_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object)
* `AUTH_LDAP_4_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string)
* `AUTH_LDAP_4_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string)
* `AUTH_LDAP_4_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object)
* `AUTH_LDAP_4_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_4_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_5_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string)
* `AUTH_LDAP_5_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string)
* `AUTH_LDAP_5_BIND_PASSWORD`: Password used to bind LDAP user account. (string)
* `AUTH_LDAP_5_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean)
* `AUTH_LDAP_5_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object)
* `AUTH_LDAP_5_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list)
* `AUTH_LDAP_5_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string)
* `AUTH_LDAP_5_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object)
* `AUTH_LDAP_5_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list)
* `AUTH_LDAP_5_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType`
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_5_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object)
* `AUTH_LDAP_5_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string)
* `AUTH_LDAP_5_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string)
* `AUTH_LDAP_5_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object)
* `AUTH_LDAP_5_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object)
* `AUTH_LDAP_5_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `AUTH_LDAP_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string, default=`""`)
* `AUTH_LDAP_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string, default=`""`)
* `AUTH_LDAP_BIND_PASSWORD`: Password used to bind LDAP user account. (string, default=`""`)
* `AUTH_LDAP_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean, default=`False`)
* `AUTH_LDAP_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object, default=`{&#x27;OPT_REFERRALS&#x27;: 0, &#x27;OPT_NETWORK_TIMEOUT&#x27;: 30}`)
* `AUTH_LDAP_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list, default=`[]`)
* `AUTH_LDAP_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string, default=`""`)
* `AUTH_LDAP_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object, default=`{}`)
* `AUTH_LDAP_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list, default=`[]`)
* `AUTH_LDAP_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType` (default)
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object, default=`OrderedDict([(&#x27;member_attr&#x27;, &#x27;member&#x27;), (&#x27;name_attr&#x27;, &#x27;cn&#x27;)])`)
* `AUTH_LDAP_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string, default=`""`)
* `AUTH_LDAP_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string, default=`""`)
* `AUTH_LDAP_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object, default=`{}`)
* `AUTH_LDAP_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_1_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string, default=`""`)
* `AUTH_LDAP_1_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string, default=`""`)
* `AUTH_LDAP_1_BIND_PASSWORD`: Password used to bind LDAP user account. (string, default=`""`)
* `AUTH_LDAP_1_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean, default=`False`)
* `AUTH_LDAP_1_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object, default=`{&#x27;OPT_REFERRALS&#x27;: 0, &#x27;OPT_NETWORK_TIMEOUT&#x27;: 30}`)
* `AUTH_LDAP_1_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list, default=`[]`)
* `AUTH_LDAP_1_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string, default=`""`)
* `AUTH_LDAP_1_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object, default=`{}`)
* `AUTH_LDAP_1_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list, default=`[]`)
* `AUTH_LDAP_1_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType` (default)
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_1_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object, default=`OrderedDict([(&#x27;member_attr&#x27;, &#x27;member&#x27;), (&#x27;name_attr&#x27;, &#x27;cn&#x27;)])`)
* `AUTH_LDAP_1_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string, default=`""`)
* `AUTH_LDAP_1_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string, default=`""`)
* `AUTH_LDAP_1_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object, default=`{}`)
* `AUTH_LDAP_1_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_1_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_2_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string, default=`""`)
* `AUTH_LDAP_2_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string, default=`""`)
* `AUTH_LDAP_2_BIND_PASSWORD`: Password used to bind LDAP user account. (string, default=`""`)
* `AUTH_LDAP_2_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean, default=`False`)
* `AUTH_LDAP_2_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object, default=`{&#x27;OPT_REFERRALS&#x27;: 0, &#x27;OPT_NETWORK_TIMEOUT&#x27;: 30}`)
* `AUTH_LDAP_2_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list, default=`[]`)
* `AUTH_LDAP_2_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string, default=`""`)
* `AUTH_LDAP_2_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object, default=`{}`)
* `AUTH_LDAP_2_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list, default=`[]`)
* `AUTH_LDAP_2_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType` (default)
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_2_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object, default=`OrderedDict([(&#x27;member_attr&#x27;, &#x27;member&#x27;), (&#x27;name_attr&#x27;, &#x27;cn&#x27;)])`)
* `AUTH_LDAP_2_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string, default=`""`)
* `AUTH_LDAP_2_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string, default=`""`)
* `AUTH_LDAP_2_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object, default=`{}`)
* `AUTH_LDAP_2_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_2_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_3_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string, default=`""`)
* `AUTH_LDAP_3_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string, default=`""`)
* `AUTH_LDAP_3_BIND_PASSWORD`: Password used to bind LDAP user account. (string, default=`""`)
* `AUTH_LDAP_3_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean, default=`False`)
* `AUTH_LDAP_3_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object, default=`{&#x27;OPT_REFERRALS&#x27;: 0, &#x27;OPT_NETWORK_TIMEOUT&#x27;: 30}`)
* `AUTH_LDAP_3_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list, default=`[]`)
* `AUTH_LDAP_3_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string, default=`""`)
* `AUTH_LDAP_3_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object, default=`{}`)
* `AUTH_LDAP_3_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list, default=`[]`)
* `AUTH_LDAP_3_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType` (default)
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_3_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object, default=`OrderedDict([(&#x27;member_attr&#x27;, &#x27;member&#x27;), (&#x27;name_attr&#x27;, &#x27;cn&#x27;)])`)
* `AUTH_LDAP_3_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string, default=`""`)
* `AUTH_LDAP_3_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string, default=`""`)
* `AUTH_LDAP_3_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object, default=`{}`)
* `AUTH_LDAP_3_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_3_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_4_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string, default=`""`)
* `AUTH_LDAP_4_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string, default=`""`)
* `AUTH_LDAP_4_BIND_PASSWORD`: Password used to bind LDAP user account. (string, default=`""`)
* `AUTH_LDAP_4_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean, default=`False`)
* `AUTH_LDAP_4_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object, default=`{&#x27;OPT_REFERRALS&#x27;: 0, &#x27;OPT_NETWORK_TIMEOUT&#x27;: 30}`)
* `AUTH_LDAP_4_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list, default=`[]`)
* `AUTH_LDAP_4_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string, default=`""`)
* `AUTH_LDAP_4_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object, default=`{}`)
* `AUTH_LDAP_4_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list, default=`[]`)
* `AUTH_LDAP_4_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType` (default)
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_4_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object, default=`OrderedDict([(&#x27;member_attr&#x27;, &#x27;member&#x27;), (&#x27;name_attr&#x27;, &#x27;cn&#x27;)])`)
* `AUTH_LDAP_4_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string, default=`""`)
* `AUTH_LDAP_4_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string, default=`""`)
* `AUTH_LDAP_4_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object, default=`{}`)
* `AUTH_LDAP_4_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_4_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_5_SERVER_URI`: URI to connect to LDAP server, such as &quot;ldap://ldap.example.com:389&quot; (non-SSL) or &quot;ldaps://ldap.example.com:636&quot; (SSL). Multiple LDAP servers may be specified by separating with spaces or commas. LDAP authentication is disabled if this parameter is empty. (string, default=`""`)
* `AUTH_LDAP_5_BIND_DN`: DN (Distinguished Name) of user to bind for all search queries. This is the system user account we will use to login to query LDAP for other user information. Refer to the documentation for example syntax. (string, default=`""`)
* `AUTH_LDAP_5_BIND_PASSWORD`: Password used to bind LDAP user account. (string, default=`""`)
* `AUTH_LDAP_5_START_TLS`: Whether to enable TLS when the LDAP connection is not using SSL. (boolean, default=`False`)
* `AUTH_LDAP_5_CONNECTION_OPTIONS`: Additional options to set for the LDAP connection.  LDAP referrals are disabled by default (to prevent certain LDAP queries from hanging with AD). Option names should be strings (e.g. &quot;OPT_REFERRALS&quot;). Refer to https://www.python-ldap.org/doc/html/ldap.html#options for possible options and values that can be set. (nested object, default=`{&#x27;OPT_REFERRALS&#x27;: 0, &#x27;OPT_NETWORK_TIMEOUT&#x27;: 30}`)
* `AUTH_LDAP_5_USER_SEARCH`: LDAP search query to find users.  Any user that matches the given pattern will be able to login to the service.  The user should also be mapped into an organization (as defined in the AUTH_LDAP_ORGANIZATION_MAP setting).  If multiple search queries need to be supported use of &quot;LDAPUnion&quot; is possible. See the documentation for details. (list, default=`[]`)
* `AUTH_LDAP_5_USER_DN_TEMPLATE`: Alternative to user search, if user DNs are all of the same format. This approach is more efficient for user lookups than searching if it is usable in your organizational environment. If this setting has a value it will be used instead of AUTH_LDAP_USER_SEARCH. (string, default=`""`)
* `AUTH_LDAP_5_USER_ATTR_MAP`: Mapping of LDAP user schema to API user attributes. The default setting is valid for ActiveDirectory but users with other LDAP configurations may need to change the values. Refer to the documentation for additional details. (nested object, default=`{}`)
* `AUTH_LDAP_5_GROUP_SEARCH`: Users are mapped to organizations based on their membership in LDAP groups. This setting defines the LDAP search query to find groups. Unlike the user search, group search does not support LDAPSearchUnion. (list, default=`[]`)
* `AUTH_LDAP_5_GROUP_TYPE`: The group type may need to be changed based on the type of the LDAP server.  Values are listed at: https://django-auth-ldap.readthedocs.io/en/stable/groups.html#types-of-groups (choice)
    - `PosixGroupType`
    - `GroupOfNamesType`
    - `GroupOfUniqueNamesType`
    - `ActiveDirectoryGroupType`
    - `OrganizationalRoleGroupType`
    - `MemberDNGroupType` (default)
    - `NestedGroupOfNamesType`
    - `NestedGroupOfUniqueNamesType`
    - `NestedActiveDirectoryGroupType`
    - `NestedOrganizationalRoleGroupType`
    - `NestedMemberDNGroupType`
    - `PosixUIDGroupType`
* `AUTH_LDAP_5_GROUP_TYPE_PARAMS`: Key value parameters to send the chosen group type init method. (nested object, default=`OrderedDict([(&#x27;member_attr&#x27;, &#x27;member&#x27;), (&#x27;name_attr&#x27;, &#x27;cn&#x27;)])`)
* `AUTH_LDAP_5_REQUIRE_GROUP`: Group DN required to login. If specified, user must be a member of this group to login via LDAP. If not set, everyone in LDAP that matches the user search will be able to login to the service. Only one require group is supported. (string, default=`""`)
* `AUTH_LDAP_5_DENY_GROUP`: Group DN denied from login. If specified, user will not be allowed to login if a member of this group.  Only one deny group is supported. (string, default=`""`)
* `AUTH_LDAP_5_USER_FLAGS_BY_GROUP`: Retrieve users from a given group. At this time, superuser and system auditors are the only groups supported. Refer to the documentation for more detail. (nested object, default=`{}`)
* `AUTH_LDAP_5_ORGANIZATION_MAP`: Mapping between organization admins/users and LDAP groups. This controls which users are placed into which organizations relative to their LDAP group memberships. Configuration details are available in the documentation. (nested object, default=`{}`)
* `AUTH_LDAP_5_TEAM_MAP`: Mapping between team members (users) and LDAP groups. Configuration details are available in the documentation. (nested object, default=`{}`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.