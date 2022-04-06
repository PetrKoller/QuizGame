package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStartQuiz(t *testing.T) {
	tests := []struct {
		name     string
		qcs      []QuestionCard
		input    string
		expected int
	}{
		{
			name:     "Simple",
			qcs:      []QuestionCard{{Question: "5+5", Answer: "10"}},
			input:    "10\n",
			expected: 1,
		},
		{
			name: "CorrectAndWrong",
			qcs:  []QuestionCard{{Question: "1+3", Answer: "4"}, {Question: "1+8", Answer: "9"}},
			input: "\t		4	\n666\n",
			expected: 1,
		},
		{
			name:     "EmptyQuestionCards",
			qcs:      []QuestionCard{},
			input:    "random input\ndasfads\n",
			expected: 0,
		},
		{
			name:     "NilQuestionCards",
			qcs:      nil,
			input:    "d\n",
			expected: 0,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result, err := StartQuiz(test.qcs, bufio.NewReader(strings.NewReader(test.input)))

			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}

}
