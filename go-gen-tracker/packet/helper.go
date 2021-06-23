package packet

func concatPacket(packets ...[]Packet) []Packet {
	var packet []Packet
	for _, p := range packets {
		packet = append(packet, p...)
	}
	return packet
}
