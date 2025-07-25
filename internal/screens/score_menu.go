package screens

import "github.com/charmbracelet/bubbles/list"

// function currently is using the game menu logic as a placeholder
func NewScoreMenu(width, height int) list.Model {
	games := []list.Item{
		MenuItem("Snake"),
		MenuItem("Tetris"),
		MenuItem("Pong"),
	}

	li := list.New(games, MenuDelegate{}, width, height)
	li.Title = "Select a High Score list to view"
	li.SetShowStatusBar(false)
	li.SetFilteringEnabled(false)
	li.Styles.Title = titleStyle
	return li
}
