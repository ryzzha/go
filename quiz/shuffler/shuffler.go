package shuffler

import (
	"math/rand"
	"quiz/quizs"
)

func Shuffle(questions []quizs.Quiz) []quizs.Quiz {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
	return questions
}
