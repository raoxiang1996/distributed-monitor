package monitor

type HostData struct {
	Hostname      string `json:"hostname"`
	Os            string `json:"os"`
	KernelVersion string `json:"kernel_version"`
	BootTime      string `json:"boot_time"`
	CpuNum        int    `json:"cpu_num"`
	Ram           int32  `json:"ram"`
	Storage       int64  `json:"storage"`
}

func (this *HostData) Parse() {

}
