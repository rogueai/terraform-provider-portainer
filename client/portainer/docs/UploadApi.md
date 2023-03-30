# \UploadApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadTLS**](UploadApi.md#UploadTLS) | **Post** /upload/tls/{certificate} | Upload TLS files


# **UploadTLS**
> UploadTLS(ctx, certificate, folder, file)
Upload TLS files

Use this environment(endpoint) to upload TLS files. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **certificate** | **string**| TLS file type. Valid values are &#39;ca&#39;, &#39;cert&#39; or &#39;key&#39;. | 
  **folder** | **string**| Folder where the TLS file will be stored. Will be created if not existing | 
  **file** | ***os.File**| The file to upload | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

