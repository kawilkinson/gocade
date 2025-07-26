package tetrisscreens

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/games/tetris/tetrisconfig"
	"github.com/kawilkinson/gocade/games/tetris/tetrisdata"
	"github.com/kawilkinson/gocade/games/tetris/tetrislogic"
	"github.com/kawilkinson/gocade/games/tetris/tetrislogic/modes/single"
	"github.com/kawilkinson/gocade/games/tetris/tutils"
	"github.com/kawilkinson/gocade/internal/utils"
)

var _ tea.Model = &SingleModel{}

type SingleModel struct {
	username        string
	game            *single.Game
	nextQueueLength int
	fallStopwatch   tetrisconfig.Stopwatch
	screen          tutils.Screen

	gameStopwatch tetrisconfig.Stopwatch

	styles   tetrisconfig.GameStyles
	help     help.Model
	keys     *tetrisconfig.TetrisKeys
	isPaused bool
	rand     *rand.Rand

	width  int
	height int
}

func NewSingleModel(input *SingleInput, cfg *tetrisconfig.Config, opts ...func(*SingleModel)) (*SingleModel, error) {
	m := &SingleModel{
		username:        input.Username,
		styles:          *tetrisconfig.CreateGameStyles(cfg.Theme),
		help:            help.New(),
		keys:            tetrisconfig.SetTetrisKeyBindings(),
		isPaused:        false,
		nextQueueLength: cfg.NextQueueLength,
		screen:          input.Screen,
		rand:            rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())),
	}

	for _, opt := range opts {
		opt(m)
	}

	var gameInput *single.Input
	switch input.Screen {
	case tutils.ScreenTetrisGame:
		gameInput = &single.Input{
			Level:         1,
			MaxLevel:      15,
			IncreaseLevel: true,
			EndOnMaxLevel: cfg.EndOnMaxLevel,

			MaxLines:      40,
			EndOnMaxLines: false,

			GhostEnabled: false,
		}
		m.gameStopwatch = tetrisconfig.NewStopwatchWithInterval(tutils.TimerUpdateInterval)

	case tutils.ScreenTetrisMenu:
		fallthrough

	default:
		return nil, fmt.Errorf("invalid single player game mode %v", input.Screen)
	}

	var err error
	m.game, err = single.NewGame(gameInput)
	if err != nil {
		return nil, fmt.Errorf("unable to create single player game: %v", err)
	}

	m.fallStopwatch = tetrisconfig.NewStopwatchWithInterval(m.game.GetDefaultFallInterval())

	return m, nil
}

func RandSource(r *rand.Rand) func(*SingleModel) {
	return func(m *SingleModel) {
		m.rand = r
	}
}

func (m *SingleModel) Init() tea.Cmd {
	cmd := m.gameStopwatch.Init()

	return tea.Batch(m.fallStopwatch.Init(), cmd)
}

