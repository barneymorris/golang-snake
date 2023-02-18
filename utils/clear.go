package utils

import (
	"github.com/inancgumus/screen"
)

func ClearTerminal() {
	// Clears the screen
	screen.Clear()
	screen.MoveTopLeft()
}
