# \CustomTemplatesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomTemplateCreate**](CustomTemplatesApi.md#CustomTemplateCreate) | **Post** /custom_templates | Create a custom template
[**CustomTemplateDelete**](CustomTemplatesApi.md#CustomTemplateDelete) | **Delete** /custom_templates/{id} | Remove a template
[**CustomTemplateFile**](CustomTemplatesApi.md#CustomTemplateFile) | **Get** /custom_templates/{id}/file | Get Template stack file content.
[**CustomTemplateInspect**](CustomTemplatesApi.md#CustomTemplateInspect) | **Get** /custom_templates/{id} | Inspect a custom template
[**CustomTemplateList**](CustomTemplatesApi.md#CustomTemplateList) | **Get** /custom_templates | List available custom templates
[**CustomTemplateUpdate**](CustomTemplatesApi.md#CustomTemplateUpdate) | **Put** /custom_templates/{id} | Update a template


# **CustomTemplateCreate**
> PortainerCustomTemplate CustomTemplateCreate(ctx, method, optional)
Create a custom template

Create a custom template. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **method** | **string**| method for creating template | 
 **optional** | ***CustomTemplatesApiCustomTemplateCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomTemplatesApiCustomTemplateCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bodyString** | [**optional.Interface of CustomtemplatesCustomTemplateFromFileContentPayload**](CustomtemplatesCustomTemplateFromFileContentPayload.md)| Required when using method&#x3D;string | 
 **bodyRepository** | [**optional.Interface of CustomtemplatesCustomTemplateFromGitRepositoryPayload**](CustomtemplatesCustomTemplateFromGitRepositoryPayload.md)| Required when using method&#x3D;repository | 
 **title** | **optional.String**| Title of the template. required when method is file | 
 **description** | **optional.String**| Description of the template. required when method is file | 
 **note** | **optional.String**| A note that will be displayed in the UI. Supports HTML content | 
 **platform** | **optional.Int32**| Platform associated to the template (1 - &#39;linux&#39;, 2 - &#39;windows&#39;). required when method is file | 
 **type_** | **optional.Int32**| Type of created stack (1 - swarm, 2 - compose), required when method is file | 
 **file** | **optional.Interface of *os.File**| required when method is file | 

### Return type

[**PortainerCustomTemplate**](portainer.CustomTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTemplateDelete**
> CustomTemplateDelete(ctx, id)
Remove a template

Remove a template. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTemplateFile**
> CustomtemplatesFileResponse CustomTemplateFile(ctx, id)
Get Template stack file content.

Retrieve the content of the Stack file for the specified custom template **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 

### Return type

[**CustomtemplatesFileResponse**](customtemplates.fileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTemplateInspect**
> PortainerCustomTemplate CustomTemplateInspect(ctx, id)
Inspect a custom template

Retrieve details about a template. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 

### Return type

[**PortainerCustomTemplate**](portainer.CustomTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTemplateList**
> []PortainerCustomTemplate CustomTemplateList(ctx, type_)
List available custom templates

List available custom templates. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | [**[]int32**](int32.md)| Template types | 

### Return type

[**[]PortainerCustomTemplate**](portainer.CustomTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTemplateUpdate**
> PortainerCustomTemplate CustomTemplateUpdate(ctx, id, body)
Update a template

Update a template. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 
  **body** | [**CustomtemplatesCustomTemplateUpdatePayload**](CustomtemplatesCustomTemplateUpdatePayload.md)| Template details | 

### Return type

[**PortainerCustomTemplate**](portainer.CustomTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

