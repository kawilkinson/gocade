package main

type screen int

const (
	MenuWidth             = 20
	MenuHeight            = 9
	GopherColor           = "#00ADD8"
	ScreenMainMenu screen = iota
	ScreenGameMenu
	ScreenScoreMenu
)
