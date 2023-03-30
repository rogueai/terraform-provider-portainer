# StacksUpdateSwarmStackPayload

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Env** | [**[]PortainerPair**](portainer.Pair.md) | A list of environment(endpoint) variables used during stack deployment | [optional] [default to null]
**Prune** | **bool** | Prune services that are no longer referenced (only available for Swarm stacks) | [optional] [default to null]
**PullImage** | **bool** | Force a pulling to current image with the original tag though the image is already the latest | [optional] [default to null]
**StackFileContent** | **string** | New content of the Stack file | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


