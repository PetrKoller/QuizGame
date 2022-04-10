package main

import (
	"bufio"
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestStartQuiz_AnsweredAllProblems(t *testing.T) {
	t.Parallel()

	timeLimit := 3 * time.Second
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
			quizStart := time.Now()

			result, err := StartQuiz(test.qcs, bufio.NewReader(strings.NewReader(test.input)), timeLimit)
			if time.Since(quizStart) > timeLimit {
				assert.Failf(t, "Quiz should've ended before time limit", "Time limit is %v", timeLimit.String())
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

type timeBlockingReadStringer struct {
	blockAfterCalls        int
	blockDurationInSeconds time.Duration
	counter                int
}

func (tb *timeBlockingReadStringer) ReadString(delim byte) (string, error) {
	tb.counter++
	if tb.counter > tb.blockAfterCalls {
		time.Sleep(tb.blockDurationInSeconds)
	}

	return "5", nil
}

func TestStartQuiz_TimeLimitPassed(t *testing.T) {
	t.Parallel()

	timeLimit := 1 * time.Second
	blockDuration := 10 * time.Second
	tests := []struct {
		name                string
		qcs                 []QuestionCard
		blockReadAfterCalls int
		expected            int
	}{
		{
			name:                "TimeLimitPassed1",
			qcs:                 []QuestionCard{{Question: "5+5", Answer: "10"}, {Question: "4+1", Answer: "5"}, {Question: "2+3", Answer: "5"}, {Question: "3 + 3", Answer: "6"}, {Question: "2 + 3", Answer: "5"}},
			blockReadAfterCalls: 3,
			expected:            2,
		},
		{
			name:                "TimeLimitPassed2",
			qcs:                 []QuestionCard{{Question: "1+4", Answer: "5"}, {Question: "1+8", Answer: "9"}, {Question: "2+3", Answer: "5"}},
			blockReadAfterCalls: 1,
			expected:            1,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			quizStart := time.Now()
			result, err := StartQuiz(test.qcs, &timeBlockingReadStringer{blockAfterCalls: test.blockReadAfterCalls, blockDurationInSeconds: blockDuration}, timeLimit)

			quizEnd := time.Since(quizStart)
			if quizEnd < timeLimit {
				assert.Failf(t, "Quiz shouldn't have ended before time limit", "Time limit is %v, quiz ended after %v", timeLimit.String(), quizEnd.String())
			}

			assert.NoError(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

type errorReadStringer struct {
	errToReturn error
}

func (e *errorReadStringer) ReadString(delim byte) (string, error) {
	return "", e.errToReturn
}

func TestStartQuiz_Error(t *testing.T) {
	t.Parallel()

	correct, err := StartQuiz([]QuestionCard{{Question: "5+5", Answer: "10"}}, &errorReadStringer{errToReturn: errors.New("Unexpected error")}, 10*time.Second)

	assert.Error(t, err)
	assert.Equal(t, 0, correct)
}
