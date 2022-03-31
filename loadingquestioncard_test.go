package main

import (
	"encoding/csv"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadQuestionCards_Success(t *testing.T) {
	expectedCards := []QuestionCard{
		{
			Question: "2+2",
			Answer:   "4",
		},
		{
			Question: "3+3",
			Answer:   "6",
		},
		{
			Question: "1+1",
			Answer:   "2",
		},
	}

	var in string

	for _, card := range expectedCards {
		in += fmt.Sprintf("%v,%v\n", card.Question, card.Answer)
	}

	res, err := LoadQuestionCards(csv.NewReader(strings.NewReader(in)))

	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedCards, res)
}

func TestLoadQuestionCards_Err(t *testing.T) {

}
