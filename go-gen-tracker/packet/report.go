package packet

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pudjamansyurin/go-gen-tracker/decoder"
)

var Report = concatPacket(
	Header,
	vcuPacket(),
	hbarPacket(),

	netPacket(),
	gpsPacket(),
	memsPacket(),
	remotePacket(),
	fingerPacket(),
	audioPacket(),

	hmi1Packet(),
	bmsPacket(),
	mcuPacket(),
	taskPacket(),
)

func vcuPacket() []Packet {
	return []Packet{
		{
			Group:     "packet",
			Field:     "frameID",
			Title:     "Frame ID",
			Required:  true,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => `${config.frames[vf]} (${vf})`,
		},
		{
			Group:     "packet.datetime",
			Field:     "logDatetime",
			Title:     "Log Datetime",
			Required:  true,
			Chartable: true,
			Size:      7,
			Format:    Format{DataType: reflect.Func, Func: decoder.ToUnixTime},
			// Format: (v) => Number(dayjs(parseDatetime(v), "YYMMDDHHmmss0d").unix()),
			// display: (vf) => dayjs.unix(vf).Format("ddd, DD-MM-YY HH:mm:ss"),
		},
		{
			Group:     "vcu",
			Field:     "state",
			Title:     "State",
			Required:  true,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Int8},
			// Format: (v) => HexToSignedInt8(cend(v)),
			// display: (vf) => `${getVehicleState(vf)} (${vf})`,
		},
		{
			Group:     "vcu",
			Field:     "events",
			Title:     "Events",
			Required:  true,
			Chartable: true,
			Size:      2,
			Format:    Format{DataType: reflect.Uint16},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => IntToHex(vf, 4).toUpperCase(),
		},
		{
			Group:     "vcu",
			Field:     "logBuffered",
			Title:     "Log Buffered",
			Required:  true,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "vcu",
			Field:     "batVoltage",
			Title:     "Bat. Voltage",
			Required:  true,
			Chartable: true,
			Unit:      "mVolt",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8, Scale: 18.0},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 18,
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "vcu",
			Field:     "uptime",
			Title:     "Uptime",
			Required:  true,
			Chartable: true,
			Unit:      "hour",
			Size:      4,
			Format:    Format{DataType: reflect.Uint32, Scale: 1.0 / 3600.0},
			// Format: (v) => HexToUnsignedInt(cend(v)) / 3600,
			// display: (vf) => Dot(vf, 2),
		},
	}
}

func hbarPacket() []Packet {
	return []Packet{
		{
			Group:     `hbar`,
			Field:     `hbarReverse`,
			Title:     `HBAR Reverse`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     "hbar.mode",
			Field:     "modeDrive",
			Title:     "Mode Drive",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => `${config.mode.drive[vf]} (${vf})`,
		},
		{
			Group:     "hbar.mode",
			Field:     "modeTrip",
			Title:     "Mode Trip",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => `${config.mode.trip[vf]} (${vf})`,
		},
		{
			Group:     "hbar.mode",
			Field:     "modeReport",
			Title:     "Mode Report",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => `${config.mode.report[vf]} (${vf})`,
		},
		{
			Group:     "hbar.trip",
			Field:     "tripA",
			Title:     "Trip A",
			Required:  false,
			Chartable: true,
			Unit:      "Km",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "hbar.trip",
			Field:     "tripB",
			Title:     "Trip B",
			Required:  false,
			Chartable: true,
			Unit:      "Km",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "hbar.trip",
			Field:     "odometer",
			Title:     "Odometer",
			Required:  false,
			Chartable: true,
			Unit:      "Km",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "hbar.report",
			Field:     "rangeReport",
			Title:     "Range Report",
			Required:  false,
			Chartable: true,
			Unit:      "Km",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "hbar.report",
			Field:     "averageReport",
			Title:     "Average Report",
			Required:  false,
			Chartable: true,
			Unit:      "Km/kWh",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
	}
}

func netPacket() []Packet {
	return []Packet{
		{
			Group:     "net",
			Field:     "netSignal",
			Title:     "Net Signal",
			Required:  false,
			Chartable: true,
			Unit:      "%",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "net",
			Field:     "netState",
			Title:     "Net State",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Int8},
			// Format: (v) => HexToSignedInt8(cend(v)),
			// display: (vf) => {
			//   const states = [
			//     "DOWN",
			//     "READY",
			//     "CONFIGURED",
			//     "NETWORK_ON",
			//     "GPRS_ON",
			//     "PDP_ON",
			//     "INTERNET_ON",
			//     "SERVER_ON",
			//     "MQTT_ON",
			//   ];

			//   return `${states[vf + 1]} (${vf})`;
			// },
		},
		{
			Group:     "net",
			Field:     "netIpStatus",
			Title:     "Net IpStatus",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Int8},
			// Format: (v) => HexToSignedInt8(cend(v)),
			// display: (vf) => {
			//   const ipStatus = [
			//     "UNKNOWN",
			//     "IP_INITIAL",
			//     "IP_START",
			//     "IP_CONFIG",
			//     "IP_GPRSACT",
			//     "IP_STATUS",
			//     "CONNECTING",
			//     "CONNECT_OK",
			//     "CLOSING",
			//     "CLOSED",
			//     "PDP_DEACT",
			//   ];
			//   return `${ipStatus[vf + 1]} (${vf})`;
			// },
		},
	}
}

