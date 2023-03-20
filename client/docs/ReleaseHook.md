# ReleaseHook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletePolicies** | **[]string** | DeletePolicies are the policies that indicate when to delete the hook | [optional] [default to null]
**Events** | **[]string** | Events are the events that this hook fires on. | [optional] [default to null]
**Kind** | **string** | Kind is the Kubernetes kind. | [optional] [default to null]
**LastRun** | [***ReleaseHookExecution**](release.HookExecution.md) | LastRun indicates the date/time this was last run. | [optional] [default to null]
**Manifest** | **string** | Manifest is the manifest contents. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Path** | **string** | Path is the chart-relative path to the template. | [optional] [default to null]
**Weight** | **int32** | Weight indicates the sort order for execution among similar Hook type | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


