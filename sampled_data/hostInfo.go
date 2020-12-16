package sampled_data

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"distributed-monitor/utils"
)

type HostInfo struct {
	Hostname      string `json:"hostname"`
	Os            string `json:"os"`
	KernelVersion string `json:"kernel_version"`
	BootTime      string `json:"boot_time"`
	CpuNum        int    `json:"cpu_num"`
	Ram           uint32 `json:"ram"`
	Storage       uint32 `json:"storage"`
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
	nowTime := time.Now().Unix()
	bootTime, err := strconv.ParseFloat(utils.SplitBySpace(content)[0], 64)
	startTime := nowTime - int64(bootTime)
	timeLayout := "2006-01-02 15:04:05"
	this.BootTime = time.Unix(startTime, 0).Format(timeLayout)
}

func (this *HostInfo) GetHostRamSize() {
	content, err := utils.ReadLine(utils.MEMORYINFO)
	if err != nil {
		fmt.Printf("get host ram size fail, err:%v\n", err)
	}
	content = strings.Split(content, ":")[1]
	content = utils.SplitBySpace(content)[0]
	size, _ := strconv.ParseUint(content, 10, 32)
	if err != nil {
		fmt.Printf("%T, %v\n", size, size)
	}
	this.Ram = uint32(size)

	//sysInfo := new(syscall.Sysinfo_t)
	//err := syscall.Sysinfo(sysInfo)
	//if err != nil {
	//	fmt.Printf("get host ram size fail, err:%v\n", err)
	//}
	//total := sysInfo.Totalram * uint32(syscall.Getpagesize())
	////free := sysInfo.Freeram * uint32(syscall.Getpagesize())
	////used := totall - free
	//this.Ram = total
}

func (this *HostInfo) GetHostStorageSize() {
	//fs := syscall.Statfs_t{}
	//err := syscall.Statfs("/", &fs)
	//if err != nil {
	//	fmt.Printf("get host storage size fail, err:%v\n", err)
	//}
	//total := fs.Blocks * uint64(fs.Bsize)
	//free := fs.Bfree * uint64(fs.Bsize)
	//used := total - free
	//this.Storage = total
}

func (this *HostInfo) GetCpuNum() {
	content, err := utils.ReadFile(utils.CPUINFO)
	cpus := strings.Split(content, "\n\n")
	num := len(cpus)
	if err != nil {
		fmt.Printf("get host cpu num fail, err:%v\n", err)
	}
	this.CpuNum = num
}

func getLong(status string, key string) (result int, err error) {
	return 0, nil
}

func getBootTime(stat string) (int, error) {
	return 0, nil
}

func HostInfoToJSON(hostinfo *HostInfo) ([]byte, error) {
	data, err := json.Marshal(*hostinfo)
	if err != nil {
		fmt.Println("host info to json failed, err:", err)
		return nil, err
	}
	return data, nil
}
