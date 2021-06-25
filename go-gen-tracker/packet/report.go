package packet

type ReportSimplePacket struct {
	Header HeaderPacket
	VCU    VcuPacket
	GPS    GpsPacket
}
type ReportFullPacket struct {
	ReportSimplePacket
	HBAR   HbarPacket
	NET    NetPacket
	MEMS   MemsPacket
	Remote RemotePacket
	Finger FingerPacket
	Audio  AudioPacket
	HMI1   Hmi1Packet
	BMS    BmsPacket
	MCU    McuPacket
	TASK   TaskPacket
}

type VcuPacket struct {
	FrameID     uint8   `chartable:""`
	LogDatetime [7]byte `type:"datetime" chartable:""`
	State       int8    `chartable:""`
	Events      uint16  `chartable:""`
	LogBuffered uint8   `chartable:""`
	BatVoltage  uint8   `unit:"mVolt" scale:"18.0" chartable:""`
	Uptime      uint32  `unit:"hour" scale:"0.000277" chartable:""`
}

type GpsPacket struct {
	Active    bool   `chartable:""`
	SatInUse  uint8  `unit:"Sat" chartable:""`
	HDOP      uint8  `scale:"0.1" chartable:""`
	VDOP      uint8  `scale:"0.1" chartable:""`
	Speed     uint8  `unit:"Kph" chartable:""`
	Heading   uint8  `unit:"Deg" scale:"2.0" chartable:""`
	Longitude int32  `scale:"0.0000001" chartable:""`
	Latitude  int32  `scale:"0.0000001" chartable:""`
	Altitude  uint16 `unit:"m" scale:"0.1" chartable:""`
}

type HbarPacket struct {
	Reverse bool `chartable:""`
	Mode    struct {
		Drive      uint8 `chartable:""`
		Trip       uint8 `chartable:""`
		Prediction uint8 `chartable:""`
	}
	Trip struct {
		A        uint16 `unit:"Km" chartable:""`
		B        uint16 `unit:"Km" chartable:""`
		Odometer uint16 `unit:"Km" chartable:""`
	}
	Prediction struct {
		Range      uint8 `unit:"Km" chartable:""`
		Efficiency uint8 `unit:"Km/Kwh" chartable:""`
	}
}

type NetPacket struct {
	Signal   uint8 `unit:"%" chartable:""`
	State    int8  `chartable:""`
	IpStatus int8  `chartable:""`
}

type MemsPacket struct {
	Active   bool `chartable:""`
	Detector bool `chartable:""`
	Accel    struct {
		X int16 `unit:"G" scale:"0.01" chartable:""`
		Y int16 `unit:"G" scale:"0.01" chartable:""`
		Z int16 `unit:"G" scale:"0.01" chartable:""`
	}
	Gyro struct {
		X int16 `unit:"rad/s" scale:"0.1" chartable:""`
		Y int16 `unit:"rad/s" scale:"0.1" chartable:""`
		Z int16 `unit:"rad/s" scale:"0.1" chartable:""`
	}
	Tilt struct {
		Pitch int16 `unit:"Deg" scale:"0.1" chartable:""`
		Roll  int16 `unit:"Deg" scale:"0.1" chartable:""`
	}
	Total struct {
		Accel       uint16 `unit:"G" scale:"0.01" chartable:""`
		Gyro        uint16 `unit:"rad/s" scale:"0.1" chartable:""`
		Tilt        uint16 `unit:"Deg" scale:"0.1" chartable:""`
		Temperature uint16 `unit:"C" scale:"0.1" chartable:""`
	}
}

type RemotePacket struct {
	Active bool `chartable:""`
	Nearby bool `chartable:""`
}

type FingerPacket struct {
	Verified bool  `chartable:""`
	DriverID uint8 `chartable:""`
}

type AudioPacket struct {
	Active bool  `chartable:""`
	Mute   bool  `chartable:""`
	Volume uint8 `unit:"%" chartable:""`
}
type Hmi1Packet struct {
	Active bool `chartable:""`
}

type BmsPacket struct {
	Active bool   `chartable:""`
	Run    bool   `chartable:""`
	SOC    uint8  `unit:"%" chartable:""`
	Fault  uint16 `chartable:""`
	Pack   [2]struct {
		ID          uint32
		Fault       uint16 `chartable:""`
		Voltage     uint16 `unit:"Volt" scale:"0.01" chartable:""`
		Current     uint16 `unit:"Ampere" scale:"0.1" chartable:""`
		SOC         uint8  `unit:"%" chartable:""`
		Temperature uint16 `unit:"C" chartable:""`
	}
}

type McuPacket struct {
	Active      bool   `chartable:""`
	Run         bool   `chartable:""`
	Reverse     bool   `chartable:""`
	DriveMode   uint8  `chartable:""`
	Speed       uint8  `unit:"Kph" chartable:""`
	RPM         int16  `unit:"rpm" chartable:""`
	Temperature uint16 `unit:"C" scale:"0.1" chartable:""`
	Fault       struct {
		Post uint32 `chartable:""`
		Run  uint32 `chartable:""`
	}
	Torque struct {
		Command  uint16 `unit:"Nm" scale:"0.1" chartable:""`
		Feedback uint16 `unit:"Nm" scale:"0.1" chartable:""`
	}
	DCBus struct {
		Current uint16 `unit:"A" scale:"0.1" chartable:""`
		Voltage uint16 `unit:"V" scale:"0.1" chartable:""`
	}
	Inverter struct {
		Enabled   bool  `chartable:""`
		Lockout   bool  `chartable:""`
		Discharge uint8 `chartable:""`
	}
	Template struct {
		MaxRPM    int16 `unit:"rpm" chartable:""`
		MaxSpeed  uint8 `unit:"Kph" chartable:""`
		DriveMode [3]struct {
			Discur uint16 `unit:"A" chartable:""`
			Torque uint16 `unit:"Nm" scale:"0.1" chartable:""`
		}
	}
}

type TaskPacket struct {
	Stack struct {
		Manager  uint16 `unit:"Bytes" chartable:""`
		Network  uint16 `unit:"Bytes" chartable:""`
		Reporter uint16 `unit:"Bytes" chartable:""`
		Command  uint16 `unit:"Bytes" chartable:""`
		Mems     uint16 `unit:"Bytes" chartable:""`
		Remote   uint16 `unit:"Bytes" chartable:""`
		Finger   uint16 `unit:"Bytes" chartable:""`
		Audio    uint16 `unit:"Bytes" chartable:""`
		Gate     uint16 `unit:"Bytes" chartable:""`
		CanRX    uint16 `unit:"Bytes" chartable:""`
		CanTX    uint16 `unit:"Bytes" chartable:""`
	}
	Wakeup struct {
		Manager  uint8 `unit:"s" chartable:""`
		Network  uint8 `unit:"s" chartable:""`
		Reporter uint8 `unit:"s" chartable:""`
		Command  uint8 `unit:"s" chartable:""`
		Mems     uint8 `unit:"s" chartable:""`
		Remote   uint8 `unit:"s" chartable:""`
		Finger   uint8 `unit:"s" chartable:""`
		Audio    uint8 `unit:"s" chartable:""`
		Gate     uint8 `unit:"s" chartable:""`
		CanRX    uint8 `unit:"s" chartable:""`
		CanTX    uint8 `unit:"s" chartable:""`
	}
}
