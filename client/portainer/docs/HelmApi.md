# \HelmApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HelmDelete**](HelmApi.md#HelmDelete) | **Delete** /endpoints/{id}/kubernetes/helm/{release} | Delete Helm Release
[**HelmInstall**](HelmApi.md#HelmInstall) | **Post** /endpoints/{id}/kubernetes/helm | Install Helm Chart
[**HelmList**](HelmApi.md#HelmList) | **Get** /endpoints/{id}/kubernetes/helm | List Helm Releases
[**HelmRepoSearch**](HelmApi.md#HelmRepoSearch) | **Get** /templates/helm | Search Helm Charts
[**HelmShow**](HelmApi.md#HelmShow) | **Get** /templates/helm/{command} | Show Helm Chart Information
[**HelmUserRepositoriesList**](HelmApi.md#HelmUserRepositoriesList) | **Get** /endpoints/{id}/kubernetes/helm/repositories | List a users helm repositories
[**HelmUserRepositoryCreate**](HelmApi.md#HelmUserRepositoryCreate) | **Post** /endpoints/{id}/kubernetes/helm/repositories | Create a user helm repository


# **HelmDelete**
> HelmDelete(ctx, id, release, namespace)
Delete Helm Release

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **release** | **string**| The name of the release/application to uninstall | 
  **namespace** | **string**| An optional namespace | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelmInstall**
> ReleaseRelease HelmInstall(ctx, id, payload)
Install Helm Chart

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **payload** | [**HelmInstallChartPayload**](HelmInstallChartPayload.md)| Chart details | 

### Return type

[**ReleaseRelease**](release.Release.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelmList**
> []ReleaseReleaseElement HelmList(ctx, id, namespace, filter, selector)
List Helm Releases

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **namespace** | **string**| specify an optional namespace | 
  **filter** | **string**| specify an optional filter | 
  **selector** | **string**| specify an optional selector | 

### Return type

[**[]ReleaseReleaseElement**](release.ReleaseElement.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelmRepoSearch**
> string HelmRepoSearch(ctx, repo)
Search Helm Charts

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| Helm repository URL | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelmShow**
> string HelmShow(ctx, repo, chart, command)
Show Helm Chart Information

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| Helm repository URL | 
  **chart** | **string**| Chart name | 
  **command** | **string**| chart/values/readme | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelmUserRepositoriesList**
> HelmHelmUserRepositoryResponse HelmUserRepositoriesList(ctx, id)
List a users helm repositories

Inspect a user helm repositories. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

[**HelmHelmUserRepositoryResponse**](helm.helmUserRepositoryResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HelmUserRepositoryCreate**
> PortainerHelmUserRepository HelmUserRepositoryCreate(ctx, id, payload)
Create a user helm repository

Create a user helm repository. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **payload** | [**HelmAddHelmRepoUrlPayload**](HelmAddHelmRepoUrlPayload.md)| Helm Repository | 

### Return type

[**PortainerHelmUserRepository**](portainer.HelmUserRepository.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

