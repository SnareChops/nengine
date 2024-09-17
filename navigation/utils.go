package navigation

import (
	"math"

	"github.com/SnareChops/nengine/bit"
)

func matchesMasks(node *NavNode, masks []uint64) bool {
	// Nodes with no mask are ALWAYS walkable
	if node.mask == 0 {
		return true
	}
	// If no masks, only allow non-masked nodes
	if len(masks) == 0 && node.mask == 0 {
		return true
	}
	// If any mask matches, the node is walkable
	for _, mask := range masks {
		if bit.IsSet(node.mask, mask) {
			return true
		}
	}
	return false
}

func getNeighbors(node *NavNode, grid [][]*NavNode, allowDiagonal bool, masks []uint64) []*NavNode {
	var neighbors []*NavNode
	x, y := int(node.X), int(node.Y)
	var directions [][2]int
	if allowDiagonal {
		directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}} // Right, Down, Left, Up, diagonals
	} else {
		directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up
	}

	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) && matchesMasks(grid[nx][ny], masks) {
			neighbors = append(neighbors, grid[nx][ny])
		}
	}

	return neighbors
}

func heuristic(nodeA, nodeB *NavNode) float64 {
	dx := float64(nodeB.X - nodeA.X)
	dy := float64(nodeB.Y - nodeA.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
