# SwarmSpecCaConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCertExpiry** | **int64** | The duration node certificates are issued for. | [optional] [default to null]
**ExternalCAs** | [**[]SwarmSpecCaConfigExternalCas**](SwarmSpec_CAConfig_ExternalCAs.md) | Configuration for forwarding signing requests to an external certificate authority.  | [optional] [default to null]
**SigningCACert** | **string** | The desired signing CA certificate for all swarm node TLS leaf certificates, in PEM format.  | [optional] [default to null]
**SigningCAKey** | **string** | The desired signing CA key for all swarm node TLS leaf certificates, in PEM format.  | [optional] [default to null]
**ForceRotate** | **int32** | An integer whose purpose is to force swarm to generate a new signing CA certificate and key, if none have been specified in &#x60;SigningCACert&#x60; and &#x60;SigningCAKey&#x60;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


