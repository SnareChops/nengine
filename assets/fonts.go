package assets

import (
	"image"
	"image/color"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/shaders"
	"github.com/hajimehoshi/ebiten/v2"
)

type Font struct {
	Height int
	Data   map[rune]*Char
}

type Char struct {
	Width int
	Image *ebiten.Image
}

var Font5 *Font
var Font10 *Font
var Font15 *Font
var Font20 *Font

func InitFonts() {
	// Open Fonts.png file
	file, err := Open("assets/fonts/Fonts.png")
	if err != nil {
		panic(err)
	}
	img, err := PNGDecode(file)
	if err != nil {
		panic(err)
	}

	Font5 = loadFontFromImage(1, 5, img)
	Font10 = loadFontFromImage(7, 10, img)
	Font15 = loadFontFromImage(18, 15, img)
	Font20 = loadFontFromImage(34, 20, img)
}

type ReadableImage interface {
	image.Image
	SubImage(image.Rectangle) image.Image
}

func loadFontFromImage(start int, height int, img image.Image) *Font {
	font := &Font{
		Height: height,
		Data:   map[rune]*Char{},
	}
	cursor := 0
	for ascii := rune(32); ascii < 127; ascii++ {
		cursor += 1
		begin := cursor
		// Start at the first pixel specified by start
		for {
			if _, g, _, _ := img.At(cursor, start).RGBA(); g == 0xffff {
				break
			}
			cursor += 1
		}
		rect := image.Rect(begin, start, cursor, start+height)
		font.Data[ascii] = &Char{
			Width: cursor - begin,
			Image: ebiten.NewImageFromImage(img.(ReadableImage).SubImage(rect)),
		}
	}
	return font
}

type ImageLike interface {
	*image.RGBA | *ebiten.Image
	Set(x, y int, clr color.Color)
	Bounds() image.Rectangle
}

type Text struct {
	*bounds.Position
	width   int
	height  int
	kerning int
	leading int
	lines   []*Line
}

func NewText(kerning int, leading int, lines ...*Line) *Text {
	return &Text{
		Position: bounds.Point(0, 0, 0),
		height:   -leading,
		kerning:  kerning,
		leading:  leading,
		lines:    lines,
	}
}

func (self *Text) AddLine(line *Line) {
	if line.Width() > self.width {
		self.width = line.Width()
	}
	self.height += line.height + self.leading
	self.lines = append(self.lines, line)
}

func (self *Text) Dx() int {
	w, _ := self.Size()
	return w
}

func (self *Text) Dy() int {
	_, h := self.Size()
	return h
}

func (self *Text) Size() (int, int) {
	maxWidth := 0
	totalHeight := -self.leading
	for _, line := range self.lines {
		if line.Width() > maxWidth {
			maxWidth = line.Width()
		}
		totalHeight += line.height + self.leading
	}
	return maxWidth, totalHeight
}

func (self *Text) Kerning() int {
	return self.kerning
}

func (self *Text) Leading() int {
	return self.leading
}

type Line struct {
	width   int
	height  int
	kerning int
	letters []Letter
}

// NewLine Creates a new line with the given kerning and letters
func NewLine(kerning int, letters ...Letter) *Line {
	line := &Line{width: -kerning, kerning: kerning}
	for _, letter := range letters {
		line.AddLetter(letter)
	}
	return line
}

func (self *Line) AddLetter(letter Letter) {
	self.width += letter.Image.Bounds().Dx() + self.kerning
	height := letter.Image.Bounds().Dy()
	if height > self.height {
		self.height = height
	}
	self.letters = append(self.letters, letter)
}

// Size Returns the (width, height) of the line
// Note: A line with no letters will have a negative width
func (self *Line) Size() (int, int) {
	return self.width, self.height
}

// Width Returns the width of the line
// Note: A line with no letters will have a negative width
func (self *Line) Width() int {
	return self.width
}

func (self *Line) Len() int {
	return len(self.letters)
}

func (self *Line) At(index int) Letter {
	return self.letters[index]
}

