# EndpointSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IPAMConfig** | [***EndpointIpamConfig**](EndpointIPAMConfig.md) |  | [optional] [default to null]
**Links** | **[]string** |  | [optional] [default to null]
**Aliases** | **[]string** |  | [optional] [default to null]
**NetworkID** | **string** | Unique ID of the network.  | [optional] [default to null]
**EndpointID** | **string** | Unique ID for the service endpoint in a Sandbox.  | [optional] [default to null]
**Gateway** | **string** | Gateway address for this network.  | [optional] [default to null]
**IPAddress** | **string** | IPv4 address.  | [optional] [default to null]
**IPPrefixLen** | **int32** | Mask length of the IPv4 address.  | [optional] [default to null]
**IPv6Gateway** | **string** | IPv6 gateway address.  | [optional] [default to null]
**GlobalIPv6Address** | **string** | Global IPv6 address.  | [optional] [default to null]
**GlobalIPv6PrefixLen** | **int64** | Mask length of the global IPv6 address.  | [optional] [default to null]
**MacAddress** | **string** | MAC address for the endpoint on this network.  | [optional] [default to null]
**DriverOpts** | **map[string]string** | DriverOpts is a mapping of driver options and values. These options are passed directly to the driver and are driver specific.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


