package screens

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/utils"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color(utils.GopherColor))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(utils.GopherColor))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)            // currently not used but may consider later
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1) // currently not used but may consider later
	QuitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)
