/*
name: snake game
author: cw1997
date: 2019.2.11
 */

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	SIZE = 20
	SPEED = 500
)


type Position struct {
	x int
	y int
}
type Snake struct {
	snake   *list.List
	forward string
}


var (
	world [SIZE][SIZE]string
	empty []Position
	snk = new(Snake)
	foodPosition Position
)

func render() {
	for row:=0; row< SIZE*5; row++  {
		for col:=0; col< SIZE*5; col++  {
			fmt.Print(" ")
		}
		fmt.Println()
	}

	for row:=0; row< SIZE; row++  {
		for col:=0; col< SIZE; col++  {
			switch world[row][col] {
			case "food":
				fmt.Print("*")
				break
			case "snake":
				fmt.Print("□")
				break
			default:
				fmt.Print(" ")
				empty = append(empty, Position{col, row})
				break
			}
		}
		fmt.Println()
	}
}

func input() {
	var s string
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s = input.Text()
	//fmt.Scanf("%s", &s)

	switch s {
	case "w", "s", "a", "d":
		snk.forward = s
		break
	}
}

func generateFood() Position {
	pos := empty[rand.Intn(len(empty))]
	world[pos.y][pos.x] = "food"
	foodPosition = pos
	return pos
}

func collisionDetect(gameObjectPos1, gameObjectPos2 Position) bool {
	return (gameObjectPos1.x == gameObjectPos2.x) && (gameObjectPos1.y == gameObjectPos2.y)
}

func update() {
	// get snake's head position
	headPos := snk.snake.Front().Value.(Position)

	var newHeadPos Position
	switch snk.forward {
	case "w":
		newHeadPos = Position{(headPos.x)-0, (headPos.y)-1}
	case "s":
		newHeadPos = Position{(headPos.x)-0, (headPos.y)+1}
	case "a":
		newHeadPos = Position{(headPos.x)-1, (headPos.y)-0}
	case "d":
		newHeadPos = Position{(headPos.x)+1, (headPos.y)-0}
	}

	//newHeadPos := snk.snake.Front().Value.(Position)

	// if collise , snake die
	gameover := newHeadPos.x<0 || newHeadPos.x>=SIZE || newHeadPos.y<0 || newHeadPos.y>=SIZE
	if gameover {
		fmt.Println("GAME OVER!!!")
		fmt.Println("Thanks for play my game. author: cw1997.")
		//for row:=0; row< SIZE-2; row++  {
		//	for col:=0; col< SIZE; col++  {
		//		fmt.Print(" ")
		//	}
		//	fmt.Println()
		//}
		log.Fatal("end")
		var i int
		fmt.Scanf("%d", &i)
	}

	snk.snake.PushFront(newHeadPos)


	//newFoodPos := generateFood()

	// if collided , snake eat food, length++, and generate a new food
	collisionStatus := collisionDetect(newHeadPos, foodPosition)
	// if not collided, remove the snake's tail
	if collisionStatus {
		// generate a new food
		generateFood()
	} else {
		tailPos := snk.snake.Back().Value.(Position)
		world[tailPos.y][tailPos.x] = "empty"
		snk.snake.Remove(snk.snake.Back())
	}


	//write snake position to world map
	for e := snk.snake.Front(); e != nil; e = e.Next() {
		pos := e.Value.(Position)
		world[pos.y][pos.x] = "snake"
	}
	fmt.Println()
}

func initialize() {
	snk.snake = list.New()
	snk.snake.PushFront( Position{0,0} )
	world[0][0] = "snake"

	render()
	generateFood()
}

func main() {

	initialize()


	// run input routine at GoRoutine
	go func() {
		for {
			input()
			time.Sleep(SPEED * time.Millisecond)
		}
	}()

	for {
		render()
		//input()
		update()

		time.Sleep(SPEED * time.Millisecond)
	}

}
