# RegistriesRegistryConfigurePayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authentication** | **bool** | Is authentication against this registry enabled | [default to null]
**Password** | **string** | Password used to authenticate against this registry. required when Authentication is true | [optional] [default to null]
**Region** | **string** | ECR region | [optional] [default to null]
**Tls** | **bool** | Use TLS | [optional] [default to null]
**TlscacertFile** | **[]int32** | The TLS CA certificate file | [optional] [default to null]
**TlscertFile** | **[]int32** | The TLS client certificate file | [optional] [default to null]
**TlskeyFile** | **[]int32** | The TLS client key file | [optional] [default to null]
**TlsskipVerify** | **bool** | Skip the verification of the server TLS certificate | [optional] [default to null]
**Username** | **string** | Username used to authenticate against this registry. Required when Authentication is true | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


