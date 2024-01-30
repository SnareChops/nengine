package types

type Position interface {
	Pos() Vector
	Pos2() (x, y float64)
	SetPos2(x, y float64)
	Pos3() (x, y, z float64)
	SetPos3(x, y, z float64)
	X() float64
	Y() float64
	Z() float64
}
