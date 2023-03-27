package portainer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"strconv"
	"strings"
	portainer "terraform-provider-portainer/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &stackResource{}
	_ resource.ResourceWithConfigure   = &stackResource{}
	_ resource.ResourceWithImportState = &stackResource{}
)

// NewStackResource is a helper function to simplify the provider implementation.
func NewStackResource() resource.Resource {
	return &stackResource{}
}

// stackResource is the resource implementation.
type stackResource struct {
	client *portainer.APIClient
}

// Configure adds the provider configured client to the resource.
func (r *stackResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*portainer.APIClient)
}

// Metadata returns the resource type name.
func (r *stackResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_stack"
}

// Schema defines the schema for the resource.
func (r *stackResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = stackResourceSchema()
}

func (r *stackResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create creates the resource and sets the initial Terraform state.
func (r *stackResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan Stack
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	sb := new(strings.Builder)
	jsonEncoder := json.NewEncoder(sb)
	err := jsonEncoder.Encode(plan.Env)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating stack",
			"Could not create stack, unexpected error: "+err.Error(),
		)
		return
	}
	jsonEnv := sb.String()

	bodySwarmString := portainer.StacksSwarmStackFromFileContentPayload{
		Env:              []portainer.PortainerPair{},
		FromAppTemplate:  plan.FromAppTemplate.ValueBool(),
		Name:             plan.Name.ValueString(),
		StackFileContent: plan.FileContent.ValueString(),
		SwarmID:          plan.SwarmId.ValueString(),
	}
	localVarOptionals := portainer.StacksApiStackCreateOpts{
		BodySwarmString:          optional.NewInterface(bodySwarmString),
		BodySwarmRepository:      optional.EmptyInterface(),
		BodyComposeString:        optional.EmptyInterface(),
		BodyComposeRepository:    optional.EmptyInterface(),
		BodyKubernetesString:     optional.EmptyInterface(),
		BodyKubernetesRepository: optional.EmptyInterface(),
		BodyKubernetesUrl:        optional.EmptyInterface(),
		Name:                     optional.NewString(plan.Name.ValueString()),
		SwarmID:                  optional.NewString(plan.SwarmId.ValueString()),
		Env:                      optional.NewString(jsonEnv),
		File:                     optional.EmptyInterface(),
	}

	// Create new stack
	type_ := int32(plan.Type.ValueInt64())
	// TODO we only support string stack content for now
	method := "string"
	endpointId := int32(plan.EndpointId.ValueInt64())
	res, _, err := r.client.StacksApi.StackCreate(ctx, type_, method, endpointId, &localVarOptionals)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating stack",
			"Could not create stack, unexpected error: "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(fmt.Sprint(res.Id))
	stackId, _ := strconv.Atoi(plan.ID.ValueString())
	// Get refreshed stack value from Portainer
	stack, _, err := r.client.StacksApi.StackInspect(ctx, int32(stackId))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack",
			"Could not read Portainer stack ID "+plan.ID.String()+": "+err.Error(),
		)
		return
	}

	updateStack(&plan, &stack)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *stackResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Stack
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed stack value from Portainer
	stackId, _ := strconv.Atoi(state.ID.ValueString())
	stack, _, err := r.client.StacksApi.StackInspect(ctx, int32(stackId))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack",
			"Could not read Portainer stack ID "+state.ID.String()+": "+err.Error(),
		)
		return
	}

	state.ID = types.StringValue(fmt.Sprint(stack.Id))

	updateStack(&state, &stack)

	tflog.Debug(ctx, "Fetching stack file content")
	stackFileContent, _, err := r.client.StacksApi.StackFileInspect(ctx, int32(stackId))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack file content",
			"Could not read Portainer stack ID "+state.ID.String()+": "+err.Error(),
		)
		return
	}
	state.FileContent = types.StringValue(stackFileContent.StackFileContent)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *stackResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan Stack
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	stackId, _ := strconv.Atoi(plan.ID.ValueString())
	//endpointId := int32(plan.EndpointId.ValueInt64())

	stackUpdateBody := portainer.StacksUpdateSwarmStackPayload{
		StackFileContent: plan.FileContent.ValueString(),
		Env:              []portainer.PortainerPair{},
		Prune:            true,
		PullImage:        true,
	}
	localVarOptionals := portainer.StacksApiStackUpdateOpts{
		//EndpointId: optional.NewInt32(endpointId),
	}
	res, _, err := r.client.StacksApi.StackUpdate(ctx, int32(stackId), stackUpdateBody, &localVarOptionals)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating stack",
			"Could not create stack, unexpected error: "+err.Error(),
		)
		return
	}

	plan.ID = types.StringValue(fmt.Sprint(res.Id))

	// Get refreshed stack value from Portainer
	stack, _, err := r.client.StacksApi.StackInspect(ctx, int32(stackId))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack",
			"Could not read Portainer stack ID "+plan.ID.String()+": "+err.Error(),
		)
		return
	}

	updateStack(&plan, &stack)

	stackFileContent, _, err := r.client.StacksApi.StackFileInspect(ctx, int32(stackId))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack file content",
			"Could not read Portainer stack ID "+plan.ID.String()+": "+err.Error(),
		)
		return
	}
	plan.FileContent = types.StringValue(stackFileContent.StackFileContent)

	//plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *stackResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Stack
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	stackId, _ := strconv.Atoi(state.ID.ValueString())
	//endpointId := int32(state.EndpointId.ValueInt64())
	localVarOptions := portainer.StacksApiStackDeleteOpts{
		//	EndpointId: optional.NewInt32(endpointId),
		//	External: optional.NewBool(false),
	}
	// Delete existing order
	_, err := r.client.StacksApi.StackDelete(ctx, int32(stackId), &localVarOptions)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting HashiCups Order",
			"Could not delete order, unexpected error: "+err.Error(),
		)
		return
	}
}

