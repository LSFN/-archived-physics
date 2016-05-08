package shapes

import (
	"math"

	"github.com/LSFN/physics/movement"
)

type Circle struct {
	movement.MovingPoint
	theta, radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
