# CustomtemplatesCustomTemplateFromGitRepositoryPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComposeFilePathInRepository** | **string** | Path to the Stack file inside the Git repository | [optional] [default to null]
**Description** | **string** | Description of the template | [default to null]
**Logo** | **string** | URL of the template&#39;s logo | [optional] [default to null]
**Note** | **string** | A note that will be displayed in the UI. Supports HTML content | [optional] [default to null]
**Platform** | **int32** | Platform associated to the template. Valid values are: 1 - &#39;linux&#39;, 2 - &#39;windows&#39; Required for Docker stacks | [optional] [default to null]
**RepositoryAuthentication** | **bool** | Use basic authentication to clone the Git repository | [optional] [default to null]
**RepositoryPassword** | **string** | Password used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**RepositoryReferenceName** | **string** | Reference name of a Git repository hosting the Stack file | [optional] [default to null]
**RepositoryURL** | **string** | URL of a Git repository hosting the Stack file | [default to null]
**RepositoryUsername** | **string** | Username used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**Title** | **string** | Title of the template | [default to null]
**Type_** | **int32** | Type of created stack (1 - swarm, 2 - compose) | [default to null]
**Variables** | [**[]PortainerCustomTemplateVariableDefinition**](portainer.CustomTemplateVariableDefinition.md) | Definitions of variables in the stack file | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


