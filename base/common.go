package base

const (
	CPUUSAGE = "/proc/stat"
	CPUINFO  = "/proc/cpuinfo"
	MEMORY   = "/proc/meminfo"
	DISK     = "/proc/diskstats"
	NETWORK  = "/proc/net/dev"
	PROCESS  = ""

	HOSTNAME     = "/proc/sys/kernel/hostname"
	HOSTKERNEL   = "/proc/version"
	HOSTOS       = "/etc/issue"
	HOSTBOOTTIME = "/proc/uptime"
)
