package tetrisconfig

import "github.com/charmbracelet/lipgloss"

// TODO: adjust styling to add personal touch to this Tetris game (ideally the colors a bit more closer to the Go mascot), default styling
// deserves credit to Tetrigo: https://github.com/Broderick-Westrope/tetrigo

type Theme struct {
	Colours struct {
		TetriminoCells struct { // these are all the tetrimino piece types
			I string
			O string
			T string
			S string
			Z string
			J string
			L string
		}
		EmptyCell string
		GhostCell string
	}

	CellTypes struct {
		Tetriminos string
		EmptyCell  string
		GhostCell  string
	}
}

func CreateTetrisTheme() *Theme {
	theme := new(Theme)

	theme.Colours.TetriminoCells.I = "#64C4EB"
	theme.Colours.TetriminoCells.O = "#F1D448"
	theme.Colours.TetriminoCells.T = "#A15398"
	theme.Colours.TetriminoCells.S = "#64B452"
	theme.Colours.TetriminoCells.Z = "#DC3A35"
	theme.Colours.TetriminoCells.J = "#5C65A8"
	theme.Colours.TetriminoCells.L = "#E07F3A"
	theme.Colours.EmptyCell = "#303040"
	theme.Colours.GhostCell = "white"

	theme.CellTypes.Tetriminos = "██"
	theme.CellTypes.EmptyCell = "▕ "
	theme.CellTypes.GhostCell = "░░"

	return theme
}

type GameStyles struct {
	Playfield           lipgloss.Style
	EmptyCell           lipgloss.Style
	TetriminoCellStyles map[byte]lipgloss.Style
	GhostCell           lipgloss.Style
	Hold                holdStyles
	Information         lipgloss.Style
	RowIndicator        lipgloss.Style
	Bag                 lipgloss.Style
	CellChar            cellCharacters
}

type holdStyles struct {
	View  lipgloss.Style
	Label lipgloss.Style
	Item  lipgloss.Style
}

type cellCharacters struct {
	Empty      string
	Ghost      string
	Tetriminos string
}

func CreateGameStyles(theme *Theme) *GameStyles {
	s := GameStyles{
		Playfield: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0),
		EmptyCell: lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.EmptyCell)),
		TetriminoCellStyles: map[byte]lipgloss.Style{
			'I': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.I)),
			'O': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.O)),
			'T': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.T)),
			'S': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.S)),
			'Z': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.Z)),
			'J': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.J)),
			'L': lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.TetriminoCells.L)),
		},
		GhostCell: lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colours.GhostCell)),
		Hold: holdStyles{
			View: lipgloss.NewStyle().Width(10).Height(5).
				Border(lipgloss.RoundedBorder(), true, false, true, true).
				Align(lipgloss.Center, lipgloss.Center),
			Label: lipgloss.NewStyle().Width(10).PaddingLeft(1).PaddingBottom(1),
			Item:  lipgloss.NewStyle().Width(10).Height(2).Align(lipgloss.Center, lipgloss.Center),
		},
		Information: lipgloss.NewStyle().Width(13).Align(lipgloss.Left, lipgloss.Top),
		RowIndicator: lipgloss.NewStyle().Foreground(lipgloss.Color(theme.CellTypes.EmptyCell)).
			Align(lipgloss.Left).Padding(0, 1, 0),
		Bag: lipgloss.NewStyle().PaddingTop(1),
		CellChar: cellCharacters{
			Empty:      theme.CellTypes.EmptyCell,
			Ghost:      theme.CellTypes.GhostCell,
			Tetriminos: theme.CellTypes.Tetriminos,
		},
	}

	return &s
}
