package tutils

type Screen int

const (
	ScreenTetrisMenu = Screen(iota)
	ScreenTetrisGame

	TetrisTitle = `
░██████████              ░██             ░██           
    ░██                  ░██                           
    ░██     ░███████  ░████████ ░██░████ ░██ ░███████  
    ░██    ░██    ░██    ░██    ░███     ░██░██        
    ░██    ░█████████    ░██    ░██      ░██ ░███████  
    ░██    ░██           ░██    ░██      ░██       ░██ 
    ░██     ░███████      ░████ ░██      ░██ ░███████                                                      
	`
)

var ScreenToStrMap = map[Screen]string{
	ScreenTetrisMenu: "Menu",
	ScreenTetrisGame: "Tetris",
}

func (s Screen) String() string {
	return ScreenToStrMap[s]
}