func (m *SingleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m, cmd = m.dependenciesUpdate(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, tea.Batch(cmds...)
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, tea.Batch(cmds...)
	}

	if m.game.IsGameOver() {
		m, cmd = m.gameOverUpdate(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	if m.isPaused {
		m, cmd = m.pausedUpdate(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	m, cmd = m.playingUpdate(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *SingleModel) dependenciesUpdate(msg tea.Msg) (*SingleModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	var err error

	cmd, err = utils.UpdateTypedModel(&m.gameStopwatch, msg)
	if err != nil {
		cmds = append(cmds, tutils.ErrorCmd(err))
	}
	cmds = append(cmds, cmd)

	cmd, err = utils.UpdateTypedModel(&m.fallStopwatch, msg)
	if err != nil {
		cmds = append(cmds, tutils.ErrorCmd(err))
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *SingleModel) gameOverUpdate(msg tea.Msg) (*SingleModel, tea.Cmd) {
	var cmds []tea.Cmd
	if msg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(msg, m.keys.Exit, m.keys.Hold) {
			newEntry := tetrisdata.Score{
				Name:  m.username,
				Score: m.game.GetTotalScore(),
				Lines: m.game.GetLinesCleared(),
				Level: m.game.GetLevel(),
			}

			err := tetrisdata.SaveScore(newEntry)
			if err != nil {
				cmds = append(cmds, tutils.ErrorCmd(err))
			}

			cmds = append(cmds, ChangeScreen(tutils.ScreenTetrisMenu, NewMenuInput()))
			return m, tea.Batch(cmds...)
		}
	}

	return m, nil
}

func (m *SingleModel) pausedUpdate(msg tea.Msg) (*SingleModel, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch {
		case key.Matches(msg, m.keys.Exit):
			return m, m.togglePause()

		case key.Matches(msg, m.keys.Hold):
			return m, ChangeScreen(tutils.ScreenTetrisMenu, NewMenuInput())
		}
	}

	return m, nil
}

func (m *SingleModel) playingUpdate(msg tea.Msg) (*SingleModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.playingKeyMsgUpdate(msg)

	case stopwatch.TickMsg:
		if msg.ID != m.fallStopwatch.ID() {
			break
		}

		return m, m.fallStopwatchTick()
	}

	return m, nil
}

func (m *SingleModel) playingKeyMsgUpdate(msg tea.KeyMsg) (*SingleModel, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Left):
		m.game.MoveLeft()
		return m, nil

	case key.Matches(msg, m.keys.Right):
		m.game.MoveRight()
		return m, nil

	case key.Matches(msg, m.keys.RotateClockwise):
		err := m.game.Rotate(true)
		if err != nil {
			return nil, tutils.ErrorCmd(fmt.Errorf("unable to rotate clockwise: %v", err))
		}

		return m, nil

	case key.Matches(msg, m.keys.RotateCounterClockwise):
		err := m.game.Rotate(false)
		if err != nil {
			return nil, tutils.ErrorCmd(fmt.Errorf("unable to rotate counterclockwise: %v", err))
		}

		return m, nil

	case key.Matches(msg, m.keys.HardDrop):
		gameOver, err := m.game.HardDrop()
		if err != nil {
			return nil, tutils.ErrorCmd(fmt.Errorf("unable to hard drop: %v", err))
		}

		var cmds []tea.Cmd
		if gameOver {
			cmds = append(cmds, m.triggerGameOver())
		}

		cmds = append(cmds, m.fallStopwatch.Reset())

		return m, tea.Batch(cmds...)

	case key.Matches(msg, m.keys.SoftDrop):
		m.game.ToggleSoftDrop()

		return m, m.fallStopwatchTick()

	case key.Matches(msg, m.keys.Hold):
		gameOver, err := m.game.Hold()
		if err != nil {
			return nil, tutils.ErrorCmd(fmt.Errorf("unable to hold tetrimino: %v", err))
		}

		var cmds []tea.Cmd
		if gameOver {
			cmds = append(cmds, m.triggerGameOver())
		}

		return m, tea.Batch(cmds...)

	case key.Matches(msg, m.keys.Exit):
		return m, m.togglePause()
	}

	return m, nil
}

func (m *SingleModel) fallStopwatchTick() tea.Cmd {
	gameOver, err := m.game.TickLower()
	if err != nil {
		return tutils.ErrorCmd(fmt.Errorf("unable to lower tetrimino (tick): %v", err))
	}

	if gameOver {
		return m.triggerGameOver()
	}

	m.fallStopwatch.SetInterval(m.game.GetFallInterval())

	return nil
}

func (m *SingleModel) View() string {
	matrixView, err := m.matrixView()
	if err != nil {
		return "** UNABLE TO BUILD MATRIX VIEW **"
	}

	var output = lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Right, m.holdView(), m.informationView()),
		matrixView,
		m.bagView(),
	)

	if m.game.IsGameOver() {
		output, err = utils.OverlayCenter(output, tutils.GameOverMessage, true)
		if err != nil {
			return "** FAILED TO OVERLAY GAME OVER MESSAGE **"
		}
	} else if m.isPaused {
		output, err = utils.OverlayCenter(output, tutils.PausedMessage, true)
		if err != nil {
			return "** FAILED TO OVERLAY PAUSED MESSAGE **"
		}
	}

	output = lipgloss.JoinVertical(lipgloss.Left, output, m.help.View(m.keys))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, output)
}

