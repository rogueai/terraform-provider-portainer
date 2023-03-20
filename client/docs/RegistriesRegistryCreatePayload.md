# RegistriesRegistryCreatePayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **bool** | Is authentication against this registry enabled | [default to null]
**BaseURL** | **string** | BaseURL required for ProGet registry | [optional] [default to null]
**Ecr** | [***PortainerEcrData**](portainer.EcrData.md) | ECR specific details, required when type &#x3D; 7 | [optional] [default to null]
**Gitlab** | [***PortainerGitlabRegistryData**](portainer.GitlabRegistryData.md) | Gitlab specific details, required when type &#x3D; 4 | [optional] [default to null]
**Name** | **string** | Name that will be used to identify this registry | [default to null]
**Password** | **string** | Password used to authenticate against this registry. required when Authentication is true | [optional] [default to null]
**Quay** | [***PortainerQuayRegistryData**](portainer.QuayRegistryData.md) | Quay specific details, required when type &#x3D; 1 | [optional] [default to null]
**Type_** | **int32** | Registry Type. Valid values are:  1 (Quay.io),  2 (Azure container registry),  3 (custom registry),  4 (Gitlab registry),  5 (ProGet registry),  6 (DockerHub)  7 (ECR) | [default to null]
**Url** | **string** | URL or IP address of the Docker registry | [default to null]
**Username** | **string** | Username used to authenticate against this registry. Required when Authentication is true | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


