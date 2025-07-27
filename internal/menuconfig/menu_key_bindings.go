package menuconfig

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

func SetExtraMainMenuKeys() *MainMenuKeys {
	keys := &MainMenuKeys{
		BackKey: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "go back")),
	}

	return keys
}
