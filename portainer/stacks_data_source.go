package portainer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	portainer "terraform-provider-portainer/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &stacksDataSource{}
	_ datasource.DataSourceWithConfigure = &stacksDataSource{}
)

func NewStacksDataSource() datasource.DataSource {
	return &stacksDataSource{}
}

type stacksDataSource struct {
	client *portainer.APIClient
}

func (d *stacksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stacks"
}

// Schema defines the schema for the data source.
func (d *stacksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"stacks": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"additional_files": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"auto_update": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{Computed: true},
								"job_id":   schema.StringAttribute{Computed: true},
								"webhook":  schema.StringAttribute{Computed: true},
							},
						},
						"endpoint_id": schema.Int64Attribute{
							Computed: true,
						},
						"entry_point": schema.StringAttribute{
							Computed: true,
						},
						"env": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Computed: true,
									},
									"value": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"option": schema.ObjectAttribute{
							Computed: true,
							AttributeTypes: map[string]attr.Type{
								"prune": types.BoolType,
							},
						},
						"resource_control": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"access_level":        schema.Int64Attribute{Computed: true},
								"administrators_only": schema.BoolAttribute{Computed: true},
								"id":                  schema.Int64Attribute{Computed: true},
								"owner_id":            schema.Int64Attribute{Computed: true},
								"public":              schema.BoolAttribute{Computed: true},
								"resource_id":         schema.StringAttribute{Computed: true},
								"sub_resource_ids": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"system": schema.BoolAttribute{Computed: true},
								"team_accesses": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"access_level": schema.Int64Attribute{Computed: true},
											"team_id":      schema.Int64Attribute{Computed: true},
										},
									},
								},
								"type": schema.Int64Attribute{Computed: true},
								"user_accesses": schema.ListNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"access_level": schema.Int64Attribute{Computed: true},
											"user_id":      schema.Int64Attribute{Computed: true},
										},
									},
								},
							},
						},
						"status": schema.Int64Attribute{
							Computed: true,
						},
						"swarm_id": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.Int64Attribute{
							Computed: true,
						},
						"created_by": schema.StringAttribute{
							Computed: true,
						},
						"creation_date": schema.Int64Attribute{
							Computed: true,
						},
						"from_app_template": schema.BoolAttribute{
							Computed: true,
						},
						"git_config": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"authentication": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"git_credential_id": schema.Int64Attribute{Computed: true},
										"password": schema.StringAttribute{
											Computed: true,
										},
										"username": schema.StringAttribute{Computed: true},
									},
								},
								"config_file_path": schema.StringAttribute{Computed: true},
								"config_hash":      schema.StringAttribute{Computed: true},
								"reference_name":   schema.StringAttribute{Computed: true},
								"url":              schema.StringAttribute{Computed: true},
							},
						},
						"is_compose_format": schema.BoolAttribute{
							Computed: true,
						},
						"namespace": schema.StringAttribute{
							Computed: true,
						},
						"project_path": schema.StringAttribute{
							Computed: true,
						},
						"update_date": schema.Int64Attribute{
							Computed: true,
						},
						"updated_by": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type StacksDataSourceModel struct {
	Stacks []stacksModel `tfsdk:"stacks"`
}

// stacksModel maps stacks schema data.
type stacksTeamAccessesModel struct {
	AccessLevel types.Int64 `tfsdk:"access_level"`
	TeamId      types.Int64 `tfsdk:"team_id"`
}

type stacksUserAccessesModel struct {
	AccessLevel types.Int64 `tfsdk:"access_level"`
	UserId      types.Int64 `tfsdk:"user_id"`
}

type stacksResourceControlModel struct {
	AccessLevel        types.Int64               `tfsdk:"access_level"`
	AdministratorsOnly types.Bool                `tfsdk:"administrators_only"`
	Id                 types.Int64               `tfsdk:"id"`
	OwnerId            types.Int64               `tfsdk:"owner_id"`
	Public             types.Bool                `tfsdk:"public"`
	ResourceId         types.String              `tfsdk:"resource_id"`
	SubResourceIds     []string                  `tfsdk:"sub_resource_ids"`
	System             types.Bool                `tfsdk:"system"`
	TeamAccesses       []stacksTeamAccessesModel `tfsdk:"team_accesses"`
	Type               types.Int64               `tfsdk:"type"`
	UserAccesses       []stacksUserAccessesModel `tfsdk:"user_accesses"`
}

