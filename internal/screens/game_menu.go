package screens

import "github.com/charmbracelet/bubbles/list"

func NewGameMenu(width, height int) list.Model {
	games := []list.Item{ // not final games, currently here as placeholders
		MenuItem("Snake"),
		MenuItem("Tetris"),
		MenuItem("Pong"),
	}

	li := list.New(games, MenuDelegate{}, width, height)
	li.Title = "Select a Game"
	li.SetShowStatusBar(false)
	li.SetFilteringEnabled(false)
	li.Styles.Title = titleStyle
	return li
}
