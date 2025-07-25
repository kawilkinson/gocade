package utils

type Screen int

const (
	MenuWidth   = 20
	MenuHeight  = 9
	GopherColor = "#00ADD8"

	ScreenMainMenu Screen = iota
	ScreenGameMenu
	ScreenScoreMenu
)
