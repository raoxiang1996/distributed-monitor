package sampled_data

import (
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

}

func (this *HostData) getOs() {

}

func (this *HostData) getKernelVersion() {

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
	time, err := getLong(stat, "btime ")
	return time, err
}
