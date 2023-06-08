package main

type Vector2D struct {
	x float64
	y float64
}

func (v *Vector2D) Add(vel Vector2D) Vector2D {

	return Vector2D{v.x + vel.x,
		v.y + vel.y,
	}
}
