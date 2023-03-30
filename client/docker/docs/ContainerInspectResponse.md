# ContainerInspectResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the container | [optional] [default to null]
**Created** | **string** | The time the container was created | [optional] [default to null]
**Path** | **string** | The path to the command being run | [optional] [default to null]
**Args** | **[]string** | The arguments to the command being run | [optional] [default to null]
**State** | [***ContainerState**](ContainerState.md) |  | [optional] [default to null]
**Image** | **string** | The container&#39;s image ID | [optional] [default to null]
**ResolvConfPath** | **string** |  | [optional] [default to null]
**HostnamePath** | **string** |  | [optional] [default to null]
**HostsPath** | **string** |  | [optional] [default to null]
**LogPath** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**RestartCount** | **int32** |  | [optional] [default to null]
**Driver** | **string** |  | [optional] [default to null]
**Platform** | **string** |  | [optional] [default to null]
**MountLabel** | **string** |  | [optional] [default to null]
**ProcessLabel** | **string** |  | [optional] [default to null]
**AppArmorProfile** | **string** |  | [optional] [default to null]
**ExecIDs** | **[]string** | IDs of exec instances that are running in the container. | [optional] [default to null]
**HostConfig** | [***HostConfig**](HostConfig.md) |  | [optional] [default to null]
**GraphDriver** | [***GraphDriverData**](GraphDriverData.md) |  | [optional] [default to null]
**SizeRw** | **int64** | The size of files that have been created or changed by this container.  | [optional] [default to null]
**SizeRootFs** | **int64** | The total size of all the files in this container. | [optional] [default to null]
**Mounts** | [**[]MountPoint**](MountPoint.md) |  | [optional] [default to null]
**Config** | [***ContainerConfig**](ContainerConfig.md) |  | [optional] [default to null]
**NetworkSettings** | [***NetworkSettings**](NetworkSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


