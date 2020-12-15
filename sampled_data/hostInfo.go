package sampled_data

import (
	"distributed-monitor/utils"
	"fmt"
	"strings"
)

type HostInfo struct {
	Hostname      string `json:"hostname"`
	Os            string `json:"os"`
	KernelVersion string `json:"kernel_version"`
	BootTime      string `json:"boot_time"`
	CpuNum        int    `json:"cpu_num"`
	Ram           int32  `json:"ram"`
	Storage       int64  `json:"storage"`
}

func (this *HostInfo) getHostName() {
	content, err := utils.ReadFile(utils.HOSTNAME)
	if err != nil {
		fmt.Printf("get host name fail, err:%v\n", err)
	}
	this.Hostname = content
}

func (this *HostInfo) getOS() {
	content, err := utils.ReadFile(utils.HOSTOS)
	if err != nil {
		fmt.Printf("get host Os fail, err:%v\n", err)
	}
	this.Os = strings.Split(content, "\\n \\l")[0]
}

func (this *HostInfo) getKernelVersion() {
	content, err := utils.ReadFile(utils.HOSTKERNEL)
	if err != nil {
		fmt.Printf("get host kernel version fail, err:%v\n", err)
	}
	this.KernelVersion = utils.SplitBySpace(content)[2]
}

func (this *HostInfo) getHostUpTime() {
	content, err := utils.ReadFile(utils.HOSTBOOTTIME)
	if err != nil {
		fmt.Printf("get host uptime fail, err:%v\n", err)
	}
	this.BootTime = utils.SplitBySpace(content)[0]
}

func (this *HostInfo) GetHostRamSize() {

}

func (this *HostInfo) GetHostStorageSize() {

}

func getLong(status string, key string) (result int, err error) {
	return 0, nil
}

func getBootTime(stat string) (int, error) {
	return 0, nil
}
