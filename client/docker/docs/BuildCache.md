# BuildCache

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **string** | Unique ID of the build cache record.  | [optional] [default to null]
**Parent** | **string** | ID of the parent build cache record.  &gt; **Deprecated**: This field is deprecated, and omitted if empty.  | [optional] [default to null]
**Parents** | **[]string** | List of parent build cache record IDs.  | [optional] [default to null]
**Type_** | **string** | Cache record type.  | [optional] [default to null]
**Description** | **string** | Description of the build-step that produced the build cache.  | [optional] [default to null]
**InUse** | **bool** | Indicates if the build cache is in use.  | [optional] [default to null]
**Shared** | **bool** | Indicates if the build cache is shared.  | [optional] [default to null]
**Size** | **int32** | Amount of disk space used by the build cache (in bytes).  | [optional] [default to null]
**CreatedAt** | **string** | Date and time at which the build cache was created in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.  | [optional] [default to null]
**LastUsedAt** | **string** | Date and time at which the build cache was last used in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.  | [optional] [default to null]
**UsageCount** | **int32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


