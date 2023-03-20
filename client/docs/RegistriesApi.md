# \RegistriesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryConfigure**](RegistriesApi.md#RegistryConfigure) | **Post** /registries/{id}/configure | Configures a registry
[**RegistryCreate**](RegistriesApi.md#RegistryCreate) | **Post** /registries | Create a new registry
[**RegistryDelete**](RegistriesApi.md#RegistryDelete) | **Delete** /registries/{id} | Remove a registry
[**RegistryInspect**](RegistriesApi.md#RegistryInspect) | **Get** /registries/{id} | Inspect a registry
[**RegistryList**](RegistriesApi.md#RegistryList) | **Get** /registries | List Registries
[**RegistryUpdate**](RegistriesApi.md#RegistryUpdate) | **Put** /registries/{id} | Update a registry


# **RegistryConfigure**
> RegistryConfigure(ctx, id, body)
Configures a registry

Configures a registry. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 
  **body** | [**RegistriesRegistryConfigurePayload**](RegistriesRegistryConfigurePayload.md)| Registry configuration | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryCreate**
> PortainerRegistry RegistryCreate(ctx, body)
Create a new registry

Create a new registry. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegistriesRegistryCreatePayload**](RegistriesRegistryCreatePayload.md)| Registry details | 

### Return type

[**PortainerRegistry**](portainer.Registry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryDelete**
> RegistryDelete(ctx, id)
Remove a registry

Remove a registry **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryInspect**
> PortainerRegistry RegistryInspect(ctx, id)
Inspect a registry

Retrieve details about a registry. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 

### Return type

[**PortainerRegistry**](portainer.Registry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryList**
> []PortainerRegistry RegistryList(ctx, )
List Registries

List all registries based on the current user authorizations. Will return all registries if using an administrator account otherwise it will only return authorized registries. **Access policy**: restricted

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PortainerRegistry**](portainer.Registry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryUpdate**
> PortainerRegistry RegistryUpdate(ctx, id, body)
Update a registry

Update a registry **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 
  **body** | [**RegistriesRegistryUpdatePayload**](RegistriesRegistryUpdatePayload.md)| Registry details | 

### Return type

[**PortainerRegistry**](portainer.Registry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

