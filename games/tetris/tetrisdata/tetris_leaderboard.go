package tetrisdata

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/kawilkinson/gocade/games/tetris/tutils"
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
		scoreFile = tutils.MarathonScoreFile

	case "Sprint":
		scoreFile = tutils.SprintScoreFile

	case "Ultra":
		scoreFile = tutils.UltraScoreFile
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

func LoadScores(mode string) ([]Score, error) {
	var scoreFile string
	switch mode {
	case "Marathon":
		scoreFile = tutils.MarathonScoreFile

	case "Sprint":
		scoreFile = tutils.SprintScoreFile

	case "Ultra":
		scoreFile = tutils.UltraScoreFile
	}

	file, err := os.Open(scoreFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Score{}, nil
		}

		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var scores []Score
	for _, rec := range records {
		if len(rec) != 6 {
			continue
		}

		var score int
		score, err = strconv.Atoi(rec[2])
		if err != nil {
			score = 0
		}

		var lines int
		lines, err = strconv.Atoi(rec[3])
		if err != nil {
			lines = 0
		}

		var level int
		level, err = strconv.Atoi(rec[4])
		if err != nil {
			level = 0
		}

		scores = append(scores, Score{
			Name:  rec[1],
			Score: score,
			Lines: lines,
			Level: level,
			Mode:  rec[5],
		})
	}

	return scores, nil
}
