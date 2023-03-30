# \RbacEnabledApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IsRBACEnabled**](RbacEnabledApi.md#IsRBACEnabled) | **Get** /kubernetes/{id}/rbac_enabled | Check if RBAC is enabled


# **IsRBACEnabled**
> IsRBACEnabled(ctx, id)
Check if RBAC is enabled

Check if RBAC is enabled in the current Kubernetes cluster. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

