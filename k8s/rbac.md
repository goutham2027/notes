https://kubernetes.io/docs/reference/access-authn-authz/rbac/
https://learnk8s.io/rbac-kubernetes
https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control

RBAC - Role Based Access Control

RBAC is a model designed to grant access to resources based on the roles
of individual users within an org

k8s uses 3 concepts:
1. identities
2. roles
3. bindings

RBAC API declares four kinds of k8s object:
- Role
  - a namespaced grouping of resources and allowed operations that you
    can assign to a user or a group of users using a RoleBinding
- ClusterRole
  - a cluster-level grouping of resources and allowed operations that
    you can assign to a user or group using RoleBinding or
ClusterRoleBinding
- RoleBinding
  - assign a Role or a ClusterRole to a user or a group within a
    specific namespace
- ClusterRoleBinding
  - assign a ClusterRole to a user or a group for all namespaces in the
    cluster

Define RBAC rules in ClusterRole and Role objects, and then assign those
rules with ClusterRoleBinding and RoleBinding objects.


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


### How should you store the User in the cluster?
- k8s does not have objects which represent regular user accounts
- users cannot be added to a cluster through an API call
- Instead, any actor that presents a valid cert signed by the cluster's
  Certificate Authority (CA) is considered authenticated
