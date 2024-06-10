package rendering_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/rendering"
	"github.com/stretchr/testify/assert"
)

func TestCameraSetPos(t *testing.T) {
	c := new(rendering.BasicCamera).Init(10, 10, 20, 20)
	x, y := c.Pos()
	assert.Equal(t, 5., x)
	assert.Equal(t, 5., y)
	assert.Equal(t, 0., c.MinX())
	assert.Equal(t, 0., c.MinY())
	assert.Equal(t, 10., c.MaxX())
	assert.Equal(t, 10., c.MaxY())

	c.SetPos(0, 0)
	x, y = c.Pos()
	assert.Equal(t, 5., x)
	assert.Equal(t, 5., y)
	assert.Equal(t, 0., c.MinX())
	assert.Equal(t, 0., c.MinY())
	assert.Equal(t, 10., c.MaxX())
	assert.Equal(t, 10., c.MaxY())

	c.SetPos(20, 20)
	x, y = c.Pos()
	assert.Equal(t, 15., x)
	assert.Equal(t, 15., y)
	assert.Equal(t, 10., c.MinX())
	assert.Equal(t, 10., c.MinY())
	assert.Equal(t, 20., c.MaxX())
	assert.Equal(t, 20., c.MaxY())

	c.SetPos(0, 5)
	x, y = c.Pos()
	assert.Equal(t, 5., x)
	assert.Equal(t, 5., y)
	assert.Equal(t, 0., c.MinX())
	assert.Equal(t, 0., c.MinY())
	assert.Equal(t, 10., c.MaxX())
	assert.Equal(t, 10., c.MaxY())

	c.SetPos(5, 0)
	x, y = c.Pos()
	assert.Equal(t, 5., x)
	assert.Equal(t, 5., y)
	assert.Equal(t, 0., c.MinX())
	assert.Equal(t, 0., c.MinY())
	assert.Equal(t, 10., c.MaxX())
	assert.Equal(t, 10., c.MaxY())

	c.SetPos(20, 5)
	x, y = c.Pos()
	assert.Equal(t, 15., x)
	assert.Equal(t, 5., y)
	assert.Equal(t, 10., c.MinX())
	assert.Equal(t, 0., c.MinY())
	assert.Equal(t, 20., c.MaxX())
	assert.Equal(t, 10., c.MaxY())

	c.SetPos(5, 20)
	x, y = c.Pos()
	assert.Equal(t, 5., x)
	assert.Equal(t, 15., y)
	assert.Equal(t, 0., c.MinX())
	assert.Equal(t, 10., c.MinY())
	assert.Equal(t, 10., c.MaxX())
	assert.Equal(t, 20., c.MaxY())

	c.SetPos(-100, -100)
	x, y = c.Pos()
	assert.Equal(t, 5., x)
	assert.Equal(t, 5., y)
	assert.Equal(t, 0., c.MinX())
	assert.Equal(t, 0., c.MinY())
	assert.Equal(t, 10., c.MaxX())
	assert.Equal(t, 10., c.MaxY())

	c.SetPos(100, 100)
	x, y = c.Pos()
	assert.Equal(t, 15., x)
	assert.Equal(t, 15., y)
	assert.Equal(t, 10., c.MinX())
	assert.Equal(t, 10., c.MinY())
	assert.Equal(t, 20., c.MaxX())
	assert.Equal(t, 20., c.MaxY())
}

func TestCameraView(t *testing.T) {
	c := new(rendering.BasicCamera).Init(200, 100, 3000, 2000)
	view := c.View()
	assert.Equal(t, 0, view.Min.X)
	assert.Equal(t, 0, view.Min.Y)
	assert.Equal(t, 200, view.Max.X)
	assert.Equal(t, 100, view.Max.Y)

	c = new(rendering.BasicCamera).Init(200, 100, 3000, 2000)
	c.SetPos(20, 10)
	view = c.View()
	assert.Equal(t, 0, view.Min.X)
	assert.Equal(t, 0, view.Min.Y)
	assert.Equal(t, 200, view.Max.X)
	assert.Equal(t, 100, view.Max.Y)

	c = new(rendering.BasicCamera).Init(200, 100, 3000, 2000)
	c.SetPos(2980, 1990)
	view = c.View()
	assert.Equal(t, 2800, view.Min.X)
	assert.Equal(t, 1900, view.Min.Y)
	assert.Equal(t, 3000, view.Max.X)
	assert.Equal(t, 2000, view.Max.Y)

	c = new(rendering.BasicCamera).Init(200, 100, 3000, 2000)
	c.SetPos(1500, 1000)
	view = c.View()
	assert.Equal(t, 1400, view.Min.X)
	assert.Equal(t, 950, view.Min.Y)
	assert.Equal(t, 1600, view.Max.X)
	assert.Equal(t, 1050, view.Max.Y)
}

