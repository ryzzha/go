package quizs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Quiz struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func LoadQuizs() ([]Quiz, error) {
	jsonData, err := os.ReadFile("quiz.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read quiz file: %w", err)
	}

	var quizs []Quiz

	err = json.Unmarshal(jsonData, &quizs)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return quizs, nil
}