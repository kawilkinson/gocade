package snakeconfig

import (
	"github.com/charmbracelet/bubbles/key"
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

type SwitchToGameMsg struct {
    Username string
    Screen   sutils.Screen
}

type SwitchToMenuMsg struct {}
