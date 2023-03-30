# \TagsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagCreate**](TagsApi.md#TagCreate) | **Post** /tags | Create a new tag
[**TagDelete**](TagsApi.md#TagDelete) | **Delete** /tags/{id} | Remove a tag
[**TagList**](TagsApi.md#TagList) | **Get** /tags | List tags


# **TagCreate**
> PortainerTag TagCreate(ctx, body)
Create a new tag

Create a new tag. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TagsTagCreatePayload**](TagsTagCreatePayload.md)| Tag details | 

### Return type

[**PortainerTag**](portainer.Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagDelete**
> TagDelete(ctx, id)
Remove a tag

Remove a tag. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Tag identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagList**
> []PortainerTag TagList(ctx, )
List tags

List tags. **Access policy**: authenticated

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PortainerTag**](portainer.Tag.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

