# \SystemApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemInfo**](SystemApi.md#SystemInfo) | **Get** /system/info | Retrieve system info
[**SystemNodesCount**](SystemApi.md#SystemNodesCount) | **Get** /system/nodes | Retrieve the count of nodes
[**SystemStatus**](SystemApi.md#SystemStatus) | **Get** /system/status | Check Portainer status
[**SystemUpgrade**](SystemApi.md#SystemUpgrade) | **Post** /system/upgrade | Upgrade Portainer to BE
[**SystemVersion**](SystemApi.md#SystemVersion) | **Get** /system/version | Check for portainer updates


# **SystemInfo**
> SystemSystemInfoResponse SystemInfo(ctx, )
Retrieve system info

**Access policy**: authenticated

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemSystemInfoResponse**](system.systemInfoResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemNodesCount**
> SystemNodesCountResponse SystemNodesCount(ctx, )
Retrieve the count of nodes

**Access policy**: authenticated

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemNodesCountResponse**](system.nodesCountResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemStatus**
> SystemStatus SystemStatus(ctx, )
Check Portainer status

Retrieve Portainer status **Access policy**: public

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemStatus**](system.status.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemUpgrade**
> SystemStatus SystemUpgrade(ctx, )
Upgrade Portainer to BE

Upgrade Portainer to BE **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemStatus**](system.status.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemVersion**
> SystemVersionResponse SystemVersion(ctx, )
Check for portainer updates

Check if portainer has an update available **Access policy**: authenticated

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemVersionResponse**](system.versionResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

