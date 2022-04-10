package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadQuestionCards_Success(t *testing.T) {
	t.Parallel()
	expectedCards := []QuestionCard{
		{
			Question: "5+5",
			Answer:   "10",
		},
		{
			Question: "7+3",
			Answer:   "10",
		},
		{
			Question: "1+1",
			Answer:   "2",
		},
	}

	input := "5+5,10\n7+3,10\n1+1,2"

	res, err := LoadQuestionCards(strings.NewReader(input))

	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedCards, res)
}

func TestLoadQuestionCards_Err(t *testing.T) {
	t.Parallel()

	input := "5+5,10,5\nads,fds,af,ads,fasd,dsa,,ads,\n\nasd..sdsddsa,sdafsad"

	res, err := LoadQuestionCards(strings.NewReader(input))

	assert.Nil(t, res)
	assert.ErrorIs(t, InvalidQuestionCardFormatErr, err)
}

func TestLoadQuestionCards_SpaceTrim(t *testing.T) {
	t.Parallel()
	expectedQuestion := "5 + 5"
	expectedAnswer := "10"
	input := "     5 + 5,                 10\n       5 + 5,10\n5 + 5,         10"

	qcs, err := LoadQuestionCards(strings.NewReader(input))

	assert.NoError(t, err)

	for _, qc := range qcs {
		assert.Equal(t, expectedQuestion, qc.Question)
		assert.Equal(t, expectedAnswer, qc.Answer)
	}
}
