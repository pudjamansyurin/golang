package decoder

import (
	"bufio"
	"encoding/binary"
	"errors"

	"github.com/pudjamansyurin/go-gen-tracker/packet"
)

type Report struct {
	Header packet.HeaderPacket
	Buf    *bufio.Reader
}

func (r *Report) Decode() (interface{}, error) {
	in, _ := r.Buf.Peek(1)

	if FrameID(in[0]) == FRAME_FULL {
		return r.decode(&packet.ReportFullPacket{})
	}
	return r.decode(&packet.ReportSimplePacket{})

	// data := make(M)
	// for _, v := range packet.Header {
	// 	buf := make([]byte, v.Size)
	// 	n, err := reader.Read(buf)
	// 	if n != v.Size || err != nil {
	// 		return M{}, errors.New("packet corrupted")
	// 	}

	// 	switch v.Format.DataType {
	// 	case reflect.String:
	// 		data[v.Field] = toAscii(buf)
	// 	case reflect.Uint8:
	// 		data[v.Field] = toUint8(buf)
	// 	case reflect.Uint32:
	// 		data[v.Field] = toUint32(buf)
	// 	case reflect.Interface:
	// 		if v.Format.InterfaceName == "datetime" {
	// 			data[v.Field] = toUnixTime(buf)
	// 		}
	// 	}
	// }
	// return data, nil
}

// func (r *Report) Validate() error {
// 	length := r.Buf.Size()

// 	minLength := int(unsafe.Sizeof(packet.ReportPacket{}))
// 	if length < minLength {
// 		return fmt.Errorf("less report length, %d < %d", length, minLength)
// 	}
// 	return nil
// }

func (r *Report) decode(reportPacket interface{}) (interface{}, error) {
	if err := binary.Read(r.Buf, binary.LittleEndian, reportPacket); err != nil {
		return nil, errors.New("cant decode report")
	}
	return reportPacket, nil
}
