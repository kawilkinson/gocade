package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/screens"
	"github.com/kawilkinson/gocade/internal/utils"
)

type MainMenuModels struct {
	Screen        utils.Screen
	MainMenu      list.Model
	GameMenu      list.Model
	ScoreMenu     list.Model
	LoadingBar    progress.Model
	loadingValue  float64
	pulseDotCount int
	width         int
	height        int
	SelectedGame  string
	quitting      bool
}

func CreateModels() *MainMenuModels {
	progressBar := progress.New(progress.WithGradient("#00ADD8", "#0082A8"), progress.WithWidth(40))
	keys := utils.AdditionalMainMenuKeys() // only additional since most of the defaults Bubble Tea has are good enough

	return &MainMenuModels{
		Screen:     utils.ScreenLoading,
		MainMenu:   screens.NewMainMenu(utils.MenuWidth, utils.MenuHeight),
		GameMenu:   screens.NewGameMenu(utils.MenuWidth, utils.MenuHeight, keys),
		ScoreMenu:  screens.NewScoreMenu(utils.MenuWidth, utils.MenuHeight, keys),
		LoadingBar: progressBar,
	}
}

func (m *MainMenuModels) Init() tea.Cmd {
	return tea.Batch(screens.Tick(), screens.PulseDots())
}

func (m *MainMenuModels) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c": // for quitting the program at any point
			m.quitting = true
			return m, screens.ExitGocade()

		case "b": // universal key for moving back to the main menu
			if m.Screen == utils.ScreenGameMenu || m.Screen == utils.ScreenScoreMenu {
				m.Screen = utils.ScreenMainMenu
				return m, nil
			}

		case "enter":
			switch m.Screen {

			case utils.ScreenMainMenu:
				choice := m.MainMenu.SelectedItem().(screens.MenuItem)

				switch choice {

				case "Play Game":
					m.Screen = utils.ScreenGameMenu

				case "High Scores":
					m.Screen = utils.ScreenScoreMenu

				case "Quit":
					m.quitting = true
					return m, screens.ExitGocade()
				}

			case utils.ScreenGameMenu:
				selected := m.GameMenu.SelectedItem().(screens.MenuItem)
				m.SelectedGame = string(selected)
				// todo: quit for now, but will change into starting the game and letting the arcade run in the background
				return m, tea.Quit

			case utils.ScreenScoreMenu:
				// todo: handle score menu
				return m, tea.Quit
			}
		}

	case screens.TickMsg:
		if m.Screen == utils.ScreenLoading {
			m.loadingValue += 0.001
			if m.loadingValue >= 1.0 {
				m.Screen = utils.ScreenMainMenu
			}

			return m, screens.Tick()
		}

	case screens.PulseMsg:
		if m.quitting || m.Screen == utils.ScreenLoading {
			m.pulseDotCount = (m.pulseDotCount + 1) % 4 // go in 0, 1, 2, 3 pattern for dots
			return m, screens.PulseDots()
		}

	case screens.ExitMsg:
		return m, tea.Quit
	}

	// screen updates for the main menus
	var cmd tea.Cmd
	switch m.Screen {

	case utils.ScreenLoading:
		updated, cmd := m.LoadingBar.Update(msg)
		m.LoadingBar = updated.(progress.Model)
		return m, cmd

	case utils.ScreenMainMenu:
		m.MainMenu, cmd = m.MainMenu.Update(msg)
		return m, cmd

	case utils.ScreenGameMenu:
		m.GameMenu, cmd = m.GameMenu.Update(msg)
		return m, cmd

	case utils.ScreenScoreMenu:
		m.ScoreMenu, cmd = m.ScoreMenu.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *MainMenuModels) View() string {
	if m.quitting {
		dots := strings.Repeat(".", m.pulseDotCount)
		quitText := screens.QuitTextStyle.Render(fmt.Sprintf("Exiting Gocade%s", dots))
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, quitText)
	}

	switch m.Screen {

	case utils.ScreenLoading:
		dots := strings.Repeat(".", m.pulseDotCount)

		content := lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Starting up Gocade%s\n", dots)),
			m.LoadingBar.ViewAs(m.loadingValue),
		)

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)

	case utils.ScreenMainMenu:
		gopher := screens.RenderGopher(m.width, m.height)
		mainMenu := m.MainMenu.View()

		content := lipgloss.JoinVertical(lipgloss.Center, gopher, mainMenu)

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)

	case utils.ScreenGameMenu:
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.GameMenu.View())

	case utils.ScreenScoreMenu:
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.ScoreMenu.View())

	default:
		return "Unknown screen"
	}
}
