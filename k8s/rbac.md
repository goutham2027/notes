https://kubernetes.io/docs/reference/access-authn-authz/rbac/

RBAC API declares four kinds of k8s object:
- Role
- ClusterRole
- RoleBinding
- ClusterRoleBinding


To define a role within a namespace, use a Role

To define a role cluster-wide, use a ClusterRole

### ClusterRole
* non-namespaces resource

Uses of ClusterRoles
- define permissions on namespaced resources
- be granted access across within individual namespaces or across all
  namespaces
- define permissions on cluster-scoped resources

### RoleBinding
* grants the permissions defined in a role to a user or set of users
* grants permissions within a specific namespace


