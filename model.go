package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/kawilkinson/gocade/internal/leaderboard"
	"github.com/kawilkinson/gocade/internal/menuconfig"
	"github.com/kawilkinson/gocade/internal/screens"
	"github.com/kawilkinson/gocade/internal/utils"
)

type MainMenuModels struct {
	Screen        utils.Screen
	Style         *screens.MenuStyles
	Keys          *menuconfig.MainMenuKeys
	MainMenu      list.Model
	GameMenu      list.Model
	ScoreMenu     list.Model
	Leaderboard   *leaderboard.LeaderboardModel
	LoadingBar    progress.Model
	loadingValue  float64
	pulseDotCount int
	width         int
	height        int
	SelectedGame  string
	launchTetris  bool
	quitting      bool
}

func CreateMainMenuModels() *MainMenuModels {
	progressBar := progress.New(progress.WithGradient("#00ADD8", "#0082A8"), progress.WithWidth(40))
	keys := menuconfig.SetExtraMainMenuKeys() // only additional since most of the defaults Bubble Tea has are good enough
	style := screens.CreateMenuStyle()

	return &MainMenuModels{
		Screen:     utils.ScreenLoading,
		Keys:       keys,
		Style:      style,
		MainMenu:   screens.NewMainMenu(utils.MenuWidth, utils.MenuHeight, keys, style),
		GameMenu:   screens.NewGameMenu(utils.MenuWidth, utils.MenuHeight, keys, style),
		ScoreMenu:  screens.NewScoreMenu(utils.MenuWidth, utils.MenuHeight, keys, style),
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
		switch {
		case key.Matches(msg, m.Keys.Exit):
			m.quitting = true
			return m, screens.ExitGocade()

		case key.Matches(msg, m.Keys.BackKey): // universal key for moving back to the main menu
			switch m.Screen {

			case utils.ScreenLeaderboard:
				m.Screen = utils.ScreenScoreMenu
				return m, nil

			case utils.ScreenGameMenu, utils.ScreenScoreMenu:
				m.Screen = utils.ScreenMainMenu
				return m, nil
			}

		case key.Matches(msg, m.Keys.Select):
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
				choice := m.GameMenu.SelectedItem().(screens.MenuItem)
				m.SelectedGame = string(choice)

				switch m.SelectedGame {

				case "Tetris":
					m.launchTetris = true
					return m, tea.Quit

				default: // quit out of the app for now when selecting other games
					return m, tea.Quit
				}

			case utils.ScreenScoreMenu:
				choice := m.ScoreMenu.SelectedItem().(screens.MenuItem)

				var filename string
				switch choice {

				case "Marathon Tetris":
					filename = "internal/leaderboard/data/tetris_marathon_scores.csv"

				case "Sprint Tetris":
					filename = "internal/leaderboard/data/tetris_sprint_scores.csv"

				case "Ultra Tetris":
					filename = "internal/leaderboard/data/tetris_ultra_scores.csv"
				}

				m.Leaderboard = leaderboard.NewLeaderBoardMenu(filename)
				m.Screen = utils.ScreenLeaderboard
				return m, nil
			}
		}

	// this section is for handling the loading and exits
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

	// screen updates
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

	case utils.ScreenLeaderboard:
		newModel, cmd := m.Leaderboard.Update(msg)
		if lm, ok := newModel.(*leaderboard.LeaderboardModel); ok {
			m.Leaderboard = lm
		}

		return m, cmd
	}

	return m, nil
}

func (m *MainMenuModels) View() string {
	if m.quitting {
		dots := strings.Repeat(".", m.pulseDotCount)
		quitText := m.Style.QuitTextStyle.Render(fmt.Sprintf("Exiting Gocade%s", dots))
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
		gopher := screens.RenderGopher(m.width, m.height, m.Style, utils.GopherMascot)
		mainMenu := m.MainMenu.View()

		content := lipgloss.JoinVertical(lipgloss.Center, gopher, mainMenu)

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)

	case utils.ScreenGameMenu:
		gopher := screens.RenderGopher(m.width, m.height, m.Style, utils.GopherMascotSword)
		gameMenu := m.GameMenu.View()

		content := lipgloss.JoinVertical(lipgloss.Center, gopher, gameMenu)

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)

	case utils.ScreenScoreMenu:
		gopher := screens.RenderGopher(m.width, m.height, m.Style, utils.GopherMascotSword)
		scoreMenu := m.ScoreMenu.View()

		content := lipgloss.JoinVertical(lipgloss.Center, gopher, scoreMenu)

		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)

	case utils.ScreenLeaderboard:
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, m.Leaderboard.View())

	default:
		return "** Unknown screen **"
	}
}
