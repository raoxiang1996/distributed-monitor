package base_data

import (
	"fmt"
	"testing"
)

func printMemInfo(meminfo *Meminfo) {
	fmt.Printf("MemTotal: %s\nMemFree: %s\nBuffers: %s\nCached:  %s\nSwapCached: %s\nActive: %s\n"+
		"Inactive: %s\nActive(anon): %s\nInactive(anon): %s\nActive(file): %s\nInactive(file): %s\nUnevictable: %s\n"+
		"Mlocked: %s\nSwapTotal: %s\nSwapFree: %s\nDirty: %s\nWriteback: %s\nAnonPages: %s\nMapped: %s\nShmem: %s\n"+
		"Slab: %s\nSReclaimable: %s\nSUnreclaim: %s\nKernelStack: %s\nPageTables: %s\nNFS_Unstable: %s\nBounce: %s\n"+
		"WritebackTmp: %s\nCommitLimit: %s\nCommitted_AS: %s\nVmallocTotal: %s\nVmallocUsed: %s\nVmallocChunk: %s\n"+
		"HardwareCorrupted: %s\nAnonHugePages: %s\nHugePages_Total: %s\nHugePages_Free: %s\nHugePages_Rsvd: %s\n"+
		"HugePages_Surp: %s\nHugepagesize: %s\nDirectMap4k: %s\nDirectMap4M: %s\n", meminfo.MemTotal, meminfo.MemFree,
		meminfo.Buffers, meminfo.Cached, meminfo.SwapCached, meminfo.Active, meminfo.Inactive, meminfo.Active_anon, meminfo.Inactive_anon,
		meminfo.Active_file, meminfo.Inactive_file, meminfo.Unevictable, meminfo.Mlocked, meminfo.SwapTotal, meminfo.SwapFree, meminfo.Dirty,
		meminfo.Writeback, meminfo.AnonHugePages, meminfo.Mapped, meminfo.Shmem, meminfo.Slab, meminfo.SReclaimable, meminfo.SUnreclaim,
		meminfo.KernelStack, meminfo.PageTables, meminfo.NFS_Unstable, meminfo.Bounce, meminfo.WritebackTmp, meminfo.CommitLimit,
		meminfo.Committed_AS, meminfo.VmallocTotal, meminfo.VmallocUsed, meminfo.VmallocChunk, meminfo.HardwareCorrupted, meminfo.AnonHugePages,
		meminfo.HugePages_Total, meminfo.HugePages_Free, meminfo.HugePages_Rsvd, meminfo.HugePages_Surp, meminfo.Hugepagesize, meminfo.DirectMap4k,
		meminfo.DirectMap4M)
}
func TestMemory(t *testing.T) {
	meminfo := Meminfo{}
	_ = meminfo.Parse("MemTotal:        8293720 kB\n" +
		"MemFree:         2246472 kB\n" +
		"Buffers:           34032 kB\n" +
		"Cached:           188576 kB\n" +
		"SwapCached:            0 kB\n" +
		"Active:           167556 kB\n" +
		"Inactive:         157876 kB\n" +
		"Active(anon):     103104 kB\n" +
		"Inactive(anon):    17440 kB\n" +
		"Active(file):      64452 kB\n" +
		"Inactive(file):   140436 kB\n" +
		"Unevictable:           0 kB\n" +
		"Mlocked:               0 kB\n" +
		"SwapTotal:      15510032 kB\n" +
		"SwapFree:       15118888 kB\n" +
		"Dirty:                 0 kB\n" +
		"Writeback:             0 kB\n" +
		"AnonPages:        102824 kB\n" +
		"Mapped:            71404 kB\n" +
		"Shmem:             17720 kB\n" +
		"Slab:              13868 kB\n" +
		"SReclaimable:       6744 kB\n" +
		"SUnreclaim:         7124 kB\n" +
		"KernelStack:        2848 kB\n" +
		"PageTables:         2524 kB\n" +
		"NFS_Unstable:          0 kB\n" +
		"Bounce:                0 kB\n" +
		"WritebackTmp:          0 kB\n" +
		"CommitLimit:      515524 kB\n" +
		"Committed_AS:    3450064 kB\n" +
		"VmallocTotal:     122880 kB\n" +
		"VmallocUsed:       21296 kB\n" +
		"VmallocChunk:      66044 kB\n" +
		"HardwareCorrupted:     0 kB\n" +
		"AnonHugePages:      2048 kB\n" +
		"HugePages_Total:       0\n" +
		"HugePages_Free:        0\n" +
		"HugePages_Rsvd:        0\n" +
		"HugePages_Surp:        0\n" +
		"Hugepagesize:       2048 kB\n" +
		"DirectMap4k:       12280 kB\n" +
		"DirectMap4M:      897024 kB")
	printMemInfo(&meminfo)
}
