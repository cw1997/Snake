package World

import (
	"Snake/Console"
	Food "Snake/GameObject/Food"
	Snake "Snake/GameObject/Snake"
	"fmt"

)

const (
	Height = 20
	Width  = 20
)

type World struct {
	OrginalWorld [Width][Height]string
}

func (world *World) Render(snake Snake.Snake, food Food.Food) {
	// create renderWorld
	renderWorld := world.OrginalWorld
	// set food position to renderWorld
	renderWorld[food.Position.Y][food.Position.X] = food.Color
	for i := range snake.Position {
		renderWorld[snake.Position[i].Y][snake.Position[i].X] = snake.Color
	}
	// clear console first
	Console.ClearScreen()

	// render world
	for x := range renderWorld {
		fmt.Println(renderWorld[x])
	}
}
