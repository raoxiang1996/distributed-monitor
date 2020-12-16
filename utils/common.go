package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

const (
	CPUUSAGE    = "/proc/stat"
	CPUINFO     = "/proc/cpuinfo"
	MEMORYINFO  = "/proc/meminfo"
	DISKINFO    = "/proc/diskstats"
	NETWORKINFO = "/proc/net/dev"
	PROCESSINFO = ""

	HOSTNAME     = "/proc/sys/kernel/hostname"
	HOSTKERNEL   = "/proc/version"
	HOSTOS       = "/etc/issue"
	HOSTBOOTTIME = "/proc/uptime"
)

func ReadFile(filename string) (content string, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadLine(filename string) (string, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if nil == err {
		br := bufio.NewReader(f)
		for {
			line, err := br.ReadBytes('\n')
			if err != nil {
				return "", err
			} else {
				return string(line), nil
			}
		}
	} else {
		return "", err
	}
}

func SplitBySpace(str string) []string {
	return strings.Fields(str)
}
