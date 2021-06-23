package packet

type Packet struct {
	group     string
	field     string
	title     string
	required  bool
	chartable bool
	unit      string
	header    bool
	size      int
	format    Format
}
