package snake

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/games/snake/snakeconfig"
)

type Model struct {
	current tea.Model
}

func NewModel() Model {
	return Model{
		current: NewMenuModel(nil),
	}
}

func (m Model) Init() tea.Cmd {
	return m.current.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case snakeconfig.SwitchToGameMsg:
		game := CreateSnakeGameModel()
		game.Username = msg.Username

		return Model{current: game}, game.Init()
	}

	currModel, cmd := m.current.Update(msg)
	m.current = currModel

	return m, cmd
}

func (m Model) View() string {
	return m.current.View()
}
