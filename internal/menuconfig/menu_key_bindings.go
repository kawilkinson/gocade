package menuconfig

import "github.com/charmbracelet/bubbles/key"

type MainMenuKeys struct {
	BackKey key.Binding
	Exit    key.Binding
	Select  key.Binding
}

// q for quit is already in help by default
func (k MainMenuKeys) ShortHelp() []key.Binding {
	return []key.Binding{
		k.BackKey,
		k.Select,
	}
}

// q for quit is already in help by default
func (k MainMenuKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.BackKey},
		{k.Select},
	}
}

func SetExtraMainMenuKeys() *MainMenuKeys {
	keys := &MainMenuKeys{
		BackKey: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "go back")),
		Exit:    key.NewBinding(key.WithKeys("q")),
		Select:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	}

	return keys
}
