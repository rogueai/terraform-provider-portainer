# Mount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** | Container path. | [optional] [default to null]
**Source** | **string** | Mount source (e.g. a volume name, a host path). | [optional] [default to null]
**Type_** | **string** | The mount type. Available types:  - &#x60;bind&#x60; Mounts a file or directory from the host into the container. Must exist prior to creating the container. - &#x60;volume&#x60; Creates a volume with the given name and options (or uses a pre-existing volume with the same name and options). These are **not** removed when the container is removed. - &#x60;tmpfs&#x60; Create a tmpfs with the given options. The mount source cannot be specified for tmpfs. - &#x60;npipe&#x60; Mounts a named pipe from the host into the container. Must exist prior to creating the container. - &#x60;cluster&#x60; a Swarm cluster volume  | [optional] [default to null]
**ReadOnly** | **bool** | Whether the mount should be read-only. | [optional] [default to null]
**Consistency** | **string** | The consistency requirement for the mount: &#x60;default&#x60;, &#x60;consistent&#x60;, &#x60;cached&#x60;, or &#x60;delegated&#x60;. | [optional] [default to null]
**BindOptions** | [***MountBindOptions**](Mount_BindOptions.md) |  | [optional] [default to null]
**VolumeOptions** | [***MountVolumeOptions**](Mount_VolumeOptions.md) |  | [optional] [default to null]
**TmpfsOptions** | [***MountTmpfsOptions**](Mount_TmpfsOptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


