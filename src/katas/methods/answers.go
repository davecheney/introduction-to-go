// +build ignore

package methods

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("point: x=%d, y=%d", p.X, p.Y)
}
