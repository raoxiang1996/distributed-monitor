package utils

type ProtobufTransportFormat struct {
	len          int
	nameLen      int
	typeName     []byte
	protobufData []byte
	checkSum     int
}
