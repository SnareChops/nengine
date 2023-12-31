package navigation

import "github.com/SnareChops/nengine/types"

type Navigation interface {
	Pathfind(start, end types.Position, allowDiagonals bool) NavPath
	Grid() [][]*NavNode
	NextNavGroup() int
	ActiveNavGroup() int
	Update(delta int)
}
