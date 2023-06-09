# HostConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuShares** | **int32** | An integer value representing this container&#39;s relative CPU weight versus other containers.  | [optional] [default to null]
**Memory** | **int64** | Memory limit in bytes. | [optional] [default to 0]
**CgroupParent** | **string** | Path to &#x60;cgroups&#x60; under which the container&#39;s &#x60;cgroup&#x60; is created. If the path is not absolute, the path is considered to be relative to the &#x60;cgroups&#x60; path of the init process. Cgroups are created if they do not already exist.  | [optional] [default to null]
**BlkioWeight** | **int32** | Block IO weight (relative weight). | [optional] [default to null]
**BlkioWeightDevice** | [**[]ResourcesBlkioWeightDevice**](Resources_BlkioWeightDevice.md) | Block IO weight (relative device weight) in the form:  &#x60;&#x60;&#x60; [{\&quot;Path\&quot;: \&quot;device_path\&quot;, \&quot;Weight\&quot;: weight}] &#x60;&#x60;&#x60;  | [optional] [default to null]
**BlkioDeviceReadBps** | [**[]ThrottleDevice**](ThrottleDevice.md) | Limit read rate (bytes per second) from a device, in the form:  &#x60;&#x60;&#x60; [{\&quot;Path\&quot;: \&quot;device_path\&quot;, \&quot;Rate\&quot;: rate}] &#x60;&#x60;&#x60;  | [optional] [default to null]
**BlkioDeviceWriteBps** | [**[]ThrottleDevice**](ThrottleDevice.md) | Limit write rate (bytes per second) to a device, in the form:  &#x60;&#x60;&#x60; [{\&quot;Path\&quot;: \&quot;device_path\&quot;, \&quot;Rate\&quot;: rate}] &#x60;&#x60;&#x60;  | [optional] [default to null]
**BlkioDeviceReadIOps** | [**[]ThrottleDevice**](ThrottleDevice.md) | Limit read rate (IO per second) from a device, in the form:  &#x60;&#x60;&#x60; [{\&quot;Path\&quot;: \&quot;device_path\&quot;, \&quot;Rate\&quot;: rate}] &#x60;&#x60;&#x60;  | [optional] [default to null]
**BlkioDeviceWriteIOps** | [**[]ThrottleDevice**](ThrottleDevice.md) | Limit write rate (IO per second) to a device, in the form:  &#x60;&#x60;&#x60; [{\&quot;Path\&quot;: \&quot;device_path\&quot;, \&quot;Rate\&quot;: rate}] &#x60;&#x60;&#x60;  | [optional] [default to null]
**CpuPeriod** | **int64** | The length of a CPU period in microseconds. | [optional] [default to null]
**CpuQuota** | **int64** | Microseconds of CPU time that the container can get in a CPU period.  | [optional] [default to null]
**CpuRealtimePeriod** | **int64** | The length of a CPU real-time period in microseconds. Set to 0 to allocate no time allocated to real-time tasks.  | [optional] [default to null]
**CpuRealtimeRuntime** | **int64** | The length of a CPU real-time runtime in microseconds. Set to 0 to allocate no time allocated to real-time tasks.  | [optional] [default to null]
**CpusetCpus** | **string** | CPUs in which to allow execution (e.g., &#x60;0-3&#x60;, &#x60;0,1&#x60;).  | [optional] [default to null]
**CpusetMems** | **string** | Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.  | [optional] [default to null]
**Devices** | [**[]DeviceMapping**](DeviceMapping.md) | A list of devices to add to the container. | [optional] [default to null]
**DeviceCgroupRules** | **[]string** | a list of cgroup rules to apply to the container | [optional] [default to null]
**DeviceRequests** | [**[]DeviceRequest**](DeviceRequest.md) | A list of requests for devices to be sent to device drivers.  | [optional] [default to null]
**KernelMemoryTCP** | **int64** | Hard limit for kernel TCP buffer memory (in bytes). Depending on the OCI runtime in use, this option may be ignored. It is no longer supported by the default (runc) runtime.  This field is omitted when empty.  | [optional] [default to null]
**MemoryReservation** | **int64** | Memory soft limit in bytes. | [optional] [default to null]
**MemorySwap** | **int64** | Total memory limit (memory + swap). Set as &#x60;-1&#x60; to enable unlimited swap.  | [optional] [default to null]
**MemorySwappiness** | **int64** | Tune a container&#39;s memory swappiness behavior. Accepts an integer between 0 and 100.  | [optional] [default to null]
**NanoCpus** | **int64** | CPU quota in units of 10&lt;sup&gt;-9&lt;/sup&gt; CPUs. | [optional] [default to null]
**OomKillDisable** | **bool** | Disable OOM Killer for the container. | [optional] [default to null]
**Init** | **bool** | Run an init inside the container that forwards signals and reaps processes. This field is omitted if empty, and the default (as configured on the daemon) is used.  | [optional] [default to null]
**PidsLimit** | **int64** | Tune a container&#39;s PIDs limit. Set &#x60;0&#x60; or &#x60;-1&#x60; for unlimited, or &#x60;null&#x60; to not change.  | [optional] [default to null]
**Ulimits** | [**[]ResourcesUlimits**](Resources_Ulimits.md) | A list of resource limits to set in the container. For example:  &#x60;&#x60;&#x60; {\&quot;Name\&quot;: \&quot;nofile\&quot;, \&quot;Soft\&quot;: 1024, \&quot;Hard\&quot;: 2048} &#x60;&#x60;&#x60;  | [optional] [default to null]
**CpuCount** | **int64** | The number of usable CPUs (Windows only).  On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is &#x60;CPUCount&#x60; first, then &#x60;CPUShares&#x60;, and &#x60;CPUPercent&#x60; last.  | [optional] [default to null]
**CpuPercent** | **int64** | The usable percentage of the available CPUs (Windows only).  On Windows Server containers, the processor resource controls are mutually exclusive. The order of precedence is &#x60;CPUCount&#x60; first, then &#x60;CPUShares&#x60;, and &#x60;CPUPercent&#x60; last.  | [optional] [default to null]
**IOMaximumIOps** | **int64** | Maximum IOps for the container system drive (Windows only) | [optional] [default to null]
**IOMaximumBandwidth** | **int64** | Maximum IO in bytes per second for the container system drive (Windows only).  | [optional] [default to null]
**Binds** | **[]string** | A list of volume bindings for this container. Each volume binding is a string in one of these forms:  - &#x60;host-src:container-dest[:options]&#x60; to bind-mount a host path   into the container. Both &#x60;host-src&#x60;, and &#x60;container-dest&#x60; must   be an _absolute_ path. - &#x60;volume-name:container-dest[:options]&#x60; to bind-mount a volume   managed by a volume driver into the container. &#x60;container-dest&#x60;   must be an _absolute_ path.  &#x60;options&#x60; is an optional, comma-delimited list of:  - &#x60;nocopy&#x60; disables automatic copying of data from the container   path to the volume. The &#x60;nocopy&#x60; flag only applies to named volumes. - &#x60;[ro|rw]&#x60; mounts a volume read-only or read-write, respectively.   If omitted or set to &#x60;rw&#x60;, volumes are mounted read-write. - &#x60;[z|Z]&#x60; applies SELinux labels to allow or deny multiple containers   to read and write to the same volume.     - &#x60;z&#x60;: a _shared_ content label is applied to the content. This       label indicates that multiple containers can share the volume       content, for both reading and writing.     - &#x60;Z&#x60;: a _private unshared_ label is applied to the content.       This label indicates that only the current container can use       a private volume. Labeling systems such as SELinux require       proper labels to be placed on volume content that is mounted       into a container. Without a label, the security system can       prevent a container&#39;s processes from using the content. By       default, the labels set by the host operating system are not       modified. - &#x60;[[r]shared|[r]slave|[r]private]&#x60; specifies mount   [propagation behavior](https://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt).   This only applies to bind-mounted volumes, not internal volumes   or named volumes. Mount propagation requires the source mount   point (the location where the source directory is mounted in the   host operating system) to have the correct propagation properties.   For shared volumes, the source mount point must be set to &#x60;shared&#x60;.   For slave volumes, the mount must be set to either &#x60;shared&#x60; or   &#x60;slave&#x60;.  | [optional] [default to null]
**ContainerIDFile** | **string** | Path to a file where the container ID is written | [optional] [default to null]
**LogConfig** | [***HostConfigLogConfig**](HostConfig_LogConfig.md) |  | [optional] [default to null]
**NetworkMode** | **string** | Network mode to use for this container. Supported standard values are: &#x60;bridge&#x60;, &#x60;host&#x60;, &#x60;none&#x60;, and &#x60;container:&lt;name|id&gt;&#x60;. Any other value is taken as a custom network&#39;s name to which this container should connect to.  | [optional] [default to null]
**PortBindings** | [***PortMap**](PortMap.md) |  | [optional] [default to null]
**RestartPolicy** | [***RestartPolicy**](RestartPolicy.md) |  | [optional] [default to null]
**AutoRemove** | **bool** | Automatically remove the container when the container&#39;s process exits. This has no effect if &#x60;RestartPolicy&#x60; is set.  | [optional] [default to null]
**VolumeDriver** | **string** | Driver that this container uses to mount volumes. | [optional] [default to null]
**VolumesFrom** | **[]string** | A list of volumes to inherit from another container, specified in the form &#x60;&lt;container name&gt;[:&lt;ro|rw&gt;]&#x60;.  | [optional] [default to null]
**Mounts** | [**[]Mount**](Mount.md) | Specification for mounts to be added to the container.  | [optional] [default to null]
**ConsoleSize** | **[]int32** | Initial console size, as an &#x60;[height, width]&#x60; array.  | [optional] [default to null]
**CapAdd** | **[]string** | A list of kernel capabilities to add to the container. Conflicts with option &#39;Capabilities&#39;.  | [optional] [default to null]
**CapDrop** | **[]string** | A list of kernel capabilities to drop from the container. Conflicts with option &#39;Capabilities&#39;.  | [optional] [default to null]
**CgroupnsMode** | **string** | cgroup namespace mode for the container. Possible values are:  - &#x60;\&quot;private\&quot;&#x60;: the container runs in its own private cgroup namespace - &#x60;\&quot;host\&quot;&#x60;: use the host system&#39;s cgroup namespace  If not specified, the daemon default is used, which can either be &#x60;\&quot;private\&quot;&#x60; or &#x60;\&quot;host\&quot;&#x60;, depending on daemon version, kernel support and configuration.  | [optional] [default to null]
**Dns** | **[]string** | A list of DNS servers for the container to use. | [optional] [default to null]
**DnsOptions** | **[]string** | A list of DNS options. | [optional] [default to null]
**DnsSearch** | **[]string** | A list of DNS search domains. | [optional] [default to null]
**ExtraHosts** | **[]string** | A list of hostnames/IP mappings to add to the container&#39;s &#x60;/etc/hosts&#x60; file. Specified in the form &#x60;[\&quot;hostname:IP\&quot;]&#x60;.  | [optional] [default to null]
**GroupAdd** | **[]string** | A list of additional groups that the container process will run as.  | [optional] [default to null]
**IpcMode** | **string** | IPC sharing mode for the container. Possible values are:  - &#x60;\&quot;none\&quot;&#x60;: own private IPC namespace, with /dev/shm not mounted - &#x60;\&quot;private\&quot;&#x60;: own private IPC namespace - &#x60;\&quot;shareable\&quot;&#x60;: own private IPC namespace, with a possibility to share it with other containers - &#x60;\&quot;container:&lt;name|id&gt;\&quot;&#x60;: join another (shareable) container&#39;s IPC namespace - &#x60;\&quot;host\&quot;&#x60;: use the host system&#39;s IPC namespace  If not specified, daemon default is used, which can either be &#x60;\&quot;private\&quot;&#x60; or &#x60;\&quot;shareable\&quot;&#x60;, depending on daemon version and configuration.  | [optional] [default to null]
**Cgroup** | **string** | Cgroup to use for the container. | [optional] [default to null]
**Links** | **[]string** | A list of links for the container in the form &#x60;container_name:alias&#x60;.  | [optional] [default to null]
**OomScoreAdj** | **int32** | An integer value containing the score given to the container in order to tune OOM killer preferences.  | [optional] [default to null]
**PidMode** | **string** | Set the PID (Process) Namespace mode for the container. It can be either:  - &#x60;\&quot;container:&lt;name|id&gt;\&quot;&#x60;: joins another container&#39;s PID namespace - &#x60;\&quot;host\&quot;&#x60;: use the host&#39;s PID namespace inside the container  | [optional] [default to null]
**Privileged** | **bool** | Gives the container full access to the host. | [optional] [default to null]
**PublishAllPorts** | **bool** | Allocates an ephemeral host port for all of a container&#39;s exposed ports.  Ports are de-allocated when the container stops and allocated when the container starts. The allocated port might be changed when restarting the container.  The port is selected from the ephemeral port range that depends on the kernel. For example, on Linux the range is defined by &#x60;/proc/sys/net/ipv4/ip_local_port_range&#x60;.  | [optional] [default to null]
**ReadonlyRootfs** | **bool** | Mount the container&#39;s root filesystem as read only. | [optional] [default to null]
**SecurityOpt** | **[]string** | A list of string values to customize labels for MLS systems, such as SELinux.  | [optional] [default to null]
**StorageOpt** | **map[string]string** | Storage driver options for this container, in the form &#x60;{\&quot;size\&quot;: \&quot;120G\&quot;}&#x60;.  | [optional] [default to null]
**Tmpfs** | **map[string]string** | A map of container directories which should be replaced by tmpfs mounts, and their corresponding mount options. For example:  &#x60;&#x60;&#x60; { \&quot;/run\&quot;: \&quot;rw,noexec,nosuid,size&#x3D;65536k\&quot; } &#x60;&#x60;&#x60;  | [optional] [default to null]
**UTSMode** | **string** | UTS namespace to use for the container. | [optional] [default to null]
**UsernsMode** | **string** | Sets the usernamespace mode for the container when usernamespace remapping option is enabled.  | [optional] [default to null]
**ShmSize** | **int64** | Size of &#x60;/dev/shm&#x60; in bytes. If omitted, the system uses 64MB.  | [optional] [default to null]
**Sysctls** | **map[string]string** | A list of kernel parameters (sysctls) to set in the container. For example:  &#x60;&#x60;&#x60; {\&quot;net.ipv4.ip_forward\&quot;: \&quot;1\&quot;} &#x60;&#x60;&#x60;  | [optional] [default to null]
**Runtime** | **string** | Runtime to use with this container. | [optional] [default to null]
**Isolation** | **string** | Isolation technology of the container. (Windows only)  | [optional] [default to null]
**MaskedPaths** | **[]string** | The list of paths to be masked inside the container (this overrides the default set of paths).  | [optional] [default to null]
**ReadonlyPaths** | **[]string** | The list of paths to be set as read-only inside the container (this overrides the default set of paths).  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


