package sampled_data

import (
	"distributed-monitor/utils"
	"fmt"
	"testing"

	"bou.ke/monkey"
)

func printHostInfo(hostInfo *HostInfo) {
	fmt.Printf("Hostname: %s\nOs: %s\nKernelVersion: %s\n"+
		"BootTime: %s\nCpuNum: %d\nRam: %d\nStorage %d\n", hostInfo.Hostname, hostInfo.Os,
		hostInfo.KernelVersion, hostInfo.BootTime, hostInfo.CpuNum, hostInfo.Ram, hostInfo.Storage)
}

func TestHostInfo(t *testing.T) {
	monkey.Patch(utils.ReadFile, func(file string) (string, error) {
		fmt.Println("hello world!")
		if file == utils.HOSTOS {
			return "Ubuntu 20.04.1 LTS \\n \\l", nil
		} else if file == utils.HOSTNAME {
			return "vultr.guest", nil
		} else if file == utils.HOSTKERNEL {
			return "Linux version 5.4.0-54-generic (buildd@lcy01-amd64-024) " +
				"(gcc version 9.3.0 (Ubuntu 9.3.0-17ubuntu1~20.04)) #60-Ubuntu " +
				"SMP Fri Nov 6 10:37:59 UTC 2020", nil
		} else if file == utils.HOSTBOOTTIME {
			return "4265.27 1041.06", nil
		}
		return "", nil
	})

	hostData := HostInfo{}
	hostData.getHostName()
	hostData.getKernelVersion()
	hostData.getOS()
	//hostData.getHostUpTime()
	//hostData.GetHostRamSize()
	//hostData.GetHostStorageSize()
	printHostInfo(&hostData)
}
