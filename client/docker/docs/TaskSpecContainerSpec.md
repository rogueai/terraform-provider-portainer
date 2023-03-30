# TaskSpecContainerSpec

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Image** | **string** | The image name to use for the container | [optional] [default to null]
**Labels** | **map[string]string** | User-defined key/value data. | [optional] [default to null]
**Command** | **[]string** | The command to be run in the image. | [optional] [default to null]
**Args** | **[]string** | Arguments to the command. | [optional] [default to null]
**Hostname** | **string** | The hostname to use for the container, as a valid [RFC 1123](https://tools.ietf.org/html/rfc1123) hostname.  | [optional] [default to null]
**Env** | **[]string** | A list of environment variables in the form &#x60;VAR&#x3D;value&#x60;.  | [optional] [default to null]
**Dir** | **string** | The working directory for commands to run in. | [optional] [default to null]
**User** | **string** | The user inside the container. | [optional] [default to null]
**Groups** | **[]string** | A list of additional groups that the container process will run as.  | [optional] [default to null]
**Privileges** | [***TaskSpecContainerSpecPrivileges**](TaskSpec_ContainerSpec_Privileges.md) |  | [optional] [default to null]
**TTY** | **bool** | Whether a pseudo-TTY should be allocated. | [optional] [default to null]
**OpenStdin** | **bool** | Open &#x60;stdin&#x60; | [optional] [default to null]
**ReadOnly** | **bool** | Mount the container&#39;s root filesystem as read only. | [optional] [default to null]
**Mounts** | [**[]Mount**](Mount.md) | Specification for mounts to be added to containers created as part of the service.  | [optional] [default to null]
**StopSignal** | **string** | Signal to stop the container. | [optional] [default to null]
**StopGracePeriod** | **int64** | Amount of time to wait for the container to terminate before forcefully killing it.  | [optional] [default to null]
**HealthCheck** | [***HealthConfig**](HealthConfig.md) |  | [optional] [default to null]
**Hosts** | **[]string** | A list of hostname/IP mappings to add to the container&#39;s &#x60;hosts&#x60; file. The format of extra hosts is specified in the [hosts(5)](http://man7.org/linux/man-pages/man5/hosts.5.html) man page:      IP_address canonical_hostname [aliases...]  | [optional] [default to null]
**DNSConfig** | [***TaskSpecContainerSpecDnsConfig**](TaskSpec_ContainerSpec_DNSConfig.md) |  | [optional] [default to null]
**Secrets** | [**[]TaskSpecContainerSpecSecrets**](TaskSpec_ContainerSpec_Secrets.md) | Secrets contains references to zero or more secrets that will be exposed to the service.  | [optional] [default to null]
**Configs** | [**[]TaskSpecContainerSpecConfigs**](TaskSpec_ContainerSpec_Configs.md) | Configs contains references to zero or more configs that will be exposed to the service.  | [optional] [default to null]
**Isolation** | **string** | Isolation technology of the containers running the service. (Windows only)  | [optional] [default to null]
**Init** | **bool** | Run an init inside the container that forwards signals and reaps processes. This field is omitted if empty, and the default (as configured on the daemon) is used.  | [optional] [default to null]
**Sysctls** | **map[string]string** | Set kernel namedspaced parameters (sysctls) in the container. The Sysctls option on services accepts the same sysctls as the are supported on containers. Note that while the same sysctls are supported, no guarantees or checks are made about their suitability for a clustered environment, and it&#39;s up to the user to determine whether a given sysctl will work properly in a Service.  | [optional] [default to null]
**CapabilityAdd** | **[]string** | A list of kernel capabilities to add to the default set for the container.  | [optional] [default to null]
**CapabilityDrop** | **[]string** | A list of kernel capabilities to drop from the default set for the container.  | [optional] [default to null]
**Ulimits** | [**[]ResourcesUlimits**](Resources_Ulimits.md) | A list of resource limits to set in the container. For example: &#x60;{\&quot;Name\&quot;: \&quot;nofile\&quot;, \&quot;Soft\&quot;: 1024, \&quot;Hard\&quot;: 2048}&#x60;\&quot;  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


