# es-handler
Elasticsearch handler used by Terraform, Crossplane and k8s operator


## How to use it

### Get client / handler

You can look [sample](samples/client.go)

### Manage snapshot repository

You can look [sample](samples/snapshot_repository.go)

### Manage component template

You can look [sample](samples/component_template.go)

### Manage index template

You can look [sample](samples/index_template.go)

### Manage ingest pipline

You can look [sample](samples/ingest_pipeline.go)

### Manage role

You can look [sample](samples/role.go)

### Manage role mapping

You can look [sample](samples/role_mapping.go)

### Manage user

You can look [sample](samples/user.go)

### Manage action group

You can look [sample](samples/action_group.go)

### Manage tenant

You can look [sample](samples/tenant.go)

### Manage security config

You can look [sample](samples/security_config.go)

### Manage audit config

You can look [sample](samples/audit_config.go)

### Manage index state management

You can look [sample](samples/index_state_management.go)

### Manage snapshot management

You can look [sample](samples/snapshot_management.go)

### Manage monitor policy

You can look [sample](samples/monitor_policy.go)

### Get cluster health

You can look [sample](samples/cluster_health.go)

## Contribute

PR are always welcome here!

Please use the branch `release-branch.v2*` to gettig start.

Implement the code and test that prove is work as expected.

Run test:
```bash
go test -v ./...
```