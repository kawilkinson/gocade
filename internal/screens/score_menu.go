package screens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/kawilkinson/gocade/internal/utils"
)

// function currently is using the game menu logic as a placeholder
func NewScoreMenu(width, height int) list.Model {
	games := []list.Item{
		MenuItem("Snake"),
		MenuItem("Tetris"),
		MenuItem("Pong"),
	}

	scoreMenu := list.New(games, MenuDelegate{}, width, height)
	scoreMenu.Title = "Select a High Score list to view"
	scoreMenu.AdditionalShortHelpKeys = func() []key.Binding { return []key.Binding{utils.KeyBindings} }
	scoreMenu.AdditionalFullHelpKeys = func() []key.Binding { return []key.Binding{utils.KeyBindings} }
	scoreMenu.SetShowStatusBar(false)
	scoreMenu.SetFilteringEnabled(false)
	scoreMenu.Styles.Title = titleStyle
	return scoreMenu
}
