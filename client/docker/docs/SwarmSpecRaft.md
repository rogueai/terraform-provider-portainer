# SwarmSpecRaft

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotInterval** | **int32** | The number of log entries between snapshots. | [optional] [default to null]
**KeepOldSnapshots** | **int32** | The number of snapshots to keep beyond the current snapshot.  | [optional] [default to null]
**LogEntriesForSlowFollowers** | **int32** | The number of log entries to keep around to sync up slow followers after a snapshot is created.  | [optional] [default to null]
**ElectionTick** | **int32** | The number of ticks that a follower will wait for a message from the leader before becoming a candidate and starting an election. &#x60;ElectionTick&#x60; must be greater than &#x60;HeartbeatTick&#x60;.  A tick currently defaults to one second, so these translate directly to seconds currently, but this is NOT guaranteed.  | [optional] [default to null]
**HeartbeatTick** | **int32** | The number of ticks between heartbeats. Every HeartbeatTick ticks, the leader will send a heartbeat to the followers.  A tick currently defaults to one second, so these translate directly to seconds currently, but this is NOT guaranteed.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


