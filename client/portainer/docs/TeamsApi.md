# \TeamsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamCreate**](TeamsApi.md#TeamCreate) | **Post** /team | Create a new team
[**TeamDelete**](TeamsApi.md#TeamDelete) | **Delete** /teams/{id} | Remove a team
[**TeamInspect**](TeamsApi.md#TeamInspect) | **Get** /teams/{id} | Inspect a team
[**TeamList**](TeamsApi.md#TeamList) | **Get** /teams | List teams
[**TeamUpdate**](TeamsApi.md#TeamUpdate) | **Put** /team/{id} | Update a team


# **TeamCreate**
> PortainerTeam TeamCreate(ctx, body)
Create a new team

Create a new team. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TeamsTeamCreatePayload**](TeamsTeamCreatePayload.md)| details | 

### Return type

[**PortainerTeam**](portainer.Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamDelete**
> TeamDelete(ctx, id)
Remove a team

Remove a team. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Team Id | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamInspect**
> PortainerTeam TeamInspect(ctx, id)
Inspect a team

Retrieve details about a team. Access is only available for administrator and leaders of that team. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team identifier | 

### Return type

[**PortainerTeam**](portainer.Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamList**
> []PortainerTeam TeamList(ctx, optional)
List teams

List teams. For non-administrator users, will only list the teams they are member of. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TeamsApiTeamListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TeamsApiTeamListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **onlyLedTeams** | **optional.Bool**| Only list teams that the user is leader of | 
 **environmentId** | **optional.Int32**| Identifier of the environment(endpoint) that will be used to filter the authorized teams | 

### Return type

[**[]PortainerTeam**](portainer.Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamUpdate**
> PortainerTeam TeamUpdate(ctx, id, body)
Update a team

Update a team. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team identifier | 
  **body** | [**TeamsTeamUpdatePayload**](TeamsTeamUpdatePayload.md)| Team details | 

### Return type

[**PortainerTeam**](portainer.Team.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

