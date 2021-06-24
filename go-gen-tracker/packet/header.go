package packet

type HeaderPacket struct {
	Prefix       [2]byte `type:"string" `
	Size         uint8   `unit:"Bytes" chartable:"1"`
	Vin          uint32
	SendDatetime [7]byte `type:"datetime" chartable:"1"`
}
