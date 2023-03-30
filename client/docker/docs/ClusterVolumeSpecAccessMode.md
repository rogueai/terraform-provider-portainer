# ClusterVolumeSpecAccessMode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | The set of nodes this volume can be used on at one time. - &#x60;single&#x60; The volume may only be scheduled to one node at a time. - &#x60;multi&#x60; the volume may be scheduled to any supported number of nodes at a time.  | [optional] [default to null]
**Sharing** | **string** | The number and way that different tasks can use this volume at one time. - &#x60;none&#x60; The volume may only be used by one task at a time. - &#x60;readonly&#x60; The volume may be used by any number of tasks, but they all must mount the volume as readonly - &#x60;onewriter&#x60; The volume may be used by any number of tasks, but only one may mount it as read/write. - &#x60;all&#x60; The volume may have any number of readers and writers.  | [optional] [default to null]
**MountVolume** | **interface{}** | Options for using this volume as a Mount-type volume.      Either MountVolume or BlockVolume, but not both, must be     present.   properties:     FsType:       type: \&quot;string\&quot;       description: |         Specifies the filesystem type for the mount volume.         Optional.     MountFlags:       type: \&quot;array\&quot;       description: |         Flags to pass when mounting the volume. Optional.       items:         type: \&quot;string\&quot; BlockVolume:   type: \&quot;object\&quot;   description: |     Options for using this volume as a Block-type volume.     Intentionally empty.  | [optional] [default to null]
**Secrets** | [**[]ClusterVolumeSpecAccessModeSecrets**](ClusterVolumeSpec_AccessMode_Secrets.md) | Swarm Secrets that are passed to the CSI storage plugin when operating on this volume.  | [optional] [default to null]
**AccessibilityRequirements** | [***ClusterVolumeSpecAccessModeAccessibilityRequirements**](ClusterVolumeSpec_AccessMode_AccessibilityRequirements.md) |  | [optional] [default to null]
**CapacityRange** | [***ClusterVolumeSpecAccessModeCapacityRange**](ClusterVolumeSpec_AccessMode_CapacityRange.md) |  | [optional] [default to null]
**Availability** | **string** | The availability of the volume for use in tasks. - &#x60;active&#x60; The volume is fully available for scheduling on the cluster - &#x60;pause&#x60; No new workloads should use the volume, but existing workloads are not stopped. - &#x60;drain&#x60; All workloads using this volume should be stopped and rescheduled, and no new ones should be started.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