func (m *SingleModel) matrixView() (string, error) {
	matrix, err := m.game.GetVisibleMatrix()
	if err != nil {
		return "", fmt.Errorf("unable to get visible matrix: %v", err)
	}

	var output string
	for row := range matrix {
		for col := range matrix[row] {
			output += m.renderCell(matrix[row][col])
		}

		if row < len(matrix)-1 {
			output += "\n"
		}
	}

	var rowIndicator string
	for i := 1; i <= 20; i++ {
		rowIndicator += fmt.Sprintf("%d\n", i)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.styles.Playfield.Render(output),
		m.styles.RowIndicator.Render(rowIndicator),
	), nil
}

func (m *SingleModel) informationView() string {
	width := m.styles.Information.GetWidth()

	var header string
	headerStyle := lipgloss.NewStyle().
		Width(width).
		AlignHorizontal(lipgloss.Center).
		Bold(true).
		Underline(true)

	switch {
	case m.game.IsGameOver():
		header = headerStyle.Render("GAME OVER")

	case m.isPaused:
		header = headerStyle.Render("PAUSED")

	default:
		header = headerStyle.Render("Standard")
	}

	toFixedWidth := func(title, value string) string {
		return fmt.Sprintf("%s%*s\n", title, width-(1+len(title)), value)
	}

	gameTime := m.gameStopwatch.Elapsed().Seconds()

	minutes := int(gameTime) / 60

	var timeStr string
	if minutes > 0 {
		seconds := int(gameTime) % 60
		timeStr += fmt.Sprintf("%02d:%02d", minutes, seconds)
	} else {
		timeStr += fmt.Sprintf("%06.3f", gameTime)
	}

	var output string
	output += fmt.Sprintln("Score:")
	output += fmt.Sprintf("%*d\n", width-1, m.game.GetTotalScore())
	output += fmt.Sprintln("Time:")
	output += fmt.Sprintf("%*s\n", width-1, timeStr)
	output += toFixedWidth("Lines:", strconv.Itoa(m.game.GetLinesCleared()))
	output += toFixedWidth("Level:", strconv.Itoa(m.game.GetLevel()))

	return m.styles.Information.Render(lipgloss.JoinVertical(lipgloss.Left, header, output))
}

func (m *SingleModel) holdView() string {
	label := m.styles.Hold.Label.Render("Hold:")
	item := m.styles.Hold.Item.Render(m.renderTetrimino(m.game.GetHoldTetrimino(), 1))
	output := lipgloss.JoinVertical(lipgloss.Top, label, item)

	return m.styles.Hold.View.Render(output)
}

func (m *SingleModel) bagView() string {
	output := "Next:\n"
	for i, t := range m.game.GetBagTetriminos() {
		for i >= m.nextQueueLength {
			break
		}

		output += "\n" + m.renderTetrimino(&t, 1)
	}

	return m.styles.Bag.Render(output)
}

func (m *SingleModel) renderTetrimino(t *tetrislogic.Tetrimino, background byte) string {
	var output string
	for row := range t.Cells {
		for col := range t.Cells[row] {
			if t.Cells[row][col] {
				output += m.renderCell(t.Value)
			} else {
				output += m.renderCell(background)
			}
		}

		output += "\n"
	}

	return output
}

func (m *SingleModel) renderCell(cell byte) string {
	switch cell {
	case 0:
		return m.styles.EmptyCell.Render(m.styles.CellChar.Empty)

	case 1:
		return "  "

	case 'G':
		return m.styles.GhostCell.Render(m.styles.CellChar.Ghost)

	default:
		cellStyle, ok := m.styles.TetriminoCellStyles[cell]
		if ok {
			return cellStyle.Render(m.styles.CellChar.Tetriminos)
		}
	}

	return "??"
}

func (m *SingleModel) triggerGameOver() tea.Cmd {
	m.game.EndGame()
	m.isPaused = false

	var cmds []tea.Cmd
	cmds = append(cmds, m.fallStopwatch.Stop())

	return tea.Batch(cmds...)
}

func (m *SingleModel) togglePause() tea.Cmd {
	m.isPaused = !m.isPaused

	cmd := m.gameStopwatch.Toggle()

	return tea.Batch(m.fallStopwatch.Stop(), cmd)
}
