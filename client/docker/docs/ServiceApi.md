# \ServiceApi

All URIs are relative to *http://localhost/v1.42*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceCreate**](ServiceApi.md#ServiceCreate) | **Post** /services/create | Create a service
[**ServiceDelete**](ServiceApi.md#ServiceDelete) | **Delete** /services/{id} | Delete a service
[**ServiceInspect**](ServiceApi.md#ServiceInspect) | **Get** /services/{id} | Inspect a service
[**ServiceList**](ServiceApi.md#ServiceList) | **Get** /services | List services
[**ServiceLogs**](ServiceApi.md#ServiceLogs) | **Get** /services/{id}/logs | Get service logs
[**ServiceUpdate**](ServiceApi.md#ServiceUpdate) | **Post** /services/{id}/update | Update a service


# **ServiceCreate**
> ServiceCreateResponse ServiceCreate(ctx, body, optional)
Create a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**object**](.md)|  | 
 **optional** | ***ServiceApiServiceCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServiceCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRegistryAuth** | **optional.String**| A base64url-encoded auth configuration for pulling from private registries.  Refer to the [authentication section](#section/Authentication) for details.  | 

### Return type

[**ServiceCreateResponse**](ServiceCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceDelete**
> ServiceDelete(ctx, id)
Delete a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or name of service. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInspect**
> Service ServiceInspect(ctx, id, optional)
Inspect a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or name of service. | 
 **optional** | ***ServiceApiServiceInspectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServiceInspectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **insertDefaults** | **optional.Bool**| Fill empty fields with default values. | [default to false]

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceList**
> []Service ServiceList(ctx, optional)
List services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ServiceApiServiceListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServiceListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the services list.  Available filters:  - &#x60;id&#x3D;&lt;service id&gt;&#x60; - &#x60;label&#x3D;&lt;service label&gt;&#x60; - &#x60;mode&#x3D;[\&quot;replicated\&quot;|\&quot;global\&quot;]&#x60; - &#x60;name&#x3D;&lt;service name&gt;&#x60;  | 
 **status** | **optional.Bool**| Include service status, with count of running and desired tasks.  | 

### Return type

[**[]Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceLogs**
> string ServiceLogs(ctx, id, optional)
Get service logs

Get `stdout` and `stderr` logs from a service. See also [`/containers/{id}/logs`](#operation/ContainerLogs).  **Note**: This endpoint works only for services with the `local`, `json-file` or `journald` logging drivers. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or name of the service | 
 **optional** | ***ServiceApiServiceLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServiceLogsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **details** | **optional.Bool**| Show service context and extra details provided to logs. | [default to false]
 **follow** | **optional.Bool**| Keep connection after returning logs. | [default to false]
 **stdout** | **optional.Bool**| Return logs from &#x60;stdout&#x60; | [default to false]
 **stderr** | **optional.Bool**| Return logs from &#x60;stderr&#x60; | [default to false]
 **since** | **optional.Int32**| Only return logs since this time, as a UNIX timestamp | [default to 0]
 **timestamps** | **optional.Bool**| Add timestamps to every log line | [default to false]
 **tail** | **optional.String**| Only return this number of log lines from the end of the logs. Specify as an integer or &#x60;all&#x60; to output all log lines.  | [default to all]

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/vnd.docker.raw-stream, application/vnd.docker.multiplexed-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceUpdate**
> ServiceUpdateResponse ServiceUpdate(ctx, id, body, version, optional)
Update a service

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| ID or name of service. | 
  **body** | [**object**](.md)|  | 
  **version** | **int32**| The version number of the service object being updated. This is required to avoid conflicting writes. This version number should be the value as currently set on the service *before* the update. You can find the current version by calling &#x60;GET /services/{id}&#x60;  | 
 **optional** | ***ServiceApiServiceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceApiServiceUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **registryAuthFrom** | **optional.String**| If the &#x60;X-Registry-Auth&#x60; header is not specified, this parameter indicates where to find registry authorization credentials.  | [default to spec]
 **rollback** | **optional.String**| Set to this parameter to &#x60;previous&#x60; to cause a server-side rollback to the previous service spec. The supplied spec will be ignored in this case.  | 
 **xRegistryAuth** | **optional.String**| A base64url-encoded auth configuration for pulling from private registries.  Refer to the [authentication section](#section/Authentication) for details.  | 

### Return type

[**ServiceUpdateResponse**](ServiceUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

