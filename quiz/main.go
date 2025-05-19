package main

import (
	"fmt"
	"os"
	"quiz/quizs"
	"quiz/shuffler"
	"quiz/game"
)

func main() {
	quizs, err := quizs.LoadQuizs()
	if err != nil {
		fmt.Printf("Error loading quiz: %v\n", err)
		os.Exit(1)
	}

	shuffledQuizs := shuffler.Shuffle(quizs)

	correctAnswers := game.Run(shuffledQuizs)

	fmt.Printf("Correct answers: %d", correctAnswers)
}