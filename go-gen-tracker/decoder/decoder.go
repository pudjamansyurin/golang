package decoder

type M map[string]interface{}

type FrameID uint8

const (
	FRAME_INVALID FrameID = iota
	FRAME_SIMPLE
	FRAME_FULL
)
