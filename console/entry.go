package console

import (
	"image/color"
	"strings"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Entry struct {
	*bounds.Raw
	value     string
	index     int
	cooldown  int
	repeating bool
	text      *fonts.Text
	cursor    *ebiten.Image
	image     *ebiten.Image
}

func (self *Entry) Init(w, h int) *Entry {
	self.Raw = new(bounds.Raw).Init(w, h)
	self.text = fonts.NewText(self.value, fontFace, color.White)
	self.cursor = ebiten.NewImage(3, 20)
	self.cursor.Fill(color.White)
	self.image = ebiten.NewImage(self.Size())
	self.render()
	return self
}

func (self *Entry) Update(delta int) (bool, string) {
	prev := self.value
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		self.cooldown -= delta
		if self.cooldown <= 0 && len(self.value) > 0 {
			self.value = self.value[:len(self.value)-1]
			if self.repeating {
				self.cooldown = 100
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
		self.render()
		return true, strings.TrimSpace(value)
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
		self.value += letter
	}
	// Re-render if value has changed
	if prev != self.value {
		self.text.SetValue(self.value)
		self.render()
		return true, ""
	}
	return false, ""
}

func (self *Entry) render() {
	self.image.Clear()
	fonts.DrawText(self.image, self.text, nil)
	// rendering.DrawAt(self.image, self.cursor, 0, 0)
}

func (self *Entry) Image() *ebiten.Image {
	return self.image
}
