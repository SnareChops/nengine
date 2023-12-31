package navigation

import "math"

func getNeighbors(node *NavNode, grid [][]*NavNode, allowDiagonal bool) []*NavNode {
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
		if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) {
			neighbors = append(neighbors, grid[nx][ny])
		}
	}

	return neighbors
}

func heuristic(nodeA, nodeB *NavNode) float64 {
	return math.Sqrt(math.Pow(float64(nodeB.X-nodeA.X), 2) + math.Pow(float64(nodeB.Y-nodeA.Y), 2))
}
