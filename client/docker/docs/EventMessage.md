# EventMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The type of object emitting the event | [optional] [default to null]
**Action** | **string** | The type of event | [optional] [default to null]
**Actor** | [***EventActor**](EventActor.md) |  | [optional] [default to null]
**Scope** | **string** | Scope of the event. Engine events are &#x60;local&#x60; scope. Cluster (Swarm) events are &#x60;swarm&#x60; scope.  | [optional] [default to null]
**Time** | **int64** | Timestamp of event | [optional] [default to null]
**TimeNano** | **int64** | Timestamp of event, with nanosecond accuracy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


