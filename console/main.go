package console

import (
	"image/color"
	"slices"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/rendering"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	background = color.RGBA{16, 16, 16, 255}
	dark       = color.RGBA{69, 69, 69, 255}
	mid        = color.RGBA{128, 128, 128, 255}
	light      = color.RGBA{193, 193, 193, 255}
	error      = color.RGBA{239, 167, 161, 255}
	good       = color.RGBA{128, 201, 144, 255}
	warn       = color.RGBA{167, 149, 93, 255}
	info       = color.RGBA{80, 203, 205, 255}
	cmd        = color.RGBA{117, 196, 228, 255}
)

type console struct {
	*bounds.Raw
	visible     bool
	key         ebiten.Key
	history     []string
	output      []*fonts.Text
	entry       *Entry
	cursorImage *ebiten.Image
	entryText   *fonts.Text
	start       *fonts.Text
	image       *ebiten.Image
}

var state = &console{}

func Init(key ebiten.Key) {
	state.key = key
	state.Raw = new(bounds.Raw).Init(1920-400, 1080-200)
	state.start = fonts.NewText(">", fontFace, dark)
	state.entry = new(Entry).Init(state.Dx()-20, 20, light)
	state.image = ebiten.NewImage(state.Size())
	addResult(ConsoleResultInfo, "Nengine Console, `help` for more info")
	reposition()
	render()
}

func Clear() {
	state.output = []*fonts.Text{}
	reposition()
	render()
}

func addToHistory(command string) {
	if slices.Contains(state.history, command) {
		for i, c := range state.history {
			if c == command {
				state.history = append(state.history[:i], state.history[i+1:]...)
				break
			}
		}
	}
	state.history = append(state.history, command)
}

func addCom(command string) {
	text := fonts.NewText("> "+command, fontFace, cmd)
	text.Wrap(state.Dx() - 10)
	text.SetPos2(5, 0)
	addToOutput(text)
}

func addResult(status ConsoleResult, result string) {
	var clr color.Color
	switch status {
	case ConsoleResultInfo:
		clr = info
	case ConsoleResultWarn:
		clr = warn
	case ConsoleResultError:
		clr = error
	default:
		clr = mid
	}
	text := fonts.NewText(result, fontFace, clr)
	text.Wrap(state.Dx() - 15)
	text.SetPos2(10, 0)
	addToOutput(text)
}

func addToOutput(text *fonts.Text) {
	// Add the result text to `output`
	state.output = append(state.output, text)
	// Limit output buffer to only the most recent 50 messages
	if len(state.output) > 50 {
		state.output = state.output[1:]
	}
}

func Update(delta int) {
	prev := state.visible
	// If console key pressed: show/hide
	if inpututil.IsKeyJustPressed(state.key) {
		state.visible = !state.visible
	}
	// If console was changed this frame: always capture all input
	if prev != state.visible {
		input.InputCapture()
	}
	// If console is not visible: exit early
	if !state.visible {
		return
	}
	// Capture ALL input when the console is visible
	input.InputCapture()
	// Run entry update
	updated, command := state.entry.Update(delta)
	// If command was submitted:
	if command != "" {
		// If was a clear command: Clear the console
		if command == "clear" {
			Clear()
			return
		}
		// Add the command to history
		addToHistory(command)
		// Attempt to run the command and capture the return text
		status, result := RunCommand(command)
		addCom(command)
		addResult(status, result)
		// output buffer has changed so we need to reposition all of the text
		reposition()
	}
	if updated {
		// If entry has been updated, redraw the console
		render()
	}
}

func reposition() {
	state.start.SetPos2(5, float64(state.Dy())-30)
	state.entry.SetPos2(20, float64(state.Dy())-30)
	pointer := state.Dy() - 60
	for i := len(state.output) - 1; i >= 0; i-- {
		pointer -= state.output[i].Dy()
		state.output[i].SetPos2(state.output[i].X(), float64(pointer))
		pointer -= 5
	}
}

func render() {
	state.image.Clear()
	state.image.Fill(color.Black)
	fonts.DrawText(state.image, state.start, nil)
	rendering.DrawSprite(state.image, state.entry, nil)
	for _, out := range state.output {
		fonts.DrawText(state.image, out, nil)
	}
}

func Draw(screen *ebiten.Image) {
	if state.visible {
		rendering.DrawAt(screen, state.image, 200, 0)
	}
}
