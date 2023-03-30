# PortainerEdgeStack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | **int32** |  | [optional] [default to null]
**EdgeGroups** | **[]int32** |  | [optional] [default to null]
**EntryPoint** | **string** |  | [optional] [default to null]
**Id** | **int32** | EdgeStack Identifier | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NumDeployments** | **int32** |  | [optional] [default to null]
**ProjectPath** | **string** |  | [optional] [default to null]
**Prune** | **bool** | Deprecated | [optional] [default to null]
**Status** | [**map[string]PortainerEdgeStackStatus**](portainer.EdgeStackStatus.md) |  | [optional] [default to null]
**Version** | **int32** |  | [optional] [default to null]
**DeploymentType** | **int32** | Deployment type to deploy this stack Valid values are: 0 - &#39;compose&#39;, 1 - &#39;kubernetes&#39;, 2 - &#39;nomad&#39; for compose stacks will use kompose to convert to kubernetes manifest for kubernetes environments(endpoints) kubernetes deploy type is enabled only for kubernetes environments(endpoints) nomad deploy type is enabled only for nomad environments(endpoints) | [optional] [default to null]
**ManifestPath** | **string** |  | [optional] [default to null]
**UseManifestNamespaces** | **bool** | Uses the manifest&#39;s namespaces instead of the default one | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


