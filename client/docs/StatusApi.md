# \StatusApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StatusInspect**](StatusApi.md#StatusInspect) | **Get** /status | Check Portainer status
[**StatusNodesCount**](StatusApi.md#StatusNodesCount) | **Get** /status/nodes | Retrieve the count of nodes
[**Version**](StatusApi.md#Version) | **Get** /status/version | Check for portainer updates


# **StatusInspect**
> SystemStatus StatusInspect(ctx, )
Check Portainer status

Deprecated: use the `/system/status` endpoint instead. Retrieve Portainer status **Access policy**: public

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

# **StatusNodesCount**
> SystemNodesCountResponse StatusNodesCount(ctx, )
Retrieve the count of nodes

Deprecated: use the `/system/nodes` endpoint instead. **Access policy**: authenticated

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

# **Version**
> SystemVersionResponse Version(ctx, )
Check for portainer updates

Deprecated: use the `/system/version` endpoint instead. Check if portainer has an update available **Access policy**: authenticated

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

