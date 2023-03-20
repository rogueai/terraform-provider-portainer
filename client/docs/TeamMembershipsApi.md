# \TeamMembershipsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamMembershipCreate**](TeamMembershipsApi.md#TeamMembershipCreate) | **Post** /team_memberships | Create a new team membership
[**TeamMembershipDelete**](TeamMembershipsApi.md#TeamMembershipDelete) | **Delete** /team_memberships/{id} | Remove a team membership
[**TeamMembershipList**](TeamMembershipsApi.md#TeamMembershipList) | **Get** /team_memberships | List team memberships
[**TeamMembershipUpdate**](TeamMembershipsApi.md#TeamMembershipUpdate) | **Put** /team_memberships/{id} | Update a team membership
[**TeamMemberships**](TeamMembershipsApi.md#TeamMemberships) | **Get** /teams/{id}/memberships | List team memberships


# **TeamMembershipCreate**
> PortainerTeamMembership TeamMembershipCreate(ctx, body)
Create a new team membership

Create a new team memberships. Access is only available to administrators leaders of the associated team. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TeammembershipsTeamMembershipCreatePayload**](TeammembershipsTeamMembershipCreatePayload.md)| Team membership details | 

### Return type

[**PortainerTeamMembership**](portainer.TeamMembership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipDelete**
> TeamMembershipDelete(ctx, id)
Remove a team membership

Remove a team membership. Access is only available to administrators leaders of the associated team. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| TeamMembership identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipList**
> []PortainerTeamMembership TeamMembershipList(ctx, )
List team memberships

List team memberships. Access is only available to administrators and team leaders. **Access policy**: administrator

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PortainerTeamMembership**](portainer.TeamMembership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipUpdate**
> PortainerTeamMembership TeamMembershipUpdate(ctx, id, body)
Update a team membership

Update a team membership. Access is only available to administrators or leaders of the associated team. **Access policy**: administrator or leaders of the associated team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team membership identifier | 
  **body** | [**TeammembershipsTeamMembershipUpdatePayload**](TeammembershipsTeamMembershipUpdatePayload.md)| Team membership details | 

### Return type

[**PortainerTeamMembership**](portainer.TeamMembership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMemberships**
> []PortainerTeamMembership TeamMemberships(ctx, id)
List team memberships

List team memberships. Access is only available to administrators and team leaders. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| Team Id | 

### Return type

[**[]PortainerTeamMembership**](portainer.TeamMembership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

