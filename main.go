package main

import (
	"fmt"
	"math/rand"
)

const AREA_WIDTH = 24 + 2
const AREA_HIGTH = 24 + 2

const SYMBOL_HEAD = "O"
const SYMBOL_TAIL = "="
const SYMBOL_FRUIT = "x"
const SYMBOL_WALL = "*"

func main() {
	fmt.Printf("Initing...\n")

	var area [AREA_HIGTH][AREA_WIDTH]string

	// Init area with empty string
	for i := 0; i < AREA_HIGTH; i++ {
		for j := 0; j < AREA_WIDTH; j++ {
			isYWall := j == 0 || j == AREA_HIGTH - 1
			isXWall := i == 0 || i == AREA_WIDTH - 1

			area[i][j] = " "

			if (isYWall) {
				area[i][j] = SYMBOL_WALL
			}

			if (isXWall) {
				area[i][j] = SYMBOL_WALL
			}
		}
	}

	// Generating random initial position for snake
	randomX := rand.Intn(AREA_WIDTH)
	randomY := rand.Intn(AREA_HIGTH)

	area[randomX][randomY] = SYMBOL_HEAD

	printArea(area)
}

func printArea(area [AREA_WIDTH][AREA_HIGTH]string) {
	for i := 0; i < AREA_HIGTH; i++ {
		for j := 0; j < AREA_HIGTH; j++ {
			fmt.Printf("%s ", area[i][j])
		}
		fmt.Printf(" \n")
	}
}