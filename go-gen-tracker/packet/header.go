package packet

import (
	"reflect"
)

var Header = []Packet{
	{
		group:    "packet",
		field:    "prefix",
		title:    "Prefix",
		header:   true,
		required: true,
		size:     2,
		format:   Format{DataType: reflect.String},
		// format: (v) => HexToAscii(cend(v)),
		// display: (vf) => vf,
		// formatCmd: (_) => cend(AsciiToHex(config.prefix.command)),
	},
	{
		group:     "packet",
		field:     "size",
		title:     "Size",
		header:    true,
		required:  true,
		chartable: true,
		unit:      "Bytes",
		size:      1,
		format:    Format{DataType: reflect.Uint8},
		// format: (v) => HexToUnsignedInt(cend(v)),
		// display: (vf) => Dot(vf),
		// formatCmd: (hex) => cend(IntToHex(hex.length / 2, 1 * 2)),
	},
	{
		group:    "packet",
		field:    "vin",
		title:    "VIN",
		header:   true,
		required: true,
		size:     4,
		format:   Format{DataType: reflect.Uint32},
		// format: (v) => HexToUnsignedInt(cend(v)),
		// display: (vf) => vf,
		// formatCmd: (v) => cend(IntToHex(v, 4 * 2)),
	},
	{
		group:     "packet.datetime",
		field:     "sendDatetime",
		title:     "Send Datetime",
		header:    true,
		required:  true,
		chartable: true,
		size:      7,
		format:    Format{DataType: reflect.Func, Func: toUnixTime},
		// format: (v) => Number(dayjs(parseDatetime(v), "YYMMDDHHmmss0d").unix()),
		// display: (vf) => dayjs.unix(vf).format("ddd, DD-MM-YY HH:mm:ss"),
		// formatCmd: (unix) =>
		//   buildTimestamp(dayjs.unix(unix).format("YYMMDDHHmmss0d")),
	},
}

var CommandHeader = append(Header, []Packet{
	{
		group:    "command",
		field:    "code",
		title:    "Code",
		header:   true,
		required: true,
		size:     1,
		format:   Format{DataType: reflect.Uint8},
		// format: (v) => HexToUnsignedInt(v),
		// display: (vf) => vf,
		// formatCmd: (v) => cend(IntToHex(v, 1 * 2)),
	},
	{
		group:    "command",
		field:    "subCode",
		title:    "Sub Code",
		header:   true,
		required: true,
		size:     1,
		format:   Format{DataType: reflect.Uint8},
		// format: (v) => HexToUnsignedInt(v),
		// display: (vf) => vf,
		// formatCmd: (v) => cend(IntToHex(v, 1 * 2)),
	},
}...)
