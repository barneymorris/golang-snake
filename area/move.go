package area

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"snake/constants"
	"snake/utils"
	"time"
)

func Move() {

	var pointer keyboard.Key = 0
	var prevPointer keyboard.Key = 0

	go utils.OnKeyPress(&pointer)

	for {
		time.Sleep(200 * time.Millisecond)

		if Area.CURRENT_HEAD_X == 0 || Area.CURRENT_HEAD_Y == 0 || Area.CURRENT_HEAD_X == constants.AREA_WIDTH-1 || Area.CURRENT_HEAD_Y == constants.AREA_HIGTH-1 {
			utils.YouLoss()
		}

		utils.ClearTerminal()

		fmt.Printf("Score: %d\n", Area.TAIL_SIZE*10-10)

		Area.Print()

		switch int(pointer) {
		case constants.GetKeyCodesStruct().UP:
			if int(prevPointer) != constants.GetKeyCodesStruct().DOWN {
				prevPointer = pointer

				Area.MoveHead(Area.CURRENT_HEAD_X-1, Area.CURRENT_HEAD_Y, "UP")
				break
			} else {
				pointer = prevPointer
			}

		case constants.GetKeyCodesStruct().DOWN:
			if int(prevPointer) != constants.GetKeyCodesStruct().UP {

				prevPointer = pointer

				Area.MoveHead(Area.CURRENT_HEAD_X+1, Area.CURRENT_HEAD_Y, "DOWN")
				break
			} else {
				pointer = prevPointer
			}

		case constants.GetKeyCodesStruct().LEFT:
			if int(prevPointer) != constants.GetKeyCodesStruct().RIGHT {

				prevPointer = pointer

				Area.MoveHead(Area.CURRENT_HEAD_X, Area.CURRENT_HEAD_Y-1, "RIGHT")
				break
			} else {
				pointer = prevPointer
			}

		case constants.GetKeyCodesStruct().RIGHT:
			if int(prevPointer) != constants.GetKeyCodesStruct().LEFT {

				prevPointer = pointer

				Area.MoveHead(Area.CURRENT_HEAD_X, Area.CURRENT_HEAD_Y+1, "LEFT")
				break
			} else {
				pointer = prevPointer
			}
		}

		if pointer == keyboard.Key(constants.GetKeyCodesStruct().ESC) {
			fmt.Printf("ESC key has been pressed. Now quiting...\n")
			return
		}
	}
}
