# ServiceSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the service. | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]
**TaskTemplate** | [***TaskSpec**](TaskSpec.md) |  | [optional] [default to null]
**Mode** | [***ServiceSpecMode**](ServiceSpec_Mode.md) |  | [optional] [default to null]
**UpdateConfig** | [***ServiceSpecUpdateConfig**](ServiceSpec_UpdateConfig.md) |  | [optional] [default to null]
**RollbackConfig** | [***ServiceSpecRollbackConfig**](ServiceSpec_RollbackConfig.md) |  | [optional] [default to null]
**Networks** | [**[]NetworkAttachmentConfig**](NetworkAttachmentConfig.md) | Specifies which networks the service should attach to. | [optional] [default to null]
**EndpointSpec** | [***EndpointSpec**](EndpointSpec.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


