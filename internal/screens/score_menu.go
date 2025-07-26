package screens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/kawilkinson/gocade/internal/utils"
)

// function currently is using the game menu logic as a placeholder
func NewScoreMenu(width, height int, keys *utils.MainMenuKeys) list.Model {
	games := []list.Item{
		MenuItem("Snake"),
		MenuItem("Tetris"),
		MenuItem("Pong"),
	}

	scoreMenu := list.New(games, MenuDelegate{}, width, height)
	scoreMenu.Title = "Select a High Score list to view"
	scoreMenu.SetShowStatusBar(false)
	scoreMenu.SetFilteringEnabled(false)
	scoreMenu.Styles.Title = titleStyle

	scoreMenu.AdditionalShortHelpKeys = keys.ShortHelp
	scoreMenu.AdditionalFullHelpKeys = func() []key.Binding {
		var flattenedKeys []key.Binding
		for _, row := range keys.FullHelp() {
			flattenedKeys = append(flattenedKeys, row...)
		}

		return flattenedKeys
	}

	return scoreMenu
}
