# \SslApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SSLInspect**](SslApi.md#SSLInspect) | **Get** /ssl | Inspect the ssl settings
[**SSLUpdate**](SslApi.md#SSLUpdate) | **Put** /ssl | Update the ssl settings


# **SSLInspect**
> PortainerSslSettings SSLInspect(ctx, )
Inspect the ssl settings

Retrieve the ssl settings. **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PortainerSslSettings**](portainer.SSLSettings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SSLUpdate**
> SSLUpdate(ctx, body)
Update the ssl settings

Update the ssl settings. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SslSslUpdatePayload**](SslSslUpdatePayload.md)| SSL Settings | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

