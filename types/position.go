package types

type Position interface {
	Vec() Vector
	Pos() Position
	Pos2() (x, y float64)
	SetPos2(x, y float64)
	Pos3() (x, y, z float64)
	SetPos3(x, y, z float64)
	GridAlign(h, v int)
	X() float64
	Y() float64
	Z() float64
}
