package base

import (
	"errors"
	"strconv"
	"strings"
)

const (
	NumFiledOfNetDat = 16
)

type NetAdapterData struct {
	AdapterName string   `json:"name"`
	Receive     Receive  `json:"receive"`
	Transmit    Transmit `json:"transmit"`
}

type Receive struct {
	Bytes      int `json:"bytes"`      //接口发送或接收的数据的总字节数
	Packets    int `json:"packets"`    //接口发送或接收的数据包总数
	Errs       int `json:"errs"`       //由设备驱动程序检测到的发送或接收错误的总数
	Drop       int `json:"drop"`       //设备驱动程序丢弃的数据包总数
	Fifo       int `json:"fifo"`       //FIFO缓冲区错误的数量
	Frame      int `json:"frame"`      //分组帧错误的数量
	Compressed int `json:"compressed"` //设备驱动程序发送或接收的压缩数据包数
	Multicast  int `json:"multicast"`  //设备驱动程序发送或接收的多播帧数
}

type Transmit struct {
	Bytes      int `json:"bytes"`      //接口发送或接收的数据的总字节数
	Packets    int `json:"packets"`    //接口发送或接收的数据包总数
	Errs       int `json:"errs"`       //由设备驱动程序检测到的发送或接收错误的总数
	Drop       int `json:"drop"`       //设备驱动程序丢弃的数据包总数
	Fifo       int `json:"fifo"`       //FIFO缓冲区错误的数量
	Colls      int `json:"colls"`      //接口上检测到的冲突数
	Carrier    int `json:"carrier"`    //由设备驱动程序检测到的载波损耗的数量
	Compressed int `json:"compressed"` //设备驱动程序发送或接收的压缩数据包数
}

func (this *NetAdapterData) Parse(netdata string) error {
	dataList := strings.Split(netdata, ":")
	if len(dataList) != NumFiledOfNetDat {
		return errors.New("net data parse fail")
	}
	this.AdapterName = dataList[0]
	dataList = strings.Split(dataList[1], " ")
	this.Receive.Bytes, _ = strconv.Atoi(dataList[0])
	this.Receive.Packets, _ = strconv.Atoi(dataList[1])
	this.Receive.Errs, _ = strconv.Atoi(dataList[2])
	this.Receive.Drop, _ = strconv.Atoi(dataList[3])
	this.Receive.Fifo, _ = strconv.Atoi(dataList[4])
	this.Receive.Frame, _ = strconv.Atoi(dataList[5])
	this.Receive.Compressed, _ = strconv.Atoi(dataList[6])
	this.Receive.Multicast, _ = strconv.Atoi(dataList[7])

	this.Transmit.Bytes, _ = strconv.Atoi(dataList[8])
	this.Transmit.Packets, _ = strconv.Atoi(dataList[9])
	this.Transmit.Errs, _ = strconv.Atoi(dataList[10])
	this.Transmit.Drop, _ = strconv.Atoi(dataList[11])
	this.Transmit.Fifo, _ = strconv.Atoi(dataList[12])
	this.Transmit.Colls, _ = strconv.Atoi(dataList[13])
	this.Transmit.Carrier, _ = strconv.Atoi(dataList[14])
	this.Transmit.Compressed, _ = strconv.Atoi(dataList[15])
	return nil
}
