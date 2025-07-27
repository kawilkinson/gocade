package snake

import (
	"math/rand/v2"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/games/snake/sutils"
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

	Width  int
	Height int
}

type TickMsg time.Time

func (m *SnakeGameModel) Tick() tea.Cmd {
	return tea.Tick(time.Second/10, func(time time.Time) tea.Msg {
		return TickMsg(time)
	})
}

func (m *SnakeGameModel) ChangeSnakeDirection(direction int) (tea.Model, tea.Cmd) {
	if m.Snake.HitWall(m) {
		m.GameOver = true

		return m, tea.Quit
	}

	oppDir := map[int]int{
		sutils.Up:    sutils.Down,
		sutils.Down:  sutils.Up,
		sutils.Left:  sutils.Right,
		sutils.Right: sutils.Left,
	}

	if oppDir[direction] != m.Snake.Direction {
		m.Snake.Direction = direction
	}

	return m, nil
}

func (m *SnakeGameModel) MoveSnake() (tea.Model, tea.Cmd) {
	head := m.Snake.GetSnakeHead()
	coord := Coordinate{x: head.x, y: head.y}

	switch m.Snake.Direction {
	case sutils.Up:
		coord.x--

	case sutils.Down:
		coord.x++

	case sutils.Left:
		coord.y--

	case sutils.Right:
		coord.y++
	}

	if coord.x == m.Food.x && coord.y == m.Food.y {
		m.Snake.Length++
		x := rand.IntN(m.Height-1) + 1
		y := rand.IntN(m.Width-1) + 1

		for {
			if !m.Snake.HitSelf(Coordinate{x, y}) {
				break
			}
		}

		m.Food.y = y
		m.Food.x = x
	}

	if m.Snake.HitWall(m) || m.Snake.HitSelf(coord) {
		m.GameOver = true

		return m, tea.Quit
	}

	if len(m.Snake.Body) < m.Snake.Length {
		m.Snake.Body = append(m.Snake.Body, coord)
		m.Score = m.Score + 1
	} else {
		m.Snake.Body = append(m.Snake.Body[1:], coord)
	}

	return m, m.Tick()
}

func CreateSnakeGameModel() *SnakeGameModel {
	return &SnakeGameModel{
		HorizontalLine: "#",
		VerticalLine:   "#",
		EmptySymbol:    " ",
		SnakeSymbol:    "o",
		FoodSymbol:     "$",
		Stage:          [][]string{},
		GameOver:       false,
		Score:          0,
		Food:           Food{x: 10, y: 20},
		Snake: Snake{
			Body: []Coordinate{
				{x: 1, y: 1},
				{x: 1, y: 2},
				{x: 1, y: 3},
				{x: 1, y: 4}},
			Length:    4,
			Direction: sutils.Right,
		},

		Width:  40,
		Height: 20,
	}
}

func (m *SnakeGameModel) Init() tea.Cmd {
	var x, y int

	x = rand.IntN(m.Height - 1)
	y = rand.IntN(m.Width - 1)

	m.Food.x = x
	m.Food.y = y

	return m.Tick()
}

func (m *SnakeGameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "w":
			return m.ChangeSnakeDirection(sutils.Up)

		case "down", "s":
			return m.ChangeSnakeDirection(sutils.Down)

		case "left", "a":
			return m.ChangeSnakeDirection(sutils.Left)

		case "right", "d":
			return m.ChangeSnakeDirection(sutils.Right)
		}

	case TickMsg:
		return m.MoveSnake()
	}

	return m, nil
}

// main goal with the view is to keep things super simple,
// so I'm using a function that draws all the strings of the entire game here
func (m *SnakeGameModel) View() string {
	return m.RenderGame()
}
