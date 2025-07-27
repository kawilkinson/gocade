package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/internal/utils"
)

type ChangeScreenMsg struct {
	Screen utils.Screen
}

func SwitchScreens(screen utils.Screen) tea.Cmd {
	return func() tea.Msg {
		return ChangeScreenMsg{
			Screen: screen,
		}
	}
}
