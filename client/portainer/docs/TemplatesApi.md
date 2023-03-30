# \TemplatesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateFile**](TemplatesApi.md#TemplateFile) | **Post** /templates/file | Get a template&#39;s file
[**TemplateList**](TemplatesApi.md#TemplateList) | **Get** /templates | List available templates


# **TemplateFile**
> TemplatesFileResponse TemplateFile(ctx, body)
Get a template's file

Get a template's file **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TemplatesFilePayload**](TemplatesFilePayload.md)| File details | 

### Return type

[**TemplatesFileResponse**](templates.fileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateList**
> TemplatesListResponse TemplateList(ctx, )
List available templates

List available templates. **Access policy**: authenticated

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TemplatesListResponse**](templates.listResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

