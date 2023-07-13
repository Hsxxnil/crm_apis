package auth

const _CASBIN_RULES = `
[
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/crm/v1.0/*",
    "v2": "GET"
  },
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/crm/v1.0/*",
    "v2": "POST"
  },
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/crm/v1.0/*",
    "v2": "PATCH"
  },
  {
    "ptype": "p",
    "v0": "admin",
    "v1": "/crm/v1.0/*",
    "v2": "DELETE"
  }
]
`
