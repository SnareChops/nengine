package nengine

import (
	"fmt"
	"time"

	"github.com/SnareChops/nengine/fonts"
	"github.com/SnareChops/nengine/types"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Scene     types.Scene
	width     int
	height    int
	prev      int64
	terminate bool
	seed      int64
	reload    ebiten.Key

	update     *DebugTimer
	draw       *DebugTimer
	reloadTime *DebugTimer
}

func NewGame(width, height int, seed int64, debug bool, reload ebiten.Key) *Game {
	game := &Game{
		width:  width,
		height: height,
		reload: reload,
	}
	if debug {
		EnableDebug(fonts.Arial12)
		DebugStat("TPS", func() string {
			return fmt.Sprintf("%0.2f", ebiten.ActualTPS())
		})
		DebugStat("FPS", func() string {
			return fmt.Sprintf("%0.2f", ebiten.ActualFPS())
		})
		DebugStat("Seed", func() string {
			return fmt.Sprintf("%d", seed)
		})
		game.update = NewDebugTimer("Update")
		game.draw = NewDebugTimer("Draw")
		if game.seed == 0 {
			game.seed = time.Now().UnixMicro()
		}
		if reload != 0 {
			fmt.Printf("Setting reload to %d\n", reload)
			game.reloadTime = NewDebugTimer("Reload")
		}
	}
	return game
}

func (self *Game) Seed() int64 {
	return self.seed
}

func (self *Game) LoadScene(scene types.Scene) {
	self.Scene = scene
	if initable, ok := scene.(types.Initable); ok {
		initable.Init(self)
	}
}

func (self *Game) Update() error {
	if self.reload != 0 && IsKeyJustPressed(self.reload) {
		if scene, ok := self.Scene.(types.Reloadable); ok {
			if self.reloadTime != nil {
				self.reloadTime.Start()
			}
			scene.Reload()
			if self.reloadTime != nil {
				self.reloadTime.End()
			}
		}
	}

	if self.update != nil {
		self.update.Start()
	}
	now := time.Now().UnixMilli()
	self.Scene.Update(int(now - self.prev))
	self.prev = now
	if self.update != nil {
		self.update.End()
	}
	if self.terminate {
		return ebiten.Termination
	}
	return nil
}

func (self *Game) Draw(screen *ebiten.Image) {
	if self.draw != nil {
		self.draw.Start()
	}
	self.Scene.Draw(screen)
	if self.draw != nil {
		self.draw.End()
	}
}

func (self *Game) Layout(w, h int) (int, int) {
	return self.width, self.height
}

func (self *Game) Terminate() {
	self.terminate = true
}
