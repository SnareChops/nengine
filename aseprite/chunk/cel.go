package chunk

import (
	"image"
)

type CelType uint16

const (
	CelTypeRaw               CelType = 0
	CelTypeLinked            CelType = 1
	CelTypeCompressedImage   CelType = 2
	CelTypeCompressedTilemap CelType = 3
)

type Cel struct {
	Index   uint16  // 0-1
	X       int16   // 2-3
	Y       int16   // 4-5
	Opacity byte    // 6
	Type    CelType // 7-8
	ZIndex  int16   // 9-10
	Image   image.NRGBA
	Link    uint16
	Tilemap byte // TODO
}

// func (chunk Chunk) AsCel() (Cel, error) {
// 	reader := bytes.NewReader(chunk.Data[2:6])
// 	pos := struct{ X, Y int16 }{}
// 	err := binary.Read(reader, binary.LittleEndian, &pos)
// 	if err != nil {
// 		return Cel{}, err
// 	}
// 	reader = bytes.NewReader(chunk.Data[9:11])
// 	var zindex int16
// 	err = binary.Read(reader, binary.LittleEndian, &zindex)
// 	cel := Cel{
// 		Index:   binary.LittleEndian.Uint16(chunk.Data[0:2]),
// 		X:       pos.X,
// 		Y:       pos.Y,
// 		Opacity: chunk.Data[6],
// 		Type:    CelType(binary.LittleEndian.Uint16(chunk.Data[7:9])),
// 		ZIndex:  zindex,
// 	}
// 	switch cel.Type {
// 	case CelTypeRaw:
// 		cel.Image = extractRawImage(chunk.Data[16:])
// 	}
// }

// func extractRawImage(data []byte) image.NRGBA {
// 	width := binary.LittleEndian.Uint16(data[0:2])
// 	height := binary.LittleEndian.Uint16(data[2:4])

// }