type Letter struct {
	Image *ebiten.Image
	Char  rune
}

func DrawString(dest *ebiten.Image, x, y int, line *Line, color color.Color) {
	cursor := 0
	for _, letter := range line.letters {
		options := &ebiten.DrawRectShaderOptions{}
		options.GeoM.Translate(float64(x+cursor), float64(y))
		r, g, b, a := colorToFloats(color)
		options.Uniforms = map[string]interface{}{
			"FontColor": []float32{r, g, b, a},
		}
		options.Images = [4]*ebiten.Image{letter.Image}
		dest.DrawRectShader(letter.Image.Bounds().Dx(), letter.Image.Bounds().Dy(), shaders.FontShader, options)
		cursor += letter.Image.Bounds().Dx() + line.kerning
	}
}

func DrawStringBlock(dest *ebiten.Image, x, y int, text *Text, color color.Color) {
	h := 0
	for _, line := range text.lines {
		DrawString(dest, x, y+h, line, color)
		h += line.height + text.leading
	}
}

func DrawText(dest *ebiten.Image, text *Text, color color.Color) {
	x, y := text.Pos2()
	DrawStringBlock(dest, int(x), int(y), text, color)
}

func colorToFloats(clr color.Color) (float32, float32, float32, float32) {
	r, g, b, a := clr.RGBA()
	return float32(r) / 65535, float32(g) / 65535, float32(b) / 65535, float32(a) / 65535
}

func ImagesForString(text string, kerning int, leading int, font *Font) *Text {
	result := NewText(kerning, leading)
	line := NewLine(kerning)
	for _, char := range text {
		if char == 10 {
			result.AddLine(line)
			line = NewLine(kerning)
			continue
		}
		if fontChar, ok := font.Data[char]; ok {
			line.AddLetter(Letter{Image: fontChar.Image, Char: char})
		}
	}
	if line.Len() > 0 {
		result.AddLine(line)
	}
	return result
}

// MeasureString Measures the width and height necessary to draw the given
// string with the provided images and kerning
// returns (w, h float64)
// func MeasureString(kerning, leading int, lines [][]Letter) (int, int) {
// 	maxWidth := 0
// 	height := -leading
// 	h := 0
// 	for _, line := range lines {
// 		width := -kerning
// 		for _, letter := range line {
// 			width += kerning + letter.Image.Bounds().Dx()
// 			if letter.Image.Bounds().Dy() > 0 {
// 				h = letter.Image.Bounds().Dy()
// 			}
// 		}
// 		if width > maxWidth {
// 			maxWidth = width
// 		}
// 		height += leading + h
// 	}
// 	return maxWidth, height
// }

func Wrap(width int, text *Text) *Text {
	result := NewText(text.kerning, text.leading)
	currLine := NewLine(text.kerning)
	// currWidth := -kerning
	// currHeight := 0
	// maxWidth := 0
	// maxheight := 0
	for _, line := range text.lines {
		for _, letter := range line.letters {
			// Add the width of this letter plus the kerning to the currWidth
			// currWidth += kerning + letter.Image.Bounds().Dx()
			// Add the letter to the currLine
			currLine.AddLetter(letter)
			// Check if adding this letter would escape the bounds
			if currLine.Width() > width {
				// If so, reverse the currLine back to the last space
				// and push currLine into the result
				for i := currLine.Len() - 1; i > 0; i-- {
					// If char is a space
					if currLine.At(i).Char == 32 {
						// Flush all letter preceeding this letter to the result
						// TODO: Possibly want to encapulate the sub-slice behaviour in a Line object function
						result.AddLine(NewLine(text.kerning, currLine.letters[0:i]...))
						// Set current line to the remainder of the items in currLine
						currLine = NewLine(text.kerning, currLine.letters[i+1:]...)
						break
					}
				}
			}
		}
		// At the end of each input line, flush the curr buffers if they contain anything
		if len(currLine.letters) > 0 {
			result.AddLine(currLine)
			currLine = NewLine(text.kerning)
		}
	}
	return result
}
