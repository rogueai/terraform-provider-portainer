# Resources

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


