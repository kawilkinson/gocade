package sutils

import "github.com/kawilkinson/gocade/internal/utils"

type Screen int

const (
	TimeInterval = 80

	SnakeScoreFile = "internal/leaderboard/data/snake_scores.csv"



	// screens for snake game
	SnakeMenu = Screen(iota)
	SnakeGame

	// directions for snake to go
	Left = 1 + iota
	Right
	Up
	Down

	// text to render for the game
	HelpMessage     = "\n\nUse arrow keys or wasd keys to nagivate the snake.\n\n\nq, esc, or ctrl+c quits the game\n"
	GameOverMessage = "\n\nGame Over!\n"

	SnakeTitle = `
	/  ___|           | |       
	\ '--. _ __   __ _| | _____ 
	'--. \ '_ \ / _' | |/ / _ \
	/\__/ / | | | (_| |   <  __/
	\____/|_| |_|\__,_|_|\_\___|
	`
)

// helper function to ensure large ASCII text always shows correctly in Snake
func RenderLargeText(ascii string) string {
	normalizedTitle := utils.NormalizeWidth(ascii)
	return normalizedTitle
}
