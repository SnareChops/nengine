package navigation

import (
	"container/heap"
	"math"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
)

// NavMesh represents a navigation grid for pathfinding
type NavMesh struct {
	grid      [][]*NavNode
	navGroups int
	group     int
	active    int
}

// Grid returns the navigation grid for the NavMesh
func (self *NavMesh) Grid() [][]*NavNode {
	return self.grid
}

func (self *NavMesh) NextNavGroup() int {
	self.group += 1
	if self.group >= self.navGroups {
		self.group = 0
	}
	return self.group
}

func (self *NavMesh) ActiveNavGroup() int {
	return self.active
}

func (self *NavMesh) Update(delta int) {
	self.active += 1
	if self.active >= self.navGroups {
		self.active = 0
	}
}

// Pathfind uses the NavMesh to find a path from the start to end Vector
// Optionally allowing diagonal movement between the nodes
func (self *NavMesh) Pathfind(start, end types.Position, allowDiagonals bool, obstacles []types.Bounds) NavPath {
	// Find closest nodes to start and end positions
	var startNode *NavNode
	var startDist float64 = math.MaxFloat64
	var endNode *NavNode
	var endDist float64 = math.MaxFloat64
	for i := range self.grid {
		for j := range self.grid[i] {
			if dist := utils.DistanceBetween(start, self.grid[i][j].Position); dist < startDist {
				startDist = dist
				startNode = self.grid[i][j]
			}
			if dist := utils.DistanceBetween(end, self.grid[i][j].Position); dist < endDist {
				endDist = dist
				endNode = self.grid[i][j]
			}
		}
	}
	// Calculate path
	path := self.AStar(startNode, endNode, allowDiagonals, obstacles)
	// Append ending vector to path
	path = append(path, end)
	return path
}

// AStar runs the A* algoritm on the NavMesh and returns a path between the provided nodes
// Optionally allowing diagonal movement between nodes
// Note: This is exposed, but is really only intended to be used internally
// Prefer using the Pathfind() method instead
func (self *NavMesh) AStar(start, end *NavNode, allowDiagonal bool, obstacles []types.Bounds) NavPath {
	defer self.reset()
	openSet := priorityQueue{}
	closedSet := make(map[*NavNode]bool)
	start.G = 0
	start.H = heuristic(start, end)
	start.F = start.G + start.H
	heap.Push(&openSet, start)

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*NavNode)

		if current == end {
			var path []types.Position
			for current != nil {
				path = append([]types.Position{current.Position}, path...)
				current = current.parent
			}
			return path
		}

		closedSet[current] = true
		for _, neighbor := range getNeighbors(current, self.grid, allowDiagonal) {
			if _, ok := closedSet[neighbor]; ok {
				continue
			}
			var gtg = true
			for _, ob := range obstacles {
				if ob.IsWithin(neighbor.Pos2()) {
					gtg = false
				}
			}
			if !gtg {
				continue
			}

			tentativeGScore := current.G + heuristic(current, neighbor)
			if tentativeGScore < neighbor.G {
				neighbor.parent = current
				neighbor.G = tentativeGScore
				neighbor.H = heuristic(neighbor, end)
				neighbor.F = neighbor.G + neighbor.H
				if neighbor.index >= 0 {
					heap.Fix(&openSet, neighbor.index)
				} else {
					heap.Push(&openSet, neighbor)
				}
			}
		}
	}
	return nil
}

func (self *NavMesh) reset() {
	for i := range self.grid {
		for j := range self.grid[i] {
			node := self.grid[i][j]
			node.parent = nil
			node.F = 0
			node.G = math.Inf(1)
			node.H = 0
			node.index = -1
		}
	}
}

func (self *NavMesh) Init(width, height, hspacing, vspacing, hoffset, voffset int) *NavMesh {
	var xCount int
	if hspacing == 0 && hoffset == 0 {
		xCount = width
	} else {
		xCount = int(math.Ceil(float64(width-hoffset) / float64(hspacing)))
	}
	var yCount int
	if vspacing == 0 && voffset == 0 {
		yCount = height
	} else {
		yCount = int(math.Ceil(float64(height-voffset) / float64(vspacing)))
	}
	grid := make([][]*NavNode, xCount)
	for i := range grid {
		grid[i] = make([]*NavNode, yCount)
		for j := range grid[i] {
			grid[i][j] = &NavNode{
				Position: bounds.Point(float64(i*hspacing+hoffset), float64(j*vspacing+voffset)),
				X:        i,
				Y:        j,
				G:        math.Inf(1),
				index:    -1,
			}
		}
	}
	self.navGroups = 10
	self.grid = grid
	return self
}

func (self *NavMesh) InitSimple(width, height int) *NavMesh {
	grid := make([][]*NavNode, width)
	for i := range grid {
		grid[i] = make([]*NavNode, height)
		for j := range grid[i] {
			grid[i][j] = &NavNode{
				Position: bounds.Point(float64(i), float64(j)),
				X:        i,
				Y:        j,
				G:        math.Inf(1),
				index:    -1,
			}
		}
	}
	self.navGroups = 10
	self.grid = grid
	return self
}
