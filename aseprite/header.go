package aseprite

import (
	"encoding/binary"
	"io"
)

type Header struct {
	FileSize       uint32 // 0-3
	Frames         uint16 // 6-7
	WidthInPixels  uint16 // 8-9
	HeightInPixels uint16 // 10-11
	ColorDepth     uint16 // 12-13
	Flags          uint32 // 14-17
	Speed          uint16 // Deprecated 18-19
	PaletteEntry   byte   // 28
	NumberOfColors uint16 // 32-33
	PixelWidth     byte   // 34
	PixelHeight    byte   // 35
	GridX          int16  // 36-37
	GridY          int16  // 38-39
	GridWidth      uint16 // 40-41
	GridHeight     uint16 // 42-43
}

func ReadHeader(reader io.Reader) (Header, error) {
	buffer := make([]byte, 128)
	_, err := reader.Read(buffer)
	if err != nil {
		return Header{}, err
	}
	return Header{
		FileSize:       binary.LittleEndian.Uint32(buffer[0:4]),
		Frames:         binary.LittleEndian.Uint16(buffer[6:8]),
		WidthInPixels:  binary.LittleEndian.Uint16(buffer[8:10]),
		HeightInPixels: binary.LittleEndian.Uint16(buffer[10:12]),
		ColorDepth:     binary.LittleEndian.Uint16(buffer[12:14]),
		Flags:          binary.LittleEndian.Uint32(buffer[14:18]),
		Speed:          binary.LittleEndian.Uint16(buffer[18:20]),
		PaletteEntry:   buffer[28],
	}, nil
}
