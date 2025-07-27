package tutils

import (
	"time"

	"github.com/kawilkinson/gocade/internal/utils"
)

type Screen int

const (
	TimerUpdateInterval = time.Millisecond * 13

	MarathonScoreFile = "internal/leaderboard/data/tetris_marathon_scores.csv"
	SprintScoreFile = "internal/leaderboard/data/tetris_sprint_scores.csv"
	UltraScoreFile = "internal/leaderboard/data/tetris_ultra_scores.csv"

	// ASCII text for styling purposes
	TetrisTitle = `
 _____    _        _     
|_   _|  | |      (_)    
  | | ___| |_ _ __ _ ___ 
  | |/ _ \ __| '__| / __|
  | |  __/ |_| |  | \__ \
  \_/\___|\__|_|  |_|___/                                                    
	`

	PausedMessage = `
______                        _ 
| ___ \                      | |
| |_/ /_ _ _   _ ___  ___  __| |
|  __/ _' | | | / __|/ _ \/ _' |
| | | (_| | |_| \__ \  __/ (_| |
\_|  \__,_|\__,_|___/\___|\__,_|
	`

	GameOverMessage = `
 _____                        _____                
|  __ \                      |  _  |               
| |  \/ __ _ _ __ ___   ___  | | | |_   _____ _ __ 
| | __ / _' | '_ ' _ \ / _ \ | | | \ \ / / _ \ '__|
| |_\ \ (_| | | | | | |  __/ \ \_/ /\ V /  __/ |   
 \____/\__,_|_| |_| |_|\___|  \___/  \_/ \___|_| 
	`
)

// helper function to ensure large ASCII text always shows correctly in Tetris
func RenderLargeText(ascii string) string {
	normalizedTitle := utils.NormalizeWidth(ascii)
	return normalizedTitle
}
