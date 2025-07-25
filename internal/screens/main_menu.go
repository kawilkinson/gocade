package screens

import "github.com/charmbracelet/bubbles/list"

func NewMainMenu(width, height int) list.Model {
	items := []list.Item{
		MenuItem("Play Game"),
		MenuItem("High Scores"),
		MenuItem("Quit"),
	}

	l := list.New(items, MenuDelegate{}, width, height)
	l.Title = "Gocade"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	return l
}
