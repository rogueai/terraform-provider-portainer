# MountPoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | The mount type:  - &#x60;bind&#x60; a mount of a file or directory from the host into the container. - &#x60;volume&#x60; a docker volume with the given &#x60;Name&#x60;. - &#x60;tmpfs&#x60; a &#x60;tmpfs&#x60;. - &#x60;npipe&#x60; a named pipe from the host into the container. - &#x60;cluster&#x60; a Swarm cluster volume  | [optional] [default to null]
**Name** | **string** | Name is the name reference to the underlying data defined by &#x60;Source&#x60; e.g., the volume name.  | [optional] [default to null]
**Source** | **string** | Source location of the mount.  For volumes, this contains the storage location of the volume (within &#x60;/var/lib/docker/volumes/&#x60;). For bind-mounts, and &#x60;npipe&#x60;, this contains the source (host) part of the bind-mount. For &#x60;tmpfs&#x60; mount points, this field is empty.  | [optional] [default to null]
**Destination** | **string** | Destination is the path relative to the container root (&#x60;/&#x60;) where the &#x60;Source&#x60; is mounted inside the container.  | [optional] [default to null]
**Driver** | **string** | Driver is the volume driver used to create the volume (if it is a volume).  | [optional] [default to null]
**Mode** | **string** | Mode is a comma separated list of options supplied by the user when creating the bind/volume mount.  The default is platform-specific (&#x60;\&quot;z\&quot;&#x60; on Linux, empty on Windows).  | [optional] [default to null]
**RW** | **bool** | Whether the mount is mounted writable (read-write).  | [optional] [default to null]
**Propagation** | **string** | Propagation describes how mounts are propagated from the host into the mount point, and vice-versa. Refer to the [Linux kernel documentation](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt) for details. This field is not used on Windows.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


