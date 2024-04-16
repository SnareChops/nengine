package chunk

type ColorProfileType uint16

const (
	NoProfile ColorProfileType = 0
	UseSRGB   ColorProfileType = 1
	UseICC    ColorProfileType = 2
)

type ColorProfileFlags uint16

const (
	ColorProfileFlagsSpecialFixedGamma ColorProfileFlags = 1
)

type ColorProfile struct {
	Type  ColorProfileType // 0-1
	Flags ColorProfileFlags
}
