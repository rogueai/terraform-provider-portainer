# \AuthApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateUser**](AuthApi.md#AuthenticateUser) | **Post** /auth | Authenticate
[**Logout**](AuthApi.md#Logout) | **Post** /auth/logout | Logout
[**ValidateOAuth**](AuthApi.md#ValidateOAuth) | **Post** /auth/oauth/validate | Authenticate with OAuth


# **AuthenticateUser**
> AuthAuthenticateResponse AuthenticateUser(ctx, body)
Authenticate

**Access policy**: public Use this environment(endpoint) to authenticate against Portainer using a username and password.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthAuthenticatePayload**](AuthAuthenticatePayload.md)| Credentials used for authentication | 

### Return type

[**AuthAuthenticateResponse**](auth.authenticateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Logout**
> Logout(ctx, )
Logout

**Access policy**: authenticated

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

# **ValidateOAuth**
> AuthAuthenticateResponse ValidateOAuth(ctx, body)
Authenticate with OAuth

**Access policy**: public

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthOauthPayload**](AuthOauthPayload.md)| OAuth Credentials used for authentication | 

### Return type

[**AuthAuthenticateResponse**](auth.authenticateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

