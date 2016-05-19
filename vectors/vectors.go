package vectors

type Vector []float64

func NewVector(v ...float64) Vector {
	return v
}

func (v Vector) Add(v2 Vector) Vector {
	var vr Vector = make([]float64, len(v))
	for i := range vr {
		vr[i] = v[i] + v2[i]
	}
	return vr
}

func (v Vector) Subtract(v2 Vector) Vector {
	var vr Vector = make([]float64, len(v))
	for i := range vr {
		vr[i] = v[i] - v2[i]
	}
	return vr
}

func (v Vector) Multiply(k float64) Vector {
	var vr Vector = make([]float64, len(v))
	for i := range vr {
		vr[i] = v[i] * k
	}
	return vr
}

func (v Vector) DotProduct(v2 Vector) float64 {
	dotProduct := 0.0
	for i := range v {
		dotProduct += v[i] * v2[i]
	}
	return dotProduct
}
