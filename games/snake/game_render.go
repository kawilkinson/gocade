package snake

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/games/snake/sutils"
	"github.com/kawilkinson/gocade/internal/utils"
)

func RenderStage(m *SnakeGameModel) {
    m.Stage = make([][]string, 0, m.Height)

    // Top border
    m.Stage = append(m.Stage,
        strings.Split(m.VerticalLine+strings.Repeat(m.HorizontalLine, m.Width-2)+m.VerticalLine, ""))

    // Middle rows
    for i := 0; i < m.Height-2; i++ {
        m.Stage = append(m.Stage,
            strings.Split(m.VerticalLine+strings.Repeat(m.EmptySymbol, m.Width-2)+m.VerticalLine, ""))
    }

    // Bottom border
    m.Stage = append(m.Stage,
        strings.Split(m.VerticalLine+strings.Repeat(m.HorizontalLine, m.Width-2)+m.VerticalLine, ""))
}

func RenderSnake(m *SnakeGameModel) {
	for _, b := range m.Snake.Body {
		m.Stage[b.x][b.y] = m.SnakeSymbol
	}
}

func RenderFood(m *SnakeGameModel) {
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

	return gameOverStyle.Render(sutils.GameOverMessage)
}

func (m *SnakeGameModel) RenderGame() string {
	var strBuilder strings.Builder
	var strStage strings.Builder

	strBuilder.WriteString(RenderTitle())
	strBuilder.WriteByte('\n')

	RenderStage(m)
	RenderSnake(m)
	RenderFood(m)

	for _, row := range m.Stage {
		strStage.WriteString(strings.Join(row, "") + "\n")
	}

	strBuilder.WriteString(strStage.String())
	strBuilder.WriteByte('\n')

	strBuilder.WriteString(RenderScore(m.Score))
	strBuilder.WriteByte('\n')

	if m.GameOver {
		strBuilder.WriteString(RenderGameOver())
	}

	strBuilder.WriteString(RenderHelp(sutils.HelpMessage))
	strBuilder.WriteByte('\n')
	strBuilder.WriteByte('\n')

	return strBuilder.String()
}
