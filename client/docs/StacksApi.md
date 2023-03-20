# \StacksApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StackAssociate**](StacksApi.md#StackAssociate) | **Put** /stacks/{id}/associate | Associate an orphaned stack to a new environment(endpoint)
[**StackCreate**](StacksApi.md#StackCreate) | **Post** /stacks | Deploy a new stack
[**StackDelete**](StacksApi.md#StackDelete) | **Delete** /stacks/{id} | Remove a stack
[**StackFileInspect**](StacksApi.md#StackFileInspect) | **Get** /stacks/{id}/file | Retrieve the content of the Stack file for the specified stack
[**StackGitRedeploy**](StacksApi.md#StackGitRedeploy) | **Put** /stacks/{id}/git/redeploy | Redeploy a stack
[**StackInspect**](StacksApi.md#StackInspect) | **Get** /stacks/{id} | Inspect a stack
[**StackList**](StacksApi.md#StackList) | **Get** /stacks | List stacks
[**StackMigrate**](StacksApi.md#StackMigrate) | **Post** /stacks/{id}/migrate | Migrate a stack to another environment(endpoint)
[**StackStart**](StacksApi.md#StackStart) | **Post** /stacks/{id}/start | Starts a stopped Stack
[**StackStop**](StacksApi.md#StackStop) | **Post** /stacks/{id}/stop | Stops a stopped Stack
[**StackUpdate**](StacksApi.md#StackUpdate) | **Put** /stacks/{id} | Update a stack
[**StackUpdateGit**](StacksApi.md#StackUpdateGit) | **Post** /stacks/{id}/git | Update a stack&#39;s Git configs
[**WebhookInvoke**](StacksApi.md#WebhookInvoke) | **Post** /stacks/webhooks/{webhookID} | Webhook for triggering stack updates from git


# **StackAssociate**
> PortainerStack StackAssociate(ctx, id, endpointId, swarmId, orphanedRunning)
Associate an orphaned stack to a new environment(endpoint)

**Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
  **endpointId** | **int32**| Stacks created before version 1.18.0 might not have an associated environment(endpoint) identifier. Use this optional parameter to set the environment(endpoint) identifier used by the stack. | 
  **swarmId** | **int32**| Swarm identifier | 
  **orphanedRunning** | **bool**| Indicates whether the stack is orphaned | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackCreate**
> PortainerCustomTemplate StackCreate(ctx, type_, method, endpointId, optional)
Deploy a new stack

Deploy a new stack into a Docker environment(endpoint) specified via the environment(endpoint) identifier. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **int32**| Stack deployment type. Possible values: 1 (Swarm stack), 2 (Compose stack) or 3 (Kubernetes stack). | 
  **method** | **string**| Stack deployment method. Possible values: file, string, repository or url. | 
  **endpointId** | **int32**| Identifier of the environment(endpoint) that will be used to deploy the stack | 
 **optional** | ***StacksApiStackCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **bodySwarmString** | [**optional.Interface of StacksSwarmStackFromFileContentPayload**](StacksSwarmStackFromFileContentPayload.md)| Required when using method&#x3D;string and type&#x3D;1 | 
 **bodySwarmRepository** | [**optional.Interface of StacksSwarmStackFromGitRepositoryPayload**](StacksSwarmStackFromGitRepositoryPayload.md)| Required when using method&#x3D;repository and type&#x3D;1 | 
 **bodyComposeString** | [**optional.Interface of StacksComposeStackFromFileContentPayload**](StacksComposeStackFromFileContentPayload.md)| Required when using method&#x3D;string and type&#x3D;2 | 
 **bodyComposeRepository** | [**optional.Interface of StacksComposeStackFromGitRepositoryPayload**](StacksComposeStackFromGitRepositoryPayload.md)| Required when using method&#x3D;repository and type&#x3D;2 | 
 **bodyKubernetesString** | [**optional.Interface of StacksKubernetesStringDeploymentPayload**](StacksKubernetesStringDeploymentPayload.md)| Required when using method&#x3D;string and type&#x3D;3 | 
 **bodyKubernetesRepository** | [**optional.Interface of StacksKubernetesGitDeploymentPayload**](StacksKubernetesGitDeploymentPayload.md)| Required when using method&#x3D;repository and type&#x3D;3 | 
 **bodyKubernetesUrl** | [**optional.Interface of StacksKubernetesManifestUrlDeploymentPayload**](StacksKubernetesManifestUrlDeploymentPayload.md)| Required when using method&#x3D;url and type&#x3D;3 | 
 **name** | **optional.String**| Name of the stack. required when method is file | 
 **swarmID** | **optional.String**| Swarm cluster identifier. Required when method equals file and type equals 1. required when method is file | 
 **env** | **optional.String**| Environment(Endpoint) variables passed during deployment, represented as a JSON array [{&#39;name&#39;: &#39;name&#39;, &#39;value&#39;: &#39;value&#39;}]. Optional, used when method equals file and type equals 1. | 
 **file** | **optional.Interface of *os.File**| Stack file. required when method is file | 

### Return type

[**PortainerCustomTemplate**](portainer.CustomTemplate.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json, multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackDelete**
> StackDelete(ctx, id, optional)
Remove a stack

Remove a stack. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
 **optional** | ***StacksApiStackDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **external** | **optional.Bool**| Set to true to delete an external stack. Only external Swarm stacks are supported | 
 **endpointId** | **optional.Int32**| Environment(Endpoint) identifier used to remove an external stack (required when external is set to true) | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackFileInspect**
> StacksStackFileResponse StackFileInspect(ctx, id)
Retrieve the content of the Stack file for the specified stack

Get Stack file content. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 

### Return type

[**StacksStackFileResponse**](stacks.stackFileResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackGitRedeploy**
> PortainerStack StackGitRedeploy(ctx, id, body, optional)
Redeploy a stack

Pull and redeploy a stack via Git **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
  **body** | [**StacksStackGitRedployPayload**](StacksStackGitRedployPayload.md)| Git configs for pull and redeploy a stack | 
 **optional** | ***StacksApiStackGitRedeployOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackGitRedeployOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointId** | **optional.Int32**| Stacks created before version 1.18.0 might not have an associated environment(endpoint) identifier. Use this optional parameter to set the environment(endpoint) identifier used by the stack. | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackInspect**
> PortainerStack StackInspect(ctx, id)
Inspect a stack

Retrieve details about a stack. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackList**
> []PortainerStack StackList(ctx, optional)
List stacks

List all stacks based on the current user authorizations. Will return all stacks if using an administrator account otherwise it will only return the list of stacks the user have access to. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StacksApiStackListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filters to process on the stack list. Encoded as JSON (a map[string]string). For example, {&#39;SwarmID&#39;: &#39;jpofkc0i9uo9wtx1zesuk649w&#39;} will only return stacks that are part of the specified Swarm cluster. Available filters: EndpointID, SwarmID. | 

### Return type

[**[]PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackMigrate**
> PortainerStack StackMigrate(ctx, id, body, optional)
Migrate a stack to another environment(endpoint)

Migrate a stack from an environment(endpoint) to another environment(endpoint). It will re-create the stack inside the target environment(endpoint) before removing the original stack. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
  **body** | [**StacksStackMigratePayload**](StacksStackMigratePayload.md)| Stack migration details | 
 **optional** | ***StacksApiStackMigrateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackMigrateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointId** | **optional.Int32**| Stacks created before version 1.18.0 might not have an associated environment(endpoint) identifier. Use this optional parameter to set the environment(endpoint) identifier used by the stack. | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackStart**
> PortainerStack StackStart(ctx, id)
Starts a stopped Stack

Starts a stopped Stack. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackStop**
> PortainerStack StackStop(ctx, id)
Stops a stopped Stack

Stops a stopped Stack. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackUpdate**
> PortainerStack StackUpdate(ctx, id, body, optional)
Update a stack

Update a stack, only for file based stacks. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
  **body** | [**StacksUpdateSwarmStackPayload**](StacksUpdateSwarmStackPayload.md)| Stack details | 
 **optional** | ***StacksApiStackUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointId** | **optional.Int32**| Stacks created before version 1.18.0 might not have an associated environment(endpoint) identifier. Use this optional parameter to set the environment(endpoint) identifier used by the stack. | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackUpdateGit**
> PortainerStack StackUpdateGit(ctx, id, body, optional)
Update a stack's Git configs

Update the Git settings in a stack, e.g., RepositoryReferenceName and AutoUpdate **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
  **body** | [**StacksStackGitUpdatePayload**](StacksStackGitUpdatePayload.md)| Git configs for pull and redeploy a stack | 
 **optional** | ***StacksApiStackUpdateGitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StacksApiStackUpdateGitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointId** | **optional.Int32**| Stacks created before version 1.18.0 might not have an associated environment(endpoint) identifier. Use this optional parameter to set the environment(endpoint) identifier used by the stack. | 

### Return type

[**PortainerStack**](portainer.Stack.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WebhookInvoke**
> WebhookInvoke(ctx, webhookID)
Webhook for triggering stack updates from git

**Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **webhookID** | **string**| Stack identifier | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

