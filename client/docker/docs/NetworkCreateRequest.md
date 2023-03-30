# NetworkCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The network&#39;s name. | [default to null]
**CheckDuplicate** | **bool** | Check for networks with duplicate names. Since Network is primarily keyed based on a random ID and not on the name, and network name is strictly a user-friendly alias to the network which is uniquely identified using ID, there is no guaranteed way to check for duplicates. CheckDuplicate is there to provide a best effort checking of any networks which has the same name but it is not guaranteed to catch all name collisions.  | [optional] [default to null]
**Driver** | **string** | Name of the network driver plugin to use. | [optional] [default to null]
**Internal** | **bool** | Restrict external access to the network. | [optional] [default to null]
**Attachable** | **bool** | Globally scoped network is manually attachable by regular containers from workers in swarm mode.  | [optional] [default to null]
**Ingress** | **bool** | Ingress network is the network which provides the routing-mesh in swarm mode.  | [optional] [default to null]
**IPAM** | [***Ipam**](IPAM.md) | Optional custom IP scheme for the network. | [optional] [default to null]
**EnableIPv6** | **bool** | Enable IPv6 on the network. | [optional] [default to null]
**Options** | **map[string]string** | Network specific options to be used by the drivers. | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


