package model

import "github.com/hashicorp/terraform-plugin-framework/types"

type Stacks struct {
	Stacks []Stack `tfsdk:"stacks"`
}

type StackTeamAccesses struct {
	AccessLevel types.Int64 `tfsdk:"access_level"`
	TeamId      types.Int64 `tfsdk:"team_id"`
}

type StackUserAccesses struct {
	AccessLevel types.Int64 `tfsdk:"access_level"`
	UserId      types.Int64 `tfsdk:"user_id"`
}

type StackResourceControl struct {
	AccessLevel        types.Int64         `tfsdk:"access_level"`
	AdministratorsOnly types.Bool          `tfsdk:"administrators_only"`
	Id                 types.Int64         `tfsdk:"id"`
	OwnerId            types.Int64         `tfsdk:"owner_id"`
	Public             types.Bool          `tfsdk:"public"`
	ResourceId         types.String        `tfsdk:"resource_id"`
	SubResourceIds     []string            `tfsdk:"sub_resource_ids"`
	System             types.Bool          `tfsdk:"system"`
	TeamAccesses       []StackTeamAccesses `tfsdk:"team_accesses"`
	Type               types.Int64         `tfsdk:"type"`
	UserAccesses       []StackUserAccesses `tfsdk:"user_accesses"`
}

type StackGitAuthentication struct {
	GitCredentialID types.Int64  `tfsdk:"git_credential_id"`
	Password        types.String `tfsdk:"password"`
	Username        types.String `tfsdk:"username"`
}

type StackGitAuthenticationConfig struct {
	Authentication StackGitAuthentication `tfsdk:"authentication"`
	ConfigFilePath types.String           `tfsdk:"config_file_path"`
	ConfigHash     types.String           `tfsdk:"config_hash"`
	ReferenceName  types.String           `tfsdk:"reference_name"`
	Url            types.String           `tfsdk:"url"`
}

type StackOption struct {
	Prune types.Bool `tfsdk:"prune"`
}

// Stack maps stacks schema data.
type Stack struct {
	AdditionalFiles []types.String               `tfsdk:"additional_files"`
	AutoUpdate      StackAutoUpdate              `tfsdk:"auto_update"`
	EndpointId      types.Int64                  `tfsdk:"endpoint_id"`
	EntryPoint      types.String                 `tfsdk:"entry_point"`
	Env             []StackEnvModel              `tfsdk:"env"`
	Id              types.Int64                  `tfsdk:"id"`
	Name            types.String                 `tfsdk:"name"`
	Option          StackOption                  `tfsdk:"option"`
	ResourceControl StackResourceControl         `tfsdk:"resource_control"`
	Status          types.Int64                  `tfsdk:"status"`
	SwarmId         types.String                 `tfsdk:"swarm_id"`
	Type            types.Int64                  `tfsdk:"type"`
	CreatedBy       types.String                 `tfsdk:"created_by"`
	CreationDate    types.Int64                  `tfsdk:"creation_date"`
	FromAppTemplate types.Bool                   `tfsdk:"from_app_template"`
	GitConfig       StackGitAuthenticationConfig `tfsdk:"git_config"`
	IsComposeFormat types.Bool                   `tfsdk:"is_compose_format"`
	Namespace       types.String                 `tfsdk:"namespace"`
	ProjectPath     types.String                 `tfsdk:"project_path"`
	UpdateDate      types.Int64                  `tfsdk:"update_date"`
	UpdatedBy       types.String                 `tfsdk:"updated_by"`
}

type StackAutoUpdate struct {
	Interval types.String `tfsdk:"interval"`
	JobID    types.String `tfsdk:"job_id"`
	Webhook  types.String `tfsdk:"webhook"`
}

type StackEnvModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
