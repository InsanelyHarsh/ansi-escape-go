package main

type Ball struct {
	Item
	velocity Vector
	Speed    float64
}

func (b *Ball) GetNextPosition() Vector {
	return Vector{
		X: b.position.X + b.velocity.X,
		Y: b.position.Y + b.velocity.Y,
	}
}
