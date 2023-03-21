package portainer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	portainer "terraform-provider-portainer/client"
	"terraform-provider-portainer/portainer/model"
	"terraform-provider-portainer/portainer/schema"
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
	resp.Schema = schema.StacksSchema()
}

// Read refreshes the Terraform state with the latest data.
func (d *stacksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state model.Stacks

	stacks, _, err := d.client.StacksApi.StackList(ctx, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Portainer Stack",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, stack := range stacks {
		stackState := model.Stack{
			AdditionalFiles: nil,
			AutoUpdate:      model.StackAutoUpdate{},
			EndpointId:      types.Int64Value(int64(stack.EndpointId)),
			EntryPoint:      types.StringValue(stack.EntryPoint),
			Env:             nil,
			Id:              types.Int64Value(int64(stack.Id)),
			Name:            types.StringValue(stack.Name),
			Option:          model.StackOption{},
			ResourceControl: model.StackResourceControl{},
			Status:          types.Int64Value(int64(stack.Status)),
			SwarmId:         types.StringValue(stack.SwarmId),
			Type:            types.Int64Value(int64(stack.Type_)),
			CreatedBy:       types.StringValue(stack.CreatedBy),
			CreationDate:    types.Int64Value(int64(stack.CreationDate)),
			FromAppTemplate: types.BoolValue(stack.FromAppTemplate),
			GitConfig:       model.StackGitAuthenticationConfig{},
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
			stackState.Env = append(stackState.Env, model.StackEnvModel{
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
				stackState.ResourceControl.TeamAccesses = append(stackState.ResourceControl.TeamAccesses, model.StackTeamAccesses{
					AccessLevel: types.Int64Value(int64(teamAccess.AccessLevel)),
					TeamId:      types.Int64Value(int64(teamAccess.TeamId)),
				})
			}
			for _, userAccess := range stack.ResourceControl.UserAccesses {
				stackState.ResourceControl.UserAccesses = append(stackState.ResourceControl.UserAccesses, model.StackUserAccesses{
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
				stackState.GitConfig.Authentication = model.StackGitAuthentication{
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
func (d *stacksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	tflog.Info(ctx, "Configuring HashiCups client")

	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*portainer.APIClient)
}
