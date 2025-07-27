package screens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/menuconfig"
	"github.com/kawilkinson/gocade/internal/utils"
)

func NewGameMenu(width, height int, keys *menuconfig.MainMenuKeys, style *MenuStyles) list.Model {
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

func RenderGopher(width, height int, style *MenuStyles) string {
	gopher := style.GopherStyle.Render(utils.NormalizeWidth(utils.GopherMascot))
	return lipgloss.Place(width, lipgloss.Height(utils.GopherMascot), lipgloss.Center, lipgloss.Top, gopher)
}
