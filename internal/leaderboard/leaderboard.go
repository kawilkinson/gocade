package leaderboard

import (
	"encoding/csv"
	"os"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type LeaderboardModel struct {
	help  help.Model
	table *table.Table

	width  int
	height int

	err error
}

func NewLeaderBoardMenu(filename string) *LeaderboardModel {
	data, err := LoadScores(filename)
	if err != nil {
		data = [][]string{{"No scores found"}}
	}

	t := table.New().
		Headers("Name", "Score", "Lines", "Level", "Mode").
		Border(lipgloss.NormalBorder()).
		BorderRow(true).
		Rows(data...).
		Width(60)

	return &LeaderboardModel{
		help:  help.New(),
		table: t,
		err:   err,
	}
}

func (m *LeaderboardModel) Init() tea.Cmd {
	return nil
}

func (m *LeaderboardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table = m.table.Width(m.width - 4)
	}

	return m, nil
}

func (m *LeaderboardModel) View() string {
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.table.Render())
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
