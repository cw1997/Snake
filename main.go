/*
name: snake game
author: cw1997
date: 2019.2.11
*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"Snake/Console"
	GameObject "Snake/GameObject"
	Food "Snake/GameObject/Food"
	Snake "Snake/GameObject/Snake"
	World "Snake/GameObject/World"
)

// declare const
const (
	SIZE  = 20
	SPEED = 500
)

// declare
var (
	GameWorld World.World 
	snake Snake.Snake = Snake.Snake{
		Body: 1,
		Position: []GameObject.Position{
			{
				X: 10,
				Y: 10,
			},
		},
		Direction: Snake.Direction{
			Horizontal: 1,
			Vertical:   1,
			Forward:    Snake.Right,
		},
		Color: "⚪",
	}
	food Food.Food = Food.Food{
		Color: "🟥",
	}
	State GameState
)

func Update() {
	for State.GameON {
		food.Alive = snake.Ate(food)
		if !food.Alive {
			food.Generate()
		}
		go snake.ChangeDirection()
		// change snake position first
		State.GameON = snake.ChangePosition()
		if State.GameON {
			// render world
			GameWorld.Render(snake, food)
			// sleep a duration to render again
			time.Sleep(SPEED / 2 * time.Millisecond)
		}
	}
	if !State.GameON{
		Console.ClearScreen()
		fmt.Println("Game Over")
	}
}

func Start() {
	for i :=0; i <20;i++ {
		for j := 0;j<20;j++{
			GameWorld.OrginalWorld[i][j] = "⬛"
		}
	}
	State.GameON = true
}

// program enter point
func main() {
	rand.Seed(time.Now().UnixNano())
	Start()
	Update()
}
