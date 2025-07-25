package screens

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kawilkinson/gocade/internal/utils"
)

type TickMsg struct{}
type PulseMsg struct{}
type ExitMsg struct{}

func Tick() tea.Cmd {
	return tea.Tick(time.Millisecond*utils.LoadingTickSpeed, func(time.Time) tea.Msg {
		return TickMsg{}
	})
}

func PulseDots() tea.Cmd {
	return tea.Tick(time.Millisecond*utils.LoadingDotPulseSpeed, func(time.Time) tea.Msg {
		return PulseMsg{}
	})
}

// helper function to help DRY the code for exiting the program
func ExitGocade() tea.Cmd {
	return tea.Batch(
		PulseDots(),
		tea.Tick(
			time.Second*utils.ExitScreenTickSpeed,
			func(time.Time) tea.Msg {
				return ExitMsg{}
			}),
	)
}
