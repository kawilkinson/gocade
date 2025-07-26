package tetris

import (
	"errors"
	"fmt"
	"reflect"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/games/tetris/tetrisconfig"
	"github.com/kawilkinson/gocade/games/tetris/tetrisscreens"
	"github.com/kawilkinson/gocade/games/tetris/tutils"
	"github.com/kawilkinson/gocade/internal/utils"
)

type Input struct {
	mode     tutils.Screen
	switchIn tetrisscreens.ChangeScreenInput
	cfg      *tetrisconfig.Config
}

func NewInput(screen tutils.Screen, switchIn tetrisscreens.ChangeScreenInput, cfg *tetrisconfig.Config) *Input {
	return &Input{
		mode:     screen,
		switchIn: switchIn,
		cfg:      cfg,
	}
}

var _ tea.Model = &Model{}

type Model struct {
	child tea.Model
	cfg   *tetrisconfig.Config

	width  int
	height int

	ExitError error
}

func NewModel(in *Input) (*Model, error) {
	m := &Model{
		cfg: in.cfg,
	}

	err := m.setChild(in.mode, in.switchIn)
	if err != nil {
		return nil, fmt.Errorf("setting child model: %w", err)
	}

	return m, nil
}

func (m *Model) Init() tea.Cmd {
	return m.initChild()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tutils.ErrorMsg:
		m.ExitError = msg
		return m, tea.Quit

	case tetrisscreens.ChangeScreenMsg:
		err := m.setChild(msg.Target, msg.Input)
		if err != nil {
			return m, tutils.ErrorCmd(fmt.Errorf("setting child model: %w", err))
		}
		cmd := m.initChild()

		return m, cmd

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.child, cmd = m.child.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	return m.child.View()
}

func (m *Model) setChild(screen tutils.Screen, changeInput tetrisscreens.ChangeScreenInput) error {
	if rv := reflect.ValueOf(changeInput); !rv.IsValid() || rv.IsNil() {
		return errors.New("switchInput is not valid")
	}

	switch screen {
	case tutils.ScreenTetrisMenu:
		menuInput, ok := changeInput.(*tetrisscreens.TetrisMenuInput)
		if !ok {
			return fmt.Errorf("switchInput is not a TetrisMenuInput: %v", utils.ErrInvalidTypeAssertion)
		}
		m.child = tetrisscreens.CreateTetrisMenuModel(menuInput)

	case tutils.ScreenTetrisGame:
		singleInput, ok := changeInput.(*tetrisscreens.SingleInput)
		if !ok {
			return fmt.Errorf("switchInput is not a SingleInput: %v", utils.ErrInvalidTypeAssertion)
		}
		child, err := tetrisscreens.NewSingleModel(singleInput, m.cfg)
		if err != nil {
			return fmt.Errorf("error when creating single model: %w", err)
		}
		m.child = child

	default:
		return errors.New("invalid screen")
	}
	
	return nil
}

func (m *Model) initChild() tea.Cmd {
	var cmds []tea.Cmd
	cmd := m.child.Init()
	cmds = append(cmds, cmd)
	m.child, cmd = m.child.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
