package constants

type KeyCodesStruct struct {
	ESC   int
	UP    int
	DOWN  int
	LEFT  int
	RIGHT int
}

func GetKeyCodesStruct() KeyCodesStruct {
	values := KeyCodesStruct{ESC: 27, UP: 65517, LEFT: 65515, RIGHT: 65514, DOWN: 65516}
	return values
}

const AREA_WIDTH = 24 + 2
const AREA_HIGTH = 24 + 2

const SYMBOL_HEAD = "O"
const SYMBOL_TAIL = "="
const SYMBOL_FRUIT = "x"
const SYMBOL_WALL = "*"
