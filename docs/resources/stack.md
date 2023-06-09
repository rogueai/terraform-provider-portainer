---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "portainer_stack Resource - portainer"
subcategory: ""
description: |-
  
---

# portainer_stack (Resource)



## Example Usage

```terraform
# Creates an example stack.
resource "portainer_stack" "example" {
  type              = 1
  endpoint_id       = 2
  name              = "example"
  file_content      = file("docker-compose.yml")
  swarm_id          = "my-swarm-id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `endpoint_id` (Number)
- `file_content` (String)
- `name` (String)
- `swarm_id` (String)
- `type` (Number)

### Optional

- `additional_files` (List of String)
- `auto_update` (Attributes) (see [below for nested schema](#nestedatt--auto_update))
- `env` (Attributes List) (see [below for nested schema](#nestedatt--env))
- `git_config` (Attributes) (see [below for nested schema](#nestedatt--git_config))
- `namespace` (String)
- `option` (Object) (see [below for nested schema](#nestedatt--option))
- `resource_control` (Attributes) (see [below for nested schema](#nestedatt--resource_control))

### Read-Only

- `created_by` (String)
- `creation_date` (Number)
- `entry_point` (String)
- `from_app_template` (Boolean)
- `id` (String) The ID of this resource.
- `is_compose_format` (Boolean)
- `project_path` (String)
- `status` (Number)
- `update_date` (Number)
- `updated_by` (String)

<a id="nestedatt--auto_update"></a>
### Nested Schema for `auto_update`

Read-Only:

- `interval` (String)
- `job_id` (String)
- `webhook` (String)


<a id="nestedatt--env"></a>
### Nested Schema for `env`

Required:

- `name` (String)
- `value` (String)


<a id="nestedatt--git_config"></a>
### Nested Schema for `git_config`

Required:

- `reference_name` (String)
- `url` (String)

Read-Only:

- `authentication` (Attributes) (see [below for nested schema](#nestedatt--git_config--authentication))
- `config_file_path` (String)
- `config_hash` (String)

<a id="nestedatt--git_config--authentication"></a>
### Nested Schema for `git_config.authentication`

Optional:

- `git_credential_id` (Number)
- `password` (String)
- `username` (String)



<a id="nestedatt--option"></a>
### Nested Schema for `option`

Optional:

- `prune` (Boolean)


<a id="nestedatt--resource_control"></a>
### Nested Schema for `resource_control`

Optional:

- `access_level` (Number)
- `administrators_only` (Boolean)
- `id` (Number)
- `owner_id` (Number)
- `public` (Boolean)
- `resource_id` (String)
- `sub_resource_ids` (List of String)
- `system` (Boolean)
- `team_accesses` (Attributes List) (see [below for nested schema](#nestedatt--resource_control--team_accesses))
- `type` (Number)
- `user_accesses` (Attributes List) (see [below for nested schema](#nestedatt--resource_control--user_accesses))

<a id="nestedatt--resource_control--team_accesses"></a>
### Nested Schema for `resource_control.team_accesses`

Required:

- `access_level` (Number)
- `team_id` (Number)


<a id="nestedatt--resource_control--user_accesses"></a>
### Nested Schema for `resource_control.user_accesses`

Required:

- `access_level` (Number)
- `user_id` (Number)

## Import

Import is supported using the following syntax:

```shell
# Stack can be imported by specifying the numeric identifier.
terraform import portainer_stack.example 123
```
