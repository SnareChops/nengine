package navigation

import "github.com/SnareChops/nengine/types"

// NavNode represents a node in the NavMesh
type NavNode struct {
	types.Position
	X, Y    int
	F, G, H float64
	parent  *NavNode
	index   int
	mask    uint64
}

func (node *NavNode) Is(mask uint64) bool {
	if mask == 0 {
		return node.mask == 0
	}
	return node.mask&mask == mask
}
