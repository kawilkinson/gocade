package screens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/kawilkinson/gocade/internal/utils"
)

func NewGameMenu(width, height int, keys *utils.MainMenuKeys, style *MenuStyles) list.Model {
	games := []list.Item{ // not final games, currently here as placeholders
		MenuItem("Snake"),
		MenuItem("Tetris"),
		MenuItem("Pong"),
	}

	delegate := MenuDelegate{Styles: style}

	gameMenu := list.New(games, delegate, width, height)
	gameMenu.Title = "Select a Game"
	gameMenu.SetShowStatusBar(false)
	gameMenu.SetFilteringEnabled(false)
	gameMenu.Styles.Title = style.TitleStyle

	gameMenu.AdditionalShortHelpKeys = keys.ShortHelp
	gameMenu.AdditionalFullHelpKeys = func() []key.Binding {
		var flattenedKeys []key.Binding
		for _, row := range keys.FullHelp() {
			flattenedKeys = append(flattenedKeys, row...)
		}

		return flattenedKeys
	}

	return gameMenu
}
