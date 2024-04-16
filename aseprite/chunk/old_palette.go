package chunk

import (
	"encoding/binary"
	"image/color"
)

type OldPaletteChunk []OldPalettePacket

type OldPalettePacket struct {
	Skip   byte
	Colors []color.Color
}

func (chunk Chunk) AsOldPalette() OldPaletteChunk {
	packets := []OldPalettePacket{}
	var pointer uint16 = 2
	for range binary.LittleEndian.Uint16(chunk.Data[0:2]) {
		skip := chunk.Data[pointer]
		colors := []color.Color{}
		count := uint16(chunk.Data[pointer+1])
		if count == 0 {
			count = 256
		}
		for c := range count {
			offset := pointer + 3*uint16(c) + 1
			colors = append(colors, color.RGBA{
				R: chunk.Data[offset+1],
				G: chunk.Data[offset+2],
				B: chunk.Data[offset+3],
				A: 255,
			})
		}
		packets = append(packets, OldPalettePacket{Skip: skip, Colors: colors})
	}
	return packets
}
