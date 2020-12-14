package sampled_data

import (
	"distributed-monitor/utils"
	"fmt"
	"strconv"
	"strings"
)

type HostData struct {
	Hostname      string `json:"hostname"`
	Os            string `json:"os"`
	KernelVersion string `json:"kernel_version"`
	BootTime      string `json:"boot_time"`
	CpuNum        int    `json:"cpu_num"`
	Ram           int32  `json:"ram"`
	Storage       int64  `json:"storage"`
}

func (this *HostData) getHostName() {
	content, err := utils.ReadFile(utils.HOSTNAME)
	if err != nil {
		fmt.Printf("get host uptime fail")
	}
	this.BootTime = strings.Split(content, "\\n \\l")[0]
}

func (this *HostData) getOS() {
	content, err := utils.ReadFile(utils.HOSTOS)
	if err != nil {
		fmt.Printf("get host uptime fail")
	}
	this.BootTime = content
}

func (this *HostData) getKernelVersion() {
	content, err := utils.ReadFile(utils.HOSTKERNEL)
	if err != nil {
		fmt.Printf("get host uptime fail")
	}
	this.BootTime = utils.SplitBySpace(content)[0]
}

func (this *HostData) getHostUpTime() {
	content, err := utils.ReadFile(utils.HOSTBOOTTIME)
	if err != nil {
		fmt.Printf("get host uptime fail")
	}
	this.BootTime = utils.SplitBySpace(content)[0]
}

func (this *HostData) GetHostRamSize() {

}

func (this *HostData) GetHostStorageSize() {

}

func getLong(status string, key string) (result int, err error) {
	pos := strings.Index(status, key)
	if pos != -1 || pos != 0 {
		s := status[pos+len(key):]
		result, err = strconv.Atoi(s)
	}
	return
}

func getBootTime(stat string) (int, error) {
	return 0, nil
}
