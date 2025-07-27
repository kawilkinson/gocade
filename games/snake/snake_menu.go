package snake

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/games/snake/snakeconfig"
	"github.com/kawilkinson/gocade/games/snake/sutils"
)

var _ tea.Model = &MenuModel{}

type MenuModel struct {
	form                   *huh.Form
	hasAnnouncedCompletion bool
	keys                   *snakeconfig.MenuKeyMap
	formData               *MenuFormData

	width  int
	height int
}

type MenuFormData struct {
	Username string
	Screen   sutils.Screen
}

func NewMenuModel(_ *snakeconfig.MenuInput) *MenuModel {
	formData := new(MenuFormData)
	keys := snakeconfig.SetSnakeMenuKeys()

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
				huh.NewSelect[sutils.Screen]().Value(&formData.Screen).
					Title("Hit Enter to Start").
					Options(
						huh.NewOption("Start Game", sutils.SnakeGame),
					),
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
		formWidth = min(formWidth, lipgloss.Width(sutils.RenderLargeText(sutils.SnakeTitle)))
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

	switch m.formData.Screen {
	case sutils.SnakeGame:
		return func() tea.Msg {
			return snakeconfig.SwitchToGameMsg{
				Username: m.formData.Username,
				Screen:   m.formData.Screen,
			}
		}

	default:
		return sutils.ErrorCmd(fmt.Errorf("invalid option for starting game %q", m.formData.Screen))
	}
}

func (m *MenuModel) View() string {
	title := sutils.RenderLargeText(sutils.SnakeTitle)
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
