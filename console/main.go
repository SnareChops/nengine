package console

import (
	"image/color"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/rendering"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Console struct {
	*bounds.Raw
	input       *input.Input
	order       int
	visible     bool
	key         ebiten.Key
	output      []*fonts.Text
	entry       *Entry
	cursorImage *ebiten.Image
	entryText   *fonts.Text
	start       *fonts.Text
	image       *ebiten.Image
}

func (self *Console) Init(order int, key ebiten.Key, input *input.Input) *Console {
	self.order = order
	self.key = key
	self.input = input
	self.Raw = new(bounds.Raw).Init(1920-400, 1080-200)
	self.start = fonts.NewText(">", fontFace, color.White)
	self.entry = new(Entry).Init(self.Dx()-20, 20)
	self.image = ebiten.NewImage(self.Size())
	self.reposition()
	self.render()
	return self
}

func (self *Console) Order() int {
	return self.order
}

func (self *Console) Clear() {
	self.output = []*fonts.Text{}
	self.reposition()
	self.render()
}

func (self *Console) Update(delta int) {
	prev := self.visible
	if inpututil.IsKeyJustPressed(self.key) {
		self.visible = !self.visible
	}
	if prev != self.visible {
		self.input.InputCapture()
	}
	if !self.visible {
		return
	}
	// Capture ALL input when the console is visible
	self.input.InputCapture()
	updated, command := self.entry.Update(delta)
	if command != "" {
		if command == "clear" {
			self.Clear()
			return
		}
		result := RunCommand(command)
		text := fonts.NewText(result, fontFace, color.White)
		text.Wrap(self.Dy() - 10)
		text.SetPos2(5, 0)
		self.output = append(self.output, text)
		if len(self.output) > 50 {
			self.output = self.output[1:]
		}
		self.reposition()
	}
	if updated {
		self.render()
	}
}

func (self *Console) reposition() {
	self.start.SetPos2(5, float64(self.Dy())-30)
	self.entry.SetPos2(20, float64(self.Dy())-30)
	pointer := self.Dy() - 60
	for i := len(self.output) - 1; i >= 0; i-- {
		pointer -= self.output[i].Dy()
		self.output[i].SetPos2(self.output[i].X(), float64(pointer))
		pointer -= 5
	}
}

func (self *Console) render() {
	self.image.Clear()
	self.image.Fill(color.Black)
	fonts.DrawText(self.image, self.start, nil)
	rendering.DrawSprite(self.image, self.entry, nil)
	for _, out := range self.output {
		fonts.DrawText(self.image, out, nil)
	}
}

func (self *Console) Draw(screen *ebiten.Image) {
	if self.visible {
		rendering.DrawAt(screen, self.image, 200, 0)
	}
}
