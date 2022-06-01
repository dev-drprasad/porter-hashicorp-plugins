package config

type NetworkSettings struct {
	DNS string `toml:"dns,omitempty"`
	Network string `toml:"network,omitempty"`
	Aliases []string `toml:"aliases,omitempty"`
	HostEntries []string `toml:"host_entries,omitempty"`
	MacAddress string `toml:"mac_address,omitempty"`
	IP string `toml:"ip,omitempty"`
	IP6 string `toml:"ip6,omitempty"`
	LinkIPs []string `toml:"link_local_ips,omitempty"`
}

type Config struct {
	Volumes  [][]string `toml:"Volumes,omitempty"`
	Devices []string `toml:"devices,omitempty"`
	Name string `toml:"name,omitempty"`
	NetworkSettings  NetworkSettings `toml:"network_settings,omitempty"`
	SecurityOpts [][]string `toml:"security_opts,omitempty"`
	Flags []string `toml:"flags,omitempty"`
	Memory string `toml:"memory,omitempty"`
	Swap string `toml:"swap,omitempty"`
	KernelMemory string `toml:"kernel_memory,omitempty"`
	CPUS string `toml:"cpus,omitempty"`
	Capabilities []string `toml:"capabilities,omitempty"`
	LoggingDriver string `toml:"logging_driver,omitempty"`
	CMD string `toml:"cmd,omitempty"`
	EntryPoint string `toml:"entrypoint,omitempty"`
	Expose []string `toml:"expose,omitempty"`
	Ports [][]string `toml:"ports,omitempty"`
	Link []string `toml:"link,omitempty"`
	ENV [][]string `toml:"env,omitempty"`
	TMPFS []string `toml:"tmpfs,omitempty"`
	User string `toml:"user,omitempty"`
	WorkDir string `toml:"workdir,omitempty"`
	
}
