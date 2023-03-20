# EdgestacksSwarmStackFromFileContentPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | **int32** | Deployment type to deploy this stack Valid values are: 0 - &#39;compose&#39;, 1 - &#39;kubernetes&#39;, 2 - &#39;nomad&#39; for compose stacks will use kompose to convert to kubernetes manifest for kubernetes environments(endpoints) kubernetes deploy type is enabled only for kubernetes environments(endpoints) nomad deploy type is enabled only for nomad environments(endpoints) | [optional] [default to null]
**EdgeGroups** | **[]int32** | List of identifiers of EdgeGroups | [optional] [default to null]
**Name** | **string** | Name of the stack | [default to null]
**Registries** | **[]int32** | List of Registries to use for this stack | [optional] [default to null]
**StackFileContent** | **string** | Content of the Stack file | [default to null]
**UseManifestNamespaces** | **bool** | Uses the manifest&#39;s namespaces instead of the default one | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


