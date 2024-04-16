package chunk

import (
	"encoding/binary"
	"io"
)

type ChunkType uint16

const (
	ChunkTypeOldPalette   ChunkType = 0x0004
	ChunkTypeOldPalette2  ChunkType = 0x0011
	ChunkTypeLayer        ChunkType = 0x2004
	ChunkTypeCel          ChunkType = 0x2005
	ChunkTypeCelExtra     ChunkType = 0x2006
	ChunkTypeColorProfile ChunkType = 0x2007
	ChunkExternalFiles    ChunkType = 0x2008
	ChunkTypeMask         ChunkType = 0x2016 // Deprecated
	ChunkTypeTags         ChunkType = 0x2018
	ChunkTypePalette      ChunkType = 0x2019
	ChunkTypeUserData     ChunkType = 0x2020
	ChunkTypeSlice        ChunkType = 0x2022
	ChunkTypeTileset      ChunkType = 0x2023
)

type Chunk struct {
	Size uint32
	Type ChunkType
	Data []byte
}

func ReadChunk(reader io.Reader) (Chunk, error) {
	header := make([]byte, 6)
	_, err := reader.Read(header)
	if err != nil {
		return Chunk{}, err
	}
	chunk := Chunk{
		Size: binary.LittleEndian.Uint32(header[0:4]),
		Type: ChunkType(binary.LittleEndian.Uint16(header[4:6])),
	}
	buffer := make([]byte, chunk.Size-6)
	_, err = reader.Read(buffer)
	if err != nil {
		return Chunk{}, err
	}
	chunk.Data = buffer
	return chunk, nil
}
