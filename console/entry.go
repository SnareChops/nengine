package console

import (
	"image/color"
	"strings"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Entry struct {
	*bounds.Raw
	value       string
	index       int
	cooldown    int
	repeating   bool
	text        *fonts.Text
	cursor      types.Image
	cursorIndex int
	image       types.Image
}

func (self *Entry) Init(w, h int, color color.Color) *Entry {
	self.Raw = new(bounds.Raw).Init(w, h)
	self.text = fonts.NewText(self.value, fontFace, color)
	self.cursor = image.NewImage(3, 20)
	self.cursor.Fill(color)
	self.image = image.NewImage(self.Size())
	self.render()
	return self
}

func (self *Entry) SetValue(value string) {
	self.value = value
	self.text.SetValue(value)
	self.index = len(value) - 1
	self.cursorIndex = len(self.value)

	self.render()
}

func (self *Entry) Update(delta int) (bool, string) {
	prev := self.value
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		self.cooldown -= delta
		if self.cooldown <= 0 && len(self.value) > 0 {
			self.value = self.value[:self.cursorIndex-1] + self.value[self.cursorIndex:]
			self.cursorIndex -= 1
			if self.repeating {
				self.cooldown = 75
			} else {
				self.cooldown = 500
				self.repeating = true
			}
		}
		if prev != self.value {
			self.text.SetValue(self.value)
			self.render()
			return true, ""
		}
		return false, ""
	} else {
		self.cooldown = 0
		self.repeating = false
	}
	// If enter key is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		value := self.value
		self.value = ""
		self.text.SetValue(self.value)
		self.cursorIndex = 0
		self.render()
		return true, strings.TrimSpace(value)
	}
	//If left arrow is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if self.cursorIndex > 0 {
			self.cursorIndex -= 1
		}
	}
	//If right arrow is pressed
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if self.cursorIndex < len(self.value) {
			self.cursorIndex += 1
		}
	}
	// Handle keypresses
	keys := inpututil.AppendJustPressedKeys([]ebiten.Key{})
	for _, key := range keys {
		if key == ebiten.KeySpace {
			self.value += " "
			continue
		}
		letter := ebiten.KeyName(key)
		if ebiten.IsKeyPressed(ebiten.KeyShift) {
			letter = input.KeyToUpper(key, letter)
			letter = strings.ToUpper(letter)
		}
		self.value = self.value[:self.cursorIndex] + letter + self.value[self.cursorIndex:]
	}
	// Re-render if value has changed
	if prev != self.value {
		self.text.SetValue(self.value)
		self.cursorIndex += 1

		self.render()
		return true, ""
	}
	return false, ""
}

func (self *Entry) getCursorPosition() int {
	index := self.cursorIndex

	substr := self.value[index:]
	substrText := fonts.GetStringWidth(substr, fontFace)

	return self.text.Dx() - substrText
}

func (self *Entry) render() {
	self.image.Clear()
	fonts.DrawText(self.image, self.text, nil)
}

func (self *Entry) Image() types.Image {
	return self.image
}
