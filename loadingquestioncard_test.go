package main

import (
	"github.com/stretchr/testify/assert"
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

	res, err := LoadQuestionCards("testingData/correctFormat.csv")

	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedCards, res)
}

func TestLoadQuestionCards_Err(t *testing.T) {
	t.Parallel()
	res, err := LoadQuestionCards("testingData/invalidFormat.csv")

	assert.Nil(t, res)
	assert.ErrorIs(t, InvalidQuestionCardFormatErr, err)
}

func TestLoadQuestionCards_SpaceTrim(t *testing.T) {
	t.Parallel()
	expectedQuestion := "5 + 5"
	expectedAnswer := "10"

	qcs, err := LoadQuestionCards("testingData/spaceTrim.csv")

	assert.NoError(t, err)

	for _, qc := range qcs {
		assert.Equal(t, expectedQuestion, qc.Question)
		assert.Equal(t, expectedAnswer, qc.Answer)
	}
}
