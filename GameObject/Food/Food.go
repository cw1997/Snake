package Food

import (
	GameObject "Snake/GameObject"
	"math/rand"
)

const (
	RANGE = 20
)

type Food struct {
	Position GameObject.Position
	Color    string
	Alive    bool
}

// generate food coordinate
func (food *Food) Generate() {
	food.Position.X = rand.Intn(RANGE)
	food.Position.Y = rand.Intn(RANGE)
	food.Alive = true
}
