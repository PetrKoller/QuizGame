package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

type QuestionCard struct {
	Question string
	Answer   string
}

var InvalidQuestionCardFormatErr = errors.New("invalid question card csv format")

func LoadQuestionCards(fileName string) ([]QuestionCard, error) {
	var questionCards []QuestionCard

	csvFile, err := os.Open(fileName)
	defer csvFile.Close()
	if err != nil {
		log.Fatal()
	}

	r := csv.NewReader(csvFile)

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
			Question: strings.TrimSpace(record[0]),
			Answer:   strings.TrimSpace(record[1]),
		})
	}

	return questionCards, nil
}
