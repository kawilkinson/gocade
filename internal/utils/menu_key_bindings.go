package utils

import "github.com/charmbracelet/bubbles/key"

type MainMenuKeys struct {
	BackKey key.Binding
}

func (k MainMenuKeys) ShortHelp() []key.Binding {
	return []key.Binding{
		k.BackKey,
	}
}

func (k MainMenuKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.BackKey},
	}
}

func AdditionalMainMenuKeys() *MainMenuKeys {
	keys := &MainMenuKeys{
		BackKey: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "return to prev screen")),
	}

	return keys
}
