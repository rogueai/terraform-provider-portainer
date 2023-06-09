package portainer

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	data ProviderData
}

func (d *stacksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stacks"
}

// Schema defines the schema for the data source.
func (d *stacksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = stacksDataSourceSchema()
}

// Configure adds the provider configured client to the data source.
func (d *stacksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Portainer client")

	if req.ProviderData == nil {
		return
	}

	data := req.ProviderData.(ProviderData)
	d.data.portainerClient = data.portainerClient
}

// Read refreshes the Terraform state with the latest data.
func (d *stacksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state Stacks

	stacks, _, err := d.data.portainerClient.StacksApi.StackList(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Portainer Stack",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, stack := range stacks {
		stackState := Stack{}

		stackState.AdditionalFiles = nil
		stackState.AutoUpdate = nil
		stackState.EndpointId = types.Int64Value(int64(stack.EndpointId))
		stackState.EntryPoint = types.StringValue(stack.EntryPoint)
		stackState.Env = nil
		stackState.ID = types.StringValue(fmt.Sprint(stack.Id))
		stackState.Name = types.StringValue(stack.Name)
		stackState.Option = nil
		stackState.ResourceControl = &StackResourceControl{}
		stackState.Status = types.Int64Value(int64(stack.Status))
		stackState.SwarmId = types.StringValue(stack.SwarmId)
		stackState.Type = types.Int64Value(int64(stack.Type_))
		stackState.CreatedBy = types.StringValue(stack.CreatedBy)
		stackState.CreationDate = types.Int64Value(int64(stack.CreationDate))
		stackState.FromAppTemplate = types.BoolValue(stack.FromAppTemplate)
		stackState.GitConfig = &StackGitAuthenticationConfig{}
		stackState.IsComposeFormat = types.BoolValue(stack.IsComposeFormat)
		stackState.Namespace = types.StringValue(stack.Namespace)
		stackState.ProjectPath = types.StringValue(stack.ProjectPath)
		stackState.UpdateDate = types.Int64Value(int64(stack.UpdateDate))
		stackState.UpdatedBy = types.StringValue(stack.UpdatedBy)

		for _, additionalFile := range stack.AdditionalFiles {
			stackState.AdditionalFiles = append(stackState.AdditionalFiles, types.StringValue(additionalFile))
		}

		if stack.AutoUpdate != nil {
			autoUpdate := StackAutoUpdate{
				Interval: types.StringValue(stack.AutoUpdate.Interval),
				JobID:    types.StringValue(stack.AutoUpdate.JobID),
				Webhook:  types.StringValue(stack.AutoUpdate.Webhook),
			}
			stackState.AutoUpdate = &autoUpdate
		}

		if stack.Option != nil {
			option := StackOption{
				Prune: types.BoolValue(stack.Option.Prune),
			}
			stackState.Option = &option
		}

		for _, env := range stack.Env {
			stackState.Env = append(stackState.Env, StackEnv{
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
				stackState.ResourceControl.SubResourceIds = append(stackState.ResourceControl.SubResourceIds, types.StringValue(subResourceId))
			}
			stackState.ResourceControl.System = types.BoolValue(stack.ResourceControl.System)
			for _, teamAccess := range stack.ResourceControl.TeamAccesses {
				stackState.ResourceControl.TeamAccesses = append(stackState.ResourceControl.TeamAccesses, StackTeamAccesses{
					AccessLevel: types.Int64Value(int64(teamAccess.AccessLevel)),
					TeamId:      types.Int64Value(int64(teamAccess.TeamId)),
				})
			}
			for _, userAccess := range stack.ResourceControl.UserAccesses {
				stackState.ResourceControl.UserAccesses = append(stackState.ResourceControl.UserAccesses, StackUserAccesses{
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
				stackState.GitConfig.Authentication = StackGitAuthentication{
					GitCredentialID: types.Int64Value(int64(stack.GitConfig.Authentication.GitCredentialID)),
					Password:        types.StringValue(stack.GitConfig.Authentication.Password),
					Username:        types.StringValue(stack.GitConfig.Authentication.Username),
				}
			}
		}

		state.Stacks = append(state.Stacks, stackState)
	}

	// required for acceptance testing
	state.ID = types.StringValue("placeholder")

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func stacksDataSourceSchema() schema.Schema {
	return schema.Schema{
		Description: "Fetches the list of stacks.",
		Attributes: map[string]schema.Attribute{
			// required for acceptance testing
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"stacks": schema.ListNestedAttribute{
				Description: "List of stacks.",
				Computed:    true,
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
						"id": schema.StringAttribute{
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
						"file_content": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}
