package portainer

import (
	"context"
	"encoding/json"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strings"
	portainer "terraform-provider-portainer/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &stackResource{}
	_ resource.ResourceWithConfigure = &stackResource{}
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
	//var file *os.File
	//if !plan.File.IsNull() {
	//	file, err = os.Open(plan.File.ValueString())
	//	if err != nil {
	//		resp.Diagnostics.AddError(
	//			"Error creating stack",
	//			"Could not create stack, unexpected error: "+err.Error(),
	//		)
	//		return
	//	}
	//}

	bodySwarmString := portainer.StacksSwarmStackFromFileContentPayload{
		Env:              nil,
		FromAppTemplate:  plan.FromAppTemplate.ValueBool(),
		Name:             plan.Name.ValueString(),
		StackFileContent: plan.FileContent.ValueString(),
		SwarmID:          plan.SwarmId.ValueString(),
	}
	//bodySwamRepository := portainer.StacksSwarmStackFromGitRepositoryPayload{}
	//bodyComposeStringSchema := portainer.StacksComposeStackFromFileContentPayload{}
	//bodyComposeRepositorySchema := portainer.StacksComposeStackFromGitRepositoryPayload{}
	//bodyKubernetesStringSchema := portainer.StacksKubernetesStringDeploymentPayload{}
	//bodyKubernetesRepositorySchema := portainer.StacksKubernetesGitDeploymentPayload{}
	// bodyKubernetesUrlSchema := portainer.StacksKubernetesManifestUrlDeploymentPayload{}
	localVarOptionals := portainer.StacksApiStackCreateOpts{
		BodySwarmString:          optional.NewInterface(bodySwarmString),
		BodySwarmRepository:      optional.EmptyInterface(), //optional.EmptyInterface()NewInterface(bodySwamRepository),
		BodyComposeString:        optional.EmptyInterface(), //optional.NewInterface(bodyComposeStringSchema),
		BodyComposeRepository:    optional.EmptyInterface(), //optional.NewInterface(bodyComposeRepositorySchema),
		BodyKubernetesString:     optional.EmptyInterface(), //optional.NewInterface(bodyKubernetesStringSchema),
		BodyKubernetesRepository: optional.EmptyInterface(), //optional.NewInterface(bodyKubernetesRepositorySchema),
		BodyKubernetesUrl:        optional.EmptyInterface(), //optional.NewInterface(bodyKubernetesUrlSchema),
		Name:                     optional.NewString(plan.Name.ValueString()),
		SwarmID:                  optional.NewString(plan.SwarmId.ValueString()),
		Env:                      optional.NewString(jsonEnv),
		File:                     optional.EmptyInterface(), //optional.NewInterface(file),
	}

	// Create new stack
	type_ := int32(plan.Type.ValueInt64())
	//method := plan.Method.ValueString()
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

	plan.ID = types.Int64Value(int64(res.Id))
	// Get refreshed stack value from Portainer
	stack, _, err := r.client.StacksApi.StackInspect(ctx, int32(plan.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack",
			"Could not read Portainer stack ID "+plan.ID.String()+": "+err.Error(),
		)
		return
	}

	plan.CreatedBy = types.StringValue(stack.CreatedBy)
	plan.CreationDate = types.Int64Value(int64(stack.CreationDate))
	plan.UpdatedBy = types.StringValue(stack.UpdatedBy)
	plan.UpdateDate = types.Int64Value(int64(stack.UpdateDate))
	plan.Status = types.Int64Value(int64(stack.Status))
	plan.ProjectPath = types.StringValue(stack.ProjectPath)
	plan.IsComposeFormat = types.BoolValue(stack.IsComposeFormat)
	plan.FromAppTemplate = types.BoolValue(stack.FromAppTemplate)

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
	stack, _, err := r.client.StacksApi.StackInspect(ctx, int32(state.ID.ValueInt64()))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Stack",
			"Could not read Portainer stack ID "+state.ID.String()+": "+err.Error(),
		)
		return
	}

	state.CreatedBy = types.StringValue(stack.CreatedBy)
	state.CreationDate = types.Int64Value(int64(stack.CreationDate))
	state.UpdatedBy = types.StringValue(stack.UpdatedBy)
	state.UpdateDate = types.Int64Value(int64(stack.UpdateDate))
	state.Status = types.Int64Value(int64(stack.Status))
	state.ProjectPath = types.StringValue(stack.ProjectPath)
	state.IsComposeFormat = types.BoolValue(stack.IsComposeFormat)
	state.FromAppTemplate = types.BoolValue(stack.FromAppTemplate)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *stackResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *stackResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

