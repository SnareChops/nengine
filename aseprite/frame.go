package aseprite

import (
	"encoding/binary"
	"io"

	"github.com/SnareChops/nengine/aseprite/chunk"
)

type Frame struct {
	Header FrameHeader
	Chunks []chunk.Chunk
}

func ReadFrame(reader io.Reader) (Frame, error) {
	header, err := ReadFrameHeader(reader)
	if err != nil {
		return Frame{}, err
	}
	count := header.NewNumberOfChunks
	if count == 0 {
		count = uint32(header.OldNumberOfChunks)
	}
	chunks := []chunk.Chunk{}
	for range count {
		chunk, err := chunk.ReadChunk(reader)
		if err != nil {
			return Frame{}, err
		}
		chunks = append(chunks, chunk)
	}
	return Frame{Header: header, Chunks: chunks}, nil
}

type FrameHeader struct {
	Size              uint32 // 0-3
	OldNumberOfChunks uint16 // 6-7
	Duration          uint16 // 8-9
	NewNumberOfChunks uint32 // 12-15
}

func ReadFrameHeader(reader io.Reader) (FrameHeader, error) {
	buffer := make([]byte, 16)
	_, err := reader.Read(buffer)
	if err != nil {
		return FrameHeader{}, err
	}
	return FrameHeader{
		Size:              binary.LittleEndian.Uint32(buffer[0:4]),
		OldNumberOfChunks: binary.LittleEndian.Uint16(buffer[6:8]),
		Duration:          binary.LittleEndian.Uint16(buffer[8:10]),
		NewNumberOfChunks: binary.LittleEndian.Uint32(buffer[12:16]),
	}, nil
}
