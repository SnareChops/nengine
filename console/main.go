package console

import (
	"image/color"
	"slices"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/image"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/rendering"
	"github.com/SnareChops/nengine/types"
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
	cursorImage types.Image
	entryText   *fonts.Text
	start       *fonts.Text
	image       types.Image
	hint        types.Image

	cont ConsoleContinueFunc
}

var state = &console{}

func Init(key ebiten.Key) {
	state.key = key
	state.Raw = new(bounds.Raw).Init(1920-400, 1080-200)
	state.start = fonts.NewText(">", fontFace, dark)
	state.entry = new(Entry).Init(state.Dx()-20, 20, light)
	state.image = image.NewImage(state.Size())
	addResult(NewConsoleResult(ResultInfo, "Nengine Console, `help` for more info", nil))
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

func addResult(result ConsoleResult) {
	var clr color.Color
	switch result.code {
	case ResultInfo:
		clr = info
	case ResultWarn:
		clr = warn
	case ResultError:
		clr = error
	default:
		clr = mid
	}
	text := fonts.NewText(result.message, fontFace, clr)
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

func setHint(message string) {
	text := fonts.NewText(message, fontFace, info)
	text.Wrap(state.Dx() - 20)
	text.SetPos2(float64(state.Dx())/2-float64(text.Dx())/2, 10)
	state.hint = image.NewImage(state.Dx(), text.Dy()+20)
	state.hint.Fill(background)
	fonts.DrawText(state.hint, text, nil)
}

func setContinue(fn ConsoleContinueFunc) {
	state.cont = fn
}

func Update(delta int) {
	// If console has a hint: Clear if ESC or console key has been pressed
	if state.hint != nil && (inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(state.key)) {
		state.hint = nil
	}
	// If console has a continue function:
	if state.cont != nil {
		// If escape key or console key is pressed: Clear continue and hint
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) || inpututil.IsKeyJustPressed(state.key) {
			state.cont = nil
		}
		// If left mouse pressed:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			// Capture the input
			input.InputCapture()
			// Clear hint
			state.hint = nil
			// Run continue function
			result := state.cont(ebiten.CursorPosition())
			// Clear continue function pointer
			state.cont = nil
			// Handle the result of the continue function
			handleResult(result)
			state.visible = false
		}
	}

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
		// Add command to
		addCom(command)
		// Attempt to run the command and capture the return text
		RunCommand(command)
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
	state.image.Fill(background)
	fonts.DrawText(state.image, state.start, nil)
	rendering.DrawSprite(state.image, state.entry, nil)
	for _, out := range state.output {
		fonts.DrawText(state.image, out, nil)
	}
}

func Draw(screen types.Image) {
	if state.hint != nil {
		rendering.DrawAt(screen, state.hint, 200, 0)
		return
	}
	if state.visible {
		rendering.DrawAt(screen, state.image, 200, 0)
	}
}
