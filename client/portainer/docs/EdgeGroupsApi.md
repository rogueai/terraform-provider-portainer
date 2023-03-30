# \EdgeGroupsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EdgeGroupCreate**](EdgeGroupsApi.md#EdgeGroupCreate) | **Post** /edge_groups | Create an EdgeGroup
[**EdgeGroupDelete**](EdgeGroupsApi.md#EdgeGroupDelete) | **Delete** /edge_groups/{id} | Deletes an EdgeGroup
[**EdgeGroupInspect**](EdgeGroupsApi.md#EdgeGroupInspect) | **Get** /edge_groups/{id} | Inspects an EdgeGroup
[**EdgeGroupList**](EdgeGroupsApi.md#EdgeGroupList) | **Get** /edge_groups | list EdgeGroups
[**EgeGroupUpdate**](EdgeGroupsApi.md#EgeGroupUpdate) | **Put** /edge_groups/{id} | Updates an EdgeGroup


# **EdgeGroupCreate**
> PortainerEdgeGroup EdgeGroupCreate(ctx, body)
Create an EdgeGroup

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EdgegroupsEdgeGroupCreatePayload**](EdgegroupsEdgeGroupCreatePayload.md)| EdgeGroup data | 

### Return type

[**PortainerEdgeGroup**](portainer.EdgeGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeGroupDelete**
> EdgeGroupDelete(ctx, id)
Deletes an EdgeGroup

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EdgeGroup Id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeGroupInspect**
> PortainerEdgeGroup EdgeGroupInspect(ctx, id)
Inspects an EdgeGroup

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EdgeGroup Id | 

### Return type

[**PortainerEdgeGroup**](portainer.EdgeGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EdgeGroupList**
> []EdgegroupsDecoratedEdgeGroup EdgeGroupList(ctx, )
list EdgeGroups

**Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]EdgegroupsDecoratedEdgeGroup**](edgegroups.decoratedEdgeGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EgeGroupUpdate**
> PortainerEdgeGroup EgeGroupUpdate(ctx, id, body)
Updates an EdgeGroup

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EdgeGroup Id | 
  **body** | [**EdgegroupsEdgeGroupUpdatePayload**](EdgegroupsEdgeGroupUpdatePayload.md)| EdgeGroup data | 

### Return type

[**PortainerEdgeGroup**](portainer.EdgeGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

