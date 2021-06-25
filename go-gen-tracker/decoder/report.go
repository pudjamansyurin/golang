package decoder

import (
	"bytes"
	"encoding/binary"

	"github.com/pudjamansyurin/go-gen-tracker/packet"
	"github.com/pudjamansyurin/go-gen-tracker/util"
)

type Report struct {
	Reader *bytes.Reader
}

func (r *Report) Decode() (interface{}, error) {
	simpleFrame := r.Reader.Len() == binary.Size(packet.ReportSimplePacket{})

	var decoded interface{}
	if simpleFrame {
		decoded, _ = r.decode(&packet.ReportSimplePacket{})
	} else {
		decoded, _ = r.decode(&packet.ReportFullPacket{})
	}

	util.Debug(decoded)

	return nil, nil

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
// 	length := r.Reader.Size()

// 	minLength := int(unsafe.Sizeof(packet.ReportPacket{}))
// 	if length < minLength {
// 		return fmt.Errorf("less report length, %d < %d", length, minLength)
// 	}
// 	return nil
// }

func (r *Report) decode(dst interface{}) (interface{}, error) {
	packet.GetTag(dst)

	// if err := binary.Read(r.Reader, binary.LittleEndian, packet); err != nil {
	// 	return nil, errors.New("cant decode packet")
	// }
	// return packet, nil
	return nil, nil
}
