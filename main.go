package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int64("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	csvFile, err := os.Open(*filename)
	defer csvFile.Close()
	if err != nil {
		log.Fatal()
	}

	questionCards, err := LoadQuestionCards(csvFile)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	correctAnswers, err := StartQuiz(questionCards, bufio.NewReader(os.Stdin), time.Duration(*timeLimit)*time.Second)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Printf("You got right %v out of %v problems\n", correctAnswers, len(questionCards))
}
