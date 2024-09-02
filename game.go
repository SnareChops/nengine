package nengine

import (
	"fmt"
	"log"
	"time"

	"github.com/SnareChops/nengine/console"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/input"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type BasicGame struct {
	Scene        types.Scene
	width        int
	height       int
	prev         int64
	terminate    bool
	reload       ebiten.Key
	loadComplete chan types.Scene
	update       *DebugTimer
	draw         *DebugTimer
	reloadTime   *DebugTimer
}

func NewGame(width, height int, consoleKey ebiten.Key, reload ebiten.Key) *BasicGame {
	game := &BasicGame{
		width:  width,
		height: height,
		reload: reload,
	}
	console.Init(consoleKey)
	game.update = NewDebugTimer("Update")
	game.draw = NewDebugTimer("Draw")
	if reload != 0 {
		fmt.Printf("Setting reload to %d\n", reload)
		game.reloadTime = NewDebugTimer("Reload")
	}
	return game
}

func (self *BasicGame) LoadScene(scene types.Scene) {
	if destroyable, ok := self.Scene.(types.Destroyable); ok {
		destroyable.Destroy()
	}
	self.Scene = nil
	if loadable, ok := scene.(types.Loadable); ok {
		log.Println("Scene is loadable, loading...")
		self.loadComplete = make(chan types.Scene)
		self.Scene = loadable.Load(self.loadComplete, self)
	} else {
		log.Println("Scene not loadable")
		self.Scene = scene
	}
	if initable, ok := self.Scene.(types.Initable); ok {
		log.Println("Scene in initable, initializing...")
		initable.Init(self)
	}
	input.Reset()
}

func (self *BasicGame) Update() error {

	// Update the input state
	input.Update()
	// If reload triggered: Send reload signal to scene
	if self.reload != 0 && IsKeyJustPressed(self.reload) {
		if scene, ok := self.Scene.(types.Reloadable); ok {
			input.InputCapture()
			log.Println("Scene Reloading...")
			if self.reloadTime != nil {
				self.reloadTime.Start()
			}
			scene.Reload()
			if self.reloadTime != nil {
				self.reloadTime.End()
			}
		}
	}
	// Calculate time between frames
	now := time.Now().UnixMilli()
	delta := int(now - self.prev)

	self.update.Start()
	// Update console state
	console.Update(delta)
	// If not a delta of 0: update scene
	if self.prev != 0 {
		self.Scene.Update(delta)
	}
	self.prev = now
	self.update.End()
	// If termination requested: Send terminate signal to ebiten
	if self.terminate {
		return ebiten.Termination
	}
	// If async loading: check if async load has completed
	if self.loadComplete != nil {
		select {
		case scene := <-self.loadComplete:
			self.Scene = scene
			self.loadComplete = nil
		default:
		}
	}
	// Update debug state
	debug.Update()
	return nil
}

func (self *BasicGame) Draw(screen *ebiten.Image) {
	self.draw.Start()
	if self.Scene != nil {
		self.Scene.Draw(screen)
	}
	// Draw input layer
	input.Draw(screen)
	// Draw console layer
	console.Draw(screen)
	// Draw debug info
	self.draw.End()
	debug.Draw(screen)
}

func (self *BasicGame) Layout(w, h int) (int, int) {
	return self.width, self.height
}

func (self *BasicGame) Terminate() {
	self.terminate = true
}
