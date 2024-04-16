package chunk

import "encoding/binary"

type LayerFlag uint16

const (
	FlagVisible          LayerFlag = 1
	FlagEditable         LayerFlag = 2
	FlagLockMovement     LayerFlag = 4
	FlagBackground       LayerFlag = 8
	FlagPreferLinkedCels LayerFlag = 16
	FlagReferenceLayer   LayerFlag = 32
	FlagCollapsed        LayerFlag = 64
)

type LayerType uint16

const (
	LayerTypeNormal  LayerType = 0
	LayerTypeGroup   LayerType = 1
	LayerTypeTilemap LayerType = 2
)

type BlendMode uint16

const (
	BlendModeNormal     BlendMode = 0
	BlendModeMultiply   BlendMode = 1
	BlendModeScreen     BlendMode = 2
	BlendModeOverlay    BlendMode = 3
	BlendModeDarken     BlendMode = 4
	BlendModeLighten    BlendMode = 5
	BlendModeColorDodge BlendMode = 6
	BlendModeColorBurn  BlendMode = 7
	BlendModeHardLight  BlendMode = 8
	BlendModeSoftLight  BlendMode = 9
	BlendModeDifference BlendMode = 10
	BlendModeExclusion  BlendMode = 11
	BlendModeHue        BlendMode = 12
	BlendModeSaturation BlendMode = 13
	BlendModeColor      BlendMode = 14
	BlendModeLuminosity BlendMode = 15
	BlendModeAddition   BlendMode = 16
	BlendModeSubtract   BlendMode = 17
	BlendModeDivide     BlendMode = 18
)

type Layer struct {
	Flags         LayerFlag // 0-1
	Type          LayerType // 2-3
	ChildLevel    uint16    // 4-5
	DefaultWidth  uint16    // 6-7 ignored
	DefaultHeight uint16    // 8-9 ignored
	BlendMode     BlendMode // 10-11
	Opacity       byte      // 12
	Name          string    // 15 - varies
	TilesetIndex  uint32    // only if TypeGroup
}

func (chunk Chunk) AsLayer() Layer {
	layer := Layer{
		Flags:         LayerFlag(binary.LittleEndian.Uint16(chunk.Data[0:2])),
		Type:          LayerType(binary.LittleEndian.Uint16(chunk.Data[2:4])),
		ChildLevel:    binary.LittleEndian.Uint16(chunk.Data[4:6]),
		DefaultWidth:  binary.LittleEndian.Uint16(chunk.Data[6:8]),
		DefaultHeight: binary.LittleEndian.Uint16(chunk.Data[8:10]),
		BlendMode:     BlendMode(binary.LittleEndian.Uint16(chunk.Data[10:12])),
		Opacity:       chunk.Data[12],
	}
	var name string
	var index uint32
	if layer.Type == LayerTypeGroup {
		name = string(chunk.Data[15 : len(chunk.Data)-4])
		index = binary.LittleEndian.Uint32(chunk.Data[len(chunk.Data)-4:])
	} else {
		name = string(chunk.Data[15:len(chunk.Data)])
	}
	layer.Name = name
	layer.TilesetIndex = index
	return layer
}
