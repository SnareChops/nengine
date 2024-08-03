package navigation

import (
	"container/heap"
	"math"

	"github.com/SnareChops/nengine/bounds"
	"github.com/SnareChops/nengine/debug"
	"github.com/SnareChops/nengine/types"
	"github.com/SnareChops/nengine/utils"
)

var pathfindTimer = debug.NewFrameTimer("Pathfind", true)

// NavMesh represents a navigation grid for pathfinding
type NavMesh struct {
	grid      [][]*NavNode
	navGroups int
	group     int
	active    int
	hspacing  int
	vspacing  int
	hoffset   int
	voffset   int
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

// MaskNode Sets the mask on the node at the provided index position
func (self *NavMesh) MaskNode(i, j, mask int) {
	self.grid[i][j].mask = mask
}

// MaskNodeAt Sets the mask on the node at the provided world position
func (self *NavMesh) MaskNodeAt(x, y float64, mask int) {
	i := int(x*float64(self.hspacing) + float64(self.hoffset))
	j := int(y*float64(self.vspacing) + float64(self.voffset))
	self.grid[i][j].mask = mask
}

// MaskNodesWithin Sets the mask on the nodes that collide with the provided Box
func (self *NavMesh) MaskNodesWithin(box types.Box, mask int) {
	for x := range self.grid {
		for y := range self.grid[x] {
			if self.grid[x][y] != nil && utils.IsWithin(box, x*self.hspacing+self.hoffset, y*self.vspacing+self.voffset) {
				self.grid[x][y].mask = mask
			}
		}
	}
}

func (self *NavMesh) Update(delta int) {
	self.active += 1
	if self.active >= self.navGroups {
		self.active = 0
	}
}

func (self *NavMesh) ClosestNode(pos types.Position, masks ...int) *NavNode {
	x := int(pos.X()-float64(self.hoffset)) / self.hspacing
	y := int(pos.Y()-float64(self.voffset)) / self.vspacing
	min := math.Inf(1)
	var closest *NavNode
	var iteration int = 1
	for {
		for i := range 2 * iteration {
			for j := range 2 * iteration {
				if !matchesMasks(self.grid[i+x][j+y], masks) {
					continue
				}
				gridPos := self.grid[i+x][j+y].Position
				dist := utils.DistanceBetweenPoints(pos.X()-float64(self.hspacing*(iteration-1)), pos.Y()-float64(self.vspacing*(iteration-1)), gridPos.X(), gridPos.Y())
				if dist < min {
					min = dist
					closest = self.grid[i+x][j+y]
				}
			}
		}
		if closest != nil {
			break
		}
		iteration += 1
	}
	return closest
}

// Pathfind uses the NavMesh to find a path from the start to end Vector
// Optionally allowing diagonal movement between the nodes
func (self *NavMesh) Pathfind(start, end types.Position, allowDiagonals bool, masks ...int) NavPath {
	pathfindTimer.Start()
	defer pathfindTimer.End()
	// Find closest nodes to start and end positions
	startNode := self.ClosestNode(start, masks...)
	endNode := self.ClosestNode(end, masks...)
	// Calculate path
	path := self.AStar(startNode, endNode, allowDiagonals, masks...)
	// Append ending vector to path
	if len(path) > 0 {
		path = append(path, end)
		return path[1:]
	}
	return []types.Position{}
}

// AStar runs the A* algoritm on the NavMesh and returns a path between the provided nodes
// Optionally allowing diagonal movement between nodes
// Note: This is exposed, but is really only intended to be used internally
// Prefer using the Pathfind() method instead
func (self *NavMesh) AStar(start, end *NavNode, allowDiagonal bool, masks ...int) NavPath {
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
		for _, neighbor := range getNeighbors(current, self.grid, allowDiagonal, masks) {
			if _, ok := closedSet[neighbor]; ok {
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
	self.hspacing = hspacing
	self.vspacing = vspacing
	self.hoffset = hoffset
	self.voffset = voffset
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
