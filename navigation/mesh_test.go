package navigation_test

import (
	"fmt"
	"testing"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/navigation"
	"github.com/SnareChops/nengine/types"
	"github.com/stretchr/testify/assert"
)

func TestPathfind(t *testing.T) {
	// Test the A* algorithm
	collider := new(bounds.Raw).Init(90, 90)
	collider.SetPos2(110, 110)
	mesh := new(navigation.NavMesh).Init(320, 320, 64, 64, 32, 32)
	x, y := mesh.Grid()[0][0].Pos2()
	assert.Equal(t, 32., x)
	assert.Equal(t, 32., y)

	x, y = mesh.Grid()[0][1].Pos2()
	assert.Equal(t, 32., x)
	assert.Equal(t, 32.+64, y)

	x, y = mesh.Grid()[1][0].Pos2()
	assert.Equal(t, 32.+64, x)
	assert.Equal(t, 32., y)

	start := bounds.Point(32, 224)
	end := bounds.Point(288, 160)
	path := mesh.Pathfind(start, end, true, []types.Bounds{collider})
	println("Path:")
	for _, vec := range path {
		x, y := vec.Pos2()
		fmt.Printf("(%.2f, %.2f)\n", x, y)
	}
}

func TestLargePathfind(t *testing.T) {
	height := 2208 * 4
	width := 976 * 4
	hspacing := 16 * 4
	vspacing := 16 * 4
	hoffset := hspacing / 2
	voffset := vspacing / 2
	mesh := new(navigation.NavMesh).Init(height, width, hspacing, vspacing, hoffset, voffset)

	start := bounds.Point(32, 32)
	end := bounds.Point(8768, 3840)

	path := mesh.Pathfind(start, end, true, []types.Bounds{})
	for _, vec := range path {
		x, y := vec.Pos2()
		fmt.Printf("%.2f, %.2f\n", x, y)
	}

}
