package base_data

import (
	"errors"
	"strconv"

	"distributed-monitor/utils"
)

const NumFileOfDiskInfo = 18

type DiskInfo struct {
	DiskName  string `json:"disk_name"`
	DiskTotal int    `json:"disk_total""`
	DiskFree  int    `json:"disk_free"`
}

type DiskData struct {
	Dev_no                  int    //设备编号
	Num                     int    //设备号
	Dev_name                string //设备名
	Read_completion_times   int64  //读完成次数
	Merge_reads_completed   int64  //合并读完成次数
	Read_sectors_times      int64  //读取的扇区数，成功读取的扇区总数
	Read_milliseconds       int64  //读操作花费毫秒数
	Write_completion_times  int64  //写完成的次数，成功写完成的总次数
	Merge_writes_completed  int64  //合并写完成次数
	Write_sectors_times     int64  //写扇区的次数
	Write_milliseconds      int64  //写操作花费的毫秒数
	IO_requests             int64  //正在处理的输入/输出请求数
	IO_milliseconds         int64  //输入/输出操作花费的毫秒数
	IO_eighted_milliseconds int64  //输入/输出操作花费的加权毫秒
}

func (this *DiskData) Parse(diskdata string) error {
	dataList := utils.SplitBySpace(diskdata)
	if len(dataList) != NumFileOfDiskInfo {
		return errors.New("disk data parse fail")
	}

	this.Dev_no, _ = strconv.Atoi(dataList[0])
	this.Num, _ = strconv.Atoi(dataList[1])
	this.Dev_name = dataList[2]
	this.Read_completion_times, _ = strconv.ParseInt(dataList[3], 10, 64)
	this.Merge_reads_completed, _ = strconv.ParseInt(dataList[4], 10, 64)
	this.Read_sectors_times, _ = strconv.ParseInt(dataList[5], 10, 64)
	this.Read_milliseconds, _ = strconv.ParseInt(dataList[6], 10, 64)
	this.Write_completion_times, _ = strconv.ParseInt(dataList[7], 10, 64)
	this.Merge_writes_completed, _ = strconv.ParseInt(dataList[8], 10, 64)
	this.Write_sectors_times, _ = strconv.ParseInt(dataList[9], 10, 64)
	this.Write_milliseconds, _ = strconv.ParseInt(dataList[10], 10, 64)
	this.IO_requests, _ = strconv.ParseInt(dataList[11], 10, 64)
	this.IO_milliseconds, _ = strconv.ParseInt(dataList[12], 10, 64)
	this.IO_eighted_milliseconds, _ = strconv.ParseInt(dataList[13], 10, 64)
	return nil
}
