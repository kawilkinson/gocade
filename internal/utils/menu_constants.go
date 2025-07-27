package utils

import (
	"strings"
)

type Screen int

const (
	MenuWidth  = 40
	MenuHeight = 12

	ExitScreenTickSpeed  = 2   // seconds
	LoadingTickSpeed     = 2   // milliseconds
	LoadingDotPulseSpeed = 500 // milliseconds

	// different types of screens for the arcade menus
	ScreenLoading = Screen(iota)
	ScreenMainMenu
	ScreenGameMenu
	ScreenScoreMenu
	ScreenLeaderboard

	// Colors for styling
	GopherColor = "#00ADD8"

	// Art for arcade machine
	GopherMascot = `         
	         ⡀⡠⢠⠰⢐⠔⡘⡐⡑⢌⢒⠰⡐⠤⢠⢀
	⡀⡄⢄⢄⡠⠢⡡⡑⠌⠔⢌⠢⠨⢂⢊⠄⢅⠢⠑⠐⠑⠔⢡⢑⠢⢔⠔⡂⢆⢄
⠀⢀⠎⡂⣮⡰⡨⢂⠕⠀⠀⠀⠀⠀⠈⠅⢕⠠⡡⠃⠀⠀⠀⠀⠀⠀⠂⡅⢅⠳⣮⢐⠌⡢
⠀⠰⡁⡪⡻⠡⡂⠅⣰⡶⢦⠀⠀⠀⠀⠈⡢⠡⡂⢴⡿⡣⠀⠀⠀⠀⠀⢪⢐⠡⢱⢐⢡⠊
⠀⠀⠑⠦⡃⠕⡨⢂⠈⠛⠂⠀⠀⠀⠀⠠⡊⢔⠡⡀⠋⠁⠀⠀⠀⠀⡠⢃⠔⡨⠂⡇⠁
⠀⠀⠀⡘⠄⢕⠠⡡⠣⡠⡀⣀⢀⢄⠢⢣⢾⣷⣿⡎⡢⠄⠤⡐⢄⢃⠪⢐⠌⡐⠅⡪
⠀⠀⠀⡪⠨⢂⠅⡢⠡⢂⢊⠔⡐⠔⢅⠁⠡⡁⠅⡐⢈⢪⢈⠢⠡⡂⠅⢅⢊⠔⠡⢊⡂
⠀⠀⠀⡪⠨⢂⢊⠄⢅⠕⡐⠌⠔⡡⠡⡃⠁⢘⠀⠸⡐⠡⡂⠅⠕⡐⢅⢑⢐⠌⡊⠔⡅
⠀⠀⠀⢜⢈⠢⢂⠅⢅⢂⠪⠨⡨⢐⠡⢊⠤⠪⡠⡘⠄⢕⠐⠅⢕⠨⡐⠔⡁⡢⠊⠔⡅
⠀⠀⠀⢨⠂⢅⠅⡊⢔⠐⢅⢑⢐⠡⢊⠔⡨⢂⢂⢊⠌⡂⢅⠕⡁⡢⠊⠔⡁⡢⢑⠡⡂
⠀⠀⠀⠰⡡⠡⢊⢐⠡⢊⠔⡁⡢⠡⡡⢂⠢⡑⠄⠕⡨⢐⠡⢂⢊⠄⠕⡡⢊⢐⠔⠡⠅
⠀⠀⢀⢈⠆⢅⢑⢐⠡⠡⡂⡊⢄⠕⡠⢑⠨⡐⡡⢑⢐⠡⢊⠔⡐⡡⢑⠄⠕⡐⠌⢌⢅⡀⡀
⠠⡊⢀⢂⠇⢅⠢⠑⠌⢌⠢⠨⡂⠌⢔⢁⢊⠔⡐⡡⢂⠕⡐⡡⢂⢊⠔⡨⠨⢂⠅⢅⢅⠄⢌⢂
⠈⠢⠂⢱⠑⢄⠅⢕⠡⡑⡨⢂⢊⠌⡂⡢⠡⢂⠢⡂⢅⠢⢂⠢⠡⡂⡊⢄⠅⢅⠪⢐⢱⠑⠐⠁
⠀⠀⠀⢨⠨⢂⢊⠔⡨⢐⠄⠕⡠⢑⢐⠌⢌⠢⠡⡂⢅⠪⢐⠡⡑⡐⠌⡂⠅⢅⠪⢐⢸
⠀⠀⠀⢸⠈⢔⢐⠡⢂⠅⡊⠌⠔⡁⡢⢑⢐⠡⡑⠄⠕⡈⡢⢑⢐⠌⢌⠢⠡⡑⠌⡂⡪
⠀⠀⠀⢪⢈⢂⠢⠡⡡⢊⢐⠅⠕⡨⢐⠡⢂⢑⢐⠅⠕⡐⠌⢔⢐⠌⡂⡪⠨⡐⡡⢂⢪
⠀⠀⠀⢱⠐⡡⠊⢌⠔⡰⢐⠌⢌⠔⡁⡪⢐⠡⢂⢊⠌⡂⢅⠅⡂⡊⠔⡐⡡⢂⠢⠡⡪
⠀⠀⠀⠸⡐⡐⠅⢅⠢⢂⠅⡊⠔⡨⢐⢐⠡⢊⠔⡐⡡⠨⢂⠅⡊⠔⡡⢂⢊⠄⢅⠕⡂
⠀⠀⠀⠈⡆⢌⠌⡢⠡⠡⢊⢐⠅⡊⠔⡁⡪⠐⠌⢔⠨⠨⡂⡊⠔⡡⢂⠅⡢⠡⠡⡪
⠀⠀⠀⠀⠘⡔⠨⡐⡡⢑⠡⢂⢊⠄⢕⠐⢔⠡⢃⠅⡊⠌⢔⢈⠢⢂⠅⡊⠄⢕⢑⠁
⠀⠀⠀⠀⠀⠘⢌⡂⡢⠡⢊⠔⡐⢅⠢⠡⡡⢊⢐⠌⠔⡡⠡⡂⠅⢅⠊⢔⢡⠣⠁
⠀⠀⠀⠀⠀⠠⢊⠐⠈⡂⡇⣅⢂⠪⢐⠌⡊⠔⡐⡡⠨⡨⢐⠡⡨⡨⡢⠕⢅⠂⡈⡢
     ⠘⢄⢅⠎⠈⠀⠀⠉⠘⠂⠃⠒⠕⠢⠪⠒⠘⠂⠃⠁⠁⠀⠀⠈⠂⠆⠎⠂⠀
`
)

// helper function to ensure the art for the gopher is always aligned properly
func NormalizeWidth(ascii string) string {
	lines := strings.Split(ascii, "\n")
	max := 0

	for _, line := range lines {
		if len([]rune(line)) > max {
			max = len([]rune(line))
		}
	}

	for i, line := range lines {
		pad := max - len([]rune(line))
		if pad > 0 {
			lines[i] = line + strings.Repeat(" ", pad)
		}
	}

	return strings.Join(lines, "\n")
}
