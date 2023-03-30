# VolumeCreateOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The new volume&#39;s name. If not specified, Docker generates a name.  | [optional] [default to null]
**Driver** | **string** | Name of the volume driver to use. | [optional] [default to null]
**DriverOpts** | **map[string]string** | A mapping of driver options and values. These options are passed directly to the driver and are driver specific.  | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]
**ClusterVolumeSpec** | [***ClusterVolumeSpec**](ClusterVolumeSpec.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


