package screens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/kawilkinson/gocade/internal/menuconfig"
)

// function currently is using the game menu logic as a placeholder
func NewScoreMenu(width, height int, keys *menuconfig.MainMenuKeys, style *MenuStyles) list.Model {
	games := []list.Item{
		MenuItem("Marathon Tetris"),
		MenuItem("Sprint Tetris"),
		MenuItem("Ultra Tetris"),
	}

	delegate := MenuDelegate{Styles: style}

	scoreMenu := list.New(games, delegate, width, height)
	scoreMenu.Title = "Select a High Score list to view"
	scoreMenu.SetShowStatusBar(false)
	scoreMenu.SetFilteringEnabled(false)
	scoreMenu.Styles.Title = style.TitleStyle

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
