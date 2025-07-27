package tetrisconfig

import "github.com/charmbracelet/bubbles/key"

type TetrisKeys struct {
	HardDrop               key.Binding
	SoftDrop               key.Binding
	Left                   key.Binding
	Right                  key.Binding
	RotateClockwise        key.Binding
	RotateCounterClockwise key.Binding
	Exit                   key.Binding
	HardExit               key.Binding
	Hold                   key.Binding
	Help                   key.Binding
}

func (k TetrisKeys) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Exit,
		k.Help,
	}
}

func (k TetrisKeys) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.HardDrop, k.SoftDrop, k.Left, k.Right},
		{k.RotateClockwise, k.RotateCounterClockwise},
		{k.Exit, k.Hold},
	}
}

func SetTetrisKeyBindings() *TetrisKeys {
	keys := &TetrisKeys{
		HardDrop:               key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "hard drop")),
		SoftDrop:               key.NewBinding(key.WithKeys("s"), key.WithHelp("s", "toggle soft drop")),
		Left:                   key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "move left")),
		Right:                  key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "move right")),
		RotateClockwise:        key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "rotate tetrimino clockwise")),
		RotateCounterClockwise: key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "rotate tetrimino counterclockwise")),
		Exit:                   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "pause")),
		HardExit:               key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl+c", "force exit")),
		Hold:                   key.NewBinding(key.WithKeys(" ", "enter"), key.WithHelp("space/enter", "hold a tetrimino")),
		Help:                   key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "help")),
	}

	return keys
}
