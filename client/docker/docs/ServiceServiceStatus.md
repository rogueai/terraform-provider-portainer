# ServiceServiceStatus

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunningTasks** | **int32** | The number of tasks for the service currently in the Running state.  | [optional] [default to null]
**DesiredTasks** | **int32** | The number of tasks for the service desired to be running. For replicated services, this is the replica count from the service spec. For global services, this is computed by taking count of all tasks for the service with a Desired State other than Shutdown.  | [optional] [default to null]
**CompletedTasks** | **int32** | The number of tasks for a job that are in the Completed state. This field must be cross-referenced with the service type, as the value of 0 may mean the service is not in a job mode, or it may mean the job-mode service has no tasks yet Completed.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


