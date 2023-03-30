# TaskSpecRestartPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **string** | Condition for restart. | [optional] [default to null]
**Delay** | **int64** | Delay between restart attempts. | [optional] [default to null]
**MaxAttempts** | **int64** | Maximum attempts to restart a given container before giving up (default value is 0, which is ignored).  | [optional] [default to 0]
**Window** | **int64** | Windows is the time window used to evaluate the restart policy (default value is 0, which is unbounded).  | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


