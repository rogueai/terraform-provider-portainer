# \SettingsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsInspect**](SettingsApi.md#SettingsInspect) | **Get** /settings | Retrieve Portainer settings
[**SettingsPublic**](SettingsApi.md#SettingsPublic) | **Get** /settings/public | Retrieve Portainer public settings
[**SettingsUpdate**](SettingsApi.md#SettingsUpdate) | **Put** /settings | Update Portainer settings


# **SettingsInspect**
> PortainerSettings SettingsInspect(ctx, )
Retrieve Portainer settings

Retrieve Portainer settings. **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PortainerSettings**](portainer.Settings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPublic**
> SettingsPublicSettingsResponse SettingsPublic(ctx, )
Retrieve Portainer public settings

Retrieve public settings. Returns a small set of settings that are not reserved to administrators only. **Access policy**: public

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SettingsPublicSettingsResponse**](settings.publicSettingsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsUpdate**
> PortainerSettings SettingsUpdate(ctx, body)
Update Portainer settings

Update Portainer settings. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsSettingsUpdatePayload**](SettingsSettingsUpdatePayload.md)| New settings | 

### Return type

[**PortainerSettings**](portainer.Settings.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

