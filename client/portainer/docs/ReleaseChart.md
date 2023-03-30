# ReleaseChart

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Files** | [**[]ReleaseFile**](release.File.md) | Files are miscellaneous files in a chart archive, e.g. README, LICENSE, etc. | [optional] [default to null]
**Lock** | [***ReleaseLock**](release.Lock.md) | Lock is the contents of Chart.lock. | [optional] [default to null]
**Metadata** | [***ReleaseMetadata**](release.Metadata.md) | Metadata is the contents of the Chartfile. | [optional] [default to null]
**Schema** | **[]int32** | Schema is an optional JSON schema for imposing structure on Values | [optional] [default to null]
**Templates** | [**[]ReleaseFile**](release.File.md) | Templates for this chart. | [optional] [default to null]
**Values** | [**map[string]ErrorUnknown**](.md) | Values are default config for this chart. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


