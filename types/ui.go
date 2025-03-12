package types

type ButtonState int

const (
	ButtonStateHovered ButtonState = 1 << iota
	ButtonStateClicked
	ButtonStateJustClicked
	ButtonStateJustHovered
	ButtonStateDisabled
)

type Button interface {
	Bounds
	Is(ButtonState) bool
	ButtonState() ButtonState
	Enable()
	Disable()
	Update(x, y int) bool
}

type CheckboxState int

const (
	CheckboxStateChecked CheckboxState = 1 << iota
	CheckboxStateJustChanged
	CheckboxStateHovered
)

type Checkbox interface {
	Bounds
	Is(CheckboxState) bool
	SetChecked(bool)
	Update(x, y int) bool
}

type IntBox interface {
	Bounds
	Content() int
	SetContent(int)
	Update(x, y, delta int)
}

type PercentBar interface {
	Sprite
	SetValue(float64)
}

type TextBox interface {
	Bounds
	Content() string
	SetContent(string)
	IsFocused() bool
	Update(x, y, delta int)
}
