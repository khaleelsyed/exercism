package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fd FodderCalculator, numberOfCows int) (float64, error) {
	totalFodder, err := fd.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fd.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (totalFodder / float64(numberOfCows)) * fatteningFactor, nil
}

func ValidateInputAndDivideFood(fd FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {
		return DivideFood(fd, numberOfCows)
	}
	return 0, errors.New("invalid number of cows")
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows > 0 {
		return nil
	} else if numberOfCows == 0 {
		return newInvalidCowsError(numberOfCows, "no cows don't need food")
	}
	return newInvalidCowsError(numberOfCows, "there are no negative cows")
}

type InvalidCowsError struct {
	numberOfCows  int
	customMessage string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.customMessage)
}

func newInvalidCowsError(numberOfCows int, customMessage string) error {
	return &InvalidCowsError{
		numberOfCows:  numberOfCows,
		customMessage: customMessage,
	}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
