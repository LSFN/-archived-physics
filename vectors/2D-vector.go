package vectors

import (
	"math"
)

type Vector2 struct {
	x, y float64
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}

func (v Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}
}

func (v Vector2) Multiply(k float64) Vector2 {
	return Vector2{
		x: v.x * k,
		y: v.y * k,
	}
}

func (v Vector2) DotProduct(v2 Vector2) float64 {
	return v.x*v2.x + v.y*v2.y
}
