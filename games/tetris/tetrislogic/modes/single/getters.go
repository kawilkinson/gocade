package single

import (
	"time"

	"github.com/kawilkinson/gocade/games/tetris/tetrislogic"
)

func (g *Game) IsGameOver() bool {
	return g.gameOver
}

// Massive credits to Tetrigo for the game logic: https://github.com/Broderick-Westrope/tetrigo

func (g *Game) GetVisibleMatrix() (tetrislogic.Matrix, error) {
	matrix := g.matrix.DeepCopy()

	if g.ghostTet != nil {
		err := matrix.AddTetrimino(g.ghostTet)
		if err != nil {
			return nil, err
		}
	}

	if err := matrix.AddTetrimino(g.tetInPlay); err != nil {
		return nil, err
	}

	return matrix.GetVisible(), nil
}

func (g *Game) GetBagTetriminos() []tetrislogic.Tetrimino {
	return g.nextQueue.GetElements()
}

func (g *Game) GetHoldTetrimino() *tetrislogic.Tetrimino {
	return g.holdQueue
}

func (g *Game) GetTotalScore() int {
	return g.scoring.Total()
}

func (g *Game) GetLevel() int {
	return g.scoring.Level()
}

func (g *Game) GetLinesCleared() int {
	return g.scoring.Lines()
}

func (g *Game) GetDefaultFallInterval() time.Duration {
	return g.fall.DefaultInterval
}
