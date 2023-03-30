# \IntelApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProfile**](IntelApi.md#CreateProfile) | **Post** /fdo/profiles | creates a new FDO Profile
[**DeleteProfile**](IntelApi.md#DeleteProfile) | **Delete** /fdo/profiles/{id} | deletes a FDO Profile
[**DeviceAction**](IntelApi.md#DeviceAction) | **Post** /open_amt/{id}/devices/{deviceId}/action | Execute out of band action on an AMT managed device
[**DeviceFeatures**](IntelApi.md#DeviceFeatures) | **Post** /open_amt/{id}/devices_features/{deviceId} | Enable features on an AMT managed device
[**Duplicate**](IntelApi.md#Duplicate) | **Post** /fdo/profiles/{id}/duplicate | duplicated an existing FDO Profile
[**FdoConfigure**](IntelApi.md#FdoConfigure) | **Post** /fdo | Enable Portainer&#39;s FDO capabilities
[**FdoConfigureDevice**](IntelApi.md#FdoConfigureDevice) | **Post** /fdo/configure/{guid} | configures an FDO device
[**FdoListAll**](IntelApi.md#FdoListAll) | **Get** /fdo/list | List all known FDO vouchers
[**FdoProfileInspect**](IntelApi.md#FdoProfileInspect) | **Get** /fdo/profiles/{id} | retrieves a given FDO profile information and content
[**FdoProfileList**](IntelApi.md#FdoProfileList) | **Get** /fdo/profiles | retrieves all FDO profiles
[**FdoRegisterDevice**](IntelApi.md#FdoRegisterDevice) | **Post** /fdo/register | register an FDO device
[**OpenAMTActivate**](IntelApi.md#OpenAMTActivate) | **Post** /open_amt/{id}/activate | Activate OpenAMT device and associate to agent endpoint
[**OpenAMTConfigure**](IntelApi.md#OpenAMTConfigure) | **Post** /open_amt | Enable Portainer&#39;s OpenAMT capabilities
[**OpenAMTDevices**](IntelApi.md#OpenAMTDevices) | **Get** /open_amt/{id}/devices | Fetch OpenAMT managed devices information for endpoint
[**OpenAMTHostInfo**](IntelApi.md#OpenAMTHostInfo) | **Get** /open_amt/{id}/info | Request OpenAMT info from a node
[**UpdateProfile**](IntelApi.md#UpdateProfile) | **Put** /fdo/profiles/{id} | updates an existing FDO Profile


# **CreateProfile**
> CreateProfile(ctx, )
creates a new FDO Profile

creates a new FDO Profile **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProfile**
> DeleteProfile(ctx, )
deletes a FDO Profile

deletes a FDO Profile **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeviceAction**
> DeviceAction(ctx, body)
Execute out of band action on an AMT managed device

Execute out of band action on an AMT managed device **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenamtDeviceActionPayload**](OpenamtDeviceActionPayload.md)| Device Action | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeviceFeatures**
> DeviceFeatures(ctx, body)
Enable features on an AMT managed device

Enable features on an AMT managed device **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenamtDeviceFeaturesPayload**](OpenamtDeviceFeaturesPayload.md)| Device Features | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Duplicate**
> Duplicate(ctx, )
duplicated an existing FDO Profile

duplicated an existing FDO Profile **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FdoConfigure**
> FdoConfigure(ctx, body)
Enable Portainer's FDO capabilities

Enable Portainer's FDO capabilities **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FdoFdoConfigurePayload**](FdoFdoConfigurePayload.md)| FDO Settings | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FdoConfigureDevice**
> FdoConfigureDevice(ctx, body)
configures an FDO device

configures an FDO device **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FdoDeviceConfigurePayload**](FdoDeviceConfigurePayload.md)| Device Configuration | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FdoListAll**
> FdoListAll(ctx, )
List all known FDO vouchers

List all known FDO vouchers **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FdoProfileInspect**
> FdoProfileInspect(ctx, )
retrieves a given FDO profile information and content

retrieves a given FDO profile information and content **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FdoProfileList**
> FdoProfileList(ctx, )
retrieves all FDO profiles

retrieves all FDO profiles **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FdoRegisterDevice**
> FdoRegisterDevice(ctx, )
register an FDO device

register an FDO device **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenAMTActivate**
> OpenAMTActivate(ctx, id)
Activate OpenAMT device and associate to agent endpoint

Activate OpenAMT device and associate to agent endpoint **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenAMTConfigure**
> OpenAMTConfigure(ctx, body)
Enable Portainer's OpenAMT capabilities

Enable Portainer's OpenAMT capabilities **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OpenamtOpenAmtConfigurePayload**](OpenamtOpenAmtConfigurePayload.md)| OpenAMT Settings | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenAMTDevices**
> OpenAMTDevices(ctx, id)
Fetch OpenAMT managed devices information for endpoint

Fetch OpenAMT managed devices information for endpoint **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OpenAMTHostInfo**
> OpenAMTHostInfo(ctx, )
Request OpenAMT info from a node

Request OpenAMT info from a node **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProfile**
> UpdateProfile(ctx, )
updates an existing FDO Profile

updates an existing FDO Profile **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

