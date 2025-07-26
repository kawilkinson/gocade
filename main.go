package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mainMenu := CreateMainMenuModels()

	p := tea.NewProgram(mainMenu, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running Gocade: %v\n", err)
		os.Exit(1)
	}
}
