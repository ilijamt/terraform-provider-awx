Use this endpoint to give a user permission to a resource or an organization.
The needed data is the user, the role definition, and the object id.
The object must be of the type specified in the role definition.
The type given in the role definition and the provided object_id are used
to look up the resource.

After creation, the assignment cannot be edited, but can be deleted to
remove those permissions.