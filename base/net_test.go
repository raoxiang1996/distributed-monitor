package base

import (
	"fmt"
	"testing"
)

func printNetAdapterData(netAdapterData *NetAdapterData) {
	fmt.Printf("AdapterName: %s\n"+
		"Bytes: %d\nPackets: %d\nErrs: %d\nDrop: %d\n"+
		"Fifo: %d\nFrame: %d\nCompressed: %d\nMulticast: %d\n"+
		"Bytes: %d\nPackets: %d\nErrs: %d\nDrop: %d\nFifo: %d\n"+
		"Colls: %d\nCarrier: %d\nCompressed: %d\n",
		netAdapterData.AdapterName, netAdapterData.Receive.Bytes, netAdapterData.Receive.Packets,
		netAdapterData.Receive.Errs, netAdapterData.Receive.Drop, netAdapterData.Receive.Fifo,
		netAdapterData.Receive.Frame, netAdapterData.Receive.Compressed, netAdapterData.Receive.Multicast,
		netAdapterData.Transmit.Bytes, netAdapterData.Transmit.Packets, netAdapterData.Transmit.Errs,
		netAdapterData.Transmit.Drop, netAdapterData.Transmit.Fifo, netAdapterData.Transmit.Colls,
		netAdapterData.Transmit.Carrier, netAdapterData.Transmit.Compressed)

}

func TestNet(t *testing.T) {
	netAdapterData := NetAdapterData{}
	netAdapterData.Parse("eth0:       1       2    3    0    0     0          0         0        0       0    0    0    0     7       8          9")
	printNetAdapterData(&netAdapterData)
	//netAdapterData.Parse("eth1:       2       3    4    0    0     0          0         0        0       0    0    0    0     6       7          8")
	//printNetAdapterData(&netAdapterData)
	//netAdapterData.Parse("lo:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0")
	//printNetAdapterData(&netAdapterData)
	//netAdapterData.Parse("wifi0:       0       0    0    0    0     0          0         0        0       0    0    0    0     0       0          0")
	//printNetAdapterData(&netAdapterData)
}
