/*
 * PortainerCE API
 *
 * Portainer API is an HTTP API served by Portainer. It is used by the Portainer UI and everything you can do with the UI can be done using the HTTP API. Examples are available at https://documentation.portainer.io/api/api-examples/ You can find out more about Portainer at [http://portainer.io](http://portainer.io) and get some support on [Slack](http://portainer.io/slack/).  # Authentication  Most of the API environments(endpoints) require to be authenticated as well as some level of authorization to be used. Portainer API uses JSON Web Token to manage authentication and thus requires you to provide a token in the **Authorization** header of each request with the **Bearer** authentication mechanism.  Example:  ``` Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGUiOjEsImV4cCI6MTQ5OTM3NjE1NH0.NJ6vE8FY1WG6jsRQzfMqeatJ4vh2TWAeeYfDhP71YEE ```  # Security  Each API environment(endpoint) has an associated access policy, it is documented in the description of each environment(endpoint).  Different access policies are available:  - Public access - Authenticated access - Restricted access - Administrator access  ### Public access  No authentication is required to access the environments(endpoints) with this access policy.  ### Authenticated access  Authentication is required to access the environments(endpoints) with this access policy.  ### Restricted access  Authentication is required to access the environments(endpoints) with this access policy. Extra-checks might be added to ensure access to the resource is granted. Returned data might also be filtered.  ### Administrator access  Authentication as well as an administrator role are required to access the environments(endpoints) with this access policy.  # Execute Docker requests  Portainer **DO NOT** expose specific environments(endpoints) to manage your Docker resources (create a container, remove a volume, etc...).  Instead, it acts as a reverse-proxy to the Docker HTTP API. This means that you can execute Docker requests **via** the Portainer HTTP API.  To do so, you can use the `/endpoints/{id}/docker` Portainer API environment(endpoint) (which is not documented below due to Swagger limitations). This environment(endpoint) has a restricted access policy so you still need to be authenticated to be able to query this environment(endpoint). Any query on this environment(endpoint) will be proxied to the Docker API of the associated environment(endpoint) (requests and responses objects are the same as documented in the Docker API).  # Private Registry  Using private registry, you will need to pass a based64 encoded JSON string ‘{\"registryId\":\\<registryID value\\>}’ inside the Request Header. The parameter name is \"X-Registry-Auth\". \\<registryID value\\> - The registry ID where the repository was created.  Example:  ``` eyJyZWdpc3RyeUlkIjoxfQ== ```  **NOTE**: You can find more information on how to query the Docker API in the [Docker official documentation](https://docs.docker.com/engine/api/v1.30/) as well as in [this Portainer example](https://documentation.portainer.io/api/api-examples/).
 *
 * API version: 2.17.1
 * Contact: info@portainer.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package portainer

type SettingsPublicSettingsResponse struct {
	// Active authentication method for the Portainer instance. Valid values are: 1 for internal, 2 for LDAP, or 3 for oauth
	AuthenticationMethod int32 `json:"AuthenticationMethod,omitempty"`
	// Whether edge compute features are enabled
	EnableEdgeComputeFeatures bool `json:"EnableEdgeComputeFeatures,omitempty"`
	// Whether telemetry is enabled
	EnableTelemetry bool `json:"EnableTelemetry,omitempty"`
	// Supported feature flags
	Features map[string]bool `json:"Features,omitempty"`
	// URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string
	LogoURL string `json:"LogoURL,omitempty"`
	// The URL used for oauth login
	OAuthLoginURI string `json:"OAuthLoginURI,omitempty"`
	// The URL used for oauth logout
	OAuthLogoutURI string `json:"OAuthLogoutURI,omitempty"`
	// The minimum required length for a password of any user when using internal auth mode
	RequiredPasswordLength int32 `json:"RequiredPasswordLength,omitempty"`
	// Show the Kompose build option (discontinued in 2.18)
	ShowKomposeBuildOption bool `json:"ShowKomposeBuildOption,omitempty"`
	// Whether team sync is enabled
	TeamSync bool                                `json:"TeamSync,omitempty"`
	Edge     *SettingsPublicSettingsResponseEdge `json:"edge,omitempty"`
	// Whether AMT is enabled
	IsAMTEnabled bool `json:"isAMTEnabled,omitempty"`
	// Whether FDO is enabled
	IsFDOEnabled bool `json:"isFDOEnabled,omitempty"`
	// The expiry of a Kubeconfig
	KubeconfigExpiry string `json:"kubeconfigExpiry,omitempty"`
}
