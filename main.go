package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int64("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	questionCards, err := LoadQuestionCards(*filename)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	correctAnswers, err := StartQuiz(questionCards, bufio.NewReader(os.Stdin), time.Duration(*timeLimit)*time.Second)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("You got right %v out of %v problems\n", correctAnswers, len(questionCards))
}

type ReadStringer interface {
	ReadString(delim byte) (string, error)
}

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
