package leaderboard

import (
	"encoding/csv"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/menuconfig"
	"github.com/kawilkinson/gocade/internal/screens"
	"github.com/kawilkinson/gocade/internal/utils"
)

type LeaderboardModel struct {
	keys *menuconfig.LeaderboardKeyBindings

	help  help.Model
	table table.Model

	width  int
	height int

	err error
}

func NewLeaderBoardMenu(filename string) *LeaderboardModel {
	data, err := LoadScores(filename)
    if err != nil {
        data = [][]string{
            {"No score", "-", "-", "-", "-"}, 
        }
    } else if len(data) == 0 {
        data = [][]string{
            {"No score", "-", "-", "-", "-"}, 
        }
    }

	return &LeaderboardModel{
		keys:  menuconfig.DefaultLeaderboardKeyBindings(),
		help:  help.New(),
		table: createLeaderboardTable(data),
		err:   err,
	}
}

func (m *LeaderboardModel) Init() tea.Cmd {
	return nil
}

func (m *LeaderboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Back):
			return m, screens.SwitchScreens(utils.ScreenScoreMenu)

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m *LeaderboardModel) View() string {
	output := m.table.View() + "\n" + m.help.View(m.keys)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, output)
}

func LoadScores(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	scores, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func createLeaderboardTable(scores [][]string) table.Model {
	cols := []table.Column{
		{Title: "Name", Width: 10},
		{Title: "Score", Width: 10},
		{Title: "Lines", Width: 10},
		{Title: "Level", Width: 5},
		{Title: "Mode", Width: 10},
	}

	rows := make([]table.Row, len(scores))

	for i, s := range scores {
		name := s[0]
		score := s[1]
		lines := s[2]
		level := s[3]
		mode := s[4]

		rows[i] = table.Row{
			name,
			score,
			lines,
			level,
			mode,
		}
	}

	style := table.DefaultStyles()
	style.Header = style.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(utils.GopherColor)).
		BorderBottom(true).
		Bold(false)

	style.Selected = style.Selected.
		Foreground(lipgloss.Color(utils.DarkerGopherColor)).
		Bold(false)

	t := table.New(
		table.WithColumns(cols),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithStyles(style),
	)

	return t
}
