package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	for {
		arcadeModels := CreateMainMenuModels()
		p := tea.NewProgram(arcadeModels, tea.WithAltScreen())

		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running Gocade: %v\n", err)
			os.Exit(1)
		}

		if arcadeModels.launchTetris {
			runTetris()
			continue
		}

		if arcadeModels.launchSnake {
			runSnake()
			continue
		}

		break
	}
}
