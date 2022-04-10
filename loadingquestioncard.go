package main

import (
	"encoding/csv"
	"errors"
	"io"
	"strings"
)

type QuestionCard struct {
	Question string
	Answer   string
}

var InvalidQuestionCardFormatErr = errors.New("invalid question card csv format")

// LoadQuestionCards loads cards from given io.Reader parameter reader and returns slice of QuestionCard.
//Expected data format from where reader reads is CSV, default delimiter is ','.
//Error is returned, if data format is not CSV or question card format is not fulfilled ('question,answer') or any other error is being encountered.
func LoadQuestionCards(reader io.Reader) ([]QuestionCard, error) {
	var questionCards []QuestionCard
	csvReader := csv.NewReader(reader)

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) != 2 {
			return nil, InvalidQuestionCardFormatErr
		}
		questionCards = append(questionCards, QuestionCard{
			Question: strings.TrimSpace(record[0]),
			Answer:   strings.TrimSpace(record[1]),
		})
	}

	return questionCards, nil
}
