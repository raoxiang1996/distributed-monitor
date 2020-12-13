package base

import (
	"errors"
	"strings"
)

const NumFileOfMemInfo = 42

type Meminfo struct {
	MemTotal          string `json:"mem_total"`
	MemFree           string `json:"mem_free"`
	Buffers           string `json:"buffers"`
	Cached            string `json:"cached"`
	SwapCached        string `json:"swap_cached"`
	Active            string `json:"active"`
	Inactive          string `json:"inactive"`
	Active_anon       string `json:"active_anon"`
	Inactive_anon     string `json:"inactive_anon"`
	Active_file       string `json:"active_file"`
	Inactive_file     string `json:"inactive_file"`
	Unevictable       string `json:"unevictable"`
	Mlocked           string `json:"mlocked"`
	SwapTotal         string `json:"swap_total"`
	SwapFreestring    string `json:"swap_freestring"`
	Dirty             string `json:"dirty"`
	Writeback         string `json:"writeback"`
	AnonPages         string `json:"anon_pages"`
	Mapped            string `json:"mapped"`
	Shmem             string `json:"shmem"`
	Slab              string `json:"slab"`
	SReclaimable      string `json:"s_reclaimable"`
	SUnreclaim        string `json:"s_unreclaim"`
	KernelStack       string `json:"kernel_stack"`
	PageTables        string `json:"page_tables"`
	NFS_Unstable      string `json:"nfs_unstable"`
	Bounce            string `json:"bounce"`
	WritebackTmp      string `json:"writeback_tmp"`
	CommitLimit       string `json:"commit_limit"`
	Committed_AS      string `json:"committed_as"`
	VmallocTotal      string `json:"vmalloc_total"`
	VmallocUsed       string `json:"vmalloc_used"`
	VmallocChunk      string `json:"vmalloc_chunk"`
	HardwareCorrupted string `json:"hardware_corrupted"`
	AnonHugePages     string `json:"anon_huge_pages"`
	HugePages_Total   string `json:"huge_pages_total"`
	HugePages_Free    string `json:"huge_pages_free"`
	HugePages_Rsvd    string `json:"huge_pages_rsvd"`
	HugePages_Surp    string `json:"huge_pages_surp"`
	Hugepagesize      string `json:"hugepagesize"`
	DirectMap4k       string `json:"direct_map_4_k"`
	DirectMap4M       string `json:"direct_map_4_m"`
}

func (this *Meminfo) Parse(meminfo string) error {
	dataList := strings.Split(meminfo, "\n")
	if len(dataList) != NumFileOfMemInfo {
		return errors.New("memory info parse fail")
	}
	this.MemTotal = strings.Split(dataList[0], ":")[1]
	this.MemFree = strings.Split(dataList[1], ":")[1]
	this.Buffers = strings.Split(dataList[2], ":")[1]
	this.Cached = strings.Split(dataList[3], ":")[1]
	this.SwapCached = strings.Split(dataList[4], ":")[1]
	this.Active = strings.Split(dataList[5], ":")[1]
	this.Inactive = strings.Split(dataList[6], ":")[1]
	this.Active_anon = strings.Split(dataList[7], ":")[1]
	this.Inactive_anon = strings.Split(dataList[8], ":")[1]
	this.Active_file = strings.Split(dataList[9], ":")[1]
	this.Inactive_file = strings.Split(dataList[10], ":")[1]
	this.Unevictable = strings.Split(dataList[11], ":")[1]
	this.Mlocked = strings.Split(dataList[12], ":")[1]
	this.SwapTotal = strings.Split(dataList[13], ":")[1]
	this.SwapFreestring = strings.Split(dataList[14], ":")[1]
	this.Dirty = strings.Split(dataList[15], ":")[1]
	this.Writeback = strings.Split(dataList[16], ":")[1]
	this.AnonPages = strings.Split(dataList[17], ":")[1]
	this.Mapped = strings.Split(dataList[18], ":")[1]
	this.Shmem = strings.Split(dataList[19], ":")[1]
	this.Slab = strings.Split(dataList[20], ":")[1]
	this.SReclaimable = strings.Split(dataList[21], ":")[1]
	this.SUnreclaim = strings.Split(dataList[22], ":")[1]
	this.KernelStack = strings.Split(dataList[23], ":")[1]
	this.PageTables = strings.Split(dataList[24], ":")[1]
	this.NFS_Unstable = strings.Split(dataList[25], ":")[1]
	this.Bounce = strings.Split(dataList[26], ":")[1]
	this.WritebackTmp = strings.Split(dataList[27], ":")[1]
	this.CommitLimit = strings.Split(dataList[28], ":")[1]
	this.Committed_AS = strings.Split(dataList[29], ":")[1]
	this.VmallocTotal = strings.Split(dataList[30], ":")[1]
	this.VmallocUsed = strings.Split(dataList[31], ":")[1]
	this.VmallocChunk = strings.Split(dataList[32], ":")[1]
	this.HardwareCorrupted = strings.Split(dataList[33], ":")[1]
	this.AnonHugePages = strings.Split(dataList[34], ":")[1]
	this.HugePages_Total = strings.Split(dataList[35], ":")[1]
	this.HugePages_Free = strings.Split(dataList[36], ":")[1]
	this.HugePages_Rsvd = strings.Split(dataList[37], ":")[1]
	this.HugePages_Surp = strings.Split(dataList[38], ":")[1]
	this.Hugepagesize = strings.Split(dataList[39], ":")[1]
	this.DirectMap4k = strings.Split(dataList[40], ":")[1]
	this.DirectMap4M = strings.Split(dataList[41], ":")[1]
	return nil
}
