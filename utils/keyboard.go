package utils

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func OnKeyPress(p *keyboard.Key) {
	fmt.Println("Press ESC to quit")

	for {
		_, key, err := keyboard.GetKey()

		if err != nil {
			return
		}

		*p = key

		if key == keyboard.KeyEsc {
			break
		}
	}
}
