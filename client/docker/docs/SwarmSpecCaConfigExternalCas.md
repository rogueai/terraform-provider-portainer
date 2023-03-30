# SwarmSpecCaConfigExternalCas

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Protocol for communication with the external CA (currently only &#x60;cfssl&#x60; is supported).  | [optional] [default to null]
**URL** | **string** | URL where certificate signing requests should be sent.  | [optional] [default to null]
**Options** | **map[string]string** | An object with key/value pairs that are interpreted as protocol-specific options for the external CA driver.  | [optional] [default to null]
**CACert** | **string** | The root CA certificate (in PEM format) this external CA uses to issue TLS certificates (assumed to be to the current swarm root CA certificate if not provided).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


