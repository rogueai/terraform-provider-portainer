# TaskSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PluginSpec** | [***TaskSpecPluginSpec**](TaskSpec_PluginSpec.md) |  | [optional] [default to null]
**ContainerSpec** | [***TaskSpecContainerSpec**](TaskSpec_ContainerSpec.md) |  | [optional] [default to null]
**NetworkAttachmentSpec** | [***TaskSpecNetworkAttachmentSpec**](TaskSpec_NetworkAttachmentSpec.md) |  | [optional] [default to null]
**Resources** | [***TaskSpecResources**](TaskSpec_Resources.md) |  | [optional] [default to null]
**RestartPolicy** | [***TaskSpecRestartPolicy**](TaskSpec_RestartPolicy.md) |  | [optional] [default to null]
**Placement** | [***TaskSpecPlacement**](TaskSpec_Placement.md) |  | [optional] [default to null]
**ForceUpdate** | **int32** | A counter that triggers an update even if no relevant parameters have been changed.  | [optional] [default to null]
**Runtime** | **string** | Runtime is the type of runtime specified for the task executor.  | [optional] [default to null]
**Networks** | [**[]NetworkAttachmentConfig**](NetworkAttachmentConfig.md) | Specifies which networks the service should attach to. | [optional] [default to null]
**LogDriver** | [***TaskSpecLogDriver**](TaskSpec_LogDriver.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


