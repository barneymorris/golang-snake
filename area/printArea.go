package area

import (
	"fmt"
	"snake/constants"
	"snake/utils"
)

func printArea(area [constants.AREA_WIDTH][constants.AREA_HIGTH]string) {
	for i := 0; i < constants.AREA_WIDTH; i++ {
		for j := 0; j < constants.AREA_HIGTH; j++ {
			fmt.Printf("%s ", area[i][j])
		}
		fmt.Printf(" \n")
	}

	utils.ClearTerminal()
}
