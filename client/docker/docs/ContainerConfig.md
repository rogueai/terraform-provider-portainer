# ContainerConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The hostname to use for the container, as a valid RFC 1123 hostname.  | [optional] [default to null]
**Domainname** | **string** | The domain name to use for the container.  | [optional] [default to null]
**User** | **string** | The user that commands are run as inside the container. | [optional] [default to null]
**AttachStdin** | **bool** | Whether to attach to &#x60;stdin&#x60;. | [optional] [default to null]
**AttachStdout** | **bool** | Whether to attach to &#x60;stdout&#x60;. | [optional] [default to null]
**AttachStderr** | **bool** | Whether to attach to &#x60;stderr&#x60;. | [optional] [default to null]
**ExposedPorts** | **map[string]interface{}** | An object mapping ports to an empty object in the form:  &#x60;{\&quot;&lt;port&gt;/&lt;tcp|udp|sctp&gt;\&quot;: {}}&#x60;  | [optional] [default to null]
**Tty** | **bool** | Attach standard streams to a TTY, including &#x60;stdin&#x60; if it is not closed.  | [optional] [default to null]
**OpenStdin** | **bool** | Open &#x60;stdin&#x60; | [optional] [default to null]
**StdinOnce** | **bool** | Close &#x60;stdin&#x60; after one attached client disconnects | [optional] [default to null]
**Env** | **[]string** | A list of environment variables to set inside the container in the form &#x60;[\&quot;VAR&#x3D;value\&quot;, ...]&#x60;. A variable without &#x60;&#x3D;&#x60; is removed from the environment, rather than to have an empty value.  | [optional] [default to null]
**Cmd** | **[]string** | Command to run specified as a string or an array of strings.  | [optional] [default to null]
**Healthcheck** | [***HealthConfig**](HealthConfig.md) |  | [optional] [default to null]
**ArgsEscaped** | **bool** | Command is already escaped (Windows only) | [optional] [default to null]
**Image** | **string** | The name (or reference) of the image to use when creating the container, or which was used when the container was created.  | [optional] [default to null]
**Volumes** | **map[string]interface{}** | An object mapping mount point paths inside the container to empty objects.  | [optional] [default to null]
**WorkingDir** | **string** | The working directory for commands to run in. | [optional] [default to null]
**Entrypoint** | **[]string** | The entry point for the container as a string or an array of strings.  If the array consists of exactly one empty string (&#x60;[\&quot;\&quot;]&#x60;) then the entry point is reset to system default (i.e., the entry point used by docker when there is no &#x60;ENTRYPOINT&#x60; instruction in the &#x60;Dockerfile&#x60;).  | [optional] [default to null]
**NetworkDisabled** | **bool** | Disable networking for the container. | [optional] [default to null]
**MacAddress** | **string** | MAC address of the container. | [optional] [default to null]
**OnBuild** | **[]string** | &#x60;ONBUILD&#x60; metadata that were defined in the image&#39;s &#x60;Dockerfile&#x60;.  | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value metadata. | [optional] [default to null]
**StopSignal** | **string** | Signal to stop a container as a string or unsigned integer.  | [optional] [default to null]
**StopTimeout** | **int32** | Timeout to stop a container in seconds. | [optional] [default to null]
**Shell** | **[]string** | Shell for when &#x60;RUN&#x60;, &#x60;CMD&#x60;, and &#x60;ENTRYPOINT&#x60; uses a shell.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


