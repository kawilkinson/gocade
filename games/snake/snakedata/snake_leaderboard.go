package snakedata

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/kawilkinson/gocade/games/snake/sutils"
)

type Score struct {
	Username string
	Score    int
}

func SaveScore(s Score) error {
	scoreFile := sutils.SnakeScoreFile

	file, err := os.OpenFile(scoreFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		s.Username,
		strconv.Itoa(s.Score),
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
