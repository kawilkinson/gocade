package screens

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/kawilkinson/gocade/internal/menuconfig"
)

func NewMainMenu(width, height int, keys *menuconfig.MainMenuKeys, style *MenuStyles) list.Model {
	items := []list.Item{
		MenuItem("Play Game"),
		MenuItem("High Scores"),
		MenuItem("Quit"),
	}

	delegate := MenuDelegate{Styles: style}

	mainMenu := list.New(items, delegate, width, height)
	mainMenu.Title = "Gocade"
	mainMenu.SetShowStatusBar(false)
	mainMenu.SetFilteringEnabled(false)
	mainMenu.Styles.Title = style.TitleStyle
	return mainMenu
}
