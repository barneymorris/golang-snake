package utils

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"snake/constants"
)

func OnKeyPress(p *keyboard.Key) {
	fmt.Println("Press ESC to quit")

	for {
		_, key, err := keyboard.GetKey()

		if err != nil {
			return
		}

		i := int(key)
		ESC := constants.GetKeyCodesStruct().ESC
		LEFT := constants.GetKeyCodesStruct().LEFT
		RIGHT := constants.GetKeyCodesStruct().RIGHT
		UP := constants.GetKeyCodesStruct().UP
		DOWN := constants.GetKeyCodesStruct().DOWN

		if i == ESC || i == LEFT || i == RIGHT || i == UP || i == DOWN {
			*p = key
		}

		if key == keyboard.KeyEsc {
			break
		}
	}
}
