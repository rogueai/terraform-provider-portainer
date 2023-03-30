package portainer

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	dockerApi "terraform-provider-portainer/client/docker"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &secretResource{}
	_ resource.ResourceWithConfigure = &secretResource{}
)

// NewSecretResource is a helper function to simplify the provider implementation.
func NewSecretResource() resource.Resource {
	return &secretResource{}
}

// secretResource is the resource implementation.
type secretResource struct {
	data ProviderData
}

// Configure adds the provider configured client to the resource.
func (r *secretResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	data := req.ProviderData.(ProviderData)
	r.data = data
}

// Metadata returns the resource type name.
func (r *secretResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

// Schema defines the schema for the resource.
func (r *secretResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = secretResourceSchema()
}

func (r *secretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

// Create creates the resource and sets the initial Terraform state.
func (r *secretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan Secret
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new secret
	labels := map[string]string{}
	_ = plan.Labels.ElementsAs(ctx, labels, false)
	secretCreateOpts := dockerApi.SecretApiSecretCreateOpts{
		Body: optional.NewInterface(dockerApi.SecretSpec{
			Name:   plan.Name.ValueString(),
			Data:   plan.Data.ValueString(),
			Labels: labels,
		}),
	}

	endpointId := int32(plan.EndpointId.ValueInt64())
	r.data.dockerClient.ChangeBasePath(r.data.host + "/endpoints/" + fmt.Sprint(endpointId) + "/docker")
	res, _, err := r.data.dockerClient.SecretApi.SecretCreate(ctx, &secretCreateOpts)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating secret",
			"Could not create secret, unexpected error: "+err.Error(),
		)
		return
	}
	plan.ID = types.StringValue(fmt.Sprint(res.Id))

	// Get refreshed secret value from Portainer
	secretId := plan.ID.ValueString()
	secret, _, err := r.data.dockerClient.SecretApi.SecretInspect(ctx, secretId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Secret",
			"Could not read Portainer secret ID "+plan.ID.String()+": "+err.Error(),
		)
		return
	}
	plan.Name = types.StringValue(secret.Spec.Name)
	plan.Version = types.Int64Value(int64(secret.Version.Index))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *secretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Secret
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed secret value from Portainer
	endpointId := int32(state.EndpointId.ValueInt64())
	r.data.dockerClient.ChangeBasePath(r.data.host + "/endpoints/" + fmt.Sprint(endpointId) + "/docker")
	secretId := state.ID.ValueString()
	secret, _, err := r.data.dockerClient.SecretApi.SecretInspect(ctx, secretId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Secret",
			"Could not read Portainer secret ID "+state.ID.String()+": "+err.Error(),
		)
		return
	}
	state.ID = types.StringValue(secret.ID)
	state.Name = types.StringValue(secret.Spec.Name)
	state.Version = types.Int64Value(int64(secret.Version.Index))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *secretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan Secret
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	secretId := plan.ID.ValueString()
	endpointId := int32(plan.EndpointId.ValueInt64())

	r.data.dockerClient.ChangeBasePath(r.data.host + "/endpoints/" + fmt.Sprint(endpointId) + "/docker")

	labels := map[string]string{}
	_ = plan.Labels.ElementsAs(ctx, labels, false)
	secretUpdateOpts := dockerApi.SecretApiSecretUpdateOpts{
		Body: optional.NewInterface(dockerApi.SecretSpec{
			Name:   plan.Name.ValueString(),
			Data:   plan.Data.ValueString(),
			Labels: labels,
		}),
	}
	_, err := r.data.dockerClient.SecretApi.SecretUpdate(ctx, secretId, plan.Version.ValueInt64(), &secretUpdateOpts)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Secret",
			"Could not create secret, unexpected error: "+err.Error(),
		)
		return
	}
	// Get refreshed secret value from Portainer
	secret, _, err := r.data.dockerClient.SecretApi.SecretInspect(ctx, secretId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Portainer Secret",
			"Could not read Portainer secret ID "+plan.ID.String()+": "+err.Error(),
		)
		return
	}
	plan.ID = types.StringValue(secret.ID)
	plan.Name = types.StringValue(secret.Spec.Name)
	plan.Version = types.Int64Value(int64(secret.Version.Index))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *secretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Secret
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	secretId := state.ID.ValueString()
	endpointId := int32(state.EndpointId.ValueInt64())
	r.data.dockerClient.ChangeBasePath(r.data.host + "/endpoints/" + fmt.Sprint(endpointId) + "/docker")
	// Delete existing secret
	_, err := r.data.dockerClient.SecretApi.SecretDelete(ctx, secretId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Portainer Secret",
			"Could not delete secret, unexpected error: "+err.Error(),
		)
		return
	}
}

func secretResourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"endpoint_id": schema.Int64Attribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"data": schema.StringAttribute{
				Required:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"labels": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"version": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}
