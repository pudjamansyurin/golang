package packet

import "reflect"

type Format struct {
	DataType reflect.Kind
	Scale    float32
	Func     func(b []byte) interface{}
}

type Packet struct {
	Group     string
	Field     string
	Title     string
	Required  bool
	Chartable bool
	Unit      string
	Header    bool
	Size      int
	Format    Format
}

func concatPacket(packets ...[]Packet) []Packet {
	var packet []Packet
	for _, p := range packets {
		packet = append(packet, p...)
	}
	return packet
}
