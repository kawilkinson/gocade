package snake

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type SnakeGameModel struct {
	HorizontalLine string
	VerticalLine   string
	EmptySymbol    string
	SnakeSymbol    string
	FoodSymbol     string
	Stage          [][]string
	Snake          Snake
	GameOver       bool
	Score          int
	Food           Food

	Width          int
	Height         int
}

type TickMsg time.Time

func (m *SnakeGameModel) ticket() tea.Cmd {
	return tea.Tick(time.Second / 10, func(time time.Time) tea.Msg {
		return TickMsg(time)
	})
}
