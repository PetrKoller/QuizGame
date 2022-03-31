package main

import (
	"encoding/csv"
	"errors"
	"io"
)

type QuestionCard struct {
	Question string
	Answer   string
}

var InvalidQuestionCardFormatErr = errors.New("invalid question card csv format")

func LoadQuestionCards(r *csv.Reader) ([]QuestionCard, error) {
	var questionCards []QuestionCard

	for {
		record, err := r.Read()
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
			Question: record[0],
			Answer:   record[1],
		})
	}

	return questionCards, nil
}
