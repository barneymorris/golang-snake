package utils

import (
	"fmt"
	"os"
)

func YouLoss() {
	ClearTerminal()
	fmt.Printf("\nYou loss! Game is over!\n")
	os.Exit(0)
}
