package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	mainMenu := CreateMainMenu()

	if _, err := tea.NewProgram(mainMenu).Run(); err != nil {
		fmt.Printf("Error running Gocade: %v\n", err)
		os.Exit(1)
	}
}
