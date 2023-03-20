# PortainerCustomTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedByUserId** | **int32** | User identifier who created this template | [optional] [default to null]
**Description** | **string** | Description of the template | [optional] [default to null]
**EntryPoint** | **string** | Path to the Stack file | [optional] [default to null]
**Id** | **int32** | CustomTemplate Identifier | [optional] [default to null]
**Logo** | **string** | URL of the template&#39;s logo | [optional] [default to null]
**Note** | **string** | A note that will be displayed in the UI. Supports HTML content | [optional] [default to null]
**Platform** | **int32** | Platform associated to the template. Valid values are: 1 - &#39;linux&#39;, 2 - &#39;windows&#39; | [optional] [default to null]
**ProjectPath** | **string** | Path on disk to the repository hosting the Stack file | [optional] [default to null]
**ResourceControl** | [***PortainerResourceControl**](portainer.ResourceControl.md) |  | [optional] [default to null]
**Title** | **string** | Title of the template | [optional] [default to null]
**Type_** | **int32** | Type of created stack (1 - swarm, 2 - compose) | [optional] [default to null]
**Variables** | [**[]PortainerCustomTemplateVariableDefinition**](portainer.CustomTemplateVariableDefinition.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


