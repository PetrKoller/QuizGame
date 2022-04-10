package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

// ReadStringer is interface that wraps ReadString method from struct bufio.Reader.
//
// ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter.
//If ReadString encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF).
//ReadString returns err != nil if and only if the returned data does not end in delim
type ReadStringer interface {
	ReadString(delim byte) (string, error)
}

// StartQuiz starts whole quiz where you specify question cards deck, reader from where the answers will be processed and maximal duration for which the quiz runs.
// If all question cards are displayed and answered or quiz runs out of time or EOF is passed, number of correct answers is returned.
// If error is met during the quiz, it ends and error is returned
func StartQuiz(questionCards []QuestionCard, reader ReadStringer, durationInSeconds time.Duration) (int, error) {
	correctAns := 0
	errCh := make(chan error)
	ansCh := make(chan string)
	timer := time.NewTimer(durationInSeconds)

	for i, qc := range questionCards {
		fmt.Printf("Problem nr. %v: %v = \n", i+1, qc.Question)
		go func() {
			ans, err := reader.ReadString('\n')
			if err != nil {
				errCh <- err
			}
			ansCh <- ans
		}()

		select {
		case ans := <-ansCh:
			if strings.TrimSpace(ans) == qc.Answer {
				correctAns++
			}
		case <-timer.C:
			fmt.Println("Out of time")
			return correctAns, nil
		case err := <-errCh:
			if err == io.EOF {
				fmt.Println("Called eof")
				return correctAns, nil
			}

			return 0, err
		}
	}

	fmt.Println("You have answered all the questions")
	return correctAns, nil
}
