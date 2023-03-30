# HealthConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Test** | **[]string** | The test to perform. Possible values are:  - &#x60;[]&#x60; inherit healthcheck from image or parent image - &#x60;[\&quot;NONE\&quot;]&#x60; disable healthcheck - &#x60;[\&quot;CMD\&quot;, args...]&#x60; exec arguments directly - &#x60;[\&quot;CMD-SHELL\&quot;, command]&#x60; run command with system&#39;s default shell  | [optional] [default to null]
**Interval** | **int64** | The time to wait between checks in nanoseconds. It should be 0 or at least 1000000 (1 ms). 0 means inherit.  | [optional] [default to null]
**Timeout** | **int64** | The time to wait before considering the check to have hung. It should be 0 or at least 1000000 (1 ms). 0 means inherit.  | [optional] [default to null]
**Retries** | **int32** | The number of consecutive failures needed to consider a container as unhealthy. 0 means inherit.  | [optional] [default to null]
**StartPeriod** | **int64** | Start period for the container to initialize before starting health-retries countdown in nanoseconds. It should be 0 or at least 1000000 (1 ms). 0 means inherit.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


