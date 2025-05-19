package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/quizs"
	"strings"
)

func Run(questions []quizs.Quiz) (correctAnswers uint) {
	fmt.Println("Start Country Quiz!")

	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}
	}

	return correctAnswers
}

func askQuestion(q quizs.Quiz) bool {
	fmt.Printf("\nEnter the capital of %s: ", q.Country)

	if getUserInput() == strings.ToLower(q.Capital) {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Printf("Incorrect! The correct answer is %s.\n", q.Capital)
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Your answer: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading the entered text! Please try again...")
			continue
		}

		return strings.ToLower(strings.TrimSpace(result))
	}
}
