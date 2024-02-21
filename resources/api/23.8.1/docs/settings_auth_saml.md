# Retrieve a Setting:

Make GET request to this resource to retrieve a single setting
record containing the following fields:

* `SAML_AUTO_CREATE_OBJECTS`: When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login. (boolean)
* `SOCIAL_AUTH_SAML_CALLBACK_URL`: Register the service as a service provider (SP) with each identity provider (IdP) you have configured. Provide your SP Entity ID and this ACS URL for your application. (string)
* `SOCIAL_AUTH_SAML_METADATA_URL`: If your identity provider (IdP) allows uploading an XML metadata file, you can download one from this URL. (string)
* `SOCIAL_AUTH_SAML_SP_ENTITY_ID`: The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service. (string)
* `SOCIAL_AUTH_SAML_SP_PUBLIC_CERT`: Create a keypair to use as a service provider (SP) and include the certificate content here. (string)
* `SOCIAL_AUTH_SAML_SP_PRIVATE_KEY`: Create a keypair to use as a service provider (SP) and include the private key content here. (string)
* `SOCIAL_AUTH_SAML_ORG_INFO`: Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax. (nested object)
* `SOCIAL_AUTH_SAML_TECHNICAL_CONTACT`: Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax. (nested object)
* `SOCIAL_AUTH_SAML_SUPPORT_CONTACT`: Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax. (nested object)
* `SOCIAL_AUTH_SAML_ENABLED_IDPS`: Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax. (nested object)
* `SOCIAL_AUTH_SAML_SECURITY_CONFIG`: A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings (nested object)
* `SOCIAL_AUTH_SAML_SP_EXTRA`: A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting. (nested object)
* `SOCIAL_AUTH_SAML_EXTRA_DATA`: A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value. (list)
* `SOCIAL_AUTH_SAML_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object)
* `SOCIAL_AUTH_SAML_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object)
* `SOCIAL_AUTH_SAML_ORGANIZATION_ATTR`: Used to translate user organization membership. (nested object)
* `SOCIAL_AUTH_SAML_TEAM_ATTR`: Used to translate user team membership. (nested object)
* `SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR`: Used to map super users and system auditors from SAML. (nested object)





# Update a Setting:

Make a PUT or PATCH request to this resource to update this
setting.  The following fields may be modified:


* `SAML_AUTO_CREATE_OBJECTS`: When enabled (the default), mapped Organizations and Teams will be created automatically on successful SAML login. (boolean, default=`True`)


* `SOCIAL_AUTH_SAML_SP_ENTITY_ID`: The application-defined unique identifier used as the audience of the SAML service provider (SP) configuration. This is usually the URL for the service. (string, default=`""`)
* `SOCIAL_AUTH_SAML_SP_PUBLIC_CERT`: Create a keypair to use as a service provider (SP) and include the certificate content here. (string, required)
* `SOCIAL_AUTH_SAML_SP_PRIVATE_KEY`: Create a keypair to use as a service provider (SP) and include the private key content here. (string, required)
* `SOCIAL_AUTH_SAML_ORG_INFO`: Provide the URL, display name, and the name of your app. Refer to the documentation for example syntax. (nested object, required)
* `SOCIAL_AUTH_SAML_TECHNICAL_CONTACT`: Provide the name and email address of the technical contact for your service provider. Refer to the documentation for example syntax. (nested object, required)
* `SOCIAL_AUTH_SAML_SUPPORT_CONTACT`: Provide the name and email address of the support contact for your service provider. Refer to the documentation for example syntax. (nested object, required)
* `SOCIAL_AUTH_SAML_ENABLED_IDPS`: Configure the Entity ID, SSO URL and certificate for each identity provider (IdP) in use. Multiple SAML IdPs are supported. Some IdPs may provide user data using attribute names that differ from the default OIDs. Attribute names may be overridden for each IdP. Refer to the Ansible documentation for additional details and syntax. (nested object, default=`{}`)
* `SOCIAL_AUTH_SAML_SECURITY_CONFIG`: A dict of key value pairs that are passed to the underlying python-saml security setting https://github.com/onelogin/python-saml#settings (nested object, default=`{&#x27;requestedAuthnContext&#x27;: False}`)
* `SOCIAL_AUTH_SAML_SP_EXTRA`: A dict of key value pairs to be passed to the underlying python-saml Service Provider configuration setting. (nested object, default=`None`)
* `SOCIAL_AUTH_SAML_EXTRA_DATA`: A list of tuples that maps IDP attributes to extra_attributes. Each attribute will be a list of values, even if only 1 value. (list, default=`None`)
* `SOCIAL_AUTH_SAML_ORGANIZATION_MAP`: Mapping to organization admins/users from social auth accounts. This setting
controls which users are placed into which organizations based on their
username and email address. Configuration details are available in the
documentation. (nested object, default=`None`)
* `SOCIAL_AUTH_SAML_TEAM_MAP`: Mapping of team members (users) from social auth accounts. Configuration
details are available in the documentation. (nested object, default=`None`)
* `SOCIAL_AUTH_SAML_ORGANIZATION_ATTR`: Used to translate user organization membership. (nested object, default=`{}`)
* `SOCIAL_AUTH_SAML_TEAM_ATTR`: Used to translate user team membership. (nested object, default=`{}`)
* `SOCIAL_AUTH_SAML_USER_FLAGS_BY_ATTR`: Used to map super users and system auditors from SAML. (nested object, default=`{}`)






For a PUT request, include **all** fields in the request.



For a PATCH request, include only the fields that are being modified.



# Delete a Setting:

Make a DELETE request to this resource to delete this setting.