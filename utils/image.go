package utils

import "github.com/hajimehoshi/ebiten/v2"

// ImageChunk represents a chunk of an image that
// can be reassembled to the full image again
// Note: This is usually used to split a large image into
// smaller pieces so it can be consumed by the Renderer
type ImageChunk struct {
	X      int
	Y      int
	Width  int
	Height int
	Image  *ebiten.Image
}

func ChunkImage(img *ebiten.Image, size int) []*ImageChunk {
	// Obtain the size of the image
	width, height := img.Size()

	var newImages []*ImageChunk
	var subImg *ebiten.Image

	// Loop over the image by segments of 4000px
	for y := 0; y < height; y += size {
		for x := 0; x < width; x += size {

			var subWidth, subHeight int

			if x+size > width {
				subWidth = width - x
			} else {
				subWidth = size
			}

			if y+size > height {
				subHeight = height - y
			} else {
				subHeight = size
			}

			subImg = ebiten.NewImage(subWidth, subHeight)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(-x), float64(-y))

			// Draw the sub-image onto the new image
			subImg.DrawImage(img, op)

			// Add the new image to the slice
			newImages = append(newImages, &ImageChunk{
				X:      x,
				Y:      y,
				Width:  subWidth,
				Height: subHeight,
				Image:  subImg,
			})
		}
	}

	return newImages
}
