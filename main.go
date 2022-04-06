package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\")")
	flag.Parse()

	questionCards, err := LoadQuestionCards(*fileName)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	correctAnswers, err := StartQuiz(questionCards, bufio.NewReader(os.Stdin))

	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("You got right %v out of %v problems\n", correctAnswers, len(questionCards))
}

func StartQuiz(questionCards []QuestionCard, reader *bufio.Reader) (int, error) {
	correctAns := 0
	for i, qc := range questionCards {
		fmt.Printf("Problem nr. %v: %v = \n", i+1, qc.Question)
		ans, err := reader.ReadString('\n')

		if err == io.EOF {
			return correctAns, nil
		} else if err != nil {
			return 0, err
		}

		if strings.TrimSpace(ans) == qc.Answer {
			correctAns++
		}
	}

	return correctAns, nil
}