func gpsPacket() []Packet {
	return []Packet{
		{
			Group:     `gps`,
			Field:     `gpsActive`,
			Title:     `GPS Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     "gps",
			Field:     "gpsSatInUse",
			Title:     "GPS Sat. in use",
			Required:  false,
			Chartable: true,
			Unit:      "Sat.",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "gps",
			Field:     "gpsHDOP",
			Title:     "GPS HDOP",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     "gps",
			Field:     "gpsVDOP",
			Title:     "GPS VDOP",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     "gps",
			Field:     "gpsSpeed",
			Title:     "GPS Speed",
			Required:  false,
			Chartable: true,
			Unit:      "Km/hr",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "gps",
			Field:     "gpsHeading",
			Title:     "GPS Heading",
			Required:  false,
			Chartable: true,
			Unit:      "Deg",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8, Scale: 2.0},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 2,
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "gps",
			Field:     "gpsLongitude",
			Title:     "GPS Longitude",
			Required:  false,
			Chartable: true,
			Size:      4,
			Format:    Format{DataType: reflect.Int32, Scale: 0.0000001},
			// Format: (v) => HexToSignedInt32(cend(v)) * 0.0000001,
			// display: (vf) => parseFloat(vf.toFixed(7)),
		},
		{
			Group:     "gps",
			Field:     "gpsLatitude",
			Title:     "GPS Latitude",
			Required:  false,
			Chartable: true,
			Size:      4,
			Format:    Format{DataType: reflect.Int32, Scale: 0.0000001},
			// Format: (v) => HexToSignedInt32(cend(v)) * 0.0000001,
			// display: (vf) => parseFloat(vf.toFixed(7)),
		},
		{
			Group:     "gps",
			Field:     "gpsAltitude",
			Title:     "GPS Altitude",
			Required:  false,
			Chartable: true,
			Unit:      "m",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
	}
}

func memsPacket() []Packet {
	data := []Packet{
		{
			Group:     `mems`,
			Field:     `memsActive`,
			Title:     `MEMS Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `mems`,
			Field:     `memsDetector`,
			Title:     `MEMS Detector`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "Enabled" : "Disabled"),
		},
	}

	axis := []string{"X", "Y", "Z"}
	for _, v := range axis {
		data = append(data, Packet{
			Group:     "mems.accel",
			Field:     fmt.Sprintf("memsAccel%s", v),
			Title:     fmt.Sprintf("MEMS Accel %s", v),
			Required:  false,
			Chartable: true,
			Unit:      "G",
			Size:      2,
			Format:    Format{DataType: reflect.Int16, Scale: 0.01},
			// Format: (v) => HexToSignedInt16(cend(v)) * 0.01,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		})
	}
	for _, v := range axis {
		data = append(data, Packet{
			Group:     "mems.gyro",
			Field:     fmt.Sprintf("memsGyro%s", v),
			Title:     fmt.Sprintf("MEMS Gyro %s", v),
			Required:  false,
			Chartable: true,
			Unit:      "rad/s",
			Size:      2,
			Format:    Format{DataType: reflect.Int16, Scale: 0.1},
			// Format: (v) => HexToSignedInt16(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		})
	}
	for _, v := range []string{"Pitch", "Roll"} {
		data = append(data, Packet{
			Group:     "mems.tilt",
			Field:     fmt.Sprintf("memsTilt%s", v),
			Title:     fmt.Sprintf("MEMS Tilt %s", v),
			Required:  false,
			Chartable: true,
			Unit:      "Deg",
			Size:      2,
			Format:    Format{DataType: reflect.Int16, Scale: 0.1},
			// Format: (v) => HexToSignedInt16(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		})
	}

	return data
}

func remotePacket() []Packet {
	return []Packet{
		{
			Group:     `remote`,
			Field:     `remoteActive`,
			Title:     `Remote Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `remote`,
			Field:     `remoteNearby`,
			Title:     `Remote Nearby`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
	}
}

func fingerPacket() []Packet {
	return []Packet{
		{
			Group:     `finger`,
			Field:     `fingerVerified`,
			Title:     `Finger Verified`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     "finger",
			Field:     "fingerDriverID",
			Title:     "Finger Driver ID",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => {
			//   if (vf === 0) return "NONE";
			//   return IntToHex(vf, 2).toUpperCase();
			// },
		},
	}
}

func audioPacket() []Packet {
	return []Packet{
		{
			Group:     `audio`,
			Field:     `audioActive`,
			Title:     `Audio Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `audio`,
			Field:     `audioMute`,
			Title:     `Audio Mute`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `audio`,
			Field:     `audioVolume`,
			Title:     `Audio Volume`,
			Required:  false,
			Chartable: true,
			Unit:      "%",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
	}
}

func hmi1Packet() []Packet {
	return []Packet{
		{
			Group:     `hmi1`,
			Field:     `hmi1Active`,
			Title:     `HMI1 Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
	}
}

func bmsPacket() []Packet {
	data := []Packet{
		{
			Group:     `bms`,
			Field:     `bmsActive`,
			Title:     `BMS Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `bms`,
			Field:     `bmsRun`,
			Title:     `BMS Run`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `bms`,
			Field:     `bmsSoc`,
			Title:     `BMS SOC`,
			Required:  false,
			Chartable: true,
			Unit:      "%",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     `bms`,
			Field:     `bmsFault`,
			Title:     `BMS Fault`,
			Required:  false,
			Chartable: true,
			Size:      2,
			Format:    Format{DataType: reflect.Uint16},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => IntToHex(vf, 4).toUpperCase(),
		},
	}

	for _, v := range []string{"One", "Two"} {
		data = append(data, []Packet{
			{
				Group:    fmt.Sprintf("bms.%s", v),
				Field:    fmt.Sprintf("bms%sId", v),
				Title:    fmt.Sprintf("BMS %s ID", v),
				Required: false,
				Size:     4,
				// Format: (v) => HexToUnsignedInt(cend(v)),
				// display: (vf) => {
				//   if (vf === 0xffffffff) return "NONE";
				//   return IntToHex(vf, 8).toUpperCase();
				// },
			},
			{
				Group:     fmt.Sprintf("bms.%s", v),
				Field:     fmt.Sprintf("bms%sFault", v),
				Title:     fmt.Sprintf("BMS %s Fault", v),
				Required:  false,
				Chartable: true,
				Size:      2,
				// Format: (v) => HexToUnsignedInt(cend(v)),
				// display: (vf) => IntToHex(vf, 4).toUpperCase(),
			},
			{
				Group:     fmt.Sprintf("bms.%s", v),
				Field:     fmt.Sprintf("bms%sVoltage", v),
				Title:     fmt.Sprintf("BMS %s Voltage", v),
				Required:  false,
				Chartable: true,
				Unit:      "Volt",
				Size:      2,
				// Format: (v) => HexToUnsignedInt(cend(v)) * 0.01,
				// display: (vf) => parseFloat(vf.toFixed(2)),
			},
			{
				Group:     fmt.Sprintf("bms.%s", v),
				Field:     fmt.Sprintf("bms%sCurrent", v),
				Title:     fmt.Sprintf("BMS %s Current", v),
				Required:  false,
				Chartable: true,
				Unit:      "Ampere",
				Size:      2,
				// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
				// display: (vf) => parseFloat(vf.toFixed(2)),
			},
			{
				Group:     fmt.Sprintf("bms.%s", v),
				Field:     fmt.Sprintf("bms%sSoc", v),
				Title:     fmt.Sprintf("BMS %s SoC", v),
				Required:  false,
				Chartable: true,
				Unit:      "%",
				Size:      1,
				// Format: (v) => HexToUnsignedInt(cend(v)),
				// display: (vf) => Dot(vf),
			},
			{
				Group:     fmt.Sprintf("bms.%s", v),
				Field:     fmt.Sprintf("bms%sTemperature", v),
				Title:     fmt.Sprintf("BMS %s Temperature", v),
				Required:  false,
				Chartable: true,
				Unit:      "Celcius",
				Size:      2,
				// Format: (v) => HexToUnsignedInt(cend(v)),
				// display: (vf) => parseFloat(vf.toFixed(2)),
			},
		}...)
	}
	return data
}

func mcuPacket() []Packet {
	data := []Packet{
		{
			Group:     `mcu`,
			Field:     `mcuActive`,
			Title:     `MCU Active`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `mcu`,
			Field:     `mcuRun`,
			Title:     `MCU Run`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `mcu`,
			Field:     `mcuReverse`,
			Title:     `MCU Reverse`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     "mcu",
			Field:     "mcuDriveMode",
			Title:     "MCU Drive Mode",
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => `${config.mode.drive[vf]} (${vf})`,
		},
		{
			Group:     `mcu`,
			Field:     `mcuSpeed`,
			Title:     `MCU Speed`,
			Required:  false,
			Chartable: true,
			Unit:      "km/hr",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     `mcu`,
			Field:     `mcuRpm`,
			Title:     `MCU RPM`,
			Required:  false,
			Chartable: true,
			Unit:      "rpm",
			Size:      2,
			Format:    Format{DataType: reflect.Int16},
			// Format: (v) => HexToSignedInt16(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     `mcu`,
			Field:     `mcuTemperature`,
			Title:     `MCU Temperature`,
			Required:  false,
			Chartable: true,
			Unit:      "Celcius",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     `mcu.fault`,
			Field:     `mcuFaultPost`,
			Title:     `MCU Fault Post`,
			Required:  false,
			Chartable: true,
			Size:      4,
			Format:    Format{DataType: reflect.Uint32},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => IntToHex(vf, 8).toUpperCase(),
		},
		{
			Group:     `mcu.fault`,
			Field:     `mcuFaultRun`,
			Title:     `MCU Fault Run`,
			Required:  false,
			Chartable: true,
			Size:      4,
			Format:    Format{DataType: reflect.Uint32},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => IntToHex(vf, 8).toUpperCase(),
		},
		{
			Group:     `mcu.torque`,
			Field:     `mcuTorqueCommand`,
			Title:     `MCU Torque Command`,
			Required:  false,
			Chartable: true,
			Unit:      "Nm",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     `mcu.torque`,
			Field:     `mcuTorqueFeedback`,
			Title:     `MCU Torque Feedback`,
			Required:  false,
			Chartable: true,
			Unit:      "Nm",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     `mcu.dcbus`,
			Field:     `mcuDcbusCurrent`,
			Title:     `MCU DCBus Current`,
			Required:  false,
			Chartable: true,
			Unit:      "A",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     `mcu.dcbus`,
			Field:     `mcuDcbusVoltage`,
			Title:     `MCU DCBus Voltage`,
			Required:  false,
			Chartable: true,
			Unit:      "V",
			Size:      2,
			Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
			// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
			// display: (vf) => parseFloat(vf.toFixed(2)),
		},
		{
			Group:     `mcu.inverter`,
			Field:     `mcuInverterEnabled`,
			Title:     `MCU Inverter Enabled`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `mcu.inverter`,
			Field:     `mcuInverterLockout`,
			Title:     `MCU Inverter Lockout`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => (vf ? "YES" : "NO"),
		},
		{
			Group:     `mcu.inverter`,
			Field:     `mcuInverterDischarge`,
			Title:     `MCU Inverter Discharge`,
			Required:  false,
			Chartable: true,
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => {
			//   const states = [
			//     "DISABLED",
			//     "ENABLED",
			//     "CHECK",
			//     "OCCURING",
			//     "COMPLETED",
			//   ];
			//   return `${states[vf]} (${vf})`;
			// },
		},
		{
			Group:     `mcu.template`,
			Field:     `mcuTemplateMaxRpm`,
			Title:     `MCU Template Max RPM`,
			Required:  false,
			Chartable: true,
			Unit:      "rpm",
			Size:      2,
			Format:    Format{DataType: reflect.Int16},
			// Format: (v) => HexToSignedInt16(cend(v)),
			// display: (vf) => Dot(vf),
		},
		{
			Group:     "mcu.template",
			Field:     "mcuTemplateMaxSpeed",
			Title:     "MCU Template Max Speed",
			Required:  false,
			Chartable: true,
			Unit:      "Km/hr",
			Size:      1,
			Format:    Format{DataType: reflect.Uint8},
			// Format: (v) => HexToUnsignedInt(cend(v)),
			// display: (vf) => Dot(vf),
		},
	}

	driveMode := []string{"ECONOMY", "STANDARD", "SPORT"}
	for _, v := range driveMode {
		data = append(data, []Packet{
			{
				Group:     fmt.Sprintf("mcu.template.%s", v),
				Field:     fmt.Sprintf("mcuTemplate%sDiscur", v),
				Title:     fmt.Sprintf("MCU Template %s Discur", v),
				Required:  false,
				Chartable: true,
				Unit:      "A",
				Size:      2,
				Format:    Format{DataType: reflect.Uint16},
				// Format: (v) => HexToUnsignedInt(cend(v)),
				// display: (vf) => Dot(vf),
			},
			{
				Group:     fmt.Sprintf("mcu.template.%s", v),
				Field:     fmt.Sprintf("mcuTemplate%sTorque", v),
				Title:     fmt.Sprintf("MCU Template %s Torque", v),
				Required:  false,
				Chartable: true,
				Unit:      "Nm",
				Size:      2,
				Format:    Format{DataType: reflect.Uint16, Scale: 0.1},
				// Format: (v) => HexToUnsignedInt(cend(v)) * 0.1,
				// display: (vf) => parseFloat(vf.toFixed(2)),
			},
		}...)
	}

	return data
}

func taskPacket() []Packet {
	tasks := []string{
		"managerTask",
		"networkTask",
		"reporterTask",
		"commandTask",
		"memsTask",
		"remoteTask",
		"fingerTask",
		"audioTask",
		"gateTask",
		"canRxTask",
		"canTxTask",
	}

	data := []Packet{}
	for _, v := range tasks {
		data = append(data,
			Packet{
				Group:     "task.stack",
				Field:     fmt.Sprintf("%sStack", v),
				Title:     fmt.Sprintf("%s stack", strings.Title(v)),
				Required:  false,
				Chartable: true,
				Unit:      "Bytes",
				Size:      2,
				Format:    Format{DataType: reflect.Uint16},
				// Format: (v) => HexToUnsignedInt(cend(v)),
				// display: (vf) => Dot(vf),
			},
		)
	}
	for _, v := range tasks {
		data = append(data,
			Packet{
				Group:     "task.wakeup",
				Field:     fmt.Sprintf("%sWakeup", v),
				Title:     fmt.Sprintf("%s wakeup", strings.Title(v)),
				Required:  false,
				Chartable: true,
				Unit:      "s",
				Size:      1,
				Format:    Format{DataType: reflect.Uint8},
				//       Format: (v) => HexToUnsignedInt(cend(v)),
				//       display: (vf) => {
				//         let prefix = vf < 255 ? "<" : ">";
				//         return `${prefix} ${Dot(vf)}`;
				//       },
			},
		)
	}

	return data
}
