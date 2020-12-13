package base

import (
	"errors"
	"strconv"
	"strings"
)

const (
	NumFiledOfCpuDat = 10
	NumFileOfCpuInfo = 24
)

type CpuData struct {
	User       int `json:"user"`
	Nice       int `json:"nice"`
	System     int `json:"system"`
	Idle       int `json:"idle"`
	Iowait     int `json:"iowait"`
	Irq        int `json:"irq"`
	Softtirq   int `json:"softtirq"`
	Stealsolen int `json:"stealsolen"`
	Guest      int `json:"guest"`
	Guest_nice int `json:"guest_nice"`
}

type CpuInfo struct {
	processor       string `json:"processor"`
	vendor_id       string `json:"vendor_id"`
	cpu_family      string `json:"cpu_family"`
	model           string `json:"model"`
	model_name      string `json:"model_name"`
	stepping        string `json:"stepping"`
	microcode       string `json:"microcode"`
	cpu_MHz         string `json:"cpu_m_hz"`
	cache_size      string `json:"cache_size"`
	physical_id     string `json:"physical_id"`
	siblings        string `json:"siblings"`
	core_id         string `json:"core_id"`
	cpu_cores       string `json:"cpu_cores"`
	apicid          string `json:"apicid"`
	initial_apicid  string `json:"initial_apicid"`
	fpu             string `json:"fpu"`
	fpu_exception   string `json:"fpu_exception"`
	cpuid_level     string `json:"cpuid_level"`
	wp              string `json:"wp"`
	flags           string `json:"flags"`
	bogomips        string `json:"bogomips"`
	clflush_size    string `json:"clflush_size"`
	cache_alignment string `json:"cache_alignment"`
	address_sizes   string `json:"address_sizes"`
}

func (this *CpuData) Parse(stat string) error {
	dataList := split(stat, " ")
	if len(dataList) != NumFiledOfCpuDat {
		return errors.New("cpu data parse fail")
	}

	this.User, _ = strconv.Atoi(dataList[0])
	this.Nice, _ = strconv.Atoi(dataList[1])
	this.System, _ = strconv.Atoi(dataList[2])
	this.Idle, _ = strconv.Atoi(dataList[3])
	this.Iowait, _ = strconv.Atoi(dataList[4])
	this.Irq, _ = strconv.Atoi(dataList[5])
	this.Softtirq, _ = strconv.Atoi(dataList[6])
	this.Stealsolen, _ = strconv.Atoi(dataList[7])
	this.Guest, _ = strconv.Atoi(dataList[8])
	this.Guest_nice, _ = strconv.Atoi(dataList[9])
	return nil
}

func (this *CpuInfo) Parse(cpuinfo string) error {
	dataList := strings.Split(cpuinfo, "\n\n")
	if len(dataList) != NumFileOfCpuInfo {
		return errors.New("cpu info parse fail")
	}
	this.processor = strings.Split(dataList[0], ":")[1]
	this.vendor_id = strings.Split(dataList[1], ":")[1]
	this.cpu_family = strings.Split(dataList[2], ":")[1]
	this.model = strings.Split(dataList[3], ":")[1]
	this.model_name = strings.Split(dataList[4], ":")[1]
	this.stepping = strings.Split(dataList[5], ":")[1]
	this.microcode = strings.Split(dataList[6], ":")[1]
	this.cpu_MHz = strings.Split(dataList[7], ":")[1]
	this.cache_size = strings.Split(dataList[8], ":")[1]
	this.physical_id = strings.Split(dataList[9], ":")[1]
	this.siblings = strings.Split(dataList[10], ":")[1]
	this.core_id = strings.Split(dataList[11], ":")[1]
	this.cpu_cores = strings.Split(dataList[12], ":")[1]
	this.apicid = strings.Split(dataList[13], ":")[1]
	this.initial_apicid = strings.Split(dataList[14], ":")[1]
	this.fpu = strings.Split(dataList[15], ":")[1]
	this.fpu_exception = strings.Split(dataList[16], ":")[1]
	this.cpuid_level = strings.Split(dataList[17], ":")[1]
	this.wp = strings.Split(dataList[18], ":")[1]
	this.flags = strings.Split(dataList[19], ":")[1]
	this.bogomips = strings.Split(dataList[20], ":")[1]
	this.clflush_size = strings.Split(dataList[21], ":")[1]
	this.cache_alignment = strings.Split(dataList[22], ":")[1]
	this.address_sizes = strings.Split(dataList[23], ":")[1]
	return nil
}
