# ServiceSpecUpdateConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parallelism** | **int64** | Maximum number of tasks to be updated in one iteration (0 means unlimited parallelism).  | [optional] [default to null]
**Delay** | **int64** | Amount of time between updates, in nanoseconds. | [optional] [default to null]
**FailureAction** | **string** | Action to take if an updated task fails to run, or stops running during the update.  | [optional] [default to null]
**Monitor** | **int64** | Amount of time to monitor each updated task for failures, in nanoseconds.  | [optional] [default to null]
**MaxFailureRatio** | **float32** | The fraction of tasks that may fail during an update before the failure action is invoked, specified as a floating point number between 0 and 1.  | [optional] [default to null]
**Order** | **string** | The order of operations when rolling out an updated task. Either the old task is shut down before the new task is started, or the new task is started before the old task is shut down.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


