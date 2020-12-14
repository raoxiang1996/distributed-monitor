package base_data

import (
	"errors"
	"strings"

	"distributed-monitor/utils"
)

const NumFileOfMemInfo = 42
const Unit = "KB"

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
	SwapFree          string `json:"swap_freestring"`
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
	this.MemTotal = utils.SplitBySpace(strings.Split(dataList[0], ":")[1])[0]
	this.MemFree = utils.SplitBySpace(strings.Split(dataList[1], ":")[1])[0]
	this.Buffers = utils.SplitBySpace(strings.Split(dataList[2], ":")[1])[0]
	this.Cached = utils.SplitBySpace(strings.Split(dataList[3], ":")[1])[0]
	this.SwapCached = utils.SplitBySpace(strings.Split(dataList[4], ":")[1])[0]
	this.Active = utils.SplitBySpace(strings.Split(dataList[5], ":")[1])[0]
	this.Inactive = utils.SplitBySpace(strings.Split(dataList[6], ":")[1])[0]
	this.Active_anon = utils.SplitBySpace(strings.Split(dataList[7], ":")[1])[0]
	this.Inactive_anon = utils.SplitBySpace(strings.Split(dataList[8], ":")[1])[0]
	this.Active_file = utils.SplitBySpace(strings.Split(dataList[9], ":")[1])[0]
	this.Inactive_file = utils.SplitBySpace(strings.Split(dataList[10], ":")[1])[0]
	this.Unevictable = utils.SplitBySpace(strings.Split(dataList[11], ":")[1])[0]
	this.Mlocked = utils.SplitBySpace(strings.Split(dataList[12], ":")[1])[0]
	this.SwapTotal = utils.SplitBySpace(strings.Split(dataList[13], ":")[1])[0]
	this.SwapFree = utils.SplitBySpace(strings.Split(dataList[14], ":")[1])[0]
	this.Dirty = utils.SplitBySpace(strings.Split(dataList[15], ":")[1])[0]
	this.Writeback = utils.SplitBySpace(strings.Split(dataList[16], ":")[1])[0]
	this.AnonPages = utils.SplitBySpace(strings.Split(dataList[17], ":")[1])[0]
	this.Mapped = utils.SplitBySpace(strings.Split(dataList[18], ":")[1])[0]
	this.Shmem = utils.SplitBySpace(strings.Split(dataList[19], ":")[1])[0]
	this.Slab = utils.SplitBySpace(strings.Split(dataList[20], ":")[1])[0]
	this.SReclaimable = utils.SplitBySpace(strings.Split(dataList[21], ":")[1])[0]
	this.SUnreclaim = utils.SplitBySpace(strings.Split(dataList[22], ":")[1])[0]
	this.KernelStack = utils.SplitBySpace(strings.Split(dataList[23], ":")[1])[0]
	this.PageTables = utils.SplitBySpace(strings.Split(dataList[24], ":")[1])[0]
	this.NFS_Unstable = utils.SplitBySpace(strings.Split(dataList[25], ":")[1])[0]
	this.Bounce = utils.SplitBySpace(strings.Split(dataList[26], ":")[1])[0]
	this.WritebackTmp = utils.SplitBySpace(strings.Split(dataList[27], ":")[1])[0]
	this.CommitLimit = utils.SplitBySpace(strings.Split(dataList[28], ":")[1])[0]
	this.Committed_AS = utils.SplitBySpace(strings.Split(dataList[29], ":")[1])[0]
	this.VmallocTotal = utils.SplitBySpace(strings.Split(dataList[30], ":")[1])[0]
	this.VmallocUsed = utils.SplitBySpace(strings.Split(dataList[31], ":")[1])[0]
	this.VmallocChunk = utils.SplitBySpace(strings.Split(dataList[32], ":")[1])[0]
	this.HardwareCorrupted = utils.SplitBySpace(strings.Split(dataList[33], ":")[1])[0]
	this.AnonHugePages = utils.SplitBySpace(strings.Split(dataList[34], ":")[1])[0]
	this.HugePages_Total = utils.SplitBySpace(strings.Split(dataList[35], ":")[1])[0]
	this.HugePages_Free = utils.SplitBySpace(strings.Split(dataList[36], ":")[1])[0]
	this.HugePages_Rsvd = utils.SplitBySpace(strings.Split(dataList[37], ":")[1])[0]
	this.HugePages_Surp = utils.SplitBySpace(strings.Split(dataList[38], ":")[1])[0]
	this.Hugepagesize = utils.SplitBySpace(strings.Split(dataList[39], ":")[1])[0]
	this.DirectMap4k = utils.SplitBySpace(strings.Split(dataList[40], ":")[1])[0]
	this.DirectMap4M = utils.SplitBySpace(strings.Split(dataList[41], ":")[1])[0]
	return nil
}
