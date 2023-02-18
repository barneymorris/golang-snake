package constraints

import (
	"fmt"
	"os"
	"runtime"
)

func ExitIfOSIsNotWindows() {
	if runtime.GOOS != "windows" {
		err := fmt.Errorf("%s", "Avaiable only for Windows users. Exiting...")
		panic(err)
		os.Exit(1)
	}
}
