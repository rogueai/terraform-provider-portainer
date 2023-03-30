# EdgestacksSwarmStackFromGitRepositoryPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentType** | **int32** | Deployment type to deploy this stack Valid values are: 0 - &#39;compose&#39;, 1 - &#39;kubernetes&#39;, 2 - &#39;nomad&#39; for compose stacks will use kompose to convert to kubernetes manifest for kubernetes environments(endpoints) kubernetes deploy type is enabled only for kubernetes environments(endpoints) nomad deploy type is enabled only for nomad environments(endpoints) | [optional] [default to null]
**EdgeGroups** | **[]int32** | List of identifiers of EdgeGroups | [optional] [default to null]
**FilePathInRepository** | **string** | Path to the Stack file inside the Git repository | [optional] [default to null]
**Name** | **string** | Name of the stack | [default to null]
**Registries** | **[]int32** | List of Registries to use for this stack | [optional] [default to null]
**RepositoryAuthentication** | **bool** | Use basic authentication to clone the Git repository | [optional] [default to null]
**RepositoryPassword** | **string** | Password used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**RepositoryReferenceName** | **string** | Reference name of a Git repository hosting the Stack file | [optional] [default to null]
**RepositoryURL** | **string** | URL of a Git repository hosting the Stack file | [default to null]
**RepositoryUsername** | **string** | Username used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**UseManifestNamespaces** | **bool** | Uses the manifest&#39;s namespaces instead of the default one | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


