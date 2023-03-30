# \WebsocketApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebsocketAttachGet**](WebsocketApi.md#WebsocketAttachGet) | **Get** /websocket/attach | Attach a websocket
[**WebsocketExecGet**](WebsocketApi.md#WebsocketExecGet) | **Get** /websocket/exec | Execute a websocket
[**WebsocketKubernetesShellGet**](WebsocketApi.md#WebsocketKubernetesShellGet) | **Get** /websocket/kubernetes-shell | Execute a websocket on kubectl shell pod
[**WebsocketPodGet**](WebsocketApi.md#WebsocketPodGet) | **Get** /websocket/pod | Execute a websocket on pod


# **WebsocketAttachGet**
> WebsocketAttachGet(ctx, endpointId, token, optional)
Attach a websocket

If the nodeName query parameter is present, the request will be proxied to the underlying agent environment(endpoint). If the nodeName query parameter is not specified, the request will be upgraded to the websocket protocol and an AttachStart operation HTTP request will be created and hijacked. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int32**| environment(endpoint) ID of the environment(endpoint) where the resource is located | 
  **token** | **string**| JWT token used for authentication against this environment(endpoint) | 
 **optional** | ***WebsocketApiWebsocketAttachGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebsocketApiWebsocketAttachGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeName** | **optional.String**| node name | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebsocketExecGet**
> WebsocketExecGet(ctx, endpointId, token, optional)
Execute a websocket

If the nodeName query parameter is present, the request will be proxied to the underlying agent environment(endpoint). If the nodeName query parameter is not specified, the request will be upgraded to the websocket protocol and an ExecStart operation HTTP request will be created and hijacked.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int32**| environment(endpoint) ID of the environment(endpoint) where the resource is located | 
  **token** | **string**| JWT token used for authentication against this environment(endpoint) | 
 **optional** | ***WebsocketApiWebsocketExecGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebsocketApiWebsocketExecGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeName** | **optional.String**| node name | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebsocketKubernetesShellGet**
> WebsocketKubernetesShellGet(ctx, endpointId, token)
Execute a websocket on kubectl shell pod

The request will be upgraded to the websocket protocol. The request will proxy input from the client to the pod via long-lived websocket connection. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int32**| environment(endpoint) ID of the environment(endpoint) where the resource is located | 
  **token** | **string**| JWT token used for authentication against this environment(endpoint) | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebsocketPodGet**
> WebsocketPodGet(ctx, endpointId, namespace, podName, containerName, command, token)
Execute a websocket on pod

The request will be upgraded to the websocket protocol. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpointId** | **int32**| environment(endpoint) ID of the environment(endpoint) where the resource is located | 
  **namespace** | **string**| namespace where the container is located | 
  **podName** | **string**| name of the pod containing the container | 
  **containerName** | **string**| name of the container | 
  **command** | **string**| command to execute in the container | 
  **token** | **string**| JWT token used for authentication against this environment(endpoint) | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

