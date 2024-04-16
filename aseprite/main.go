package aseprite

import (
	"io"
)

type Image struct {
	Header Header
	Frames []Frame
}

func ReadImage(reader io.Reader) (Image, error) {
	header, err := ReadHeader(reader)
	if err != nil {
		return Image{}, err
	}
	frames := []Frame{}
	for range header.Frames {
		frame, err := ReadFrame(reader)
		if err != nil {
			return Image{}, err
		}
		frames = append(frames, frame)
	}
	return Image{Header: header, Frames: frames}, nil
}
