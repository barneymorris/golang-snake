package constants

type KeyCodesStruct struct {
	ESC int
}

func GetKeyCodesStruct() KeyCodesStruct {
	values := KeyCodesStruct{ESC: 27}
	return values
}

const AREA_WIDTH = 24 + 2
const AREA_HIGTH = 24 + 2

const SYMBOL_HEAD = "O"
const SYMBOL_TAIL = "="
const SYMBOL_FRUIT = "x"
const SYMBOL_WALL = "*"
