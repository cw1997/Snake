package Snake

import (
	"log"

	"Snake/GameObject"
	Food "Snake/GameObject/Food"

	"github.com/mattn/go-tty"
)

// define direction
type Direction struct {
	Horizontal int
	Vertical   int
	Forward    int
}

// define Snake
type Snake struct {
	Body      int
	Position  []GameObject.Position
	Direction Direction
	Color     string
}

// change snake coordinate
func (snake *Snake) ChangePosition() bool {
	// length > 1
	if len(snake.Position) > 1 {
		// move data from last to first
		for i := len(snake.Position); i > 0; i-- {
			if i > 1 {
				snake.Position[i-1] = snake.Position[i-2]
			}
		}
	}
	// increase first position
	// if out of index then game over
	switch snake.Direction.Forward {
	case Up:
		snake.Position[0].Y -= snake.Direction.Vertical
		return determineGameOver(snake)
	case Down:
		snake.Position[0].Y += snake.Direction.Vertical
		return determineGameOver(snake)
	case Left:
		snake.Position[0].X -= snake.Direction.Horizontal
		return determineGameOver(snake)
	case Right:
		snake.Position[0].X += snake.Direction.Horizontal
		return determineGameOver(snake)
	default:
		return true
	}
}

// if player press key then change direction or keep the direction
func (snake *Snake) ChangeDirection() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		switch string(r) {
		case "W", "w":
			// direction locker
			if snake.Direction.Forward != Up && snake.Direction.Forward != Down {
				snake.Direction.Forward = Up
			}
		case "S", "s":
			// direction locker
			if snake.Direction.Forward != Up && snake.Direction.Forward != Down {
				snake.Direction.Forward = Down
			}
		case "A", "a":
			// direction locker
			if snake.Direction.Forward != Left && snake.Direction.Forward != Right {
				snake.Direction.Forward = Left
			}
		case "D", "d":
			// direction locker
			if snake.Direction.Forward != Left && snake.Direction.Forward != Right {
				snake.Direction.Forward = Right
			}
		default:
		}
	}

}

// Ate food
func (snake *Snake) Ate(food Food.Food) bool {
	if snake.Position[0] == food.Position {
		snake.Position = append(snake.Position, snake.Position[len(snake.Position)-1])
		return false
	}
	return true
}

// determine gameover by hit the wall and hit self body
func determineGameOver(snake *Snake) bool {
	if snake.Position[0].X > 19 || snake.Position[0].X < 0 ||
		snake.Position[0].Y > 19 || snake.Position[0].Y < 0 {
		return false
	} else if len(snake.Position) > 3 {
		for _, v := range snake.Position[3:] {
			if v == snake.Position[0] {
				return false
			}
		}
		return true
	}
	return true
}
