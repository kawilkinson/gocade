package screens

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/utils"
)

type MenuStyles struct {
	TitleStyle        lipgloss.Style
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
	QuitTextStyle     lipgloss.Style
	GopherStyle       lipgloss.Style
}

func CreateMenuStyle() *MenuStyles {
	menuStyles := MenuStyles{
		TitleStyle:        lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color(utils.GopherColor)),
		ItemStyle:         lipgloss.NewStyle().PaddingLeft(4),
		SelectedItemStyle: lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(utils.GopherColor)),
		QuitTextStyle:     lipgloss.NewStyle().Bold(true),
		GopherStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00ADD8")).
			Bold(true).Padding(1, 2).
			Align(lipgloss.Center),
	}

	return &menuStyles
}
