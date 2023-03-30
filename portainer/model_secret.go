package portainer

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Secrets struct {
	ID      types.String `tfsdk:"id"` // required for acceptance testing
	Secrets []Secret     `tfsdk:"secrets"`
}

// Secret maps secrets schema data.
type Secret struct {
	ID         types.String `tfsdk:"id"`
	EndpointId types.Int64  `tfsdk:"endpoint_id"`
	Name       types.String `tfsdk:"name"`
	Data       types.String `tfsdk:"data"`
	Labels     types.Map    `tfsdk:"labels"`
	Version    types.Int64  `tfsdk:"version"`
}
