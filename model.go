package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/internal/screens"
	"github.com/kawilkinson/gocade/internal/utils"
)

type model struct {
	Screen       utils.Screen
	MainMenu     list.Model
	GameMenu     list.Model
	ScoreMenu    list.Model
	SelectedGame string
	Quitting     bool
}

func CreateModels() model {
	return model{
		Screen:    utils.ScreenMainMenu,
		MainMenu:  screens.NewMainMenu(utils.MenuWidth, utils.MenuHeight),
		GameMenu:  screens.NewGameMenu(utils.MenuWidth, utils.MenuHeight),
		ScoreMenu: screens.NewScoreMenu(utils.MenuWidth, utils.MenuHeight),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c": // for quitting the program at any point
			m.Quitting = true
			return m, tea.Quit

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
					m.Quitting = true
					return m, tea.Quit
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
	}

	var cmd tea.Cmd
	switch m.Screen {

	case utils.ScreenMainMenu:
		m.MainMenu, cmd = m.MainMenu.Update(msg)

	case utils.ScreenGameMenu:
		m.GameMenu, cmd = m.GameMenu.Update(msg)

	case utils.ScreenScoreMenu:
		m.ScoreMenu, cmd = m.ScoreMenu.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	if m.Quitting {
		return screens.QuitTextStyle.Render("Exiting Gocade...")
	}
	switch m.Screen {

	case utils.ScreenMainMenu:
		return "\n" + m.MainMenu.View()

	case utils.ScreenGameMenu:
		return "\n" + m.GameMenu.View() + "\n\nPress 'b' to go back"

	case utils.ScreenScoreMenu:
		return "\n" + m.ScoreMenu.View() + "\n\nPress 'b' to go back"

	default:
		return "Unknown screen"
	}
}
