package tetrisconfig

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

type MenuKeyMap struct {
	Exit     key.Binding
	FormKeys *huh.KeyMap
}

func SetTetrisMenuKeys() *MenuKeyMap {
	keys := &MenuKeyMap{
		Exit:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit")),
		FormKeys: huh.NewDefaultKeyMap(),
	}

	keys.FormKeys.Quit.SetEnabled(false)

	return keys
}
