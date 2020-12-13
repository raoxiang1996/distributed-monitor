package base

import (
	"io/ioutil"
	"strings"
)

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

func readFile(filename string) (content string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil

}

func splitBySpace(str string) []string {
	return strings.Fields(str)

}
