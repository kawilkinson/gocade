package screens

import (
	"strings"

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

// helper function to ensure the art for the gopher is always aligned properly
func NormalizeWidth(ascii string) string {
	lines := strings.Split(ascii, "\n")
	max := 0

	for _, line := range lines {
		if len([]rune(line)) > max {
			max = len([]rune(line))
		}
	}

	for i, line := range lines {
		pad := max - len([]rune(line))
		if pad > 0 {
			lines[i] = line + strings.Repeat(" ", pad)
		}
	}
	return strings.Join(lines, "\n")
}
