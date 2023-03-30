# PortainerEndpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AMTDeviceGUID** | **string** | The identifier of the AMT Device associated with this environment(endpoint) | [optional] [default to null]
**AuthorizedTeams** | **[]int32** |  | [optional] [default to null]
**AuthorizedUsers** | **[]int32** | Deprecated in DBVersion &#x3D;&#x3D; 18 | [optional] [default to null]
**AzureCredentials** | [***PortainerAzureCredentials**](portainer.AzureCredentials.md) |  | [optional] [default to null]
**ComposeSyntaxMaxVersion** | **string** | Maximum version of docker-compose | [optional] [default to null]
**EdgeCheckinInterval** | **int32** | The check in interval for edge agent (in seconds) | [optional] [default to null]
**EdgeID** | **string** | The identifier of the edge agent associated with this environment(endpoint) | [optional] [default to null]
**EdgeKey** | **string** | The key which is used to map the agent to Portainer | [optional] [default to null]
**Gpus** | [**[]PortainerPair**](portainer.Pair.md) |  | [optional] [default to null]
**GroupId** | **int32** | Environment(Endpoint) group identifier | [optional] [default to null]
**Id** | **int32** | Environment(Endpoint) Identifier | [optional] [default to null]
**Kubernetes** | [***PortainerKubernetesData**](portainer.KubernetesData.md) | Associated Kubernetes data | [optional] [default to null]
**Name** | **string** | Environment(Endpoint) name | [optional] [default to null]
**PostInitMigrations** | [***PortainerEndpointPostInitMigrations**](portainer.EndpointPostInitMigrations.md) | Whether we need to run any \&quot;post init migrations\&quot;. | [optional] [default to null]
**PublicURL** | **string** | URL or IP address where exposed containers will be reachable | [optional] [default to null]
**Snapshots** | [**[]PortainerDockerSnapshot**](portainer.DockerSnapshot.md) | List of snapshots | [optional] [default to null]
**Status** | **int32** | The status of the environment(endpoint) (1 - up, 2 - down) | [optional] [default to null]
**TLS** | **bool** | Deprecated fields Deprecated in DBVersion &#x3D;&#x3D; 4 | [optional] [default to null]
**TLSCACert** | **string** |  | [optional] [default to null]
**TLSCert** | **string** |  | [optional] [default to null]
**TLSConfig** | [***PortainerTlsConfiguration**](portainer.TLSConfiguration.md) |  | [optional] [default to null]
**TLSKey** | **string** |  | [optional] [default to null]
**TagIds** | **[]int32** | List of tag identifiers to which this environment(endpoint) is associated | [optional] [default to null]
**Tags** | **[]string** | Deprecated in DBVersion &#x3D;&#x3D; 22 | [optional] [default to null]
**TeamAccessPolicies** | [***PortainerTeamAccessPolicies**](portainer.TeamAccessPolicies.md) | List of team identifiers authorized to connect to this environment(endpoint) | [optional] [default to null]
**Type_** | **int32** | Environment(Endpoint) environment(endpoint) type. 1 for a Docker environment(endpoint), 2 for an agent on Docker environment(endpoint) or 3 for an Azure environment(endpoint). | [optional] [default to null]
**URL** | **string** | URL or IP address of the Docker host associated to this environment(endpoint) | [optional] [default to null]
**UserAccessPolicies** | [***PortainerUserAccessPolicies**](portainer.UserAccessPolicies.md) | List of user identifiers authorized to connect to this environment(endpoint) | [optional] [default to null]
**Agent** | [***PortainerEndpointAgent**](portainer.Endpoint_agent.md) |  | [optional] [default to null]
**Edge** | [***PortainerEndpointEdge**](portainer.Endpoint_edge.md) |  | [optional] [default to null]
**IsEdgeDevice** | **bool** | IsEdgeDevice marks if the environment was created as an EdgeDevice | [optional] [default to null]
**LastCheckInDate** | **int32** | LastCheckInDate mark last check-in date on checkin | [optional] [default to null]
**QueryDate** | **int32** | QueryDate of each query with the endpoints list | [optional] [default to null]
**SecuritySettings** | [***PortainerEndpointSecuritySettings**](portainer.EndpointSecuritySettings.md) | Environment(Endpoint) specific security settings | [optional] [default to null]
**UserTrusted** | **bool** | Whether the device has been trusted or not by the user | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


