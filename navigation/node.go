package navigation

import "github.com/SnareChops/nengine/types"

// NavNode represents a node in the NavMesh
type NavNode struct {
	types.Position
	X, Y    int
	F, G, H float64
	parent  *NavNode
	index   int
}