func TestCameraWorldToScreenPos(t *testing.T) {
	c := new(rendering.BasicCamera).Init(10, 10, 20, 20)
	x, y := c.WorldToScreenPos(2, 2)
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)

	x, y = c.WorldToScreenPos(19, 19)
	assert.Equal(t, 19, x)
	assert.Equal(t, 19, y)

	c.SetPos(8, 8)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	x, y = c.WorldToScreenPos(10, 10)
	assert.Equal(t, 7, x)
	assert.Equal(t, 7, y)

	c = new(rendering.BasicCamera).Init(1920, 1080, 5000, 6000)
	x, y = c.WorldToScreenPos(400, 500)
	assert.Equal(t, 400, x)
	assert.Equal(t, 500, y)

	c = new(rendering.BasicCamera).Init(1920, 1080, 5000, 6000)
	c.SetPos(1000, 900)
	x, y = c.WorldToScreenPos(400, 500)
	assert.Equal(t, 360, x)
	assert.Equal(t, 140, y)
}

func TestScreenToWorldPos(t *testing.T) {
	c := new(rendering.BasicCamera).Init(10, 10, 20, 20)
	x, y := c.ScreenToWorldPos(2, 2)
	assert.Equal(t, 2., x)
	assert.Equal(t, 2., y)

	x, y = c.ScreenToWorldPos(19, 19)
	assert.Equal(t, 19., x)
	assert.Equal(t, 19., y)

	c.SetPos(8, 8)
	x, y = c.ScreenToWorldPos(-1, -1)
	assert.Equal(t, 2., x)
	assert.Equal(t, 2., y)

	x, y = c.ScreenToWorldPos(7, 7)
	assert.Equal(t, 10., x)
	assert.Equal(t, 10., y)
}

func TestNewCameraWorldToScreenPosWithZoom(t *testing.T) {
	c := new(rendering.BasicCamera).Init(10, 10, 20, 20)
	c.SetZoom(2)
	x, y := c.WorldToScreenPos(2, 2)
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	c.SetPos(8, 8)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, -7, x)
	assert.Equal(t, -7, y)

	c = new(rendering.BasicCamera).Init(10, 10, 20, 20)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)

	c.SetZoom(2)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	c.SetZoom(.5)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, y)
}

func TestCameraWorldToScreenPosWithZoomAndPos(t *testing.T) {
	c := new(rendering.BasicCamera).Init(10, 10, 20, 20)
	c.SetPos(5, 5)
	x, y := c.WorldToScreenPos(2, 2)
	assert.Equal(t, 2, x)
	assert.Equal(t, 2, y)

	c.SetZoom(2)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, -1, x)
	assert.Equal(t, -1, y)

	c.SetZoom(.5)
	x, y = c.WorldToScreenPos(2, 2)
	assert.Equal(t, 1, x)
	assert.Equal(t, 1, y)
}

func TestCameraViewConsistentRectSize(t *testing.T) {
	c := new(rendering.BasicCamera).Init(200, 100, 2000, 1000)
	result := c.View()
	assert.Equal(t, 0, result.Min.X)
	assert.Equal(t, 0, result.Min.Y)
	assert.Equal(t, 200, result.Max.X)
	assert.Equal(t, 100, result.Max.Y)

	c = new(rendering.BasicCamera).Init(200, 100, 2000, 1000)
	c.SetZoom(2)
	result = c.View()
	assert.Equal(t, 50, result.Min.X)
	assert.Equal(t, 25, result.Min.Y)
	assert.Equal(t, 150, result.Max.X)
	assert.Equal(t, 75, result.Max.Y)
}

func TestCameraScreenToWorldPosWithZoom(t *testing.T) {
	c := new(rendering.BasicCamera).Init(200, 100, 2000, 1000)
	c.SetZoom(.5)

	x, y := c.ScreenToWorldPos(20, 10)
	assert.Equal(t, -60., x)
	assert.Equal(t, -30., y)

	c = new(rendering.BasicCamera).Init(200, 100, 2000, 1000)
	c.SetZoom(2)

	x, y = c.ScreenToWorldPos(20, 10)
	assert.Equal(t, 60., x)
	assert.Equal(t, 30., y)
}

func TestCameraFollow(t *testing.T) {
	camera := new(rendering.BasicCamera).Init(1920, 1080, 5000, 5000)
	bounds := new(bounds.Raw).Init(32, 32)
	bounds.SetPos2(796.00, 786.00)
	camera.Follow(bounds)
	camera.SetPos(796, 786)

	x, y := camera.Pos()
	assert.Equal(t, 960., x)
	assert.Equal(t, 786., y)

	bounds.SetPos2(800, 900)
	camera.Update(10)
	x, y = camera.Pos()
	assert.Equal(t, 960., x)
	assert.Equal(t, 900., y)

	bounds.SetPos2(1000, 900)
	camera.Update(10)
	x, y = camera.Pos()
	assert.Equal(t, 1000., x)
	assert.Equal(t, 900., y)
}
