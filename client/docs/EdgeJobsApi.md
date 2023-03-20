# \EdgeJobsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeJobCreate**](EdgeJobsApi.md#EdgeJobCreate) | **Post** /edge_jobs | Create an EdgeJob
[**EdgeJobDelete**](EdgeJobsApi.md#EdgeJobDelete) | **Delete** /edge_jobs/{id} | Delete an EdgeJob
[**EdgeJobFile**](EdgeJobsApi.md#EdgeJobFile) | **Get** /edge_jobs/{id}/file | Fetch a file of an EdgeJob
[**EdgeJobInspect**](EdgeJobsApi.md#EdgeJobInspect) | **Get** /edge_jobs/{id} | Inspect an EdgeJob
[**EdgeJobList**](EdgeJobsApi.md#EdgeJobList) | **Get** /edge_jobs | Fetch EdgeJobs list
[**EdgeJobTaskLogsInspect**](EdgeJobsApi.md#EdgeJobTaskLogsInspect) | **Get** /edge_jobs/{id}/tasks/{taskID}/logs | Fetch the log for a specifc task on an EdgeJob
[**EdgeJobTasksClear**](EdgeJobsApi.md#EdgeJobTasksClear) | **Delete** /edge_jobs/{id}/tasks/{taskID}/logs | Clear the log for a specifc task on an EdgeJob
[**EdgeJobTasksCollect**](EdgeJobsApi.md#EdgeJobTasksCollect) | **Post** /edge_jobs/{id}/tasks/{taskID}/logs | Collect the log for a specifc task on an EdgeJob
[**EdgeJobTasksList**](EdgeJobsApi.md#EdgeJobTasksList) | **Get** /edge_jobs/{id}/tasks | Fetch the list of tasks on an EdgeJob
[**EdgeJobUpdate**](EdgeJobsApi.md#EdgeJobUpdate) | **Post** /edge_jobs/{id} | Update an EdgeJob


# **EdgeJobCreate**
> PortainerEdgeGroup EdgeJobCreate(ctx, method, bodyString, bodyFile)
Create an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **method** | **string**| Creation Method | 
  **bodyString** | [**EdgejobsEdgeJobCreateFromFileContentPayload**](EdgejobsEdgeJobCreateFromFileContentPayload.md)| EdgeGroup data when method is string | 
  **bodyFile** | [**EdgejobsEdgeJobCreateFromFilePayload**](EdgejobsEdgeJobCreateFromFilePayload.md)| EdgeGroup data when method is file | 

### Return type

[**PortainerEdgeGroup**](portainer.EdgeGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobDelete**
> EdgeJobDelete(ctx, id)
Delete an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobFile**
> EdgejobsEdgeJobFileResponse EdgeJobFile(ctx, id)
Fetch a file of an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 

### Return type

[**EdgejobsEdgeJobFileResponse**](edgejobs.edgeJobFileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobInspect**
> PortainerEdgeJob EdgeJobInspect(ctx, id)
Inspect an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 

### Return type

[**PortainerEdgeJob**](portainer.EdgeJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobList**
> []PortainerEdgeJob EdgeJobList(ctx, )
Fetch EdgeJobs list

**Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PortainerEdgeJob**](portainer.EdgeJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobTaskLogsInspect**
> EdgejobsFileResponse EdgeJobTaskLogsInspect(ctx, id, taskID)
Fetch the log for a specifc task on an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 
  **taskID** | **string**| Task Id | 

### Return type

[**EdgejobsFileResponse**](edgejobs.fileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobTasksClear**
> EdgeJobTasksClear(ctx, id, taskID)
Clear the log for a specifc task on an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 
  **taskID** | **string**| Task Id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobTasksCollect**
> EdgeJobTasksCollect(ctx, id, taskID)
Collect the log for a specifc task on an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 
  **taskID** | **string**| Task Id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobTasksList**
> []EdgejobsTaskContainer EdgeJobTasksList(ctx, id)
Fetch the list of tasks on an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 

### Return type

[**[]EdgejobsTaskContainer**](edgejobs.taskContainer.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeJobUpdate**
> PortainerEdgeJob EdgeJobUpdate(ctx, id, body)
Update an EdgeJob

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| EdgeJob Id | 
  **body** | [**EdgejobsEdgeJobUpdatePayload**](EdgejobsEdgeJobUpdatePayload.md)| EdgeGroup data | 

### Return type

[**PortainerEdgeJob**](portainer.EdgeJob.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

