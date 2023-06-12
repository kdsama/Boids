package main

import "math"

type Vector2D struct {
	x float64
	y float64
}

func (v *Vector2D) Add(vel Vector2D) Vector2D {

	return Vector2D{v.x + vel.x,
		v.y + vel.y,
	}
}

func (v *Vector2D) Subtract(vel Vector2D) Vector2D {

	return Vector2D{v.x - vel.x,
		v.y - vel.y,
	}
}

func (v *Vector2D) Multiply(vel Vector2D) Vector2D {

	return Vector2D{v.x * vel.x,
		v.y * vel.y,
	}
}

// add values
func (v *Vector2D) AddV(vel float64) Vector2D {

	return Vector2D{v.x + vel,
		v.y + vel,
	}
}

func (v *Vector2D) SubtractV(vel float64) Vector2D {

	return Vector2D{v.x - vel,
		v.y - vel,
	}
}

func (v *Vector2D) MultiplyV(vel float64) Vector2D {

	return Vector2D{v.x * vel,
		v.y * vel,
	}
}

func (v *Vector2D) DivisionV(vel float64) Vector2D {

	return Vector2D{v.x / vel,
		v.y / vel,
	}
}

func (v *Vector2D) Limit(lower float64, upper float64) Vector2D {

	return Vector2D{math.Min(math.Max(v.x, lower), upper),
		math.Min(math.Max(v.y, lower), upper),
	}
}

func (v *Vector2D) Distance(v1 Vector2D) float64 {

	return math.Sqrt(math.Pow(v.x-v1.x, 2) + math.Pow(v.y-v1.y, 2))
}
