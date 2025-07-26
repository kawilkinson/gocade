package tetris

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/games/tetris/tetrisconfig"
	"github.com/kawilkinson/gocade/games/tetris/tetrisscreens"
	"github.com/kawilkinson/gocade/games/tetris/tutils"
	"github.com/kawilkinson/gocade/internal/utils"
)

type Input struct {
	mode     tetrisconfig.Mode
	switchIn tetrisconfig.SwitchModeInput
	cfg      *tetrisconfig.Config
}

func NewInput(mode tetrisconfig.Mode, switchIn tetrisconfig.SwitchModeInput, cfg *tetrisconfig.Config) *Input {
	return &Input{
		mode:     mode,
		switchIn: switchIn,
		cfg:      cfg,
	}
}

var _ tea.Model = &Model{}

type Model struct {
	child        tea.Model
	cfg          *tetrisconfig.Config
	forceQuitKey key.Binding

	width  int
	height int

	ExitError error
}

func NewModel(in *Input) (*Model, error) {
	m := &Model{
		cfg:          in.cfg,
		forceQuitKey: key.NewBinding(key.WithKeys("ctrl+c")),
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

	case tea.KeyMsg:
		if key.Matches(msg, m.forceQuitKey) {
			return m, tea.Quit
		}

	case tetrisconfig.SwitchModeMsg:
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

func (m *Model) setChild(mode tetrisconfig.Mode, switchIn tetrisconfig.SwitchModeInput) error {
	if rv := reflect.ValueOf(switchIn); !rv.IsValid() || rv.IsNil() {
		return errors.New("switchIn is not valid")
	}

	switch mode {
	case tetrisconfig.ModeMenu:
		menuIn, ok := switchIn.(*tetrisconfig.MenuInput)
		if !ok {
			return fmt.Errorf("switchIn is not a MenuInput: %w", utils.ErrInvalidTypeAssertion)
		}
		m.child = tetrisscreens.NewMenuModel(menuIn)

	case tetrisconfig.ModeMarathon, tetrisconfig.ModeSprint, tetrisconfig.ModeUltra:
		singleIn, ok := switchIn.(*tetrisconfig.SingleInput)
		if !ok {
			return fmt.Errorf("switchIn is not a SingleInput: %w", utils.ErrInvalidTypeAssertion)
		}
		child, err := tetrisscreens.NewSingleModel(singleIn, m.cfg)
		if err != nil {
			return fmt.Errorf("creating single model: %w", err)
		}
		m.child = child

	default:
		return errors.New("invalid Mode")
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
