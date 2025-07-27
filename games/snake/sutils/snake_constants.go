package sutils

const (
	TimeInterval = 80

	// directions for snake to go
	Left = 1 + iota
	Right
	Up
	Down

	// text to render for the game
	HelpMessage = "\n\nUse arrow keys or wasd keys to nagivate the snake.\n\n\nq or ctrl+c quits the game\n"
	GameOverMessage = "\n\nGame Over!\n"

)
