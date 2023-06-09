package portainer

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
	dockerApi "terraform-provider-portainer/client/docker"
	portainerApi "terraform-provider-portainer/client/portainer"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &portainerProvider{}
)

type ProviderData struct {
	host            string
	portainerClient *portainerApi.APIClient
	dockerClient    *dockerApi.APIClient
}

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &portainerProvider{}
}

// portainerProvider is the provider implementation.
type portainerProvider struct{}

// Metadata returns the provider type name.
func (p *portainerProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "portainer"
}

// Schema defines the provider-level schema for configuration data.
func (p *portainerProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Portainer provider.",
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "URI for Portainer API. May also be provided via PORTAINER_HOST environment variable.",
				Optional:    true,
			},
			"api_key": schema.StringAttribute{
				Description: "Api Key for Portainer API. May also be provided via PORTAINER_APIKEY environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// portainerProviderModel maps provider schema data to a Go type.
type portainerProviderModel struct {
	Host   types.String `tfsdk:"host"`
	ApiKey types.String `tfsdk:"api_key"`
}

// Configure prepares a portainer API client for data sources and resources.
func (p *portainerProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config portainerProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown portainer API Host",
			"The provider cannot create the Portainer API client as there is an unknown configuration value for the Portainer API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PORTAINER_HOST environment variable.",
		)
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("apiKey"),
			"Unknown Portainer API ApiKey",
			"The provider cannot create the Portainer API client as there is an unknown configuration value for the Portainer API apiKey. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PORTAINER_APiKEY environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("PORTAINER_HOST")
	apiKey := os.Getenv("PORTAINER_APIKEY")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.ApiKey.IsNull() {
		apiKey = config.ApiKey.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Portainer API Host",
			"The provider cannot create the Portainer API client as there is a missing or empty value for the Portainer API apiKey. "+
				"Set the host value in the configuration or use the PORTAINER_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("apiKey"),
			"Missing Portainer API ApiKey",
			"The provider cannot create the Portainer API client as there is a missing or empty value for the Portainer API apiKey. "+
				"Set the api_key value in the configuration or use the PORTAINER_APIKEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "portainer_host", host)
	ctx = tflog.SetField(ctx, "portainer_api_key", apiKey)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "portainer_api_key")

	tflog.Debug(ctx, "Creating Portainer client")

	// Create a new Portainer client using the configuration values
	clientConfig := portainerApi.NewConfiguration()
	clientConfig.BasePath = host
	clientConfig.DefaultHeader = map[string]string{
		"X-API-Key": apiKey,
	}
	portainerClient := portainerApi.NewAPIClient(clientConfig)

	dockerConfig := dockerApi.NewConfiguration()
	// Portainer provides a docker API proxy available at: /api/endpoints/(endpointId}/docker
	// BasePath needs to be changed before attempting to make actual requests
	dockerConfig.BasePath = host
	dockerConfig.DefaultHeader = map[string]string{
		"X-API-Key": apiKey,
	}
	dockerClient := dockerApi.NewAPIClient(dockerConfig)

	providerData := ProviderData{
		host:            host,
		portainerClient: portainerClient,
		dockerClient:    dockerClient,
	}

	// Make the Portainer client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = providerData
	resp.ResourceData = providerData

	tflog.Info(ctx, "Configured Portainer client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *portainerProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewStacksDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *portainerProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewStackResource,
		NewSecretResource,
	}
}
