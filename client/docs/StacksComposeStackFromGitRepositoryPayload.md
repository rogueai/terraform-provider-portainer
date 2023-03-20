# StacksComposeStackFromGitRepositoryPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFiles** | **[]string** | Applicable when deploying with multiple stack files | [optional] [default to null]
**AutoUpdate** | [***PortainerStackAutoUpdate**](portainer.StackAutoUpdate.md) | Optional auto update configuration | [optional] [default to null]
**ComposeFile** | **string** | Path to the Stack file inside the Git repository | [optional] [default to null]
**Env** | [**[]PortainerPair**](portainer.Pair.md) | A list of environment(endpoint) variables used during stack deployment | [optional] [default to null]
**FromAppTemplate** | **bool** | Whether the stack is from a app template | [optional] [default to null]
**Name** | **string** | Name of the stack | [default to null]
**RepositoryAuthentication** | **bool** | Use basic authentication to clone the Git repository | [optional] [default to null]
**RepositoryPassword** | **string** | Password used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**RepositoryReferenceName** | **string** | Reference name of a Git repository hosting the Stack file | [optional] [default to null]
**RepositoryURL** | **string** | URL of a Git repository hosting the Stack file | [default to null]
**RepositoryUsername** | **string** | Username used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


