package rendering_test

import (
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/rendering"
	"github.com/stretchr/testify/assert"
)

func TestCameraSetPos(t *testing.T) {
	c := new(rendering.Camera).Init(200, 100, 3000, 2000)
	view := c.CameraView()
	assert.Equal(t, 0, view.Min.X)
	assert.Equal(t, 0, view.Min.Y)
	assert.Equal(t, 200, view.Max.X)
	assert.Equal(t, 100, view.Max.Y)

	c = new(rendering.Camera).Init(200, 100, 3000, 2000)
	c.SetCameraPos(20, 10)
	view = c.CameraView()
	assert.Equal(t, 0, view.Min.X)
	assert.Equal(t, 0, view.Min.Y)
	assert.Equal(t, 200, view.Max.X)
	assert.Equal(t, 100, view.Max.Y)

	c = new(rendering.Camera).Init(200, 100, 3000, 2000)
	c.SetCameraPos(2980, 1990)
	view = c.CameraView()
	assert.Equal(t, 2800, view.Min.X)
	assert.Equal(t, 1900, view.Min.Y)
	assert.Equal(t, 3000, view.Max.X)
	assert.Equal(t, 2000, view.Max.Y)

	c = new(rendering.Camera).Init(200, 100, 3000, 2000)
	c.SetCameraPos(1500, 1000)
	view = c.CameraView()
	assert.Equal(t, 1400, view.Min.X)
	assert.Equal(t, 950, view.Min.Y)
	assert.Equal(t, 1600, view.Max.X)
	assert.Equal(t, 1050, view.Max.Y)
}

func TestCameraWorldToScreenPos(t *testing.T) {
	c := new(rendering.Camera).Init(1920, 1080, 5000, 6000)
	x, y := c.WorldToScreenPos(400, 500)
	assert.Equal(t, 400, x)
	assert.Equal(t, 500, y)

	c = new(rendering.Camera).Init(1920, 1080, 5000, 6000)
	c.SetCameraPos(1000, 900)
	x, y = c.WorldToScreenPos(400, 500)
	assert.Equal(t, 360, x)
	assert.Equal(t, 140, y)
}

func TestCameraViewConsistentRectSize(t *testing.T) {
	c := new(rendering.Camera).Init(200, 100, 2000, 1000)
	result := c.CameraView()
	assert.Equal(t, 0, result.Min.X)
	assert.Equal(t, 0, result.Min.Y)
	assert.Equal(t, 200, result.Max.X)
	assert.Equal(t, 100, result.Max.Y)

	c = new(rendering.Camera).Init(200, 100, 2000, 1000)
	c.SetCameraZoom(2)
	result = c.CameraView()
	assert.Equal(t, 0, result.Min.X)
	assert.Equal(t, 0, result.Min.Y)
	assert.Equal(t, 200, result.Max.X)
	assert.Equal(t, 100, result.Max.Y)
}

func TestCameraScreenToWorldPosWithZoom(t *testing.T) {
	c := new(rendering.Camera).Init(200, 100, 2000, 1000)
	c.SetCameraZoom(.5)

	x, y := c.ScreenToWorldPos(20, 10)
	assert.Equal(t, 40., x)
	assert.Equal(t, 20., y)

	c = new(rendering.Camera).Init(200, 100, 2000, 1000)
	c.SetCameraZoom(2)

	x, y = c.ScreenToWorldPos(20, 10)
	assert.Equal(t, 10., x)
	assert.Equal(t, 5., y)
}

type TestEntity struct {
	*bounds.Raw
}

func (self *TestEntity) Update(delta int) {

}

func TestCameraFollow(t *testing.T) {
	camera := new(rendering.Camera).Init(1920, 1080, 5000, 5000)
	entity := &TestEntity{
		Raw: new(bounds.Raw).Init(32, 32),
	}
	entity.SetPos2(796.00, 786.00)
	camera.CameraFollow(entity)
	camera.SetCameraPos(796, 786)
	x, y := camera.CameraPos()
	assert.Equal(t, 796, x)
	assert.Equal(t, 786, y)
	camera.Update(10)
	x, y = camera.CameraPos()
	assert.Equal(t, 796, x)
	assert.Equal(t, 786, y)
	camera.Update(10)
	x, y = camera.CameraPos()
	assert.Equal(t, 796, x)
	assert.Equal(t, 786, y)
}
