# SettingsPublicSettingsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | **int32** | Active authentication method for the Portainer instance. Valid values are: 1 for internal, 2 for LDAP, or 3 for oauth | [optional] [default to null]
**EnableEdgeComputeFeatures** | **bool** | Whether edge compute features are enabled | [optional] [default to null]
**EnableTelemetry** | **bool** | Whether telemetry is enabled | [optional] [default to null]
**Features** | **map[string]bool** | Supported feature flags | [optional] [default to null]
**LogoURL** | **string** | URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string | [optional] [default to null]
**OAuthLoginURI** | **string** | The URL used for oauth login | [optional] [default to null]
**OAuthLogoutURI** | **string** | The URL used for oauth logout | [optional] [default to null]
**RequiredPasswordLength** | **int32** | The minimum required length for a password of any user when using internal auth mode | [optional] [default to null]
**ShowKomposeBuildOption** | **bool** | Show the Kompose build option (discontinued in 2.18) | [optional] [default to null]
**TeamSync** | **bool** | Whether team sync is enabled | [optional] [default to null]
**Edge** | [***SettingsPublicSettingsResponseEdge**](settings.publicSettingsResponse_edge.md) |  | [optional] [default to null]
**IsAMTEnabled** | **bool** | Whether AMT is enabled | [optional] [default to null]
**IsFDOEnabled** | **bool** | Whether FDO is enabled | [optional] [default to null]
**KubeconfigExpiry** | **string** | The expiry of a Kubeconfig | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


