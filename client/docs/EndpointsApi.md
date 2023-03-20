# \EndpointsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointAssociationDelete**](EndpointsApi.md#EndpointAssociationDelete) | **Put** /endpoints/{id}/association | De-association an edge environment(endpoint)
[**EndpointCreate**](EndpointsApi.md#EndpointCreate) | **Post** /endpoints | Create a new environment(endpoint)
[**EndpointCreateGlobalKey**](EndpointsApi.md#EndpointCreateGlobalKey) | **Post** /endpoints/global-key | Create or retrieve the endpoint for an EdgeID
[**EndpointDelete**](EndpointsApi.md#EndpointDelete) | **Delete** /endpoints/{id} | Remove an environment(endpoint)
[**EndpointDockerhubStatus**](EndpointsApi.md#EndpointDockerhubStatus) | **Get** /endpoints/{id}/dockerhub/{registryId} | fetch docker pull limits
[**EndpointEdgeStatusInspect**](EndpointsApi.md#EndpointEdgeStatusInspect) | **Get** /endpoints/{id}/edge/status | Get environment(endpoint) status
[**EndpointInspect**](EndpointsApi.md#EndpointInspect) | **Get** /endpoints/{id} | Inspect an environment(endpoint)
[**EndpointList**](EndpointsApi.md#EndpointList) | **Get** /endpoints | List environments(endpoints)
[**EndpointRegistriesList**](EndpointsApi.md#EndpointRegistriesList) | **Get** /endpoints/{id}/registries | List Registries on environment
[**EndpointRegistryAccess**](EndpointsApi.md#EndpointRegistryAccess) | **Put** /endpoints/{id}/registries/{registryId} | update registry access for environment
[**EndpointSettingsUpdate**](EndpointsApi.md#EndpointSettingsUpdate) | **Put** /endpoints/{id}/settings | Update settings for an environment(endpoint)
[**EndpointSnapshot**](EndpointsApi.md#EndpointSnapshot) | **Post** /endpoints/{id}/snapshot | Snapshots an environment(endpoint)
[**EndpointSnapshots**](EndpointsApi.md#EndpointSnapshots) | **Post** /endpoints/snapshot | Snapshot all environments(endpoints)
[**EndpointUpdate**](EndpointsApi.md#EndpointUpdate) | **Put** /endpoints/{id} | Update an environment(endpoint)
[**EndpointsIdEdgeJobsJobIDLogsPost**](EndpointsApi.md#EndpointsIdEdgeJobsJobIDLogsPost) | **Post** /endpoints/{id}/edge/jobs/{jobID}/logs | Inspect an EdgeJob Log
[**EndpointsIdEdgeStacksStackIdGet**](EndpointsApi.md#EndpointsIdEdgeStacksStackIdGet) | **Get** /endpoints/{id}/edge/stacks/{stackId} | Inspect an Edge Stack for an Environment(Endpoint)


# **EndpointAssociationDelete**
> EndpointAssociationDelete(ctx, id)
De-association an edge environment(endpoint)

De-association an edge environment(endpoint). **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointCreate**
> PortainerEndpoint EndpointCreate(ctx, name, endpointCreationType, edgeTunnelServerAddress, optional)
Create a new environment(endpoint)

Create a new environment(endpoint) that will be used to manage an environment(endpoint). **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name that will be used to identify this environment(endpoint) (example: my-environment) | 
  **endpointCreationType** | **int32**| Environment(Endpoint) type. Value must be one of: 1 (Local Docker environment), 2 (Agent environment), 3 (Azure environment), 4 (Edge agent environment) or 5 (Local Kubernetes Environment | 
  **edgeTunnelServerAddress** | **string**| URL or IP address that will be used to establish a reverse tunnel | 
 **optional** | ***EndpointsApiEndpointCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsApiEndpointCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **uRL** | **optional.String**| URL or IP address of a Docker host (example: docker.mydomain.tld:2375). Defaults to local if not specified (Linux: /var/run/docker.sock, Windows: //./pipe/docker_engine) | 
 **publicURL** | **optional.String**| URL or IP address where exposed containers will be reachable. Defaults to URL if not specified (example: docker.mydomain.tld:2375) | 
 **groupID** | **optional.Int32**| Environment(Endpoint) group identifier. If not specified will default to 1 (unassigned). | 
 **tLS** | **optional.Bool**| Require TLS to connect against this environment(endpoint) | 
 **tLSSkipVerify** | **optional.Bool**| Skip server verification when using TLS | 
 **tLSSkipClientVerify** | **optional.Bool**| Skip client verification when using TLS | 
 **tLSCACertFile** | **optional.Interface of *os.File**| TLS CA certificate file | 
 **tLSCertFile** | **optional.Interface of *os.File**| TLS client certificate file | 
 **tLSKeyFile** | **optional.Interface of *os.File**| TLS client key file | 
 **azureApplicationID** | **optional.String**| Azure application ID. Required if environment(endpoint) type is set to 3 | 
 **azureTenantID** | **optional.String**| Azure tenant ID. Required if environment(endpoint) type is set to 3 | 
 **azureAuthenticationKey** | **optional.String**| Azure authentication key. Required if environment(endpoint) type is set to 3 | 
 **tagIDs** | [**optional.Interface of []int32**](int32.md)| List of tag identifiers to which this environment(endpoint) is associated | 
 **edgeCheckinInterval** | **optional.Int32**| The check in interval for edge agent (in seconds) | 
 **isEdgeDevice** | **optional.Bool**| Is Edge Device | 
 **gpus** | [**optional.Interface of []string**](string.md)| List of GPUs | 

### Return type

[**PortainerEndpoint**](portainer.Endpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointCreateGlobalKey**
> EndpointsEndpointCreateGlobalKeyResponse EndpointCreateGlobalKey(ctx, )
Create or retrieve the endpoint for an EdgeID

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EndpointsEndpointCreateGlobalKeyResponse**](endpoints.endpointCreateGlobalKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointDelete**
> EndpointDelete(ctx, id)
Remove an environment(endpoint)

Remove an environment(endpoint). **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointDockerhubStatus**
> EndpointsDockerhubStatusResponse EndpointDockerhubStatus(ctx, id, registryId)
fetch docker pull limits

get docker pull limits for a docker hub registry in the environment **Access policy**:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| endpoint ID | 
  **registryId** | **int32**| registry ID | 

### Return type

[**EndpointsDockerhubStatusResponse**](endpoints.dockerhubStatusResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointEdgeStatusInspect**
> EndpointedgeEndpointEdgeStatusInspectResponse EndpointEdgeStatusInspect(ctx, id)
Get environment(endpoint) status

environment(endpoint) for edge agent to check status of environment(endpoint) **Access policy**: restricted only to Edge environments(endpoints)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

[**EndpointedgeEndpointEdgeStatusInspectResponse**](endpointedge.endpointEdgeStatusInspectResponse.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointInspect**
> PortainerEndpoint EndpointInspect(ctx, id)
Inspect an environment(endpoint)

Retrieve details about an environment(endpoint). **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

[**PortainerEndpoint**](portainer.Endpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointList**
> []PortainerEndpoint EndpointList(ctx, optional)
List environments(endpoints)

List all environments(endpoints) based on the current user authorizations. Will return all environments(endpoints) if using an administrator or team leader account otherwise it will only return authorized environments(endpoints). **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EndpointsApiEndpointListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsApiEndpointListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.Int32**| Start searching from | 
 **limit** | **optional.Int32**| Limit results to this value | 
 **sort** | **optional.Int32**| Sort results by this value | 
 **order** | **optional.Int32**| Order sorted results by desc/asc | 
 **search** | **optional.String**| Search query | 
 **groupIds** | [**optional.Interface of []int32**](int32.md)| List environments(endpoints) of these groups | 
 **status** | [**optional.Interface of []int32**](int32.md)| List environments(endpoints) by this status | 
 **types** | [**optional.Interface of []int32**](int32.md)| List environments(endpoints) of this type | 
 **tagIds** | [**optional.Interface of []int32**](int32.md)| search environments(endpoints) with these tags (depends on tagsPartialMatch) | 
 **tagsPartialMatch** | **optional.Bool**| If true, will return environment(endpoint) which has one of tagIds, if false (or missing) will return only environments(endpoints) that has all the tags | 
 **endpointIds** | [**optional.Interface of []int32**](int32.md)| will return only these environments(endpoints) | 
 **provisioned** | **optional.Bool**| If true, will return environment(endpoint) that were provisioned | 
 **agentVersions** | [**optional.Interface of []string**](string.md)| will return only environments with on of these agent versions | 
 **edgeDevice** | **optional.Bool**| if exists true show only edge devices, false show only regular edge endpoints. if missing, will show both types (relevant only for edge endpoints) | 
 **edgeDeviceUntrusted** | **optional.Bool**| if true, show only untrusted endpoints, if false show only trusted (relevant only for edge devices, and if edgeDevice is true) | 
 **name** | **optional.String**| will return only environments(endpoints) with this name | 

### Return type

[**[]PortainerEndpoint**](portainer.Endpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointRegistriesList**
> []PortainerRegistry EndpointRegistriesList(ctx, id, optional)
List Registries on environment

List all registries based on the current user authorizations in current environment. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
 **optional** | ***EndpointsApiEndpointRegistriesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsApiEndpointRegistriesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace** | **optional.String**| required if kubernetes environment, will show registries by namespace | 

### Return type

[**[]PortainerRegistry**](portainer.Registry.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointRegistryAccess**
> EndpointRegistryAccess(ctx, id, registryId, body)
update registry access for environment

**Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **registryId** | **int32**| Registry identifier | 
  **body** | [**EndpointsRegistryAccessPayload**](EndpointsRegistryAccessPayload.md)| details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointSettingsUpdate**
> PortainerEndpoint EndpointSettingsUpdate(ctx, id, body)
Update settings for an environment(endpoint)

Update settings for an environment(endpoint). **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **body** | [**EndpointsEndpointSettingsUpdatePayload**](EndpointsEndpointSettingsUpdatePayload.md)| Environment(Endpoint) details | 

### Return type

[**PortainerEndpoint**](portainer.Endpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointSnapshot**
> EndpointSnapshot(ctx, id)
Snapshots an environment(endpoint)

Snapshots an environment(endpoint) **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointSnapshots**
> EndpointSnapshots(ctx, )
Snapshot all environments(endpoints)

Snapshot all environments(endpoints) **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointUpdate**
> PortainerEndpoint EndpointUpdate(ctx, id, body)
Update an environment(endpoint)

Update an environment(endpoint). **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Environment(Endpoint) identifier | 
  **body** | [**EndpointsEndpointUpdatePayload**](EndpointsEndpointUpdatePayload.md)| Environment(Endpoint) details | 

### Return type

[**PortainerEndpoint**](portainer.Endpoint.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointsIdEdgeJobsJobIDLogsPost**
> EndpointsIdEdgeJobsJobIDLogsPost(ctx, id, jobID)
Inspect an EdgeJob Log

**Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| environment(endpoint) Id | 
  **jobID** | **string**| Job Id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointsIdEdgeStacksStackIdGet**
> EndpointedgeConfigResponse EndpointsIdEdgeStacksStackIdGet(ctx, id, stackId)
Inspect an Edge Stack for an Environment(Endpoint)

**Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| environment(endpoint) Id | 
  **stackId** | **string**| EdgeStack Id | 

### Return type

[**EndpointedgeConfigResponse**](endpointedge.configResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

