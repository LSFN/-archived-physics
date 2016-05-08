package movement

import (
	"github.com/LSFN/physics/vectors"
)

type Positionable interface {
	GetPosition() vectors.Vector2
}
