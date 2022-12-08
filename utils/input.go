package utils

import (
	"errors"
	"fmt"
)

func InputAfterQuestion(question string) (string, error) {
	var input string
	fmt.Print(question)
	_, err := fmt.Scan(&input)
	if err != nil {
		return "", errors.New("Invalid input")
	}
	return input, nil
}
