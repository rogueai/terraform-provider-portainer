# ClusterVolume

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | The Swarm ID of this volume. Because cluster volumes are Swarm objects, they have an ID, unlike non-cluster volumes. This ID can be used to refer to the Volume instead of the name.  | [optional] [default to null]
**Version** | [***ObjectVersion**](ObjectVersion.md) |  | [optional] [default to null]
**CreatedAt** | **string** |  | [optional] [default to null]
**UpdatedAt** | **string** |  | [optional] [default to null]
**Spec** | [***ClusterVolumeSpec**](ClusterVolumeSpec.md) |  | [optional] [default to null]
**Info** | [***ClusterVolumeInfo**](ClusterVolume_Info.md) |  | [optional] [default to null]
**PublishStatus** | [**[]ClusterVolumePublishStatus**](ClusterVolume_PublishStatus.md) | The status of the volume as it pertains to its publishing and use on specific nodes  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


