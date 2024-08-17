package navigation_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/SnareChops/nengine"
	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/navigation"
	"github.com/stretchr/testify/assert"
)

// func TestPathfind(t *testing.T) {
// 	// Test the A* algorithm
// 	collider := new(bounds.Raw).Init(90, 90)
// 	collider.SetPos2(110, 110)
// 	mesh := new(navigation.NavMesh).Init(320, 320, 64, 64, 32, 32)
// 	x, y := mesh.Grid()[0][0].Pos2()
// 	assert.Equal(t, 32., x)
// 	assert.Equal(t, 32., y)

// 	x, y = mesh.Grid()[0][1].Pos2()
// 	assert.Equal(t, 32., x)
// 	assert.Equal(t, 32.+64, y)

// 	x, y = mesh.Grid()[1][0].Pos2()
// 	assert.Equal(t, 32.+64, x)
// 	assert.Equal(t, 32., y)

// 	start := bounds.Point(32, 224)
// 	end := bounds.Point(288, 160)
// 	path := mesh.Pathfind(start, end, true, []types.Collidable{collider})
// }

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

	path := mesh.Pathfind(start, end, true)
	for _, vec := range path {
		x, y := vec.Pos2()
		fmt.Printf("%.2f, %.2f\n", x, y)
	}
}

func TestNoPathFound(t *testing.T) {
	// Test the A* algorithm
	collider := new(bounds.Raw).Init(64, 320)
	collider.SetPos2(128, 0)
	mesh := new(navigation.NavMesh).Init(320, 320, 64, 64, 32, 32)
	mesh.MaskNodesWithin(collider, 1)

	start := bounds.Point(64, 64)
	end := bounds.Point(256, 256)
	path := mesh.Pathfind(start, end, true, 0)
	assert.Equal(t, 0, len(path))
}

func TestPathfindWithNodeMasks(t *testing.T) {
	MASK_1 := 1 << 1
	MASK_2 := 1 << 2
	MASK_3 := 1 << 3
	collider := bounds.NewBox(64, 320)
	collider.SetPos2(128, 0)
	mesh := new(navigation.NavMesh).Init(320, 320, 64, 64, 32, 32)
	mesh.MaskNodesWithin(collider, MASK_1|MASK_2|MASK_3)

	start := bounds.Point(64, 64)
	end := bounds.Point(256, 256)

	// No path
	path := mesh.Pathfind(start, end, true)
	assert.Equal(t, 0, len(path))

	// Path matching MASK_1
	path = mesh.Pathfind(start, end, true, MASK_1)
	assert.Equal(t, 4, len(path))

	// Path matching MASK_2|MASK_3
	path = mesh.Pathfind(start, end, true, MASK_2|MASK_3)
	assert.Equal(t, 4, len(path))
}

func TestPathfindBetweenMaskedAreas(t *testing.T) {
	MASK_1 := 1 << 1
	box1 := bounds.NewBox(64, 64)
	box2 := bounds.NewBox(64, 64)
	box1.SetPos2(128, 64)
	box2.SetPos2(256, 192)

	mesh := new(navigation.NavMesh).Init(1000, 1000, 64, 64, 32, 32)
	mesh.MaskNodesWithin(box1, MASK_1)
	mesh.MaskNodesWithin(box2, MASK_1)

	path := mesh.ExactPath(nengine.Point(box1.Mid()), nengine.Point(box2.Mid()), true, MASK_1)

	assert.Equal(t, 3, len(path))
	assert.Equal(t, box1.MidX(), path[0].X())
	assert.Equal(t, box1.MidY(), path[0].Y())
	assert.Equal(t, box2.MidX(), path[2].X())
	assert.Equal(t, box2.MidY(), path[2].Y())
}

func TestClosestNodes(t *testing.T) {
	mesh := new(navigation.NavMesh).Init(640, 640, 64, 64, 32, 32)
	pos := bounds.Point(220, 365)

	node := mesh.ClosestNode(pos, 0)
	assert.Equal(t, 3, node.X)
	assert.Equal(t, 5, node.Y)
}

func TestClosestNodesWithMask(t *testing.T) {
	mesh := new(navigation.NavMesh).Init(640, 640, 64, 64, 32, 32)
	pos := bounds.Point(220, 365)

	node := mesh.ClosestNode(pos, 1)
	assert.Equal(t, 3, node.X)
	assert.Equal(t, 5, node.Y)

	mesh.MaskNodesWithin(new(bounds.Raw).Init(640, 640), 1)
	node = mesh.ClosestNode(pos, 1)
	assert.Equal(t, 3, node.X)
	assert.Equal(t, 5, node.Y)
}

func TestClosestExpandingSearch(t *testing.T) {
	mesh := new(navigation.NavMesh).Init(640, 640, 64, 64, 32, 32)
	pos := bounds.Point(220, 365)

	b := new(bounds.Raw).Init(65, 65)
	b.SetPos2(float64((220/64)*64), float64((365/64)*64))
	mesh.MaskNodesWithin(b, 1)
	node := mesh.ClosestNode(pos, 0)
	log.Println(node.Position)
	assert.Equal(t, 3, node.X)
	assert.Equal(t, 6, node.Y)
}
