# \ImageApi

All URIs are relative to *http://localhost/v1.42*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildPrune**](ImageApi.md#BuildPrune) | **Post** /build/prune | Delete builder cache
[**ImageBuild**](ImageApi.md#ImageBuild) | **Post** /build | Build an image
[**ImageCommit**](ImageApi.md#ImageCommit) | **Post** /commit | Create a new image from a container
[**ImageCreate**](ImageApi.md#ImageCreate) | **Post** /images/create | Create an image
[**ImageDelete**](ImageApi.md#ImageDelete) | **Delete** /images/{name} | Remove an image
[**ImageGet**](ImageApi.md#ImageGet) | **Get** /images/{name}/get | Export an image
[**ImageGetAll**](ImageApi.md#ImageGetAll) | **Get** /images/get | Export several images
[**ImageHistory**](ImageApi.md#ImageHistory) | **Get** /images/{name}/history | Get the history of an image
[**ImageInspect**](ImageApi.md#ImageInspect) | **Get** /images/{name}/json | Inspect an image
[**ImageList**](ImageApi.md#ImageList) | **Get** /images/json | List Images
[**ImageLoad**](ImageApi.md#ImageLoad) | **Post** /images/load | Import images
[**ImagePrune**](ImageApi.md#ImagePrune) | **Post** /images/prune | Delete unused images
[**ImagePush**](ImageApi.md#ImagePush) | **Post** /images/{name}/push | Push an image
[**ImageSearch**](ImageApi.md#ImageSearch) | **Get** /images/search | Search images
[**ImageTag**](ImageApi.md#ImageTag) | **Post** /images/{name}/tag | Tag an image


# **BuildPrune**
> BuildPruneResponse BuildPrune(ctx, optional)
Delete builder cache

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiBuildPruneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiBuildPruneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keepStorage** | **optional.Int64**| Amount of disk space in bytes to keep for cache | 
 **all** | **optional.Bool**| Remove all types of build cache | 
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the list of build cache objects.  Available filters:  - &#x60;until&#x3D;&lt;duration&gt;&#x60;: duration relative to daemon&#39;s time, during which build cache was not used, in Go&#39;s duration format (e.g., &#39;24h&#39;) - &#x60;id&#x3D;&lt;id&gt;&#x60; - &#x60;parent&#x3D;&lt;id&gt;&#x60; - &#x60;type&#x3D;&lt;string&gt;&#x60; - &#x60;description&#x3D;&lt;string&gt;&#x60; - &#x60;inuse&#x60; - &#x60;shared&#x60; - &#x60;private&#x60;  | 

### Return type

[**BuildPruneResponse**](BuildPruneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageBuild**
> ImageBuild(ctx, optional)
Build an image

Build an image from a tar archive with a `Dockerfile` in it.  The `Dockerfile` specifies how the image is built from the tar archive. It is typically in the archive's root, but can be at a different path or have a different name by specifying the `dockerfile` parameter. [See the `Dockerfile` reference for more information](/engine/reference/builder/).  The Docker daemon performs a preliminary validation of the `Dockerfile` before starting the build, and returns an error if the syntax is incorrect. After that, each instruction is run one-by-one until the ID of the new image is output.  The build is canceled if the client drops the connection by quitting or being killed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImageBuildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageBuildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inputStream** | **optional.String**| A tar archive compressed with one of the following algorithms: identity (no compression), gzip, bzip2, xz. | 
 **dockerfile** | **optional.String**| Path within the build context to the &#x60;Dockerfile&#x60;. This is ignored if &#x60;remote&#x60; is specified and points to an external &#x60;Dockerfile&#x60;. | [default to Dockerfile]
 **t** | **optional.String**| A name and optional tag to apply to the image in the &#x60;name:tag&#x60; format. If you omit the tag the default &#x60;latest&#x60; value is assumed. You can provide several &#x60;t&#x60; parameters. | 
 **extrahosts** | **optional.String**| Extra hosts to add to /etc/hosts | 
 **remote** | **optional.String**| A Git repository URI or HTTP/HTTPS context URI. If the URI points to a single text file, the file’s contents are placed into a file called &#x60;Dockerfile&#x60; and the image is built from that file. If the URI points to a tarball, the file is downloaded by the daemon and the contents therein used as the context for the build. If the URI points to a tarball and the &#x60;dockerfile&#x60; parameter is also specified, there must be a file with the corresponding path inside the tarball. | 
 **q** | **optional.Bool**| Suppress verbose build output. | [default to false]
 **nocache** | **optional.Bool**| Do not use the cache when building the image. | [default to false]
 **cachefrom** | **optional.String**| JSON array of images used for build cache resolution. | 
 **pull** | **optional.String**| Attempt to pull the image even if an older image exists locally. | 
 **rm** | **optional.Bool**| Remove intermediate containers after a successful build. | [default to true]
 **forcerm** | **optional.Bool**| Always remove intermediate containers, even upon failure. | [default to false]
 **memory** | **optional.Int32**| Set memory limit for build. | 
 **memswap** | **optional.Int32**| Total memory (memory + swap). Set as &#x60;-1&#x60; to disable swap. | 
 **cpushares** | **optional.Int32**| CPU shares (relative weight). | 
 **cpusetcpus** | **optional.String**| CPUs in which to allow execution (e.g., &#x60;0-3&#x60;, &#x60;0,1&#x60;). | 
 **cpuperiod** | **optional.Int32**| The length of a CPU period in microseconds. | 
 **cpuquota** | **optional.Int32**| Microseconds of CPU time that the container can get in a CPU period. | 
 **buildargs** | **optional.String**| JSON map of string pairs for build-time variables. Users pass these values at build-time. Docker uses the buildargs as the environment context for commands run via the &#x60;Dockerfile&#x60; RUN instruction, or for variable expansion in other &#x60;Dockerfile&#x60; instructions. This is not meant for passing secret values.  For example, the build arg &#x60;FOO&#x3D;bar&#x60; would become &#x60;{\&quot;FOO\&quot;:\&quot;bar\&quot;}&#x60; in JSON. This would result in the query parameter &#x60;buildargs&#x3D;{\&quot;FOO\&quot;:\&quot;bar\&quot;}&#x60;. Note that &#x60;{\&quot;FOO\&quot;:\&quot;bar\&quot;}&#x60; should be URI component encoded.  [Read more about the buildargs instruction.](/engine/reference/builder/#arg)  | 
 **shmsize** | **optional.Int32**| Size of &#x60;/dev/shm&#x60; in bytes. The size must be greater than 0. If omitted the system uses 64MB. | 
 **squash** | **optional.Bool**| Squash the resulting images layers into a single layer. *(Experimental release only.)* | 
 **labels** | **optional.String**| Arbitrary key/value labels to set on the image, as a JSON map of string pairs. | 
 **networkmode** | **optional.String**| Sets the networking mode for the run commands during build. Supported standard values are: &#x60;bridge&#x60;, &#x60;host&#x60;, &#x60;none&#x60;, and &#x60;container:&lt;name|id&gt;&#x60;. Any other value is taken as a custom network&#39;s name or ID to which this container should connect to.  | 
 **contentType** | **optional.String**|  | [default to application/x-tar]
 **xRegistryConfig** | **optional.String**| This is a base64-encoded JSON object with auth configurations for multiple registries that a build may refer to.  The key is a registry URL, and the value is an auth configuration object, [as described in the authentication section](#section/Authentication). For example:  &#x60;&#x60;&#x60; {   \&quot;docker.example.com\&quot;: {     \&quot;username\&quot;: \&quot;janedoe\&quot;,     \&quot;password\&quot;: \&quot;hunter2\&quot;   },   \&quot;https://index.docker.io/v1/\&quot;: {     \&quot;username\&quot;: \&quot;mobydock\&quot;,     \&quot;password\&quot;: \&quot;conta1n3rize14\&quot;   } } &#x60;&#x60;&#x60;  Only the registry domain name (and port if not the default 443) are required. However, for legacy reasons, the Docker Hub registry must be specified with both a &#x60;https://&#x60; prefix and a &#x60;/v1/&#x60; suffix even though Docker will prefer to use the v2 registry API.  | 
 **platform** | **optional.String**| Platform in the format os[/arch[/variant]] | [default to ]
 **target** | **optional.String**| Target build stage | [default to ]
 **outputs** | **optional.String**| BuildKit output configuration | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageCommit**
> IdResponse ImageCommit(ctx, optional)
Create a new image from a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImageCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageCommitOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerConfig** | [**optional.Interface of ContainerConfig**](ContainerConfig.md)| The container configuration | 
 **container** | **optional.String**| The ID or name of the container to commit | 
 **repo** | **optional.String**| Repository name for the created image | 
 **tag** | **optional.String**| Tag name for the create image | 
 **comment** | **optional.String**| Commit message | 
 **author** | **optional.String**| Author of the image (e.g., &#x60;John Hannibal Smith &lt;hannibal@a-team.com&gt;&#x60;) | 
 **pause** | **optional.Bool**| Whether to pause the container before committing | [default to true]
 **changes** | **optional.String**| &#x60;Dockerfile&#x60; instructions to apply while committing | 

### Return type

[**IdResponse**](IdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageCreate**
> ImageCreate(ctx, optional)
Create an image

Create an image by either pulling it from a registry or importing it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImageCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromImage** | **optional.String**| Name of the image to pull. The name may include a tag or digest. This parameter may only be used when pulling an image. The pull is cancelled if the HTTP connection is closed. | 
 **fromSrc** | **optional.String**| Source to import. The value may be a URL from which the image can be retrieved or &#x60;-&#x60; to read the image from the request body. This parameter may only be used when importing an image. | 
 **repo** | **optional.String**| Repository name given to an image when it is imported. The repo may include a tag. This parameter may only be used when importing an image. | 
 **tag** | **optional.String**| Tag or digest. If empty when pulling an image, this causes all tags for the given image to be pulled. | 
 **message** | **optional.String**| Set commit message for imported image. | 
 **inputImage** | **optional.String**| Image content if the value &#x60;-&#x60; has been specified in fromSrc query parameter | 
 **xRegistryAuth** | **optional.String**| A base64url-encoded auth configuration.  Refer to the [authentication section](#section/Authentication) for details.  | 
 **changes** | [**optional.Interface of []string**](string.md)| Apply &#x60;Dockerfile&#x60; instructions to the image that is created, for example: &#x60;changes&#x3D;ENV DEBUG&#x3D;true&#x60;. Note that &#x60;ENV DEBUG&#x3D;true&#x60; should be URI component encoded.  Supported &#x60;Dockerfile&#x60; instructions: &#x60;CMD&#x60;|&#x60;ENTRYPOINT&#x60;|&#x60;ENV&#x60;|&#x60;EXPOSE&#x60;|&#x60;ONBUILD&#x60;|&#x60;USER&#x60;|&#x60;VOLUME&#x60;|&#x60;WORKDIR&#x60;  | 
 **platform** | **optional.String**| Platform in the format os[/arch[/variant]].  When used in combination with the &#x60;fromImage&#x60; option, the daemon checks if the given image is present in the local image cache with the given OS and Architecture, and otherwise attempts to pull the image. If the option is not set, the host&#39;s native OS and Architecture are used. If the given image does not exist in the local image cache, the daemon attempts to pull the image with the host&#39;s native OS and Architecture. If the given image does exists in the local image cache, but its OS or architecture does not match, a warning is produced.  When used with the &#x60;fromSrc&#x60; option to import an image from an archive, this option sets the platform information for the imported image. If the option is not set, the host&#39;s native OS and Architecture are used for the imported image.  | [default to ]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: text/plain, application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageDelete**
> []ImageDeleteResponseItem ImageDelete(ctx, name, optional)
Remove an image

Remove an image, along with any untagged parent images that were referenced by that image.  Images can't be removed if they have descendant images, are being used by a running container or are being used by a build. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or ID | 
 **optional** | ***ImageApiImageDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| Remove the image even if it is being used by stopped containers or has other tags | [default to false]
 **noprune** | **optional.Bool**| Do not delete untagged parent images | [default to false]

### Return type

[**[]ImageDeleteResponseItem**](ImageDeleteResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageGet**
> string ImageGet(ctx, name)
Export an image

Get a tarball containing all images and metadata for a repository.  If `name` is a specific name and tag (e.g. `ubuntu:latest`), then only that image (and its parents) are returned. If `name` is an image ID, similarly only that image (and its parents) are returned, but with the exclusion of the `repositories` file in the tarball, as there were no image names referenced.  ### Image tarball format  An image tarball contains one directory per image layer (named using its long ID), each containing these files:  - `VERSION`: currently `1.0` - the file format version - `json`: detailed layer information, similar to `docker inspect layer_id` - `layer.tar`: A tarfile containing the filesystem changes in this layer  The `layer.tar` file contains `aufs` style `.wh..wh.aufs` files and directories for storing attribute changes and deletions.  If the tarball defines a repository, the tarball should also include a `repositories` file at the root that contains a list of repository and tag names mapped to layer IDs.  ```json {   \"hello-world\": {     \"latest\": \"565a9d68a73f6706862bfe8409a7f659776d4d60a8d096eb4a3cbce6999cc2a1\"   } } ``` 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or ID | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/x-tar

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageGetAll**
> string ImageGetAll(ctx, optional)
Export several images

Get a tarball containing all images and metadata for several image repositories.  For each value of the `names` parameter: if it is a specific name and tag (e.g. `ubuntu:latest`), then only that image (and its parents) are returned; if it is an image ID, similarly only that image (and its parents) are returned and there would be no names referenced in the 'repositories' file for this image ID.  For details on the format, see the [export image endpoint](#operation/ImageGet). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImageGetAllOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageGetAllOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **names** | [**optional.Interface of []string**](string.md)| Image names to filter by | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/x-tar

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageHistory**
> []HistoryResponseItem ImageHistory(ctx, name)
Get the history of an image

Return parent layers of an image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or ID | 

### Return type

[**[]HistoryResponseItem**](HistoryResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageInspect**
> ImageInspect ImageInspect(ctx, name)
Inspect an image

Return low-level information about an image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or id | 

### Return type

[**ImageInspect**](ImageInspect.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageList**
> []ImageSummary ImageList(ctx, optional)
List Images

Returns a list of images on the server. Note that it uses a different, smaller representation of an image than inspecting a single image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImageListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Show all images. Only images from a final layer (no children) are shown by default. | [default to false]
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the images list.  Available filters:  - &#x60;before&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;,  &#x60;&lt;image id&gt;&#x60; or &#x60;&lt;image@digest&gt;&#x60;) - &#x60;dangling&#x3D;true&#x60; - &#x60;label&#x3D;key&#x60; or &#x60;label&#x3D;\&quot;key&#x3D;value\&quot;&#x60; of an image label - &#x60;reference&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;) - &#x60;since&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;,  &#x60;&lt;image id&gt;&#x60; or &#x60;&lt;image@digest&gt;&#x60;)  | 
 **sharedSize** | **optional.Bool**| Compute and show shared size as a &#x60;SharedSize&#x60; field on each image. | [default to false]
 **digests** | **optional.Bool**| Show digest information as a &#x60;RepoDigests&#x60; field on each image. | [default to false]

### Return type

[**[]ImageSummary**](ImageSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageLoad**
> ImageLoad(ctx, optional)
Import images

Load a set of images and tags into a repository.  For details on the format, see the [export image endpoint](#operation/ImageGet). 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImageLoadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageLoadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **imagesTarball** | **optional.String**| Tar archive containing images | 
 **quiet** | **optional.Bool**| Suppress progress details during load. | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagePrune**
> ImagePruneResponse ImagePrune(ctx, optional)
Delete unused images

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImageApiImagePruneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImagePruneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filters to process on the prune list, encoded as JSON (a &#x60;map[string][]string&#x60;). Available filters:  - &#x60;dangling&#x3D;&lt;boolean&gt;&#x60; When set to &#x60;true&#x60; (or &#x60;1&#x60;), prune only    unused *and* untagged images. When set to &#x60;false&#x60;    (or &#x60;0&#x60;), all unused images are pruned. - &#x60;until&#x3D;&lt;string&gt;&#x60; Prune images created before this timestamp. The &#x60;&lt;timestamp&gt;&#x60; can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. &#x60;10m&#x60;, &#x60;1h30m&#x60;) computed relative to the daemon machine’s time. - &#x60;label&#x60; (&#x60;label&#x3D;&lt;key&gt;&#x60;, &#x60;label&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;, &#x60;label!&#x3D;&lt;key&gt;&#x60;, or &#x60;label!&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;) Prune images with (or without, in case &#x60;label!&#x3D;...&#x60; is used) the specified labels.  | 

### Return type

[**ImagePruneResponse**](ImagePruneResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagePush**
> ImagePush(ctx, name, xRegistryAuth, optional)
Push an image

Push an image to a registry.  If you wish to push an image on to a private registry, that image must already have a tag which references the registry. For example, `registry.example.com/myimage:latest`.  The push is cancelled if the HTTP connection is closed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or ID. | 
  **xRegistryAuth** | **string**| A base64url-encoded auth configuration.  Refer to the [authentication section](#section/Authentication) for details.  | 
 **optional** | ***ImageApiImagePushOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImagePushOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **optional.String**| The tag to associate with the image on the registry. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageSearch**
> []ImageSearchResponseItem ImageSearch(ctx, term, optional)
Search images

Search for an image on Docker Hub.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **term** | **string**| Term to search | 
 **optional** | ***ImageApiImageSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Maximum number of results to return | 
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the images list. Available filters:  - &#x60;is-automated&#x3D;(true|false)&#x60; - &#x60;is-official&#x3D;(true|false)&#x60; - &#x60;stars&#x3D;&lt;number&gt;&#x60; Matches images that has at least &#39;number&#39; stars.  | 

### Return type

[**[]ImageSearchResponseItem**](ImageSearchResponseItem.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageTag**
> ImageTag(ctx, name, optional)
Tag an image

Tag an image so that it becomes part of a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Image name or ID to tag. | 
 **optional** | ***ImageApiImageTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImageApiImageTagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repo** | **optional.String**| The repository to tag in. For example, &#x60;someuser/someimage&#x60;. | 
 **tag** | **optional.String**| The name of the new tag. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, text/plain
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

