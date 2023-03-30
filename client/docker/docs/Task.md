# Task

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The ID of the task. | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]
**CreatedAt** | **string** |  | [optional] [default to null]
**UpdatedAt** | **string** |  | [optional] [default to null]
**Name** | **string** | Name of the task. | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]
**Spec** | [***TaskSpec**](TaskSpec.md) |  | [optional] [default to null]
**ServiceID** | **string** | The ID of the service this task is part of. | [optional] [default to null]
**Slot** | **int32** |  | [optional] [default to null]
**NodeID** | **string** | The ID of the node that this task is on. | [optional] [default to null]
**AssignedGenericResources** | [***GenericResources**](GenericResources.md) |  | [optional] [default to null]
**Status** | [***TaskStatus**](Task_Status.md) |  | [optional] [default to null]
**DesiredState** | [***TaskState**](TaskState.md) |  | [optional] [default to null]
**JobIteration** | [***ObjectVersion**](ObjectVersion.md) | If the Service this Task belongs to is a job-mode service, contains the JobIteration of the Service this Task was created for. Absent if the Task was created for a Replicated or Global Service.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


