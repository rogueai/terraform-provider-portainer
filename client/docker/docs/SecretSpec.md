# SecretSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User-defined name of the secret. | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]
**Data** | **string** | Base64-url-safe-encoded ([RFC 4648](https://tools.ietf.org/html/rfc4648#section-5)) data to store as secret.  This field is only used to _create_ a secret, and is not returned by other endpoints.  | [optional] [default to null]
**Driver** | [***Driver**](Driver.md) | Name of the secrets driver used to fetch the secret&#39;s value from an external secret store.  | [optional] [default to null]
**Templating** | [***Driver**](Driver.md) | Templating driver, if applicable  Templating controls whether and how to evaluate the config payload as a template. If no driver is set, no templating is used.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


