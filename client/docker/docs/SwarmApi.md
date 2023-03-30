# \SwarmApi

All URIs are relative to *http://localhost/v1.42*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SwarmInit**](SwarmApi.md#SwarmInit) | **Post** /swarm/init | Initialize a new swarm
[**SwarmInspect**](SwarmApi.md#SwarmInspect) | **Get** /swarm | Inspect swarm
[**SwarmJoin**](SwarmApi.md#SwarmJoin) | **Post** /swarm/join | Join an existing swarm
[**SwarmLeave**](SwarmApi.md#SwarmLeave) | **Post** /swarm/leave | Leave a swarm
[**SwarmUnlock**](SwarmApi.md#SwarmUnlock) | **Post** /swarm/unlock | Unlock a locked manager
[**SwarmUnlockkey**](SwarmApi.md#SwarmUnlockkey) | **Get** /swarm/unlockkey | Get the unlock key
[**SwarmUpdate**](SwarmApi.md#SwarmUpdate) | **Post** /swarm/update | Update a swarm


# **SwarmInit**
> string SwarmInit(ctx, body)
Initialize a new swarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwarmInitRequest**](SwarmInitRequest.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwarmInspect**
> Swarm SwarmInspect(ctx, )
Inspect swarm

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Swarm**](Swarm.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwarmJoin**
> SwarmJoin(ctx, body)
Join an existing swarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwarmJoinRequest**](SwarmJoinRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwarmLeave**
> SwarmLeave(ctx, optional)
Leave a swarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SwarmApiSwarmLeaveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwarmApiSwarmLeaveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **force** | **optional.Bool**| Force leave swarm, even if this is the last manager or that it will break the cluster.  | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwarmUnlock**
> SwarmUnlock(ctx, body)
Unlock a locked manager

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwarmUnlockRequest**](SwarmUnlockRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwarmUnlockkey**
> UnlockKeyResponse SwarmUnlockkey(ctx, )
Get the unlock key

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UnlockKeyResponse**](UnlockKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SwarmUpdate**
> SwarmUpdate(ctx, body, version, optional)
Update a swarm

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SwarmSpec**](SwarmSpec.md)|  | 
  **version** | **int64**| The version number of the swarm object being updated. This is required to avoid conflicting writes.  | 
 **optional** | ***SwarmApiSwarmUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SwarmApiSwarmUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rotateWorkerToken** | **optional.Bool**| Rotate the worker join token. | [default to false]
 **rotateManagerToken** | **optional.Bool**| Rotate the manager join token. | [default to false]
 **rotateManagerUnlockKey** | **optional.Bool**| Rotate the manager unlock key. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

