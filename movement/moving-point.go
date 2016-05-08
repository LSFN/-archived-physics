package movement

import (
	"github.com/LSFN/physics/vectors"
)

type MovingPoint struct {
	position     vectors.Vector2
	velocity     vectors.Vector2
	acceleration vectors.Vector2
}

/**
Calculate the position and velocity of this moving point after t time units
*/
func (m *MovingPoint) Move(t float64) {
	m.position = m.position.Add(m.velocity.Multiply(t))
	m.velocity = m.velocity.Add(m.acceleration.Multiply(t))
}
