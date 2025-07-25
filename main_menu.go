package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color(GopherColor))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(GopherColor))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type menuItem string

func (i menuItem) FilterValue() string { return "" }

type menuDelegate struct{}

func (d menuDelegate) Height() int                             { return 1 }
func (d menuDelegate) Spacing() int                            { return 0 }
func (d menuDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d menuDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(menuItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	screen       screen // screen is an int defined in constants.go
	mainMenu     list.Model
	gameMenu     list.Model
	scoreMenu    list.Model
	list         list.Model
	choice       string
	quitting     bool
	selectedGame string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "b":
			if m.screen == ScreenGameMenu {
				m.screen = ScreenMainMenu
				return m, nil
			}

		case "enter":
			switch m.screen {

			case ScreenMainMenu:
				choice := m.mainMenu.SelectedItem().(menuItem)

				switch choice {

				case "Play Game":
					m.screen = ScreenGameMenu

				case "High Scores":
					return m, tea.Quit

				case "Quit":
					m.quitting = true
					return m, tea.Quit
				}

			case ScreenGameMenu:
				m.selectedGame = string(m.gameMenu.SelectedItem().(menuItem))
				m.choice = fmt.Sprintf("Launching game: %s\n", m.selectedGame)
				return m, tea.Quit

			}
		}
	}

	var cmd tea.Cmd
	switch m.screen {
	case ScreenMainMenu:
		m.mainMenu, cmd = m.mainMenu.Update(msg)

	case ScreenGameMenu:
		m.gameMenu, cmd = m.gameMenu.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return quitTextStyle.Render("Exiting Gocade...")
	}

	switch m.choice {
	case "High Scores":

	}

	switch m.screen {
	case ScreenMainMenu:
		return "\n" + m.mainMenu.View()

	case ScreenGameMenu:
		return "\n" + m.gameMenu.View() + "\n\nPress 'b' to go back"
	}

	return "\n" + m.list.View()
}

func CreateMainMenu() model {
	menuItems := []list.Item{
		menuItem("Play Game"),
		menuItem("High Scores"),
		menuItem("Quit"),
	}

	games := []list.Item{ // these aren't the final games to be included, added these for testing
		menuItem("Snake"),
		menuItem("Tetris"),
		menuItem("Pong"),
	}

	mainList := list.New(menuItems, menuDelegate{}, MenuWidth, MenuHeight)
	mainList.Title = "Gocade"
	mainList.SetShowStatusBar(false)
	mainList.SetFilteringEnabled(false)
	mainList.Styles.Title = titleStyle
	mainList.Styles.PaginationStyle = paginationStyle
	mainList.Styles.HelpStyle = helpStyle

	gameList := list.New(games, menuDelegate{}, MenuWidth, MenuHeight)
	gameList.Title = "Select a Game"
	gameList.SetShowStatusBar(false)
	gameList.SetFilteringEnabled(false)
	gameList.Styles.Title = titleStyle
	gameList.Styles.PaginationStyle = paginationStyle
	gameList.Styles.HelpStyle = helpStyle

	m := model{
		screen:   ScreenMainMenu,
		mainMenu: mainList,
		gameMenu: gameList,
	}

	return m
}
