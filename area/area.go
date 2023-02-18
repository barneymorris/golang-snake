package area

import (
	"fmt"
	"math/rand"
	"snake/constants"
	"snake/utils"
	"time"
)

type AreaStruct struct {
	CURRENT_HEAD_X int
	CURRENT_HEAD_Y int
	FRUIT_X        int
	FRUINT_Y       int
	TAIL_SIZE      int
	TAIL_COORDS    [][]int
	Area           [constants.AREA_HIGTH][constants.AREA_WIDTH]string
}

var Area AreaStruct

func (Area *AreaStruct) InitArea() {
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

	rand.Seed(time.Now().UnixNano())

	// Generating random initial position for snake
	randomX := 5
	randomY := 5

	area[randomX+1][randomY] = constants.SYMBOL_TAIL
	item := []int{randomX + 1, randomY}
	Area.TAIL_COORDS = append(Area.TAIL_COORDS, item)
	Area.TAIL_SIZE = 1

	area[randomX][randomY] = constants.SYMBOL_HEAD

	Area.CURRENT_HEAD_X = randomX
	Area.CURRENT_HEAD_Y = randomY
	Area.GenerateFruit(&area)
	Area.Area = area
}

func (Area *AreaStruct) MoveHead(x, y int, direction string) {
	oldX := Area.CURRENT_HEAD_X
	oldY := Area.CURRENT_HEAD_Y

	isFruitAte := oldX == Area.FRUIT_X && oldY == Area.FRUINT_Y

	if isFruitAte {
		Area.TAIL_SIZE++
		item := []int{oldX, oldY}
		Area.TAIL_COORDS = append(Area.TAIL_COORDS, item)
	} else {
		Area.Area[oldX][oldY] = " "
	}

	for i := 0; i < len(Area.TAIL_COORDS); i++ {
		x := Area.TAIL_COORDS[i][0]
		y := Area.TAIL_COORDS[i][1]

		Area.Area[x][y] = " "
	}

	if Area.TAIL_SIZE > 0 {
		for i := 0; i < len(Area.TAIL_COORDS); i++ {
			if i == len(Area.TAIL_COORDS)-1 {
				item := []int{oldX, oldY}
				Area.TAIL_COORDS[i] = item
			} else {
				Area.TAIL_COORDS[i] = Area.TAIL_COORDS[i+1]
			}
		}
	}

	for i := 0; i < len(Area.TAIL_COORDS); i++ {
		x := Area.TAIL_COORDS[i][0]
		y := Area.TAIL_COORDS[i][1]

		Area.Area[x][y] = constants.SYMBOL_TAIL
	}

	if Area.Area[x][y] == constants.SYMBOL_TAIL {
		utils.YouLoss()
	}

	Area.Area[x][y] = constants.SYMBOL_HEAD
	Area.CURRENT_HEAD_X = x
	Area.CURRENT_HEAD_Y = y

	if isFruitAte {
		Area.GenerateFruit(&Area.Area)
	}
}

func (Area *AreaStruct) Print() {
	for i := 0; i < constants.AREA_HIGTH; i++ {
		for j := 0; j < constants.AREA_WIDTH; j++ {
			fmt.Printf("%s ", Area.Area[i][j])
		}
		fmt.Printf(" \n")
	}
}

func (Area *AreaStruct) GenerateFruit(area *[constants.AREA_HIGTH][constants.AREA_WIDTH]string) {
	rand.Seed(time.Now().UnixNano())

	randomX := rand.Intn(constants.AREA_WIDTH - 2)
	randomY := rand.Intn(constants.AREA_HIGTH - 2)

	for area[randomX][randomY] != " " {
		randomX = rand.Intn(constants.AREA_WIDTH - 2)
		randomY = rand.Intn(constants.AREA_HIGTH - 2)
	}

	fmt.Printf("%d, %d\n", randomX, randomY)

	area[randomX][randomY] = constants.SYMBOL_FRUIT
	Area.FRUIT_X = randomX
	Area.FRUINT_Y = randomY
}
