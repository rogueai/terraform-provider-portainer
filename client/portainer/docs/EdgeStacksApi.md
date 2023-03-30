# \EdgeStacksApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeStackCreate**](EdgeStacksApi.md#EdgeStackCreate) | **Post** /edge_stacks | Create an EdgeStack
[**EdgeStackDelete**](EdgeStacksApi.md#EdgeStackDelete) | **Delete** /edge_stacks/{id} | Delete an EdgeStack
[**EdgeStackFile**](EdgeStacksApi.md#EdgeStackFile) | **Get** /edge_stacks/{id}/file | Fetches the stack file for an EdgeStack
[**EdgeStackInspect**](EdgeStacksApi.md#EdgeStackInspect) | **Get** /edge_stacks/{id} | Inspect an EdgeStack
[**EdgeStackList**](EdgeStacksApi.md#EdgeStackList) | **Get** /edge_stacks | Fetches the list of EdgeStacks
[**EdgeStackStatusDelete**](EdgeStacksApi.md#EdgeStackStatusDelete) | **Delete** /edge_stacks/{id}/status/{endpoint_id} | Delete an EdgeStack status
[**EdgeStackStatusUpdate**](EdgeStacksApi.md#EdgeStackStatusUpdate) | **Put** /edge_stacks/{id}/status | Update an EdgeStack status
[**EdgeStackUpdate**](EdgeStacksApi.md#EdgeStackUpdate) | **Put** /edge_stacks/{id} | Update an EdgeStack
[**EndpointsIdEdgeStacksStackIdGet**](EdgeStacksApi.md#EndpointsIdEdgeStacksStackIdGet) | **Get** /endpoints/{id}/edge/stacks/{stackId} | Inspect an Edge Stack for an Environment(Endpoint)


# **EdgeStackCreate**
> PortainerEdgeStack EdgeStackCreate(ctx, method, bodyString, bodyFile, bodyRepository)
Create an EdgeStack

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **method** | **string**| Creation Method | 
  **bodyString** | [**EdgestacksSwarmStackFromFileContentPayload**](EdgestacksSwarmStackFromFileContentPayload.md)| Required when using method&#x3D;string | 
  **bodyFile** | [**EdgestacksSwarmStackFromFileUploadPayload**](EdgestacksSwarmStackFromFileUploadPayload.md)| Required when using method&#x3D;file | 
  **bodyRepository** | [**EdgestacksSwarmStackFromGitRepositoryPayload**](EdgestacksSwarmStackFromGitRepositoryPayload.md)| Required when using method&#x3D;repository | 

### Return type

[**PortainerEdgeStack**](portainer.EdgeStack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackDelete**
> EdgeStackDelete(ctx, id)
Delete an EdgeStack

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeStack Id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackFile**
> EdgestacksStackFileResponse EdgeStackFile(ctx, id)
Fetches the stack file for an EdgeStack

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeStack Id | 

### Return type

[**EdgestacksStackFileResponse**](edgestacks.stackFileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackInspect**
> PortainerEdgeStack EdgeStackInspect(ctx, id)
Inspect an EdgeStack

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeStack Id | 

### Return type

[**PortainerEdgeStack**](portainer.EdgeStack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackList**
> []PortainerEdgeStack EdgeStackList(ctx, )
Fetches the list of EdgeStacks

**Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PortainerEdgeStack**](portainer.EdgeStack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackStatusDelete**
> PortainerEdgeStack EdgeStackStatusDelete(ctx, id)
Delete an EdgeStack status

Authorized only if the request is done by an Edge Environment(Endpoint)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeStack Id | 

### Return type

[**PortainerEdgeStack**](portainer.EdgeStack.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackStatusUpdate**
> PortainerEdgeStack EdgeStackStatusUpdate(ctx, id)
Update an EdgeStack status

Authorized only if the request is done by an Edge Environment(Endpoint)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeStack Id | 

### Return type

[**PortainerEdgeStack**](portainer.EdgeStack.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeStackUpdate**
> PortainerEdgeStack EdgeStackUpdate(ctx, id, body)
Update an EdgeStack

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeStack Id | 
  **body** | [**EdgestacksUpdateEdgeStackPayload**](EdgestacksUpdateEdgeStackPayload.md)| EdgeStack data | 

### Return type

[**PortainerEdgeStack**](portainer.EdgeStack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointsIdEdgeStacksStackIdGet**
> EndpointedgeConfigResponse EndpointsIdEdgeStacksStackIdGet(ctx, id, stackId)
Inspect an Edge Stack for an Environment(Endpoint)

**Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| environment(endpoint) Id | 
  **stackId** | **string**| EdgeStack Id | 

### Return type

[**EndpointedgeConfigResponse**](endpointedge.configResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

