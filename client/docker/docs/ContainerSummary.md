# ContainerSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of this container | [optional] [default to null]
**Names** | **[]string** | The names that this container has been given | [optional] [default to null]
**Image** | **string** | The name of the image used when creating this container | [optional] [default to null]
**ImageID** | **string** | The ID of the image that this container was created from | [optional] [default to null]
**Command** | **string** | Command to run when starting the container | [optional] [default to null]
**Created** | **int64** | When the container was created | [optional] [default to null]
**Ports** | [**[]Port**](Port.md) | The ports exposed by this container | [optional] [default to null]
**SizeRw** | **int64** | The size of files that have been created or changed by this container | [optional] [default to null]
**SizeRootFs** | **int64** | The total size of all the files in this container | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]
**State** | **string** | The state of this container (e.g. &#x60;Exited&#x60;) | [optional] [default to null]
**Status** | **string** | Additional human-readable status of this container (e.g. &#x60;Exit 0&#x60;) | [optional] [default to null]
**HostConfig** | [***ContainerSummaryHostConfig**](ContainerSummary_HostConfig.md) |  | [optional] [default to null]
**NetworkSettings** | [***ContainerSummaryNetworkSettings**](ContainerSummary_NetworkSettings.md) |  | [optional] [default to null]
**Mounts** | [**[]MountPoint**](MountPoint.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


