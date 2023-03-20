# \EndpointGroupsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointGroupAddEndpoint**](EndpointGroupsApi.md#EndpointGroupAddEndpoint) | **Put** /endpoint_groups/{id}/endpoints/{endpointId} | Add an environment(endpoint) to an environment(endpoint) group
[**EndpointGroupDelete**](EndpointGroupsApi.md#EndpointGroupDelete) | **Delete** /endpoint_groups/{id} | Remove an environment(endpoint) group
[**EndpointGroupDeleteEndpoint**](EndpointGroupsApi.md#EndpointGroupDeleteEndpoint) | **Delete** /endpoint_groups/{id}/endpoints/{endpointId} | Removes environment(endpoint) from an environment(endpoint) group
[**EndpointGroupList**](EndpointGroupsApi.md#EndpointGroupList) | **Get** /endpoint_groups | List Environment(Endpoint) groups
[**EndpointGroupUpdate**](EndpointGroupsApi.md#EndpointGroupUpdate) | **Put** /endpoint_groups/{id} | Update an environment(endpoint) group
[**EndpointGroupsIdGet**](EndpointGroupsApi.md#EndpointGroupsIdGet) | **Get** /endpoint_groups/{id} | Inspect an Environment(Endpoint) group
[**EndpointGroupsPost**](EndpointGroupsApi.md#EndpointGroupsPost) | **Post** /endpoint_groups | Create an Environment(Endpoint) Group


# **EndpointGroupAddEndpoint**
> EndpointGroupAddEndpoint(ctx, id, endpointId)
Add an environment(endpoint) to an environment(endpoint) group

Add an environment(endpoint) to an environment(endpoint) group **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 
  **endpointId** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupDelete**
> EndpointGroupDelete(ctx, id)
Remove an environment(endpoint) group

Remove an environment(endpoint) group. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupDeleteEndpoint**
> EndpointGroupDeleteEndpoint(ctx, id, endpointId)
Removes environment(endpoint) from an environment(endpoint) group

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 
  **endpointId** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupList**
> []PortainerEndpointGroup EndpointGroupList(ctx, )
List Environment(Endpoint) groups

List all environment(endpoint) groups based on the current user authorizations. Will return all environment(endpoint) groups if using an administrator account otherwise it will only return authorized environment(endpoint) groups. **Access policy**: restricted

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PortainerEndpointGroup**](portainer.EndpointGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupUpdate**
> PortainerEndpointGroup EndpointGroupUpdate(ctx, id, body)
Update an environment(endpoint) group

Update an environment(endpoint) group. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 
  **body** | [**EndpointgroupsEndpointGroupUpdatePayload**](EndpointgroupsEndpointGroupUpdatePayload.md)| EndpointGroup details | 

### Return type

[**PortainerEndpointGroup**](portainer.EndpointGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupsIdGet**
> PortainerEndpointGroup EndpointGroupsIdGet(ctx, id)
Inspect an Environment(Endpoint) group

Retrieve details abont an environment(endpoint) group. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) group identifier | 

### Return type

[**PortainerEndpointGroup**](portainer.EndpointGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupsPost**
> PortainerEndpointGroup EndpointGroupsPost(ctx, body)
Create an Environment(Endpoint) Group

Create a new environment(endpoint) group. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointgroupsEndpointGroupCreatePayload**](EndpointgroupsEndpointGroupCreatePayload.md)| Environment(Endpoint) Group details | 

### Return type

[**PortainerEndpointGroup**](portainer.EndpointGroup.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

