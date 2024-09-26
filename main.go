package nengine

import (
	"github.com/SnareChops/nengine/animators"
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/console"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/emitters"
	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/loaders"
	"github.com/SnareChops/nengine/navigation"
	"github.com/SnareChops/nengine/rendering"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/ui"
	"github.com/SnareChops/nengine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// Animators
type FrameByFrameAnimator = animators.FrameByFrameAnimator
type Animation = animators.Animation
type AnimationFrame = animators.AnimationFrame
type SlideAnimator = animators.SlideAnimator
type SimpleAnimator = animators.SimpleAnimator
type SimpleFrame = animators.SimpleFrame
type GeneralAnimator = animators.GeneralAnimator
type GeneralFrame = animators.GeneralFrame

var NewGeneralFrame = animators.NewGeneralFrame

// Types
type Animator = types.Animator
type Game = types.Game
type Scene = types.Scene
type Position = types.Position
type Bounds = types.Bounds
type Entity = types.Entity
type Sprite = types.Sprite
type RenderLayer = types.RenderLayer
type SpriteRenderLayer = types.SpriteRenderLayer
type Camera = types.Camera
type Vector = types.Vector
type Box = types.Box

var NewVector = types.NewVector

// Traits
type Reloadable = types.Reloadable
type Drawable = types.Drawable

// Bounds
type RawBounds = bounds.Raw
type RelativeBounds = bounds.Relative
type PhysicsBounds = bounds.Physics
type VelocityBounds = bounds.VelocityBounds
type ProjectileBounds = bounds.Projectile

const (
	TOP    = bounds.TOP
	CENTER = bounds.CENTER
	BOTTOM = bounds.BOTTOM
	LEFT   = bounds.LEFT
	RIGHT  = bounds.RIGHT
)

var Point = bounds.Point
var NewBox = bounds.NewBox

type ConsoleFunc = console.ConsoleFunc
type ConsoleResult = console.ConsoleResult

var ConsoleResultNormal = console.ConsoleResultNormal
var ConsoleResultWarn = console.ConsoleResultWarn
var ConsoleResultError = console.ConsoleResultError
var ConsoleRegister = console.ConsoleRegister

// Debug
type DebugTimer = debug.DebugTimer
type FrameTimer = debug.FrameTimer

var NewDebugTimer = debug.NewDebugTimer
var NewFrameTimer = debug.NewFrameTimer
var DebugStat = debug.DebugStat
var DebugPath = debug.DebugPath

// Emitters
type ParticleEmitter = emitters.Emitter
type ExplosiveEmitter = emitters.Explosive
type UniformEmitter = emitters.Uniform
type ProjectileEmitter = emitters.Projectile
type ParticleBase = emitters.ParticleBase
type Particle = emitters.Particle

// Fonts
type Text = fonts.Text

var NewText = fonts.NewText
var LoadFont = fonts.LoadFont
var LoadTTF = fonts.LoadTTF
var LoadOTF = fonts.LoadOTF
var Font = fonts.Font
var GetStringWidth = fonts.GetStringWidth
var GetStringHeight = fonts.GetStringHeight
var GetStringSize = fonts.GetStringSize
var DrawText = fonts.DrawText

// Input
var InputCapture = input.InputCapture
var InputUncapture = input.InputUncapture
var IsInputCaptured = input.IsInputCaptured
var CursorContent = input.CursorContent
var SetCursorContent = input.SetCursorContent
var HideCursor = input.HideCursor
var ShowCursor = input.ShowCursor
var CursorDelta = input.CursorDelta
var IsAnyMouseButtonPressed = input.IsAnyMouseButtonPressed
var KeyToUpper = input.KeyToUpper

// Loaders
type Sheet = loaders.Sheet
type Anim = loaders.Anim

var PreloadImage = loaders.PreloadImage
var PreloadSheet = loaders.PreloadSheet
var PreloadAnim = loaders.PreloadAnim
var PreloadImageAseprite = loaders.PreloadImageAseprite
var PreloadSheetAseprite = loaders.PreloadSheetAseprite
var PreloadAnimAseprite = loaders.PreloadAnimAseprite
var PreloadImagePng = loaders.PreloadImagePng
var PreloadSheetPng = loaders.PreloadSheetPng
var PreloadImageJpeg = loaders.PreloadImageJpeg
var PreloadSheetJpeg = loaders.PreloadSheetJpeg
var GetImage = loaders.GetImage
var GetSheet = loaders.GetSheet
var GetSheetCell = loaders.GetSheetCell
var GetSheetRange = loaders.GetSheetRange
var GetAnim = loaders.GetAnim

// Navigation
type NavMesh = navigation.NavMesh
type NavNode = navigation.NavNode
type NavPath = navigation.NavPath
type Navigation = navigation.Navigation

// Rendering
type Background = rendering.Background
type ParallaxBackground = rendering.ParallaxBackground
type BasicCamera = rendering.BasicCamera
type BufferedCamera = rendering.BufferedCamera
type Renderer = rendering.Renderer
type Screen = rendering.Screen
type World = rendering.World

var DrawSprite = rendering.DrawSprite
var StrokeRect = rendering.StrokeRect
var GridDraw = rendering.GridDraw

func DrawAt[T ~int | float32 | float64](dest, src *ebiten.Image, x, y T) {
	rendering.DrawAt(dest, src, x, y)
}

// UI
type ButtonState = ui.ButtonState
type Button = ui.Button
type CheckboxState = ui.CheckboxState
type Checkbox = ui.Checkbox
type IntBox = ui.IntBox
type TextBox = ui.TextBox
type PercentBar = ui.PercentBar

const (
	ButtonStateClicked       = ui.ButtonStateClicked
	ButtonStateHovered       = ui.ButtonStateHovered
	ButtonStateJustClicked   = ui.ButtonStateJustClicked
	ButtonStateJustHovered   = ui.ButtonStateJustHovered
	ButtonStateDisabled      = ui.ButtonStateDisabled
	CheckboxStateChecked     = ui.CheckboxStateChecked
	CheckboxStateJustChanged = ui.CheckboxStateJustChanged
	CheckboxStateHovered     = ui.CheckboxStateHovered
)

// Utils
type ImageChunk = utils.ImageChunk

var DirectionVector = utils.DirectionVector
var DistanceBetween = utils.DistanceBetween
var DistanceBetweenPoints = utils.DistanceBetweenPoints
var DistanceToBounds = utils.DistanceToBounds
var PointAtAngleWithDistance = utils.PointAtAngleWithDistance
var AngleBetween = utils.AngleBetween
var AngleBetweenPoints = utils.AngleBetweenPoints
var MoveTowards = utils.MoveTowards
var MoveAway = utils.MoveAway
var Lerp = utils.Lerp
var ChunkImage = utils.ChunkImage
var ChunkBounds = utils.ChunkBounds
var DoesCollide = utils.DoesCollide
var IsPosWithin = utils.IsPosWithin

func IsWithin[T ~int | ~float64](box Box, x, y T) bool {
	return utils.IsWithin(box, x, y)
}

func ScaleFactor[T ~int | ~uint | ~float32 | ~float64](fromWidth, fromHeight, toWidth, toHeight T) (float64, float64) {
	return utils.ScaleFactor(fromWidth, fromHeight, toWidth, toHeight)
}

func Clamp[T ~int | ~uint | ~float32 | ~float64](num, min, max T) T {
	return utils.Clamp(num, min, max)
}

func LinearInterpolate[T ~float32 | ~float64](min, max, percent T) T {
	return utils.LinearInterpolate(min, max, percent)
}
