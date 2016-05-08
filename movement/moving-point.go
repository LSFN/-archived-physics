package movement

import (
	"github.com/LSFN/physics/vectors"
)

type MovingPoint struct {
	Position     vectors.Vector2
	Velocity     vectors.Vector2
	Acceleration vectors.Vector2
}

/**
Calculate the position and velocity of this moving point after t time units
*/
func (m *MovingPoint) Move(t float64) {
	m.Position = m.Position.Add(m.Velocity.Multiply(t))
	m.Velocity = m.Velocity.Add(m.Acceleration.Multiply(t))
}

func (m *MovingPoint) GetPosition() vectors.Vector2 {
	return m.Position
}
