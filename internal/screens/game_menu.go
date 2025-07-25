package screens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/kawilkinson/gocade/internal/utils"
)

func NewGameMenu(width, height int) list.Model {
	games := []list.Item{ // not final games, currently here as placeholders
		MenuItem("Snake"),
		MenuItem("Tetris"),
		MenuItem("Pong"),
	}

	gameMenu := list.New(games, MenuDelegate{}, width, height)
	gameMenu.Title = "Select a Game"
	gameMenu.AdditionalShortHelpKeys = func() []key.Binding { return []key.Binding{utils.KeyBindings} }
	gameMenu.AdditionalFullHelpKeys = func() []key.Binding { return []key.Binding{utils.KeyBindings} }
	gameMenu.SetShowStatusBar(false)
	gameMenu.SetFilteringEnabled(false)
	gameMenu.Styles.Title = titleStyle
	return gameMenu
}
