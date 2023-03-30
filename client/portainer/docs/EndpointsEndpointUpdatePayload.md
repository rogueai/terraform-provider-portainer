# EndpointsEndpointUpdatePayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureApplicationID** | **string** | Azure application ID | [optional] [default to null]
**AzureAuthenticationKey** | **string** | Azure authentication key | [optional] [default to null]
**AzureTenantID** | **string** | Azure tenant ID | [optional] [default to null]
**EdgeCheckinInterval** | **int32** | The check in interval for edge agent (in seconds) | [optional] [default to null]
**Gpus** | [**[]PortainerPair**](portainer.Pair.md) | GPUs information | [optional] [default to null]
**GroupID** | **int32** | Group identifier | [optional] [default to null]
**Kubernetes** | [***PortainerKubernetesData**](portainer.KubernetesData.md) | Associated Kubernetes data | [optional] [default to null]
**Name** | **string** | Name that will be used to identify this environment(endpoint) | [optional] [default to null]
**PublicURL** | **string** | URL or IP address where exposed containers will be reachable.\\ Defaults to URL if not specified | [optional] [default to null]
**Status** | **int32** | The status of the environment(endpoint) (1 - up, 2 - down) | [optional] [default to null]
**TagIDs** | **[]int32** | List of tag identifiers to which this environment(endpoint) is associated | [optional] [default to null]
**TeamAccessPolicies** | [***PortainerTeamAccessPolicies**](portainer.TeamAccessPolicies.md) |  | [optional] [default to null]
**Tls** | **bool** | Require TLS to connect against this environment(endpoint) | [optional] [default to null]
**TlsskipClientVerify** | **bool** | Skip client verification when using TLS | [optional] [default to null]
**TlsskipVerify** | **bool** | Skip server verification when using TLS | [optional] [default to null]
**Url** | **string** | URL or IP address of a Docker host | [optional] [default to null]
**UserAccessPolicies** | [***PortainerUserAccessPolicies**](portainer.UserAccessPolicies.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


