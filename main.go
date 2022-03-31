package main

import (
	"encoding/csv"
	"log"
	"os"
)

func StartQuiz(questionCards []QuestionCard) error {
	for _, qc := range questionCards {
		log.Printf("%+v\n", qc)
	}

	return nil
}

func main() {
	fileName := "problems.csv"

	csvFile, err := os.Open(fileName)
	defer csvFile.Close()
	if err != nil {
		log.Fatal()
	}

	questionCards, err := LoadQuestionCards(csv.NewReader(csvFile))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	if err = StartQuiz(questionCards); err != nil {
		log.Fatal("Error: ", err)
	}
}
