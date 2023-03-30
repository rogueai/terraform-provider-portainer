# \DistributionApi

All URIs are relative to *http://localhost/v1.42*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistributionInspect**](DistributionApi.md#DistributionInspect) | **Get** /distribution/{name}/json | Get image information from the registry


# **DistributionInspect**
> DistributionInspect DistributionInspect(ctx, name)
Get image information from the registry

Return image digest and platform information by contacting the registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or id | 

### Return type

[**DistributionInspect**](DistributionInspect.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

