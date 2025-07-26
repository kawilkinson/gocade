package tutils

type Screen int

const (
	ScreenMenu = Screen(iota)
	ScreenTetrisGame
)

var ScreenToStrMap = map[Screen]string{
	ScreenMenu:       "Menu",
	ScreenTetrisGame: "Tetris",
}

func (s Screen) String() string {
	return ScreenToStrMap[s]
}
