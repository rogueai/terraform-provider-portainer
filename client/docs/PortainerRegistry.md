# PortainerRegistry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Stores temporary access token | [optional] [default to null]
**AccessTokenExpiry** | **int32** |  | [optional] [default to null]
**Authentication** | **bool** | Is authentication against this registry enabled | [optional] [default to null]
**AuthorizedTeams** | **[]int32** | Deprecated in DBVersion &#x3D;&#x3D; 18 | [optional] [default to null]
**AuthorizedUsers** | **[]int32** | Deprecated in DBVersion &#x3D;&#x3D; 18 | [optional] [default to null]
**BaseURL** | **string** | Base URL, introduced for ProGet registry | [optional] [default to null]
**Ecr** | [***PortainerEcrData**](portainer.EcrData.md) |  | [optional] [default to null]
**Gitlab** | [***PortainerGitlabRegistryData**](portainer.GitlabRegistryData.md) |  | [optional] [default to null]
**Id** | **int32** | Registry Identifier | [optional] [default to null]
**ManagementConfiguration** | [***PortainerRegistryManagementConfiguration**](portainer.RegistryManagementConfiguration.md) |  | [optional] [default to null]
**Name** | **string** | Registry Name | [optional] [default to null]
**Password** | **string** | Password or SecretAccessKey used to authenticate against this registry | [optional] [default to null]
**Quay** | [***PortainerQuayRegistryData**](portainer.QuayRegistryData.md) |  | [optional] [default to null]
**RegistryAccesses** | [***PortainerRegistryAccesses**](portainer.RegistryAccesses.md) |  | [optional] [default to null]
**TeamAccessPolicies** | [***PortainerTeamAccessPolicies**](portainer.TeamAccessPolicies.md) | Deprecated in DBVersion &#x3D;&#x3D; 31 | [optional] [default to null]
**Type_** | **int32** | Registry Type (1 - Quay, 2 - Azure, 3 - Custom, 4 - Gitlab, 5 - ProGet, 6 - DockerHub, 7 - ECR) | [optional] [default to null]
**URL** | **string** | URL or IP address of the Docker registry | [optional] [default to null]
**UserAccessPolicies** | [***PortainerUserAccessPolicies**](portainer.UserAccessPolicies.md) | Deprecated fields Deprecated in DBVersion &#x3D;&#x3D; 31 | [optional] [default to null]
**Username** | **string** | Username or AccessKeyID used to authenticate against this registry | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


