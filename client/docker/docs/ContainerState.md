# ContainerState

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | String representation of the container state. Can be one of \&quot;created\&quot;, \&quot;running\&quot;, \&quot;paused\&quot;, \&quot;restarting\&quot;, \&quot;removing\&quot;, \&quot;exited\&quot;, or \&quot;dead\&quot;.  | [optional] [default to null]
**Running** | **bool** | Whether this container is running.  Note that a running container can be _paused_. The &#x60;Running&#x60; and &#x60;Paused&#x60; booleans are not mutually exclusive:  When pausing a container (on Linux), the freezer cgroup is used to suspend all processes in the container. Freezing the process requires the process to be running. As a result, paused containers are both &#x60;Running&#x60; _and_ &#x60;Paused&#x60;.  Use the &#x60;Status&#x60; field instead to determine if a container&#39;s state is \&quot;running\&quot;.  | [optional] [default to null]
**Paused** | **bool** | Whether this container is paused. | [optional] [default to null]
**Restarting** | **bool** | Whether this container is restarting. | [optional] [default to null]
**OOMKilled** | **bool** | Whether this container has been killed because it ran out of memory.  | [optional] [default to null]
**Dead** | **bool** |  | [optional] [default to null]
**Pid** | **int32** | The process ID of this container | [optional] [default to null]
**ExitCode** | **int32** | The last exit code of this container | [optional] [default to null]
**Error_** | **string** |  | [optional] [default to null]
**StartedAt** | **string** | The time when this container was last started. | [optional] [default to null]
**FinishedAt** | **string** | The time when this container last exited. | [optional] [default to null]
**Health** | [***Health**](Health.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


