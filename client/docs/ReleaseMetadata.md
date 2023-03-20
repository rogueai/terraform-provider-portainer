# ReleaseMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotations** | **map[string]string** | Annotations are additional mappings uninterpreted by Helm, made available for inspection by other applications. | [optional] [default to null]
**ApiVersion** | **string** | The API Version of this chart. Required. | [optional] [default to null]
**AppVersion** | **string** | The version of the application enclosed inside of this chart. | [optional] [default to null]
**Condition** | **string** | The condition to check to enable chart | [optional] [default to null]
**Dependencies** | [**[]ReleaseDependency**](release.Dependency.md) | Dependencies are a list of dependencies for a chart. | [optional] [default to null]
**Deprecated** | **bool** | Whether or not this chart is deprecated | [optional] [default to null]
**Description** | **string** | A one-sentence description of the chart | [optional] [default to null]
**Home** | **string** | The URL to a relevant project page, git repo, or contact person | [optional] [default to null]
**Icon** | **string** | The URL to an icon file. | [optional] [default to null]
**Keywords** | **[]string** | A list of string keywords | [optional] [default to null]
**KubeVersion** | **string** | KubeVersion is a SemVer constraint specifying the version of Kubernetes required. | [optional] [default to null]
**Maintainers** | [**[]ReleaseMaintainer**](release.Maintainer.md) | A list of name and URL/email address combinations for the maintainer(s) | [optional] [default to null]
**Name** | **string** | The name of the chart. Required. | [optional] [default to null]
**Sources** | **[]string** | Source is the URL to the source code of this chart | [optional] [default to null]
**Tags** | **string** | The tags to check to enable chart | [optional] [default to null]
**Type_** | **string** | Specifies the chart type: application or library | [optional] [default to null]
**Version** | **string** | A SemVer 2 conformant version string of the chart. Required. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


