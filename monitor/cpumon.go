package monitor

type CpuInfo struct {
	processor   string `json:"processor"`
	vendor      string `json:"vendor"`
	family      string `json:"family"`
	model       string `json:"model"`
	model_name  string `json:"model_name"`
	stepping    string `json:"stepping"`
	mirocode    string `json:"mirocode"`
	cpu_MHz     string `json:"cpu_m_hz"`
	cache_size  string `json:"cache_size"`
	physical_id string `json:"physical_id"`
	siblings    string `json:"siblings"`
	core_id     string `json:"core_id"`
	name        string `json:"name"`
}

type CpuData struct {
	Runtime uint64 `json:"runtime""`
	Utime   uint64 `json:"utime""`
	Ntime   uint64 `json:"ntime""`
	Stime   uint64 `json:"stime""`
}

func (this *CpuData) Parse() {

}
