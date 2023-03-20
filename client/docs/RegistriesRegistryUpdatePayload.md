# RegistriesRegistryUpdatePayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **bool** | Is authentication against this registry enabled | [default to null]
**BaseURL** | **string** | BaseURL is used for quay registry | [optional] [default to null]
**Ecr** | [***PortainerEcrData**](portainer.EcrData.md) | ECR data | [optional] [default to null]
**Name** | **string** | Name that will be used to identify this registry | [default to null]
**Password** | **string** | Password used to authenticate against this registry. required when Authentication is true | [optional] [default to null]
**Quay** | [***PortainerQuayRegistryData**](portainer.QuayRegistryData.md) | Quay data | [optional] [default to null]
**RegistryAccesses** | [***PortainerRegistryAccesses**](portainer.RegistryAccesses.md) | Registry access control | [optional] [default to null]
**Url** | **string** | URL or IP address of the Docker registry | [default to null]
**Username** | **string** | Username used to authenticate against this registry. Required when Authentication is true | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


