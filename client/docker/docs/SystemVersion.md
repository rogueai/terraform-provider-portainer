# SystemVersion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | [***SystemVersionPlatform**](SystemVersion_Platform.md) |  | [optional] [default to null]
**Components** | [**[]SystemVersionComponents**](SystemVersion_Components.md) | Information about system components  | [optional] [default to null]
**Version** | **string** | The version of the daemon | [optional] [default to null]
**ApiVersion** | **string** | The default (and highest) API version that is supported by the daemon  | [optional] [default to null]
**MinAPIVersion** | **string** | The minimum API version that is supported by the daemon  | [optional] [default to null]
**GitCommit** | **string** | The Git commit of the source code that was used to build the daemon  | [optional] [default to null]
**GoVersion** | **string** | The version Go used to compile the daemon, and the version of the Go runtime in use.  | [optional] [default to null]
**Os** | **string** | The operating system that the daemon is running on (\&quot;linux\&quot; or \&quot;windows\&quot;)  | [optional] [default to null]
**Arch** | **string** | The architecture that the daemon is running on  | [optional] [default to null]
**KernelVersion** | **string** | The kernel version (&#x60;uname -r&#x60;) that the daemon is running on.  This field is omitted when empty.  | [optional] [default to null]
**Experimental** | **bool** | Indicates if the daemon is started with experimental features enabled.  This field is omitted when empty / false.  | [optional] [default to null]
**BuildTime** | **string** | The date and time that the daemon was compiled.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


