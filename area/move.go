package area

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"snake/constants"
	"snake/utils"
	"time"
)

func Move(area [26][26]string) {

	var pointer keyboard.Key = 0

	go utils.OnKeyPress(&pointer)

	for {
		time.Sleep(1 * time.Second)
		printArea(area)

		if pointer == keyboard.Key(constants.GetKeyCodesStruct().ESC) {
			fmt.Printf("ESC key has been pressed. Now quiting...\n")
			return
		}
	}
}
