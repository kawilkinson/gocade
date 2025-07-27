package snakeconfig

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/kawilkinson/gocade/games/snake/sutils"
)

type MenuKeyMap struct {
	Exit     key.Binding
	FormKeys *huh.KeyMap
}

func SetSnakeMenuKeys() *MenuKeyMap {
	keys := &MenuKeyMap{
		Exit:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit")),
		FormKeys: huh.NewDefaultKeyMap(),
	}

	keys.FormKeys.Quit.SetEnabled(false)

	return keys
}

type MenuInput struct{}

func NewMenuInput() *MenuInput {
	return &MenuInput{}
}

func (in *MenuInput) isSwitchModeInput() {}

type SwitchScreenMsg struct {
	Target sutils.Screen
	Input  SwitchModeInput
}

type SwitchModeInput interface {
	isSwitchModeInput()
}

func SwitchModeCmd(target sutils.Screen, in SwitchModeInput) tea.Cmd {
	return func() tea.Msg {
		return SwitchScreenMsg{
			Target: target,
			Input:  in,
		}
	}
}

type SwitchToGameMsg struct {
    Username string
    Screen   sutils.Screen
}
