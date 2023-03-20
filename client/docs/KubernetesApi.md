# \KubernetesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetKubernetesConfig**](KubernetesApi.md#GetKubernetesConfig) | **Get** /kubernetes/config | Generates kubeconfig file enabling client communication with k8s api server
[**GetKubernetesNodesLimits**](KubernetesApi.md#GetKubernetesNodesLimits) | **Get** /kubernetes/{id}/nodes_limits | Get CPU and memory limits of all nodes within k8s cluster
[**KubernetesNamespacesToggleSystem**](KubernetesApi.md#KubernetesNamespacesToggleSystem) | **Put** /kubernetes/{id}/namespaces/{namespace}/system | Toggle the system state for a namespace


# **GetKubernetesConfig**
> GetKubernetesConfig(ctx, optional)
Generates kubeconfig file enabling client communication with k8s api server

Generates kubeconfig file enabling client communication with k8s api server **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***KubernetesApiGetKubernetesConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KubernetesApiGetKubernetesConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int32**](int32.md)| will include only these environments(endpoints) | 
 **excludeIds** | [**optional.Interface of []int32**](int32.md)| will exclude these environments(endpoints) | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKubernetesNodesLimits**
> PortainerK8sNodesLimits GetKubernetesNodesLimits(ctx, id)
Get CPU and memory limits of all nodes within k8s cluster

Get CPU and memory limits of all nodes within k8s cluster **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

[**PortainerK8sNodesLimits**](portainer.K8sNodesLimits.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **KubernetesNamespacesToggleSystem**
> KubernetesNamespacesToggleSystem(ctx, id, namespace, body)
Toggle the system state for a namespace

Toggle the system state for a namespace **Access policy**: administrator or environment(endpoint) admin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **namespace** | **string**| Namespace name | 
  **body** | [**KubernetesNamespacesToggleSystemPayload**](KubernetesNamespacesToggleSystemPayload.md)| Update details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

