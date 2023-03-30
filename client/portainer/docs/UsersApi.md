# \UsersApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAdminCheck**](UsersApi.md#UserAdminCheck) | **Get** /users/admin/check | Check administrator account existence
[**UserAdminInit**](UsersApi.md#UserAdminInit) | **Post** /users/admin/init | Initialize administrator account
[**UserCreate**](UsersApi.md#UserCreate) | **Post** /users | Create a new user
[**UserDelete**](UsersApi.md#UserDelete) | **Delete** /users/{id} | Remove a user
[**UserGenerateAPIKey**](UsersApi.md#UserGenerateAPIKey) | **Post** /users/{id}/tokens | Generate an API key for a user
[**UserGetAPIKeys**](UsersApi.md#UserGetAPIKeys) | **Get** /users/{id}/tokens | Get all API keys for a user
[**UserInspect**](UsersApi.md#UserInspect) | **Get** /users/{id} | Inspect a user
[**UserList**](UsersApi.md#UserList) | **Get** /users | List users
[**UserMembershipsInspect**](UsersApi.md#UserMembershipsInspect) | **Get** /users/{id}/memberships | Inspect a user memberships
[**UserRemoveAPIKey**](UsersApi.md#UserRemoveAPIKey) | **Delete** /users/{id}/tokens/{keyID} | Remove an api-key associated to a user
[**UserUpdate**](UsersApi.md#UserUpdate) | **Put** /users/{id} | Update a user
[**UserUpdatePassword**](UsersApi.md#UserUpdatePassword) | **Put** /users/{id}/passwd | Update password for a user


# **UserAdminCheck**
> UserAdminCheck(ctx, )
Check administrator account existence

Check if an administrator account exists in the database. **Access policy**: public

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserAdminInit**
> PortainerUser UserAdminInit(ctx, body)
Initialize administrator account

Initialize the 'admin' user account. **Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersAdminInitPayload**](UsersAdminInitPayload.md)| User details | 

### Return type

[**PortainerUser**](portainer.User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserCreate**
> PortainerUser UserCreate(ctx, body)
Create a new user

Create a new Portainer user. Only administrators can create users. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UsersUserCreatePayload**](UsersUserCreatePayload.md)| User details | 

### Return type

[**PortainerUser**](portainer.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserDelete**
> UserDelete(ctx, id)
Remove a user

Remove a user. **Access policy**: administrator

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGenerateAPIKey**
> UsersAccessTokenResponse UserGenerateAPIKey(ctx, id, body)
Generate an API key for a user

Generates an API key for a user. Only the calling user can generate a token for themselves. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 
  **body** | [**UsersUserAccessTokenCreatePayload**](UsersUserAccessTokenCreatePayload.md)| details | 

### Return type

[**UsersAccessTokenResponse**](users.accessTokenResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGetAPIKeys**
> []PortainerApiKey UserGetAPIKeys(ctx, id)
Get all API keys for a user

Gets all API keys for a user. Only the calling user or admin can retrieve api-keys. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

[**[]PortainerApiKey**](portainer.APIKey.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserInspect**
> PortainerUser UserInspect(ctx, id)
Inspect a user

Retrieve details about a user. User passwords are filtered out, and should never be accessible. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

[**PortainerUser**](portainer.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserList**
> []PortainerUser UserList(ctx, optional)
List users

List Portainer users. Non-administrator users will only be able to list other non-administrator user accounts. User passwords are filtered out, and should never be accessible. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUserListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUserListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentId** | **optional.Int32**| Identifier of the environment(endpoint) that will be used to filter the authorized users | 

### Return type

[**[]PortainerUser**](portainer.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserMembershipsInspect**
> PortainerTeamMembership UserMembershipsInspect(ctx, id)
Inspect a user memberships

Inspect a user memberships. **Access policy**: restricted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

[**PortainerTeamMembership**](portainer.TeamMembership.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRemoveAPIKey**
> UserRemoveAPIKey(ctx, id, keyID)
Remove an api-key associated to a user

Remove an api-key associated to a user.. Only the calling user or admin can remove api-key. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 
  **keyID** | **int32**| Api Key identifier | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUpdate**
> PortainerUser UserUpdate(ctx, id, body)
Update a user

Update user details. A regular user account can only update his details. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 
  **body** | [**UsersUserUpdatePayload**](UsersUserUpdatePayload.md)| User details | 

### Return type

[**PortainerUser**](portainer.User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUpdatePassword**
> UserUpdatePassword(ctx, id, body)
Update password for a user

Update password for the specified user. **Access policy**: authenticated

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| identifier | 
  **body** | [**UsersUserUpdatePasswordPayload**](UsersUserUpdatePasswordPayload.md)| details | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth), [jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

