package rbac.authz

# Assigning user roles

user_roles := {
    "crab": ["admin"],
    "bob": ["cook"],
    "plankton":["villain"]
}

# Role permission assignments
role_permissions := {
    "cook":    [{"action": "read",  "object": "secret_formula"}],
    "admin":   [{"action": "read",  "object": "secret_formula"},
                {"action": "write", "object": "secret_formula"}],
    "villain": [{"action": "want",  "object": "secret_formula"}]
}

# Logic that implements RBAC.
default allow = false
allow {
    # Lookup the list of roles for the user
    roles := user_roles[input.user]
    # For each role in that list
    r := roles[_]
    # Lookup the permissions list for role r
    permissions := role_permissions[r]
    # For each permission
    p := permissions[_]
    # Check if the permission granted to r matches the user's request
    p == {"action": input.action, "object": input.object}
}