# PortainerStack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalFiles** | **[]string** | Only applies when deploying stack with multiple files | [optional] [default to null]
**AutoUpdate** | [***PortainerStackAutoUpdate**](portainer.StackAutoUpdate.md) | The auto update settings of a git stack | [optional] [default to null]
**EndpointId** | **int32** | Environment(Endpoint) identifier. Reference the environment(endpoint) that will be used for deployment | [optional] [default to null]
**EntryPoint** | **string** | Path to the Stack file | [optional] [default to null]
**Env** | [**[]PortainerPair**](portainer.Pair.md) | A list of environment(endpoint) variables used during stack deployment | [optional] [default to null]
**Id** | **int32** | Stack Identifier | [optional] [default to null]
**Name** | **string** | Stack name | [optional] [default to null]
**Option** | [***PortainerStackOption**](portainer.StackOption.md) | The stack deployment option | [optional] [default to null]
**ResourceControl** | [***PortainerResourceControl**](portainer.ResourceControl.md) |  | [optional] [default to null]
**Status** | **int32** | Stack status (1 - active, 2 - inactive) | [optional] [default to null]
**SwarmId** | **string** | Cluster identifier of the Swarm cluster where the stack is deployed | [optional] [default to null]
**Type_** | **int32** | Stack type. 1 for a Swarm stack, 2 for a Compose stack | [optional] [default to null]
**CreatedBy** | **string** | The username which created this stack | [optional] [default to null]
**CreationDate** | **int32** | The date in unix time when stack was created | [optional] [default to null]
**FromAppTemplate** | **bool** | Whether the stack is from a app template | [optional] [default to null]
**GitConfig** | [***GittypesRepoConfig**](gittypes.RepoConfig.md) | The git config of this stack | [optional] [default to null]
**IsComposeFormat** | **bool** | IsComposeFormat indicates if the Kubernetes stack is created from a Docker Compose file | [optional] [default to null]
**Namespace** | **string** | Kubernetes namespace if stack is a kube application | [optional] [default to null]
**ProjectPath** | **string** | Path on disk to the repository hosting the Stack file | [optional] [default to null]
**UpdateDate** | **int32** | The date in unix time when stack was last updated | [optional] [default to null]
**UpdatedBy** | **string** | The username which last updated this stack | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


