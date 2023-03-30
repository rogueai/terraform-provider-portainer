# SwarmInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeID** | **string** | Unique identifier of for this node in the swarm. | [optional] [default to null]
**NodeAddr** | **string** | IP address at which this node can be reached by other nodes in the swarm.  | [optional] [default to null]
**LocalNodeState** | [***LocalNodeState**](LocalNodeState.md) |  | [optional] [default to null]
**ControlAvailable** | **bool** |  | [optional] [default to null]
**Error_** | **string** |  | [optional] [default to null]
**RemoteManagers** | [**[]PeerNode**](PeerNode.md) | List of ID&#39;s and addresses of other managers in the swarm.  | [optional] [default to null]
**Nodes** | **int32** | Total number of nodes in the swarm. | [optional] [default to null]
**Managers** | **int32** | Total number of managers in the swarm. | [optional] [default to null]
**Cluster** | [***ClusterInfo**](ClusterInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


