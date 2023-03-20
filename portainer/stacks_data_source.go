package portainer

import (
	"context"
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
							ElementType: types.StringType,
						},
						"auto_update": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"interval": schema.StringAttribute{},
								"job_id":   schema.StringAttribute{},
								"webhook":  schema.StringAttribute{},
							},
						},
						// ===
						"endpoint_id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"teaser": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"price": schema.Float64Attribute{
							Computed: true,
						},
						"image": schema.StringAttribute{
							Computed: true,
						},
						"ingredients": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *stacksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
}

// Configure adds the provider configured client to the data source.
func (d *stacksDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*portainer.APIClient)
}
