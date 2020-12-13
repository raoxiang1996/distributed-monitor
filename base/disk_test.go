package base

import (
	"fmt"
	"testing"
)

func printDiskData(diskdata *DiskData) {
	fmt.Printf("Dev_no: %d\nNum: %d\nDev_name: %s\nRead_completion_times: %d\n"+
		"Merge_reads_completed: %d\nRead_sectors_times: %d\nRead_milliseconds: %d\nWrite_completion_times: %d\n"+
		"Merge_writes_completed: %d\nWrite_sectors_times: %d\nWrite_milliseconds: %d\nIO_requests: %d\b"+
		"IO_milliseconds: %d\nIO_eighted_milliseconds: %d\n", diskdata.Dev_no, diskdata.Num, diskdata.Dev_name,
		diskdata.Read_completion_times, diskdata.Merge_reads_completed, diskdata.Read_sectors_times, diskdata.Read_milliseconds,
		diskdata.Write_completion_times, diskdata.Merge_writes_completed, diskdata.Write_sectors_times, diskdata.Write_milliseconds,
		diskdata.IO_requests, diskdata.IO_milliseconds, diskdata.IO_eighted_milliseconds)
}

func TestDisk(t *testing.T) {
	diskdata := DiskData{}
	diskdata.Parse("254       0 vda 91229 32 7436534 71590 169762 137488 12319618 247129 0 54880 262364 0 0 0 0")
	printDiskData(&diskdata)
	//diskdata.Parse("254       1 vda1 90790 32 7422418 71475 161276 137488 12319618 246835 0 54804 262600 0 0 0 0")
	//printDiskData(&diskdata)
}
