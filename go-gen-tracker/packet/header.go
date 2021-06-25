package packet

type HeaderPacket struct {
	Prefix       [2]byte `type:"string"`
	Size         uint8   `unit:"Bytes" chartable:""`
	Vin          uint32
	SendDatetime [7]byte `type:"datetime" chartable:""`
}