//func bodySwarmStringSchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"env": schema.ListNestedAttribute{
//				Optional: true,
//				NestedObject: schema.NestedAttributeObject{
//					Attributes: map[string]schema.Attribute{
//						"name":  schema.StringAttribute{Optional: true},
//						"value": schema.StringAttribute{Optional: true},
//					},
//				},
//			},
//			"from_app_template":  schema.BoolAttribute{Optional: true},
//			"name":               schema.StringAttribute{Optional: true},
//			"stack_file_content": schema.StringAttribute{Optional: true},
//			"swarm_id":           schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func bodySwarmRepositorySchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"additional_files": schema.ListAttribute{
//				Optional:    true,
//				ElementType: types.StringType,
//			},
//			"auto_update": schema.SingleNestedAttribute{
//				Optional: true,
//				Attributes: map[string]schema.Attribute{
//					"interval": schema.StringAttribute{Optional: true},
//					"job_id":   schema.StringAttribute{Optional: true},
//					"webhook":  schema.StringAttribute{Optional: true},
//				},
//			},
//			"compose_file": schema.StringAttribute{Optional: true},
//			"env": schema.ListNestedAttribute{
//				Optional: true,
//				NestedObject: schema.NestedAttributeObject{
//					Attributes: map[string]schema.Attribute{
//						"name":  schema.StringAttribute{Optional: true},
//						"value": schema.StringAttribute{Optional: true},
//					},
//				},
//			},
//			"from_app_template":         schema.BoolAttribute{Optional: true},
//			"name":                      schema.StringAttribute{Optional: true},
//			"repository_authentication": schema.BoolAttribute{Optional: true},
//			"repository_password":       schema.StringAttribute{Optional: true, Sensitive: true},
//			"repository_reference_name": schema.StringAttribute{Optional: true},
//			"repository_url":            schema.StringAttribute{Optional: true},
//			"repository_username":       schema.StringAttribute{Optional: true},
//			"swarm_id":                  schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func bodyComposeStringSchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"env": schema.ListNestedAttribute{
//				Optional: true,
//				NestedObject: schema.NestedAttributeObject{
//					Attributes: map[string]schema.Attribute{
//						"name":  schema.StringAttribute{Optional: true},
//						"value": schema.StringAttribute{Optional: true},
//					},
//				},
//			},
//			"from_app_template":  schema.BoolAttribute{Optional: true},
//			"name":               schema.StringAttribute{Optional: true},
//			"stack_file_content": schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func bodyComposeRepositorySchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"additional_files": schema.ListAttribute{
//				Optional:    true,
//				ElementType: types.StringType,
//			},
//			"auto_update": schema.SingleNestedAttribute{
//				Optional: true,
//				Attributes: map[string]schema.Attribute{
//					"interval": schema.StringAttribute{Optional: true},
//					"job_id":   schema.StringAttribute{Optional: true},
//					"webhook":  schema.StringAttribute{Optional: true},
//				},
//			},
//			"compose_file": schema.StringAttribute{Optional: true},
//			"env": schema.ListNestedAttribute{
//				Optional: true,
//				NestedObject: schema.NestedAttributeObject{
//					Attributes: map[string]schema.Attribute{
//						"name":  schema.StringAttribute{Optional: true},
//						"value": schema.StringAttribute{Optional: true},
//					},
//				},
//			},
//			"from_app_template":         schema.BoolAttribute{Optional: true},
//			"name":                      schema.StringAttribute{Optional: true},
//			"repository_authentication": schema.BoolAttribute{Optional: true},
//			"repository_password":       schema.StringAttribute{Optional: true, Sensitive: true},
//			"repository_reference_name": schema.StringAttribute{Optional: true},
//			"repository_url":            schema.StringAttribute{Optional: true},
//			"repository_username":       schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func bodyKubernetesStringSchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"compose_format":     schema.BoolAttribute{Optional: true},
//			"namespace":          schema.StringAttribute{Optional: true},
//			"stack_file_content": schema.StringAttribute{Optional: true},
//			"stack_name":         schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func bodyKubernetesRepositorySchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"additional_files": schema.ListAttribute{
//				Optional:    true,
//				ElementType: types.StringType,
//			},
//			"auto_update": schema.SingleNestedAttribute{
//				Optional: true,
//				Attributes: map[string]schema.Attribute{
//					"interval": schema.StringAttribute{Optional: true},
//					"job_id":   schema.StringAttribute{Optional: true},
//					"webhook":  schema.StringAttribute{Optional: true},
//				},
//			},
//			"compose_format":            schema.BoolAttribute{Optional: true},
//			"manifest_file":             schema.StringAttribute{Optional: true},
//			"namespace":                 schema.StringAttribute{Optional: true},
//			"repository_authentication": schema.BoolAttribute{Optional: true},
//			"repository_password":       schema.StringAttribute{Optional: true, Sensitive: true},
//			"repository_reference_name": schema.StringAttribute{Optional: true},
//			"repository_url":            schema.StringAttribute{Optional: true},
//			"repository_username":       schema.StringAttribute{Optional: true},
//			"stack_name":                schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func bodyKubernetesUrlSchema() schema.SingleNestedAttribute {
//	return schema.SingleNestedAttribute{
//		Optional: true,
//		Attributes: map[string]schema.Attribute{
//			"compose_format": schema.BoolAttribute{Optional: true},
//			"manifest_url":   schema.StringAttribute{Optional: true},
//			"namespace":      schema.StringAttribute{Optional: true},
//			"stack_name":     schema.StringAttribute{Optional: true},
//		},
//	}
//}
//
//func stackResourceSchema() schema.Schema {
//	return schema.Schema{
//		Attributes: map[string]schema.Attribute{
//			"type": schema.Int64Attribute{
//				Required: true,
//			},
//			"method": schema.StringAttribute{
//				Required: true,
//			},
//			"endpoint_id": schema.Int64Attribute{
//				Required: true,
//			},
//			"body_swarm_string":          bodySwarmStringSchema(),
//			"body_swarm_repository":      bodySwarmRepositorySchema(),
//			"body_compose_string":        bodyComposeStringSchema(),
//			"body_compose_repository":    bodyComposeRepositorySchema(),
//			"body_kubernetes_string":     bodyKubernetesStringSchema(),
//			"body_kubernetes_repository": bodyKubernetesRepositorySchema(),
//			"body_kubernetes_url":        bodyKubernetesUrlSchema(),
//			"name": schema.StringAttribute{
//				Optional: true,
//			},
//			"swarm_id": schema.StringAttribute{
//				Optional: true,
//			},
//			"env": schema.ListNestedAttribute{
//				Optional: true,
//				NestedObject: schema.NestedAttributeObject{
//					Attributes: map[string]schema.Attribute{
//						"name":  schema.StringAttribute{Optional: true},
//						"value": schema.StringAttribute{Optional: true},
//					},
//				},
//			},
//			"file": schema.StringAttribute{
//				Optional: true,
//			},
//		},
//	}
//}

func stackResourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
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
				Optional: true,
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
			"id": schema.Int64Attribute{
				Computed: true,
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
						Computed:    true,
						ElementType: types.StringType,
					},
					"system": schema.BoolAttribute{Optional: true},
					"team_accesses": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_level": schema.Int64Attribute{Optional: true},
								"team_id":      schema.Int64Attribute{Optional: true},
							},
						},
					},
					"type": schema.Int64Attribute{Optional: true},
					"user_accesses": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"access_level": schema.Int64Attribute{Optional: true},
								"user_id":      schema.Int64Attribute{Optional: true},
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
