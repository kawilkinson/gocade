package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/games/snake"
	"github.com/kawilkinson/gocade/games/tetris"
	"github.com/kawilkinson/gocade/games/tetris/tetrisconfig"
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

func runTetris() {
	input := tetris.NewInput(
		tetrisconfig.ModeMenu,
		&tetrisconfig.MenuInput{},
		tetrisconfig.CreateConfig(),
	)

	tetrisModel, err := tetris.NewModel(input)
	if err != nil {
		fmt.Printf("Error starting Tetris: %v\n", err)
		os.Exit(1)
	}

	tetrisProgram := tea.NewProgram(tetrisModel, tea.WithAltScreen())
	if _, err := tetrisProgram.Run(); err != nil {
		fmt.Printf("Error running Tetris: %v\n", err)
		os.Exit(1)
	}
}

func runSnake() {
	snakeModel := snake.CreateSnakeGameModel()

	snakeProgram := tea.NewProgram(snakeModel, tea.WithAltScreen())
	if _, err := snakeProgram.Run(); err != nil {
		fmt.Printf("Error running Snake: %v\n", err)
		os.Exit(1)
	}
}
