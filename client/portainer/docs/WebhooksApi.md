# \WebhooksApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhooksGet**](WebhooksApi.md#WebhooksGet) | **Get** /webhooks | List webhooks
[**WebhooksIdDelete**](WebhooksApi.md#WebhooksIdDelete) | **Delete** /webhooks/{id} | Delete a webhook
[**WebhooksIdPut**](WebhooksApi.md#WebhooksIdPut) | **Put** /webhooks/{id} | Update a webhook
[**WebhooksPost**](WebhooksApi.md#WebhooksPost) | **Post** /webhooks | Create a webhook
[**WebhooksTokenPost**](WebhooksApi.md#WebhooksTokenPost) | **Post** /webhooks/{token} | Execute a webhook


# **WebhooksGet**
> []PortainerWebhook WebhooksGet(ctx, optional)
List webhooks

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WebhooksApiWebhooksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WebhooksApiWebhooksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **endpointID** | **optional.Int32**|  | 
 **resourceID** | **optional.String**|  | 

### Return type

[**[]PortainerWebhook**](portainer.Webhook.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhooksIdDelete**
> WebhooksIdDelete(ctx, id)
Delete a webhook

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Webhook id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhooksIdPut**
> PortainerWebhook WebhooksIdPut(ctx, body)
Update a webhook

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksWebhookUpdatePayload**](WebhooksWebhookUpdatePayload.md)| Webhook data | 

### Return type

[**PortainerWebhook**](portainer.Webhook.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhooksPost**
> PortainerWebhook WebhooksPost(ctx, body)
Create a webhook

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhooksWebhookCreatePayload**](WebhooksWebhookCreatePayload.md)| Webhook data | 

### Return type

[**PortainerWebhook**](portainer.Webhook.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhooksTokenPost**
> WebhooksTokenPost(ctx, token)
Execute a webhook

Acts on a passed in token UUID to restart the docker service **Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**| Webhook token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

