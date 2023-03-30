# SettingsSettingsUpdatePayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EdgePortainerURL** | **string** | EdgePortainerURL is the URL that is exposed to edge agents | [optional] [default to null]
**ShowKomposeBuildOption** | **bool** | Show the Kompose build option (discontinued in 2.18) | [optional] [default to null]
**AuthenticationMethod** | **int32** | Active authentication method for the Portainer instance. Valid values are: 1 for internal, 2 for LDAP, or 3 for oauth | [optional] [default to null]
**BlackListedLabels** | [**[]PortainerPair**](portainer.Pair.md) | A list of label name &amp; value that will be used to hide containers when querying containers | [optional] [default to null]
**EdgeAgentCheckinInterval** | **int32** | The default check in interval for edge agent (in seconds) | [optional] [default to null]
**EnableEdgeComputeFeatures** | **bool** | Whether edge compute features are enabled | [optional] [default to null]
**EnableTelemetry** | **bool** | Whether telemetry is enabled | [optional] [default to null]
**EnforceEdgeID** | **bool** | EnforceEdgeID makes Portainer store the Edge ID instead of accepting anyone | [optional] [default to null]
**HelmRepositoryURL** | **string** | Helm repository URL | [optional] [default to null]
**InternalAuthSettings** | [***PortainerInternalAuthSettings**](portainer.InternalAuthSettings.md) |  | [optional] [default to null]
**KubeconfigExpiry** | **string** | The expiry of a Kubeconfig | [optional] [default to null]
**KubectlShellImage** | **string** | Kubectl Shell Image | [optional] [default to null]
**Ldapsettings** | [***PortainerLdapSettings**](portainer.LDAPSettings.md) |  | [optional] [default to null]
**LogoURL** | **string** | URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string | [optional] [default to null]
**OauthSettings** | [***PortainerOAuthSettings**](portainer.OAuthSettings.md) |  | [optional] [default to null]
**SnapshotInterval** | **string** | The interval in which environment(endpoint) snapshots are created | [optional] [default to null]
**TemplatesURL** | **string** | URL to the templates that will be displayed in the UI when navigating to App Templates | [optional] [default to null]
**TrustOnFirstConnect** | **bool** | TrustOnFirstConnect makes Portainer accepting edge agent connection by default | [optional] [default to null]
**UserSessionTimeout** | **string** | The duration of a user session | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