type stacksGitAuthenticationModel struct {
	GitCredentialID types.Int64  `tfsdk:"git_credential_id"`
	Password        types.String `tfsdk:"password"`
	Username        types.String `tfsdk:"username"`
}

type stacksGitConfigModel struct {
	Authentication stacksGitAuthenticationModel `tfsdk:"authentication"`
	ConfigFilePath types.String                 `tfsdk:"config_file_path"`
	ConfigHash     types.String                 `tfsdk:"config_hash"`
	ReferenceName  types.String                 `tfsdk:"reference_name"`
	Url            types.String                 `tfsdk:"url"`
}

type stacksOptionModel struct {
	Prune types.Bool `tfsdk:"prune"`
}

type stacksModel struct {
	AdditionalFiles []types.String             `tfsdk:"additional_files"`
	AutoUpdate      stacksAutoUpdateModel      `tfsdk:"auto_update"`
	EndpointId      types.Int64                `tfsdk:"endpoint_id"`
	EntryPoint      types.String               `tfsdk:"entry_point"`
	Env             []stacksEnvModel           `tfsdk:"env"`
	Id              types.Int64                `tfsdk:"id"`
	Name            types.String               `tfsdk:"name"`
	Option          stacksOptionModel          `tfsdk:"option"`
	ResourceControl stacksResourceControlModel `tfsdk:"resource_control"`
	Status          types.Int64                `tfsdk:"status"`
	SwarmId         types.String               `tfsdk:"swarm_id"`
	Type            types.Int64                `tfsdk:"type"`
	CreatedBy       types.String               `tfsdk:"created_by"`
	CreationDate    types.Int64                `tfsdk:"creation_date"`
	FromAppTemplate types.Bool                 `tfsdk:"from_app_template"`
	GitConfig       stacksGitConfigModel       `tfsdk:"git_config"`
	IsComposeFormat types.Bool                 `tfsdk:"is_compose_format"`
	Namespace       types.String               `tfsdk:"namespace"`
	ProjectPath     types.String               `tfsdk:"project_path"`
	UpdateDate      types.Int64                `tfsdk:"update_date"`
	UpdatedBy       types.String               `tfsdk:"updated_by"`
}

// coffeesIngredientsModel maps coffee ingredients data
type stacksAutoUpdateModel struct {
	Interval types.String `tfsdk:"interval"`
	JobID    types.String `tfsdk:"job_id"`
	Webhook  types.String `tfsdk:"webhook"`
}

type stacksEnvModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// Read refreshes the Terraform state with the latest data.
func (d *stacksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state StacksDataSourceModel

	stacks, _, err := d.client.StacksApi.StackList(context.Background(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Portainer Stacks",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, stack := range stacks {
		stackState := stacksModel{
			AdditionalFiles: nil,
			AutoUpdate:      stacksAutoUpdateModel{},
			EndpointId:      types.Int64Value(int64(stack.EndpointId)),
			EntryPoint:      types.StringValue(stack.EntryPoint),
			Env:             nil,
			Id:              types.Int64Value(int64(stack.Id)),
			Name:            types.StringValue(stack.Name),
			Option:          stacksOptionModel{},
			ResourceControl: stacksResourceControlModel{},
			Status:          types.Int64Value(int64(stack.Status)),
			SwarmId:         types.StringValue(stack.SwarmId),
			Type:            types.Int64Value(int64(stack.Type_)),
			CreatedBy:       types.StringValue(stack.CreatedBy),
			CreationDate:    types.Int64Value(int64(stack.CreationDate)),
			FromAppTemplate: types.BoolValue(stack.FromAppTemplate),
			GitConfig:       stacksGitConfigModel{},
			IsComposeFormat: types.BoolValue(stack.IsComposeFormat),
			Namespace:       types.StringValue(stack.Namespace),
			ProjectPath:     types.StringValue(stack.ProjectPath),
			UpdateDate:      types.Int64Value(int64(stack.UpdateDate)),
			UpdatedBy:       types.StringValue(stack.UpdatedBy),
		}

		for _, additionalFile := range stack.AdditionalFiles {
			stackState.AdditionalFiles = append(stackState.AdditionalFiles, types.StringValue(additionalFile))
		}

		if stack.AutoUpdate != nil {
			stackState.AutoUpdate.Interval = types.StringValue(stack.AutoUpdate.Interval)
			stackState.AutoUpdate.JobID = types.StringValue(stack.AutoUpdate.JobID)
			stackState.AutoUpdate.Webhook = types.StringValue(stack.AutoUpdate.Webhook)
		}

		if stack.Option != nil {
			stackState.Option.Prune = types.BoolValue(stack.Option.Prune)
		}

		for _, env := range stack.Env {
			stackState.Env = append(stackState.Env, stacksEnvModel{
				Name:  types.StringValue(env.Name),
				Value: types.StringValue(env.Value),
			})
		}

		if stack.ResourceControl != nil {
			stackState.ResourceControl.AccessLevel = types.Int64Value(int64(stack.ResourceControl.AccessLevel))
			stackState.ResourceControl.AdministratorsOnly = types.BoolValue(stack.ResourceControl.AdministratorsOnly)
			stackState.ResourceControl.Id = types.Int64Value(int64(stack.ResourceControl.Id))
			stackState.ResourceControl.OwnerId = types.Int64Value(int64(stack.ResourceControl.OwnerId))
			stackState.ResourceControl.Public = types.BoolValue(stack.ResourceControl.Public)
			stackState.ResourceControl.ResourceId = types.StringValue(stack.ResourceControl.ResourceId)
			for _, subResourceId := range stack.ResourceControl.SubResourceIds {
				stackState.ResourceControl.SubResourceIds = append(stackState.ResourceControl.SubResourceIds, subResourceId)
			}
			stackState.ResourceControl.System = types.BoolValue(stack.ResourceControl.System)
			for _, teamAccess := range stack.ResourceControl.TeamAccesses {
				stackState.ResourceControl.TeamAccesses = append(stackState.ResourceControl.TeamAccesses, stacksTeamAccessesModel{
					AccessLevel: types.Int64Value(int64(teamAccess.AccessLevel)),
					TeamId:      types.Int64Value(int64(teamAccess.TeamId)),
				})
			}
			for _, userAccess := range stack.ResourceControl.UserAccesses {
				stackState.ResourceControl.UserAccesses = append(stackState.ResourceControl.UserAccesses, stacksUserAccessesModel{
					AccessLevel: types.Int64Value(int64(userAccess.AccessLevel)),
					UserId:      types.Int64Value(int64(userAccess.UserId)),
				})
			}
			stackState.ResourceControl.Type = types.Int64Value(int64(stack.ResourceControl.Type_))

		}

		if stack.GitConfig != nil {
			stackState.GitConfig.ConfigFilePath = types.StringValue(stack.GitConfig.ConfigFilePath)
			stackState.GitConfig.ConfigHash = types.StringValue(stack.GitConfig.ConfigHash)
			stackState.GitConfig.Url = types.StringValue(stack.GitConfig.Url)
			stackState.GitConfig.ReferenceName = types.StringValue(stack.GitConfig.ReferenceName)
			if stack.GitConfig.Authentication != nil {
				stackState.GitConfig.Authentication = stacksGitAuthenticationModel{
					GitCredentialID: types.Int64Value(int64(stack.GitConfig.Authentication.GitCredentialID)),
					Password:        types.StringValue(stack.GitConfig.Authentication.Password),
					Username:        types.StringValue(stack.GitConfig.Authentication.Username),
				}
			}
		}

		state.Stacks = append(state.Stacks, stackState)
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *stacksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*portainer.APIClient)
}
