package tetrisdata

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/kawilkinson/gocade/internal/utils"
)

type Score struct {
	Name  string
	Score int
	Lines int
	Level int
	Mode  string
}

func SaveScore(s Score, mode string) error {
	var scoreFile string
	switch mode {
	case "Marathon":
		scoreFile = utils.MarathonScoreFile

	case "Sprint":
		scoreFile = utils.SprintScoreFile

	case "Ultra":
		scoreFile = utils.UltraScoreFile
	}

	file, err := os.OpenFile(scoreFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		s.Name,
		strconv.Itoa(s.Score),
		strconv.Itoa(s.Lines),
		strconv.Itoa(s.Level),
		s.Mode,
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}