func updateStack(plan *Stack, stack *portainer.PortainerStack) {
	plan.AutoUpdate = nil
	plan.EndpointId = types.Int64Value(int64(stack.EndpointId))
	plan.EntryPoint = types.StringValue(stack.EntryPoint)
	plan.Env = nil
	plan.ID = types.StringValue(fmt.Sprint(stack.Id))
	plan.Name = types.StringValue(stack.Name)
	plan.Option = nil
	plan.ResourceControl = nil
	plan.Status = types.Int64Value(int64(stack.Status))
	plan.SwarmId = types.StringValue(stack.SwarmId)
	plan.Type = types.Int64Value(int64(stack.Type_))
	plan.CreatedBy = types.StringValue(stack.CreatedBy)
	plan.CreationDate = types.Int64Value(int64(stack.CreationDate))
	plan.FromAppTemplate = types.BoolValue(stack.FromAppTemplate)
	plan.GitConfig = nil
	plan.IsComposeFormat = types.BoolValue(stack.IsComposeFormat)
	plan.Namespace = types.StringValue(stack.Namespace)
	plan.ProjectPath = types.StringValue(stack.ProjectPath)
	plan.UpdateDate = types.Int64Value(int64(stack.UpdateDate))
	plan.UpdatedBy = types.StringValue(stack.UpdatedBy)

	for _, additionalFile := range stack.AdditionalFiles {
		plan.AdditionalFiles = append(plan.AdditionalFiles, types.StringValue(additionalFile))
	}

	if stack.AutoUpdate != nil {
		autoUpdate := StackAutoUpdate{
			Interval: types.StringValue(stack.AutoUpdate.Interval),
			JobID:    types.StringValue(stack.AutoUpdate.JobID),
			Webhook:  types.StringValue(stack.AutoUpdate.Webhook),
		}
		plan.AutoUpdate = &autoUpdate
	}

	if stack.Option != nil {
		option := StackOption{
			Prune: types.BoolValue(stack.Option.Prune),
		}
		plan.Option = &option
	}

	for _, env := range stack.Env {
		plan.Env = append(plan.Env, StackEnv{
			Name:  types.StringValue(env.Name),
			Value: types.StringValue(env.Value),
		})
	}

}

func stackResourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"additional_files": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"auto_update": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"interval": schema.StringAttribute{Computed: true},
					"job_id":   schema.StringAttribute{Computed: true},
					"webhook":  schema.StringAttribute{Computed: true},
				},
			},
			"endpoint_id": schema.Int64Attribute{
				Required: true,
			},
			"entry_point": schema.StringAttribute{
				Computed: true,
			},
			"env": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
						},
						"value": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"option": schema.ObjectAttribute{
				Optional: true,
				AttributeTypes: map[string]attr.Type{
					"prune": types.BoolType,
				},
			},
			// TODO unmapped
			"resource_control": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"access_level":        schema.Int64Attribute{Optional: true},
					"administrators_only": schema.BoolAttribute{Optional: true},
					"id":                  schema.Int64Attribute{Optional: true},
					"owner_id":            schema.Int64Attribute{Optional: true},
					"public":              schema.BoolAttribute{Optional: true},
					"resource_id":         schema.StringAttribute{Optional: true},
					"sub_resource_ids": schema.ListAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"system": schema.BoolAttribute{Optional: true},
					"team_accesses": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_level": schema.Int64Attribute{Required: true},
								"team_id":      schema.Int64Attribute{Required: true},
							},
						},
					},
					"type": schema.Int64Attribute{Optional: true},
					"user_accesses": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_level": schema.Int64Attribute{Required: true},
								"user_id":      schema.Int64Attribute{Required: true},
							},
						},
					},
				},
			},
			"status": schema.Int64Attribute{
				Computed: true,
			},
			"swarm_id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.Int64Attribute{
				Required: true,
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
			// TODO unmapped
			"git_config": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"authentication": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"git_credential_id": schema.Int64Attribute{Optional: true},
							"password": schema.StringAttribute{
								Optional: true,
							},
							"username": schema.StringAttribute{Optional: true},
						},
					},
					"config_file_path": schema.StringAttribute{Computed: true},
					"config_hash":      schema.StringAttribute{Computed: true},
					"reference_name":   schema.StringAttribute{Required: true},
					"url":              schema.StringAttribute{Required: true},
				},
			},
			"is_compose_format": schema.BoolAttribute{
				Computed: true,
			},
			"namespace": schema.StringAttribute{
				Optional: true,
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
	}
}
