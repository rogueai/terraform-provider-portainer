# EndpointedgeEndpointEdgeStatusInspectResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checkin** | **int32** | The current value of CheckinInterval | [optional] [default to null]
**Credentials** | **string** |  | [optional] [default to null]
**Port** | **int32** | The tunnel port | [optional] [default to null]
**Schedules** | [**[]EndpointedgeEdgeJobResponse**](endpointedge.edgeJobResponse.md) | List of requests for jobs to run on the environment(endpoint) | [optional] [default to null]
**Stacks** | [**[]EndpointedgeStackStatusResponse**](endpointedge.stackStatusResponse.md) | List of stacks to be deployed on the environments(endpoints) | [optional] [default to null]
**Status** | **string** | Status represents the environment(endpoint) status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


