package tetrisscreens

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

type menuKeyMap struct {
	Exit     key.Binding
	formKeys *huh.KeyMap
}

func SetTetrisMenuKeys() *menuKeyMap {
	keys := &menuKeyMap{
		Exit:     key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit")),
		formKeys: huh.NewDefaultKeyMap(),
	}

	keys.formKeys.Quit.SetEnabled(false)

	return keys
}
