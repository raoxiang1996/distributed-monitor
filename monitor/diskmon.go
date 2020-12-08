package monitor

type DiskData struct {
	DiskSize     uint64  `json:"disk_size"`
	FreeDisk     uint64  `json:"free_disk"`
	ReadSpeed    float32 `json:"read_speed"`
	WriteSpeed   float32 `json:"write_speed"`
	time_io_last int64   `json:"time_io_last"`
}

func (this *DiskData) Parse() {

}
