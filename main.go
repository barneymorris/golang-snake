package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	area2 "snake/area"
)

func main() {
	fmt.Printf("Initing...\n")
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	area2.Area.InitArea()
	area2.Move()
}
