package snake

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/utils"
)

func RenderStage(m *Model) {
	m.Stage = append(m.Stage,
		strings.Split(m.VerticalLine+strings.Repeat(m.HorizontalLine, m.Width-2)+m.VerticalLine,
			""))

	for i := 0; i < m.Height-1; i++ {
		m.Stage = append(m.Stage,
			strings.Split(m.VerticalLine+strings.Repeat(m.EmptySymbol, m.Width-2)+m.VerticalLine,
				""))
	}

	m.Stage = append(m.Stage,
		strings.Split(m.VerticalLine+strings.Repeat(m.HorizontalLine, m.Width-2)+m.VerticalLine,
			""))
}

func RenderSnake(m *Model) {
	for _, b := range m.Snake.body {
		m.Stage[b.x][b.y] = m.SnakeSymbol
	}
}

func RenderFood(m *Model) {
	m.Stage[m.Food.x][m.Food.y] = m.FoodSymbol
}

func RenderTitle() string {
	titleStyle := lipgloss.NewStyle().Bold(true).
		Foreground(lipgloss.Color(utils.GopherColor)).
		Width(40).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	return titleStyle.Render("Snake")
}

func RenderScore(score int) string {
	scoreStr := fmt.Sprintf("Score: %d", score)

	scoreStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFFFFF"))

	return scoreStyle.Render(scoreStr)
}

func RenderHelp(help string) string {
	helpStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#D5B60A"))

	return helpStyle.Render(help)
}

func RenderGameOver() string {
	gameOverStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF0000")).
		Width(40).
		AlignHorizontal(lipgloss.Center).
		MarginTop(1).
		MarginBottom(1)

	return gameOverStyle.Render("Game Over!")
}
