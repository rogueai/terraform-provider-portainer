/*
 * Docker Engine API
 *
 * The Engine API is an HTTP API served by Docker Engine. It is the API the Docker client uses to communicate with the Engine, so everything the Docker client can do can be done with the API.  Most of the client's commands map directly to API endpoints (e.g. `docker ps` is `GET /containers/json`). The notable exception is running containers, which consists of several API calls.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call. The body of the response will be JSON in the following format:  ``` {   \"message\": \"page not found\" } ```  # Versioning  The API is usually changed in each release, so API calls are versioned to ensure that clients don't break. To lock to a specific version of the API, you prefix the URL with its version, for example, call `/v1.30/info` to use the v1.30 version of the `/info` endpoint. If the API version specified in the URL is not supported by the daemon, a HTTP `400 Bad Request` error message is returned.  If you omit the version-prefix, the current version of the API (v1.42) is used. For example, calling `/info` is the same as calling `/v1.42/info`. Using the API without a version-prefix is deprecated and will be removed in a future release.  Engine releases in the near future should support this version of the API, so your client will continue to work even if it is talking to a newer Engine.  The API uses an open schema model, which means server may add extra properties to responses. Likewise, the server will ignore any extra query parameters and request body properties. When you write clients, you need to ignore additional properties in responses to ensure they do not break when talking to newer daemons.   # Authentication  Authentication for registries is handled client side. The client has to send authentication details to various endpoints that need to communicate with registries, such as `POST /images/(name)/push`. These are sent as `X-Registry-Auth` header as a [base64url encoded](https://tools.ietf.org/html/rfc4648#section-5) (JSON) string with the following structure:  ``` {   \"username\": \"string\",   \"password\": \"string\",   \"email\": \"string\",   \"serveraddress\": \"string\" } ```  The `serveraddress` is a domain/IP without a protocol. Throughout this structure, double quotes are required.  If you have already got an identity token from the [`/auth` endpoint](#operation/SystemAuth), you can just pass this instead of credentials:  ``` {   \"identitytoken\": \"9cbaf023786cd7...\" } ```
 *
 * API version: 1.42
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package docker

type SystemInfo struct {
	// Unique identifier of the daemon.  <p><br /></p>  > **Note**: The format of the ID itself is not part of the API, and > should not be considered stable.
	ID string `json:"ID,omitempty"`
	// Total number of containers on the host.
	Containers int32 `json:"Containers,omitempty"`
	// Number of containers with status `\"running\"`.
	ContainersRunning int32 `json:"ContainersRunning,omitempty"`
	// Number of containers with status `\"paused\"`.
	ContainersPaused int32 `json:"ContainersPaused,omitempty"`
	// Number of containers with status `\"stopped\"`.
	ContainersStopped int32 `json:"ContainersStopped,omitempty"`
	// Total number of images on the host.  Both _tagged_ and _untagged_ (dangling) images are counted.
	Images int32 `json:"Images,omitempty"`
	// Name of the storage driver in use.
	Driver string `json:"Driver,omitempty"`
	// Information specific to the storage driver, provided as \"label\" / \"value\" pairs.  This information is provided by the storage driver, and formatted in a way consistent with the output of `docker info` on the command line.  <p><br /></p>  > **Note**: The information returned in this field, including the > formatting of values and labels, should not be considered stable, > and may change without notice.
	DriverStatus [][]string `json:"DriverStatus,omitempty"`
	// Root directory of persistent Docker state.  Defaults to `/var/lib/docker` on Linux, and `C:\\ProgramData\\docker` on Windows.
	DockerRootDir string       `json:"DockerRootDir,omitempty"`
	Plugins       *PluginsInfo `json:"Plugins,omitempty"`
	// Indicates if the host has memory limit support enabled.
	MemoryLimit bool `json:"MemoryLimit,omitempty"`
	// Indicates if the host has memory swap limit support enabled.
	SwapLimit bool `json:"SwapLimit,omitempty"`
	// Indicates if the host has kernel memory TCP limit support enabled. This field is omitted if not supported.  Kernel memory TCP limits are not supported when using cgroups v2, which does not support the corresponding `memory.kmem.tcp.limit_in_bytes` cgroup.
	KernelMemoryTCP bool `json:"KernelMemoryTCP,omitempty"`
	// Indicates if CPU CFS(Completely Fair Scheduler) period is supported by the host.
	CpuCfsPeriod bool `json:"CpuCfsPeriod,omitempty"`
	// Indicates if CPU CFS(Completely Fair Scheduler) quota is supported by the host.
	CpuCfsQuota bool `json:"CpuCfsQuota,omitempty"`
	// Indicates if CPU Shares limiting is supported by the host.
	CPUShares bool `json:"CPUShares,omitempty"`
	// Indicates if CPUsets (cpuset.cpus, cpuset.mems) are supported by the host.  See [cpuset(7)](https://www.kernel.org/doc/Documentation/cgroup-v1/cpusets.txt)
	CPUSet bool `json:"CPUSet,omitempty"`
	// Indicates if the host kernel has PID limit support enabled.
	PidsLimit bool `json:"PidsLimit,omitempty"`
	// Indicates if OOM killer disable is supported on the host.
	OomKillDisable bool `json:"OomKillDisable,omitempty"`
	// Indicates IPv4 forwarding is enabled.
	IPv4Forwarding bool `json:"IPv4Forwarding,omitempty"`
	// Indicates if `bridge-nf-call-iptables` is available on the host.
	BridgeNfIptables bool `json:"BridgeNfIptables,omitempty"`
	// Indicates if `bridge-nf-call-ip6tables` is available on the host.
	BridgeNfIp6tables bool `json:"BridgeNfIp6tables,omitempty"`
	// Indicates if the daemon is running in debug-mode / with debug-level logging enabled.
	Debug bool `json:"Debug,omitempty"`
	// The total number of file Descriptors in use by the daemon process.  This information is only returned if debug-mode is enabled.
	NFd int32 `json:"NFd,omitempty"`
	// The  number of goroutines that currently exist.  This information is only returned if debug-mode is enabled.
	NGoroutines int32 `json:"NGoroutines,omitempty"`
	// Current system-time in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format with nano-seconds.
	SystemTime string `json:"SystemTime,omitempty"`
	// The logging driver to use as a default for new containers.
	LoggingDriver string `json:"LoggingDriver,omitempty"`
	// The driver to use for managing cgroups.
	CgroupDriver string `json:"CgroupDriver,omitempty"`
	// The version of the cgroup.
	CgroupVersion string `json:"CgroupVersion,omitempty"`
	// Number of event listeners subscribed.
	NEventsListener int32 `json:"NEventsListener,omitempty"`
	// Kernel version of the host.  On Linux, this information obtained from `uname`. On Windows this information is queried from the <kbd>HKEY_LOCAL_MACHINE\\\\SOFTWARE\\\\Microsoft\\\\Windows NT\\\\CurrentVersion\\\\</kbd> registry value, for example _\"10.0 14393 (14393.1198.amd64fre.rs1_release_sec.170427-1353)\"_.
	KernelVersion string `json:"KernelVersion,omitempty"`
	// Name of the host's operating system, for example: \"Ubuntu 16.04.2 LTS\" or \"Windows Server 2016 Datacenter\"
	OperatingSystem string `json:"OperatingSystem,omitempty"`
	// Version of the host's operating system  <p><br /></p>  > **Note**: The information returned in this field, including its > very existence, and the formatting of values, should not be considered > stable, and may change without notice.
	OSVersion string `json:"OSVersion,omitempty"`
	// Generic type of the operating system of the host, as returned by the Go runtime (`GOOS`).  Currently returned values are \"linux\" and \"windows\". A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	OSType string `json:"OSType,omitempty"`
	// Hardware architecture of the host, as returned by the Go runtime (`GOARCH`).  A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	Architecture string `json:"Architecture,omitempty"`
	// The number of logical CPUs usable by the daemon.  The number of available CPUs is checked by querying the operating system when the daemon starts. Changes to operating system CPU allocation after the daemon is started are not reflected.
	NCPU int32 `json:"NCPU,omitempty"`
	// Total amount of physical memory available on the host, in bytes.
	MemTotal int64 `json:"MemTotal,omitempty"`
	// Address / URL of the index server that is used for image search, and as a default for user authentication for Docker Hub and Docker Cloud.
	IndexServerAddress string                 `json:"IndexServerAddress,omitempty"`
	RegistryConfig     *RegistryServiceConfig `json:"RegistryConfig,omitempty"`
	GenericResources   *GenericResources      `json:"GenericResources,omitempty"`
	// HTTP-proxy configured for the daemon. This value is obtained from the [`HTTP_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable. Credentials ([user info component](https://tools.ietf.org/html/rfc3986#section-3.2.1)) in the proxy URL are masked in the API response.  Containers do not automatically inherit this configuration.
	HttpProxy string `json:"HttpProxy,omitempty"`
	// HTTPS-proxy configured for the daemon. This value is obtained from the [`HTTPS_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable. Credentials ([user info component](https://tools.ietf.org/html/rfc3986#section-3.2.1)) in the proxy URL are masked in the API response.  Containers do not automatically inherit this configuration.
	HttpsProxy string `json:"HttpsProxy,omitempty"`
	// Comma-separated list of domain extensions for which no proxy should be used. This value is obtained from the [`NO_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.  Containers do not automatically inherit this configuration.
	NoProxy string `json:"NoProxy,omitempty"`
	// Hostname of the host.
	Name string `json:"Name,omitempty"`
	// User-defined labels (key/value metadata) as set on the daemon.  <p><br /></p>  > **Note**: When part of a Swarm, nodes can both have _daemon_ labels, > set through the daemon configuration, and _node_ labels, set from a > manager node in the Swarm. Node labels are not included in this > field. Node labels can be retrieved using the `/nodes/(id)` endpoint > on a manager node in the Swarm.
	Labels []string `json:"Labels,omitempty"`
	// Indicates if experimental features are enabled on the daemon.
	ExperimentalBuild bool `json:"ExperimentalBuild,omitempty"`
	// Version string of the daemon.  > **Note**: the [standalone Swarm API](/swarm/swarm-api/) > returns the Swarm version instead of the daemon  version, for example > `swarm/1.2.8`.
	ServerVersion string `json:"ServerVersion,omitempty"`
	// URL of the distributed storage backend.   The storage backend is used for multihost networking (to store network and endpoint information) and by the node discovery mechanism.  <p><br /></p>  > **Deprecated**: This field is only propagated when using standalone Swarm > mode, and overlay networking using an external k/v store. Overlay > networks with Swarm mode enabled use the built-in raft store, and > this field will be empty.
	ClusterStore string `json:"ClusterStore,omitempty"`
	// The network endpoint that the Engine advertises for the purpose of node discovery. ClusterAdvertise is a `host:port` combination on which the daemon is reachable by other hosts.  <p><br /></p>  > **Deprecated**: This field is only propagated when using standalone Swarm > mode, and overlay networking using an external k/v store. Overlay > networks with Swarm mode enabled use the built-in raft store, and > this field will be empty.
	ClusterAdvertise string `json:"ClusterAdvertise,omitempty"`
	// List of [OCI compliant](https://github.com/opencontainers/runtime-spec) runtimes configured on the daemon. Keys hold the \"name\" used to reference the runtime.  The Docker daemon relies on an OCI compliant runtime (invoked via the `containerd` daemon) as its interface to the Linux kernel namespaces, cgroups, and SELinux.  The default runtime is `runc`, and automatically configured. Additional runtimes can be configured by the user and will be listed here.
	Runtimes map[string]Runtime `json:"Runtimes,omitempty"`
	// Name of the default OCI runtime that is used when starting containers.  The default can be overridden per-container at create time.
	DefaultRuntime string     `json:"DefaultRuntime,omitempty"`
	Swarm          *SwarmInfo `json:"Swarm,omitempty"`
	// Indicates if live restore is enabled.  If enabled, containers are kept running when the daemon is shutdown or upon daemon start if running containers are detected.
	LiveRestoreEnabled bool `json:"LiveRestoreEnabled,omitempty"`
	// Represents the isolation technology to use as a default for containers. The supported values are platform-specific.  If no isolation value is specified on daemon start, on Windows client, the default is `hyperv`, and on Windows server, the default is `process`.  This option is currently not used on other platforms.
	Isolation string `json:"Isolation,omitempty"`
	// Name and, optional, path of the `docker-init` binary.  If the path is omitted, the daemon searches the host's `$PATH` for the binary and uses the first result.
	InitBinary       string  `json:"InitBinary,omitempty"`
	ContainerdCommit *Commit `json:"ContainerdCommit,omitempty"`
	RuncCommit       *Commit `json:"RuncCommit,omitempty"`
	InitCommit       *Commit `json:"InitCommit,omitempty"`
	// List of security features that are enabled on the daemon, such as apparmor, seccomp, SELinux, user-namespaces (userns), and rootless.  Additional configuration options for each security feature may be present, and are included as a comma-separated list of key/value pairs.
	SecurityOptions []string `json:"SecurityOptions,omitempty"`
	// Reports a summary of the product license on the daemon.  If a commercial license has been applied to the daemon, information such as number of nodes, and expiration are included.
	ProductLicense string `json:"ProductLicense,omitempty"`
	// List of custom default address pools for local networks, which can be specified in the daemon.json file or dockerd option.  Example: a Base \"10.10.0.0/16\" with Size 24 will define the set of 256 10.10.[0-255].0/24 address pools.
	DefaultAddressPools []SystemInfoDefaultAddressPools `json:"DefaultAddressPools,omitempty"`
	// List of warnings / informational messages about missing features, or issues related to the daemon configuration.  These messages can be printed by the client as information to the user.
	Warnings []string `json:"Warnings,omitempty"`
}