package screens

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/utils"
)

func NewMainMenu(width, height int) list.Model {
	items := []list.Item{
		MenuItem("Play Game"),
		MenuItem("High Scores"),
		MenuItem("Quit"),
	}

	mainMenu := list.New(items, MenuDelegate{}, width, height)
	mainMenu.Title = "Gocade"
	mainMenu.SetShowStatusBar(false)
	mainMenu.SetFilteringEnabled(false)
	mainMenu.Styles.Title = titleStyle
	return mainMenu
}

func RenderGopher(width, height int) string {
	gopher := GopherStyle.Render(normalizeWidth(utils.GopherMascot))
	return lipgloss.Place(width, lipgloss.Height(utils.GopherMascot), lipgloss.Center, lipgloss.Top, gopher)
}
