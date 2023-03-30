# \ResourceControlsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceControlCreate**](ResourceControlsApi.md#ResourceControlCreate) | **Post** /resource_controls | Create a new resource control
[**ResourceControlDelete**](ResourceControlsApi.md#ResourceControlDelete) | **Delete** /resource_controls/{id} | Remove a resource control
[**ResourceControlUpdate**](ResourceControlsApi.md#ResourceControlUpdate) | **Put** /resource_controls/{id} | Update a resource control


# **ResourceControlCreate**
> PortainerResourceControl ResourceControlCreate(ctx, body)
Create a new resource control

Create a new resource control to restrict access to a Docker resource. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResourcecontrolsResourceControlCreatePayload**](ResourcecontrolsResourceControlCreatePayload.md)| Resource control details | 

### Return type

[**PortainerResourceControl**](portainer.ResourceControl.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResourceControlDelete**
> ResourceControlDelete(ctx, id)
Remove a resource control

Remove a resource control. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Resource control identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ResourceControlUpdate**
> PortainerResourceControl ResourceControlUpdate(ctx, id, body)
Update a resource control

Update a resource control **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Resource control identifier | 
  **body** | [**ResourcecontrolsResourceControlUpdatePayload**](ResourcecontrolsResourceControlUpdatePayload.md)| Resource control details | 

### Return type

[**PortainerResourceControl**](portainer.ResourceControl.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

