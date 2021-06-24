package packet

type ReportSimplePacket struct {
	VCU VcuPacket
	GPS GpsPacket
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
	FrameID     uint8
	LogDatetime [7]byte
	State       int8
	Events      uint16
	LogBuffered uint8
	BatVoltage  uint8
	Uptime      uint32
}

type GpsPacket struct {
	Active    uint8
	SatInUse  uint8
	HDOP      uint8
	VDOP      uint8
	Speed     uint8
	Heading   uint8
	Longitude int32
	Latitude  int32
	Altitude  uint16
}

type HbarPacket struct {
	Reverse uint8
	Mode    struct {
		Drive      uint8
		Trip       uint8
		Prediction uint8
	}
	Trip struct {
		A        uint16
		B        uint16
		Odometer uint16
	}
	Prediction struct {
		Range      uint8
		Efficiency uint8
	}
}

type NetPacket struct {
	Signal   uint8
	State    int8
	IpStatus int8
}

type MemsPacket struct {
	Active   uint8
	Detector uint8
	Accel    struct {
		X int16
		Y int16
		Z int16
	}
	Gyro struct {
		X int16
		Y int16
		Z int16
	}
	Tilt struct {
		Pitch int16
		Roll  int16
	}
	Total struct {
		Accel       uint16
		Gyro        uint16
		Tilt        uint16
		Temperature uint16
	}
}

type RemotePacket struct {
	Active uint8
	Nearby uint8
}

type FingerPacket struct {
	Verified uint8
	DriverID uint8
}

type AudioPacket struct {
	Active uint8
	Mute   uint8
	Volume uint8
}
type Hmi1Packet struct {
	Active uint8
}

type BmsPacket struct {
	Active uint8
	Run    uint8
	SOC    uint8
	Fault  uint16
	Pack   [2]struct {
		ID          uint32
		Fault       uint16
		Voltage     uint16
		Current     uint16
		SOC         uint8
		Temperature uint16
	}
}

type McuPacket struct {
	Active      uint8
	Run         uint8
	Reverse     uint8
	DriveMode   uint8
	Speed       uint8
	RPM         int16
	Temperature uint16
	Fault       struct {
		Post uint32
		Run  uint32
	}
	Torque struct {
		Command  uint16
		Feedback uint16
	}
	DCBus struct {
		Current uint16
		Voltage uint16
	}
	Inverter struct {
		Enabled   uint8
		Lockout   uint8
		Discharge uint8
	}
	Template struct {
		MaxRPM    int16
		MaxSpeed  uint8
		DriveMode [3]struct {
			Discur uint16
			Torque uint16
		}
	}
}

type TaskPacket struct {
	Stack struct {
		Manager  uint16
		Network  uint16
		Reporter uint16
		Command  uint16
		Mems     uint16
		Remote   uint16
		Finger   uint16
		Audio    uint16
		Gate     uint16
		CanRX    uint16
		CanTX    uint16
	}
	Wakeup struct {
		Manager  uint8
		Network  uint8
		Reporter uint8
		Command  uint8
		Mems     uint8
		Remote   uint8
		Finger   uint8
		Audio    uint8
		Gate     uint8
		CanRX    uint8
		CanTX    uint8
	}
}
