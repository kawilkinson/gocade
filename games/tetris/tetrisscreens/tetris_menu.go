package tetrisscreens

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/games/tetris/tetrisconfig"
	"github.com/kawilkinson/gocade/games/tetris/tutils"
	"github.com/kawilkinson/gocade/internal/utils"
)

var _ tea.Model = &MenuModel{}

type MenuModel struct {
	form                   *huh.Form
	hasAnnouncedCompletion bool
	keys                   *tetrisconfig.MenuKeyMap
	formData               *MenuFormData

	width  int
	height int
}

type MenuFormData struct {
	Username string
	GameMode tetrisconfig.Mode
	Level    int
}

func NewMenuModel(_ *tetrisconfig.MenuInput) *MenuModel {
	formData := new(MenuFormData)
	keys := tetrisconfig.SetTetrisMenuKeys()

	return &MenuModel{
		formData: formData,
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewInput().Value(&formData.Username).
					Title("Username:").CharLimit(100).
					Validate(func(s string) error {
						if len(s) == 0 {
							return errors.New("empty username not allowed")
						}
						return nil
					}),
				huh.NewSelect[tetrisconfig.Mode]().Value(&formData.GameMode).
					Title("Game Mode:").
					Options(
						huh.NewOption("Marathon", tetrisconfig.ModeMarathon),
						huh.NewOption("Sprint (40 Lines)", tetrisconfig.ModeSprint),
						huh.NewOption("Ultra (Time Trial)", tetrisconfig.ModeUltra),
					),
				huh.NewSelect[int]().Value(&formData.Level).
					Title("Starting Level:").
					Options(utils.HuhIntRangeOptions(1, 15)...),
			),
		).WithKeyMap(keys.FormKeys),
		keys: keys,
	}
}

func (m *MenuModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Exit) {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		formWidth := msg.Width / 2
		formWidth = min(formWidth, lipgloss.Width(tutils.RenderLargeText(tutils.TetrisTitle)))
		m.form = m.form.WithWidth(formWidth)
		return m, nil
	}

	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted && !m.hasAnnouncedCompletion {
		cmds = append(cmds, m.announceCompletion())
	}

	return m, tea.Batch(cmds...)
}

func (m *MenuModel) announceCompletion() tea.Cmd {
	m.hasAnnouncedCompletion = true

	switch m.formData.GameMode {
	case tetrisconfig.ModeMarathon, tetrisconfig.ModeSprint, tetrisconfig.ModeUltra:
		in := tetrisconfig.NewSingleInput(m.formData.GameMode, m.formData.Level, m.formData.Username)
		return tetrisconfig.SwitchModeCmd(m.formData.GameMode, in)

	case tetrisconfig.ModeMenu, tetrisconfig.ModeLeaderboard:
		fallthrough
	default:
		return tutils.ErrorCmd(fmt.Errorf("invalid mode for starting game %q", m.formData.GameMode))
	}
}

func (m *MenuModel) View() string {
	title := tutils.RenderLargeText(tutils.TetrisTitle)
	form := m.form.View()

	helpText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8")).
		Render("Press 'esc' at any time in this menu to exit the game")

	menuStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center)

	content := lipgloss.JoinVertical(lipgloss.Center, title, form)

	mainContent := menuStyle.Render(content)
	footer := lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render(helpText)

	return lipgloss.JoinVertical(lipgloss.Left, mainContent, footer)
}
