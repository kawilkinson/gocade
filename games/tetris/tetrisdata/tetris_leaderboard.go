package tetrisdata

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Score struct {
	ID    int
	Name  string
	Score int
	Lines int
	Level int
}

const scoreFile = "../../internal/leaderboards/tetris_scores.csv"

func SaveScore(s Score) error {
	file, err := os.OpenFile(scoreFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		strconv.Itoa(s.ID),
		s.Name,
		strconv.Itoa(s.Score),
		strconv.Itoa(s.Lines),
		strconv.Itoa(s.Level),
	}

	return writer.Write(record)
}

func LoadScores() ([]Score, error) {
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

		var id int
		id, err = strconv.Atoi(rec[0])
		if err != nil {
			id = 0
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
			ID:    id,
			Name:  rec[1],
			Score: score,
			Lines: lines,
			Level: level,
		})
	}

	return scores, nil
}
