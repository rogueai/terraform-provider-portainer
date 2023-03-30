# EndpointSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **string** | The mode of resolution to use for internal load balancing between tasks.  | [optional] [default to null]
**Ports** | [**[]EndpointPortConfig**](EndpointPortConfig.md) | List of exposed ports that this service is accessible on from the outside. Ports can only be provided if &#x60;vip&#x60; resolution mode is used.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


