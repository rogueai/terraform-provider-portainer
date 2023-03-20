# \EdgeApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointsIdEdgeJobsJobIDLogsPost**](EdgeApi.md#EndpointsIdEdgeJobsJobIDLogsPost) | **Post** /endpoints/{id}/edge/jobs/{jobID}/logs | Inspect an EdgeJob Log
[**EndpointsIdEdgeStacksStackIdGet**](EdgeApi.md#EndpointsIdEdgeStacksStackIdGet) | **Get** /endpoints/{id}/edge/stacks/{stackId} | Inspect an Edge Stack for an Environment(Endpoint)


# **EndpointsIdEdgeJobsJobIDLogsPost**
> EndpointsIdEdgeJobsJobIDLogsPost(ctx, id, jobID)
Inspect an EdgeJob Log

**Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| environment(endpoint) Id | 
  **jobID** | **string**| Job Id | 

### Return type

 (empty response body)

### Authorization

No authorization required

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

