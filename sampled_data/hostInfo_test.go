package sampled_data

import (
	"distributed-monitor/utils"
	"fmt"
	"testing"

	"bou.ke/monkey"
)

func printHostInfo(hostInfo *HostInfo) {
	fmt.Printf("Hostname: %s\nOs: %s\nKernelVersion: %s\n"+
		"BootTime: %s\nCpuNum: %d\nRam: %d\nStorage %d\n", hostInfo.Hostname, hostInfo.Os,
		hostInfo.KernelVersion, hostInfo.BootTime, hostInfo.CpuNum, hostInfo.Ram, hostInfo.Storage)
}

func TestHostInfo(t *testing.T) {
	monkey.Patch(utils.ReadFile, func(file string) (string, error) {
		if file == utils.HOSTOS {
			return "Ubuntu 20.04.1 LTS \\n \\l", nil
		} else if file == utils.HOSTNAME {
			return "vultr.guest", nil
		} else if file == utils.HOSTKERNEL {
			return "Linux version 5.4.0-54-generic (buildd@lcy01-amd64-024) " +
				"(gcc version 9.3.0 (Ubuntu 9.3.0-17ubuntu1~20.04)) #60-Ubuntu " +
				"SMP Fri Nov 6 10:37:59 UTC 2020", nil
		} else if file == utils.HOSTBOOTTIME {
			return "4265.27 1041.06", nil
		} else if file == utils.CPUINFO {
			return "processor       : 0\nvendor_id       : GenuineIntel\ncpu family      : 6\nmodel           : 60\nmodel name      : Intel(R) Core(TM) i5-4210H CPU @ 2.90GHz\nstepping        : 3\nmicrocode       : 0xffffffff\ncpu MHz         : 2901.000\ncache size      : 256 KB\nphysical id     : 0\nsiblings        : 4\ncore id         : 0\ncpu cores       : 2\napicid          : 0\ninitial apicid  : 0\nfpu             : yes\nfpu_exception   : yes\ncpuid level     : 6\nwp              : yes\nflags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm pni pclmulqdq est tm2 ssse3 fma cx16 xtpr pdcm pcid sse4_1 sse4_2 movbe popcnt aes xsave osxsave avx f16c rdrand hypervisor lahf_lm abm fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid ibrs ibpb stibp ssbd\nbogomips        : 5802.00\nclflush size    : 64\ncache_alignment : 64\naddress sizes   : 36 bits physical, 48 bits virtual\npower management:\n\n" +
				"processor       : 1\nvendor_id       : GenuineIntel\ncpu family      : 6\nmodel           : 60\nmodel name      : Intel(R) Core(TM) i5-4210H CPU @ 2.90GHz\nstepping        : 3\nmicrocode       : 0xffffffff\ncpu MHz         : 2901.000\ncache size      : 256 KB\nphysical id     : 0\nsiblings        : 4\ncore id         : 0\ncpu cores       : 2\napicid          : 0\ninitial apicid  : 0\nfpu             : yes\nfpu_exception   : yes\ncpuid level     : 6\nwp              : yes\nflags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm pni pclmulqdq est tm2 ssse3 fma cx16 xtpr pdcm pcid sse4_1 sse4_2 movbe popcnt aes xsave osxsave avx f16c rdrand hypervisor lahf_lm abm fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid ibrs ibpb stibp ssbd\nbogomips        : 5802.00\nclflush size    : 64\ncache_alignment : 64\naddress sizes   : 36 bits physical, 48 bits virtual\npower management:\n\n" +
				"processor       : 2\nvendor_id       : GenuineIntel\ncpu family      : 6\nmodel           : 60\nmodel name      : Intel(R) Core(TM) i5-4210H CPU @ 2.90GHz\nstepping        : 3\nmicrocode       : 0xffffffff\ncpu MHz         : 2901.000\ncache size      : 256 KB\nphysical id     : 0\nsiblings        : 4\ncore id         : 1\ncpu cores       : 2\napicid          : 0\ninitial apicid  : 0\nfpu             : yes\nfpu_exception   : yes\ncpuid level     : 6\nwp              : yes\nflags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm pni pclmulqdq est tm2 ssse3 fma cx16 xtpr pdcm pcid sse4_1 sse4_2 movbe popcnt aes xsave osxsave avx f16c rdrand hypervisor lahf_lm abm fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid ibrs ibpb stibp ssbd\nbogomips        : 5802.00\nclflush size    : 64\ncache_alignment : 64\naddress sizes   : 36 bits physical, 48 bits virtual\npower management:\n\n" +
				"processor       : 3\nvendor_id       : GenuineIntel\ncpu family      : 6\nmodel           : 60\nmodel name      : Intel(R) Core(TM) i5-4210H CPU @ 2.90GHz\nstepping        : 3\nmicrocode       : 0xffffffff\ncpu MHz         : 2901.000\ncache size      : 256 KB\nphysical id     : 0\nsiblings        : 4\ncore id         : 1\ncpu cores       : 2\napicid          : 0\ninitial apicid  : 0\nfpu             : yes\nfpu_exception   : yes\ncpuid level     : 6\nwp              : yes\nflags           : fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm pni pclmulqdq est tm2 ssse3 fma cx16 xtpr pdcm pcid sse4_1 sse4_2 movbe popcnt aes xsave osxsave avx f16c rdrand hypervisor lahf_lm abm fsgsbase tsc_adjust bmi1 avx2 smep bmi2 erms invpcid ibrs ibpb stibp ssbd\nbogomips        : 5802.00\nclflush size    : 64\ncache_alignment : 64\naddress sizes   : 36 bits physical, 48 bits virtual\npower management:", nil
		}
		return "", nil
	})
	monkey.Patch(utils.ReadLine, func(filename string) (string, error) {
		if filename == utils.MEMORYINFO {
			return "MemTotal:         489440 kB", nil
		}
		return "", nil
	})

	hostData := HostInfo{}
	hostData.getHostName()
	hostData.getKernelVersion()
	hostData.getOS()
	hostData.GetCpuNum()
	hostData.getHostUpTime()
	hostData.GetHostRamSize()
	//hostData.GetHostStorageSize()
	printHostInfo(&hostData)
	println(HostInfoToJSON(&hostData))
}
