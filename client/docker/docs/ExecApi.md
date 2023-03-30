# \ExecApi

All URIs are relative to *http://localhost/v1.42*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerExec**](ExecApi.md#ContainerExec) | **Post** /containers/{id}/exec | Create an exec instance
[**ExecInspect**](ExecApi.md#ExecInspect) | **Get** /exec/{id}/json | Inspect an exec instance
[**ExecResize**](ExecApi.md#ExecResize) | **Post** /exec/{id}/resize | Resize an exec instance
[**ExecStart**](ExecApi.md#ExecStart) | **Post** /exec/{id}/start | Start an exec instance


# **ContainerExec**
> IdResponse ContainerExec(ctx, execConfig, id)
Create an exec instance

Run a command inside a running container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **execConfig** | [**ExecConfig**](ExecConfig.md)| Exec configuration | 
  **id** | **string**| ID or name of container | 

### Return type

[**IdResponse**](IdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecInspect**
> ExecInspectResponse ExecInspect(ctx, id)
Inspect an exec instance

Return low-level information about an exec instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exec instance ID | 

### Return type

[**ExecInspectResponse**](ExecInspectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecResize**
> ExecResize(ctx, id, optional)
Resize an exec instance

Resize the TTY session used by an exec instance. This endpoint only works if `tty` was specified as part of creating and starting the exec instance. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exec instance ID | 
 **optional** | ***ExecApiExecResizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecApiExecResizeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **h** | **optional.Int32**| Height of the TTY session in characters | 
 **w** | **optional.Int32**| Width of the TTY session in characters | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecStart**
> ExecStart(ctx, id, optional)
Start an exec instance

Starts a previously set up exec instance. If detach is true, this endpoint returns immediately after starting the command. Otherwise, it sets up an interactive session with the command. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Exec instance ID | 
 **optional** | ***ExecApiExecStartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecApiExecStartOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execStartConfig** | [**optional.Interface of ExecStartConfig**](ExecStartConfig.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.docker.raw-stream, application/vnd.docker.multiplexed-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

