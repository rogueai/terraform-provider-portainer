# ExecConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachStdin** | **bool** | Attach to &#x60;stdin&#x60; of the exec command. | [optional] [default to null]
**AttachStdout** | **bool** | Attach to &#x60;stdout&#x60; of the exec command. | [optional] [default to null]
**AttachStderr** | **bool** | Attach to &#x60;stderr&#x60; of the exec command. | [optional] [default to null]
**ConsoleSize** | **[]int32** | Initial console size, as an &#x60;[height, width]&#x60; array. | [optional] [default to null]
**DetachKeys** | **string** | Override the key sequence for detaching a container. Format is a single character &#x60;[a-Z]&#x60; or &#x60;ctrl-&lt;value&gt;&#x60; where &#x60;&lt;value&gt;&#x60; is one of: &#x60;a-z&#x60;, &#x60;@&#x60;, &#x60;^&#x60;, &#x60;[&#x60;, &#x60;,&#x60; or &#x60;_&#x60;.  | [optional] [default to null]
**Tty** | **bool** | Allocate a pseudo-TTY. | [optional] [default to null]
**Env** | **[]string** | A list of environment variables in the form &#x60;[\&quot;VAR&#x3D;value\&quot;, ...]&#x60;.  | [optional] [default to null]
**Cmd** | **[]string** | Command to run, as a string or array of strings. | [optional] [default to null]
**Privileged** | **bool** | Runs the exec process with extended privileges. | [optional] [default to null]
**User** | **string** | The user, and optionally, group to run the exec process inside the container. Format is one of: &#x60;user&#x60;, &#x60;user:group&#x60;, &#x60;uid&#x60;, or &#x60;uid:gid&#x60;.  | [optional] [default to null]
**WorkingDir** | **string** | The working directory for the exec process inside the container.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


