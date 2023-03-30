# \PluginApi

All URIs are relative to *http://localhost/v1.42*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPluginPrivileges**](PluginApi.md#GetPluginPrivileges) | **Get** /plugins/privileges | Get plugin privileges
[**PluginCreate**](PluginApi.md#PluginCreate) | **Post** /plugins/create | Create a plugin
[**PluginDelete**](PluginApi.md#PluginDelete) | **Delete** /plugins/{name} | Remove a plugin
[**PluginDisable**](PluginApi.md#PluginDisable) | **Post** /plugins/{name}/disable | Disable a plugin
[**PluginEnable**](PluginApi.md#PluginEnable) | **Post** /plugins/{name}/enable | Enable a plugin
[**PluginInspect**](PluginApi.md#PluginInspect) | **Get** /plugins/{name}/json | Inspect a plugin
[**PluginList**](PluginApi.md#PluginList) | **Get** /plugins | List plugins
[**PluginPull**](PluginApi.md#PluginPull) | **Post** /plugins/pull | Install a plugin
[**PluginPush**](PluginApi.md#PluginPush) | **Post** /plugins/{name}/push | Push a plugin
[**PluginSet**](PluginApi.md#PluginSet) | **Post** /plugins/{name}/set | Configure a plugin
[**PluginUpgrade**](PluginApi.md#PluginUpgrade) | **Post** /plugins/{name}/upgrade | Upgrade a plugin


# **GetPluginPrivileges**
> []PluginPrivilege GetPluginPrivileges(ctx, remote)
Get plugin privileges

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remote** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 

### Return type

[**[]PluginPrivilege**](PluginPrivilege.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginCreate**
> PluginCreate(ctx, name, optional)
Create a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 
 **optional** | ***PluginApiPluginCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tarContext** | **optional.String**| Path to tar containing plugin rootfs and manifest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-tar
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginDelete**
> Plugin PluginDelete(ctx, name, optional)
Remove a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 
 **optional** | ***PluginApiPluginDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Disable the plugin before removing. This may result in issues if the plugin is in use by a container.  | [default to false]

### Return type

[**Plugin**](Plugin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginDisable**
> PluginDisable(ctx, name)
Disable a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginEnable**
> PluginEnable(ctx, name, optional)
Enable a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 
 **optional** | ***PluginApiPluginEnableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginEnableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeout** | **optional.Int32**| Set the HTTP client timeout (in seconds) | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginInspect**
> Plugin PluginInspect(ctx, name)
Inspect a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 

### Return type

[**Plugin**](Plugin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginList**
> []Plugin PluginList(ctx, optional)
List plugins

Returns information about installed plugins.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PluginApiPluginListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the plugin list.  Available filters:  - &#x60;capability&#x3D;&lt;capability name&gt;&#x60; - &#x60;enable&#x3D;&lt;true&gt;|&lt;false&gt;&#x60;  | 

### Return type

[**[]Plugin**](Plugin.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginPull**
> PluginPull(ctx, remote, optional)
Install a plugin

Pulls and installs a plugin. After the plugin is installed, it can be enabled using the [`POST /plugins/{name}/enable` endpoint](#operation/PostPluginsEnable). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **remote** | **string**| Remote reference for plugin to install.  The &#x60;:latest&#x60; tag is optional, and is used as the default if omitted.  | 
 **optional** | ***PluginApiPluginPullOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginPullOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| Local name for the pulled plugin.  The &#x60;:latest&#x60; tag is optional, and is used as the default if omitted.  | 
 **xRegistryAuth** | **optional.String**| A base64url-encoded auth configuration to use when pulling a plugin from a registry.  Refer to the [authentication section](#section/Authentication) for details.  | 
 **body** | [**optional.Interface of []PluginPrivilege**](PluginPrivilege.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginPush**
> PluginPush(ctx, name)
Push a plugin

Push a plugin to the registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginSet**
> PluginSet(ctx, name, optional)
Configure a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 
 **optional** | ***PluginApiPluginSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginSetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **optional.[]string**|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PluginUpgrade**
> PluginUpgrade(ctx, name, remote, optional)
Upgrade a plugin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the plugin. The &#x60;:latest&#x60; tag is optional, and is the default if omitted.  | 
  **remote** | **string**| Remote reference to upgrade to.  The &#x60;:latest&#x60; tag is optional, and is used as the default if omitted.  | 
 **optional** | ***PluginApiPluginUpgradeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PluginApiPluginUpgradeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRegistryAuth** | **optional.String**| A base64url-encoded auth configuration to use when pulling a plugin from a registry.  Refer to the [authentication section](#section/Authentication) for details.  | 
 **body** | [**optional.Interface of []PluginPrivilege**](PluginPrivilege.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

