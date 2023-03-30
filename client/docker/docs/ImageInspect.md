# ImageInspect

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID is the content-addressable ID of an image.  This identifier is a content-addressable digest calculated from the image&#39;s configuration (which includes the digests of layers used by the image).  Note that this digest differs from the &#x60;RepoDigests&#x60; below, which holds digests of image manifests that reference the image.  | [optional] [default to null]
**RepoTags** | **[]string** | List of image names/tags in the local image cache that reference this image.  Multiple image tags can refer to the same image, and this list may be empty if no tags reference the image, in which case the image is \&quot;untagged\&quot;, in which case it can still be referenced by its ID.  | [optional] [default to null]
**RepoDigests** | **[]string** | List of content-addressable digests of locally available image manifests that the image is referenced from. Multiple manifests can refer to the same image.  These digests are usually only available if the image was either pulled from a registry, or if the image was pushed to a registry, which is when the manifest is generated and its digest calculated.  | [optional] [default to null]
**Parent** | **string** | ID of the parent image.  Depending on how the image was created, this field may be empty and is only set for images that were built/created locally. This field is empty if the image was pulled from an image registry.  | [optional] [default to null]
**Comment** | **string** | Optional message that was set when committing or importing the image.  | [optional] [default to null]
**Created** | **string** | Date and time at which the image was created, formatted in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.  | [optional] [default to null]
**Container** | **string** | The ID of the container that was used to create the image.  Depending on how the image was created, this field may be empty.  | [optional] [default to null]
**ContainerConfig** | [***ContainerConfig**](ContainerConfig.md) |  | [optional] [default to null]
**DockerVersion** | **string** | The version of Docker that was used to build the image.  Depending on how the image was created, this field may be empty.  | [optional] [default to null]
**Author** | **string** | Name of the author that was specified when committing the image, or as specified through MAINTAINER (deprecated) in the Dockerfile.  | [optional] [default to null]
**Config** | [***ContainerConfig**](ContainerConfig.md) |  | [optional] [default to null]
**Architecture** | **string** | Hardware CPU architecture that the image runs on.  | [optional] [default to null]
**Variant** | **string** | CPU architecture variant (presently ARM-only).  | [optional] [default to null]
**Os** | **string** | Operating System the image is built to run on.  | [optional] [default to null]
**OsVersion** | **string** | Operating System version the image is built to run on (especially for Windows).  | [optional] [default to null]
**Size** | **int64** | Total size of the image including all layers it is composed of.  | [optional] [default to null]
**VirtualSize** | **int64** | Total size of the image including all layers it is composed of.  In versions of Docker before v1.10, this field was calculated from the image itself and all of its parent images. Docker v1.10 and up store images self-contained, and no longer use a parent-chain, making this field an equivalent of the Size field.  This field is kept for backward compatibility, but may be removed in a future version of the API.  | [optional] [default to null]
**GraphDriver** | [***GraphDriverData**](GraphDriverData.md) |  | [optional] [default to null]
**RootFS** | [***ImageInspectRootFs**](ImageInspect_RootFS.md) |  | [optional] [default to null]
**Metadata** | [***ImageInspectMetadata**](ImageInspect_Metadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


