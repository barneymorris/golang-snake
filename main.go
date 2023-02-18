package main

import (
	"fmt"
	"math/rand"
	area2 "snake/area"
	"snake/constants"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Printf("Initing...\n")

	var area [constants.AREA_HIGTH][constants.AREA_WIDTH]string

	// Init area with empty string
	for i := 0; i < constants.AREA_HIGTH; i++ {
		for j := 0; j < constants.AREA_WIDTH; j++ {
			isYWall := j == 0 || j == constants.AREA_HIGTH-1
			isXWall := i == 0 || i == constants.AREA_WIDTH-1

			area[i][j] = " "

			if isYWall {
				area[i][j] = constants.SYMBOL_WALL
			}

			if isXWall {
				area[i][j] = constants.SYMBOL_WALL
			}
		}
	}

	// Generating random initial position for snake
	randomX := rand.Intn(constants.AREA_WIDTH - 2)
	randomY := rand.Intn(constants.AREA_HIGTH - 2)

	area[randomX][randomY] = constants.SYMBOL_HEAD

	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	area2.Move(area)
}
