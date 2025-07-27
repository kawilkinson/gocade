package menuconfig

import "github.com/charmbracelet/bubbles/key"

type LeaderboardKeyBindings struct {
	Help key.Binding
	Back key.Binding
	Up   key.Binding
	Down key.Binding
}

func DefaultLeaderboardKeyBindings() *LeaderboardKeyBindings {
	return &LeaderboardKeyBindings{
		Help: key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
		Back: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "go back")),
		Up:   key.NewBinding(key.WithKeys("up"), key.WithHelp("up arrow", "move up")),
		Down: key.NewBinding(key.WithKeys("down"), key.WithHelp("down arrow", "move down")),
	}
}

func (k *LeaderboardKeyBindings) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Back,
		k.Up,
		k.Down,
		k.Help,
	}
}

func (k *LeaderboardKeyBindings) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.Back,
			k.Help,
		},
		{
			k.Up,
			k.Down,
		},
	}
}
