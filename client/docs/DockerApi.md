# \DockerApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DockerContainerGpusInspect**](DockerApi.md#DockerContainerGpusInspect) | **Get** /docker/{environmentId}/containers/{containerId}/gpus | Fetch container gpus data


# **DockerContainerGpusInspect**
> ContainersContainerGpusResponse DockerContainerGpusInspect(ctx, environmentId, containerId)
Fetch container gpus data

**Access policy**:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **environmentId** | **int32**| Environment identifier | 
  **containerId** | **int32**| Container identifier | 

### Return type

[**ContainersContainerGpusResponse**](containers.containerGpusResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

