package prime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return -1, errors.New("`n` must be greater than 0")
	}

	currentInteger := 0
	nThPrime := 0

	for nThPrime < n {
		currentInteger++
		if isPrime(currentInteger) {
			nThPrime++
		}
	}
	return currentInteger, nil
}

func isPrime(integerToBeChecked int) bool {
	if integerToBeChecked < 2 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(integerToBeChecked)))+1; i++ {
		if integerToBeChecked%i == 0 {
			return false
		}
	}
	return true
}
