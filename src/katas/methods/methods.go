package methods

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("point: x=%d, y=%d", p.X, p.Y)
}

func (p Point) GetX() int {
	return p.X
}

func (p Point) GetY() int {
	return p.Y
}
