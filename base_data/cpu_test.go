package base_data

import (
	"fmt"
	"testing"
)

func printCpuData(cpudata *CpuData) {
	fmt.Printf("Name : %s\nUser: %d\nNice: %d\nSystem: %d\nIdle: %d\nIowait： %d\n"+
		"Irq: %d\nSofttirq: %d\nStealsolen: %d\nGuest: %d\nGuest_nice: %d\n", cpudata.Name,
		cpudata.User, cpudata.Nice, cpudata.System, cpudata.Idle, cpudata.Iowait, cpudata.Irq,
		cpudata.Softtirq, cpudata.Stealsolen, cpudata.Guest, cpudata.Guest_nice)
}

func TestCpu(t *testing.T) {
	cpudata := CpuData{}
	_ = cpudata.Parse("cpu  7249 0 12561 9535 0 1115 0 0 0 0")
	printCpuData(&cpudata)
	_ = cpudata.Parse("cpu0 1509 0 5170 657 0 1001 0 0 0 0")
	printCpuData(&cpudata)
	_ = cpudata.Parse("cpu1 2014 0 1976 3346 0 43 0 0 0 0")
	printCpuData(&cpudata)
	_ = cpudata.Parse("cpu2 1848 0 2914 2575 0 39 0 0 0 0")
	printCpuData(&cpudata)
	_ = cpudata.Parse("cpu3 1878 0 2501 2957 0 32 0 0 0 0")
	printCpuData(&cpudata)
	t.Log("hello world")
}

func printCpuInfo(cpuinfo *CpuInfo) {
	fmt.Printf("processor: %s\nvendor_id: %s\ncpu_family: %s\nmodel: %s\n"+
		"model_name: %s\nstepping: %s\nmicrocode: %s\ncpu_MHz: %s\ncache_size: %s\n"+
		"physical_id: %s\nsiblings: %s\ncore_id: %s\ncpu_cores: %s\napicid: %s\ninitial_apicid: %s\n"+
		"fpu: %s\nfpu_exception: %s\ncpuid_level: %s\nwp: %s\nflags: %s\nbogomips: %s\n"+
		"clflush_size: %s\ncache_alignment: %s\naddress_sizes: %s\npower management:%s\n", cpuinfo.processor, cpuinfo.vendor_id,
		cpuinfo.cpu_family, cpuinfo.model, cpuinfo.model_name, cpuinfo.stepping, cpuinfo.microcode, cpuinfo.cpu_MHz,
		cpuinfo.cache_size, cpuinfo.physical_id, cpuinfo.siblings, cpuinfo.core_id, cpuinfo.cpu_cores, cpuinfo.apicid,
		cpuinfo.initial_apicid, cpuinfo.fpu, cpuinfo.fpu_exception, cpuinfo.cpuid_level, cpuinfo.wp, cpuinfo.flags,
		cpuinfo.bogomips, cpuinfo.clflush_size, cpuinfo.cache_alignment, cpuinfo.address_sizes, cpuinfo.power_management)
}
func TestCpuInfo(t *testing.T) {
	cpuinfo := CpuInfo{}
	_ = cpuinfo.Parse("processor       : 0\n" +
		"vendor_id       : GenuineIntel\n" +
		"cpu family      : 6\n" +
		"model           : 60\n" +
		"model name      : Intel(R) Core(TM) i5-4210H CPU @ 2.90GHz\n" +
		"stepping        : 3\n" +
		"microcode       : 0xffffffff\n" +
		"cpu MHz         : 2901.000\n" +
		"cache size      : 256 KB\n" +
		"physical id     : 0\n" +
		"siblings        : 4\n" +
		"core id         : 0\n" +
		"cpu cores       : 2\n" +
		"apicid          : 0\n" +
		"initial apicid  : 0\n" +
		"fpu             : yes\n" +
		"fpu_exception   : yes\n" +
		"cpuid level     : 6\n" +
		"wp              : yes\n" +
		"flags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm pni pclmulqdq est tm2 ssse3 fma cx16 xtpr pdcm pcid sse4_1 sse4_2 movbe popcnt aes xsave osxsave avx f16c rdrand hypervisor lahf_lm abm fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid ibrs ibpb stibp ssbd\n" +
		"bogomips        : 5802.00\n" +
		"clflush size    : 64\n" +
		"cache_alignment : 64\n" +
		"address sizes   : 36 bits physical, 48 bits virtual\n" +
		"power management:")
	printCpuInfo(&cpuinfo)
}
