package menuconfig

import "github.com/charmbracelet/bubbles/key"

type LeaderboardKeyBindings struct {
	Help key.Binding
	Back key.Binding
}

func DefaultLeaderboardKeyBindings() *LeaderboardKeyBindings {
	return &LeaderboardKeyBindings{
		Help: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		Back: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "go back")),
	}
}

func (k *LeaderboardKeyBindings) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Back,
		k.Help,
	}
}

func (k *LeaderboardKeyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Back,
			k.Help,
		},
	}
}
