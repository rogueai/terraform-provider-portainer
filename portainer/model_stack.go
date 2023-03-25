package portainer

import "github.com/hashicorp/terraform-plugin-framework/types"

type Stacks struct {
	ID     types.String `tfsdk:"id"` // required for acceptance testing
	Stacks []Stack      `tfsdk:"stacks"`
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
	SubResourceIds     []types.String      `tfsdk:"sub_resource_ids"`
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
	ID              types.String                  `tfsdk:"id"`
	AdditionalFiles []types.String                `tfsdk:"additional_files"`
	AutoUpdate      *StackAutoUpdate              `tfsdk:"auto_update"`
	EndpointId      types.Int64                   `tfsdk:"endpoint_id"`
	EntryPoint      types.String                  `tfsdk:"entry_point"`
	Env             []StackEnv                    `tfsdk:"env"`
	Name            types.String                  `tfsdk:"name"`
	Option          *StackOption                  `tfsdk:"option"`
	ResourceControl *StackResourceControl         `tfsdk:"resource_control"`
	Status          types.Int64                   `tfsdk:"status"`
	SwarmId         types.String                  `tfsdk:"swarm_id"`
	Type            types.Int64                   `tfsdk:"type"`
	CreatedBy       types.String                  `tfsdk:"created_by"`
	CreationDate    types.Int64                   `tfsdk:"creation_date"`
	FromAppTemplate types.Bool                    `tfsdk:"from_app_template"`
	GitConfig       *StackGitAuthenticationConfig `tfsdk:"git_config"`
	IsComposeFormat types.Bool                    `tfsdk:"is_compose_format"`
	Namespace       types.String                  `tfsdk:"namespace"`
	ProjectPath     types.String                  `tfsdk:"project_path"`
	UpdateDate      types.Int64                   `tfsdk:"update_date"`
	UpdatedBy       types.String                  `tfsdk:"updated_by"`
	FileContent     types.String                  `tfsdk:"file_content"`
}

type StackAutoUpdate struct {
	Interval types.String `tfsdk:"interval"`
	JobID    types.String `tfsdk:"job_id"`
	Webhook  types.String `tfsdk:"webhook"`
}

type StackEnv struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

//type StackBodySwarmString struct {
//	Env              []StackEnv   `tfsdk:"env"`
//	FromAppTemplate  types.Bool   `tfsdk:"from_app_template"`
//	Name             types.String `tfsdk:"name"`
//	FileContent types.String `tfsdk:"stack_file_content"`
//	SwarmId          types.String `tfsdk:"swarm_id"`
//}
//type StackBodySwarmRepository struct {
//	AdditionalFiles          []types.String  `tfsdk:"additional_files"`
//	AutoUpdate               StackAutoUpdate `tfsdk:"auto_update"`
//	ComposeFile              types.String    `tfsdk:"compose_file"`
//	Env                      []StackEnv      `tfsdk:"env"`
//	FromAppTemplate          types.Bool      `tfsdk:"from_app_template"`
//	Name                     types.String    `tfsdk:"name"`
//	RepositoryAuthentication types.Bool      `tfsdk:"repository_authentication"`
//	RepositoryPassword       types.String    `tfsdk:"repository_password"`
//	RepositoryReferenceName  types.String    `tfsdk:"repository_reference_name"`
//	RepositoryUrl            types.String    `tfsdk:"repository_url"`
//	RepositoryUsername       types.String    `tfsdk:"repository_username"`
//	SwarmId                  types.String    `tfsdk:"swarm_id"`
//}
//type StackBodyComposeString struct {
//	Env              []StackEnv   `tfsdk:"env"`
//	FromAppTemplate  types.Bool   `tfsdk:"from_app_template"`
//	Name             types.String `tfsdk:"name"`
//	FileContent types.String `tfsdk:"stack_file_content"`
//}
//type StackBodyComposeRepository struct {
//	AdditionalFiles          []types.String  `tfsdk:"additional_files"`
//	AutoUpdate               StackAutoUpdate `tfsdk:"auto_update"`
//	ComposeFile              types.String    `tfsdk:"compose_file"`
//	Env                      []StackEnv      `tfsdk:"env"`
//	FromAppTemplate          types.Bool      `tfsdk:"from_app_template"`
//	Name                     types.String    `tfsdk:"name"`
//	RepositoryAuthentication types.Bool      `tfsdk:"repository_authentication"`
//	RepositoryPassword       types.String    `tfsdk:"repository_password"`
//	RepositoryReferenceName  types.String    `tfsdk:"repository_reference_name"`
//	RepositoryUrl            types.String    `tfsdk:"repository_url"`
//	RepositoryUsername       types.String    `tfsdk:"repository_username"`
//}
//type StackBodyKubernetesString struct {
//	ComposeFormat    types.Bool   `tfsdk:"compose_format"`
//	Namespace        types.String `tfsdk:"namespace"`
//	FileContent types.String `tfsdk:"stack_file_content"`
//	StackName        types.String `tfsdk:"stack_name"`
//}
//type StackBodyKubernetesRepository struct {
//	AdditionalFiles          []types.String  `tfsdk:"additional_files"`
//	AutoUpdate               StackAutoUpdate `tfsdk:"auto_update"`
//	ComposeFormat            types.Bool      `tfsdk:"compose_format"`
//	ManifestFile             types.String    `tfsdk:"manifest_file"`
//	Namespace                types.String    `tfsdk:"namespace"`
//	RepositoryAuthentication types.Bool      `tfsdk:"repository_authentication"`
//	RepositoryPassword       types.String    `tfsdk:"repository_password"`
//	RepositoryReferenceName  types.String    `tfsdk:"repository_reference_name"`
//	RepositoryUrl            types.String    `tfsdk:"repository_url"`
//	RepositoryUsername       types.String    `tfsdk:"repository_username"`
//	StackName                types.String    `tfsdk:"stack_name"`
//}
//type StackBodyKubernetesUrl struct {
//	ComposeFormat types.Bool   `tfsdk:"compose_format"`
//	ManifestUrl   types.String `tfsdk:"manifest_url"`
//	Namespace     types.String `tfsdk:"namespace"`
//	StackName     types.String `tfsdk:"stack_name"`
//}
//
//type StackCreate struct {
//	Type                     types.Int64                    `tfsdk:"type"`
//	Method                   types.String                   `tfsdk:"method"`
//	EndpointId               types.Int64                    `tfsdk:"endpoint_id"`
//	BodySwarmString          *StackBodySwarmString          `tfsdk:"body_swarm_string"`
//	BodySwarmRepository      *StackBodySwarmRepository      `tfsdk:"body_swarm_repository"`
//	BodyComposeString        *StackBodyComposeString        `tfsdk:"body_compose_string"`
//	BodyComposeRepository    *StackBodyComposeRepository    `tfsdk:"body_compose_repository"`
//	BodyKubernetesString     *StackBodyKubernetesString     `tfsdk:"body_kubernetes_string"`
//	BodyKubernetesRepository *StackBodyKubernetesRepository `tfsdk:"body_kubernetes_repository"`
//	BodyKubernetesUrl        *StackBodyKubernetesUrl        `tfsdk:"body_kubernetes_url"`
//	Name                     types.String                   `tfsdk:"name"`
//	SwarmId                  types.String                   `tfsdk:"swarm_id"`
//	Env                      []StackEnv                     `tfsdk:"env"`
//	File                     types.String                   `tfsdk:"file"`
//}
