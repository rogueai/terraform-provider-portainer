# ResourcecontrolsResourceControlCreatePayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdministratorsOnly** | **bool** | Permit access to resource only to admins | [optional] [default to null]
**Public** | **bool** | Permit access to the associated resource to any user | [optional] [default to null]
**ResourceID** | **string** |  | [default to null]
**SubResourceIDs** | **[]string** | List of Docker resources that will inherit this access control | [optional] [default to null]
**Teams** | **[]int32** | List of team identifiers with access to the associated resource | [optional] [default to null]
**Type_** | **int32** | Type of Resource. Valid values are: 1 - container, 2 - service 3 - volume, 4 - network, 5 - secret, 6 - stack, 7 - config, 8 - custom template, 9 - azure-container-group | [default to null]
**Users** | **[]int32** | List of user identifiers with access to the associated resource | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


