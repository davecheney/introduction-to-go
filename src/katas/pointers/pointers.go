package pointers

import "fmt"

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("point: x=%d, y=%d", p.X, p.Y)
}

func (p *Point) Move(x, y int) {
	p.X += x
	p.Y += y
}
