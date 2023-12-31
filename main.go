package nengine

import (
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/emitters"
	"github.com/SnareChops/nengine/navigation"
	"github.com/SnareChops/nengine/rendering"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
)

// Interfaces
type IGame = types.Game
type Scene = types.Scene
type Position = types.Position
type Bounds = types.Bounds
type Entity = types.Entity
type Sprite = types.Sprite
type RenderLayer = types.RenderLayer
type SpriteRenderLayer = types.SpriteRenderLayer

// Traits
type Reloadable = types.Reloadable
type Drawable = types.Drawable

// Bounds
type RawBounds = bounds.Raw
type RelativeBounds = bounds.Relative
type PhysicsBounds = bounds.Physics

const (
	TOP    = bounds.TOP
	CENTER = bounds.CENTER
	BOTTOM = bounds.BOTTOM
	LEFT   = bounds.LEFT
	RIGHT  = bounds.RIGHT
)

func Point(x, y, z float64) Position {
	return bounds.Point(x, y, z)
}

// Debug
type DebugTimer = debug.DebugTimer
type FrameTimer = debug.FrameTimer

var NewDebugTimer = debug.NewDebugTimer
var NewFrameTimer = debug.NewFrameTimer
var EnableDebug = debug.EnableDebug
var DebugEnabled = debug.DebugEnabled
var DebugStat = debug.DebugStat
var DebugPath = debug.DebugPath
var DebugDraw = debug.DebugDraw

// Emitters
type ParticleEmitter = emitters.Emitter
type ExplosiveEmitter = emitters.Explosive
type UniformEmitter = emitters.Uniform
type ProjectileEmitter = emitters.Projectile
type ParticleBase = emitters.ParticleBase
type Particle = emitters.Particle

// Navigation
type NavMesh = navigation.NavMesh
type NavNode = navigation.NavNode
type NavPath = navigation.NavPath
type Navigation = navigation.Navigation

// Rendering
type Background = rendering.Background
type ParallaxBackground = rendering.ParallaxBackground
type Camera = rendering.Camera
type Renderer = rendering.Renderer
type Screen = rendering.Screen
type World = rendering.World

// Utils
type ImageChunk = utils.ImageChunk

var DistanceBetween = utils.DistanceBetween
var DistanceBetweenPoints = utils.DistanceBetweenPoints
var PointAtAngleWithDistance = utils.PointAtAngleWithDistance
var AngleBetweenPoints = utils.AngleBetweenPoints
var MoveTowards = utils.MoveTowards
var Lerp = utils.Lerp
var ChunkImage = utils.ChunkImage
var ChunkBounds = utils.ChunkBounds

func Clamp[T ~int | ~uint | ~float32 | ~float64](num, min, max T) T {
	return utils.Clamp(num, min, max)
}

func LinearInterpolate[T ~float32 | ~float64](min, max, percent T) T {
	return utils.LinearInterpolate(min, max, percent)
}
