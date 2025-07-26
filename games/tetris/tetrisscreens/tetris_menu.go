package tetrisscreens

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/games/tetris/tutils"
)

type TetrisMenuModel struct {
	IsCompleted bool
	form        *huh.Form
	keys        *menuKeyMap
	formData    *TetrisMenuFormData
	width       int
	height      int
}

type TetrisMenuFormData struct {
	Username   string
	TetrisGame tutils.Screen
}

type TetrisMenuInput struct{}

func NewMenuInput() *TetrisMenuInput {
	return &TetrisMenuInput{}
}

func (input *TetrisMenuInput) isChangeScreenInput() {}

type ChangeScreenMsg struct {
	Target tutils.Screen
	Input  ChangeScreenInput
}

type ChangeScreenInput interface {
	isChangeScreenInput()
}

type SingleInput struct {
	Username string
	Screen   tutils.Screen
}

func NewSingleInput(username string, screen tutils.Screen) *SingleInput {
	return &SingleInput{
		Username: username,
		Screen: screen,
	}
}

func (input *SingleInput) isChangeScreenInput() {}

func ChangeScreen(target tutils.Screen, input ChangeScreenInput) tea.Cmd {
	return func() tea.Msg {
		return ChangeScreenMsg{
			Target: target,
			Input:  input,
		}
	}
}

func CreateTetrisMenuModel(_ *TetrisMenuInput) *TetrisMenuModel {
	formData := new(TetrisMenuFormData)
	keys := SetTetrisMenuKeys()

	return &TetrisMenuModel{
		formData: formData,
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Value(&formData.Username).
					Title("Username:").CharLimit(50).
					Validate(func(str string) error {
						if len(str) == 0 {
							return errors.New("a username must be entered to continue")
						}

						return nil
					}),
				huh.NewSelect[tutils.Screen]().Value(&formData.TetrisGame).
					Title("Start Tetris when ready").
					Options(
						huh.NewOption("Play Tetris", tutils.ScreenTetrisGame),
					),
			),
		).WithKeyMap(keys.formKeys),
		keys: keys,
	}
}

func (m *TetrisMenuModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m *TetrisMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		formWidth := msg.Width / 2
		formWidth = min(formWidth, lipgloss.Width(tutils.RenderLargeText(tutils.TetrisTitle)))
		m.form = m.form.WithWidth(formWidth)
		return m, nil

	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Exit) {
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted && !m.IsCompleted {
		cmds = append(cmds, m.performCompletion())
	}

	return m, tea.Batch(cmds...)
}

func (m *TetrisMenuModel) performCompletion() tea.Cmd {
	m.IsCompleted = true

	switch m.formData.TetrisGame {
	case tutils.ScreenTetrisGame:
		input := NewSingleInput(m.formData.Username, m.formData.TetrisGame)
		return ChangeScreen(m.formData.TetrisGame, input)

	case tutils.ScreenTetrisMenu:
		fallthrough

	default:
		return tutils.ErrorCmd(fmt.Errorf("tetris game not selected, unable to start game: %v", m.formData.TetrisGame))
	}
}

func (m *TetrisMenuModel) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		tutils.RenderLargeText(tutils.TetrisTitle)+"\n",
		m.form.View(),
	)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
