package monitor

type MemData struct {
	MemSize      int32 `json:"mem_size"`
	MemFree      int32 `json:"mem_free"`
	MemAvailable int32 `json:"mem_available"`
	Buffers      int32 `json:"buffers"`
	Cached       int32 `json:"cached"`
	SwapCached   int32 `json:"swap_cached"`
	Active       int32 `json:"active"`
	SwapTotal    int32 `json:"swap_total"`
	SwapFree     int32 `json:"swap_free"`
}

func (this *MemData) Parse() {

}
