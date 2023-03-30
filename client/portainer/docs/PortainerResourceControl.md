# PortainerResourceControl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | **int32** |  | [optional] [default to null]
**AdministratorsOnly** | **bool** | Permit access to resource only to admins | [optional] [default to null]
**Id** | **int32** | ResourceControl Identifier | [optional] [default to null]
**OwnerId** | **int32** | Deprecated fields Deprecated in DBVersion &#x3D;&#x3D; 2 | [optional] [default to null]
**Public** | **bool** | Permit access to the associated resource to any user | [optional] [default to null]
**ResourceId** | **string** | Docker resource identifier on which access control will be applied.\\ In the case of a resource control applied to a stack, use the stack name as identifier | [optional] [default to null]
**SubResourceIds** | **[]string** | List of Docker resources that will inherit this access control | [optional] [default to null]
**System** | **bool** |  | [optional] [default to null]
**TeamAccesses** | [**[]PortainerTeamResourceAccess**](portainer.TeamResourceAccess.md) |  | [optional] [default to null]
**Type_** | **int32** | Type of Docker resource. Valid values are: 1- container, 2 -service 3 - volume, 4 - secret, 5 - stack, 6 - config or 7 - custom template | [optional] [default to null]
**UserAccesses** | [**[]PortainerUserResourceAccess**](portainer.UserResourceAccess.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


