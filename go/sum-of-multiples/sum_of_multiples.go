package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sequenceOfMultiples := make([][]int, len(divisors))
	for i := range divisors {
		sequenceOfMultiples[i] = make([]int, 0)
	}

	for i, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for num := 0; num < limit; num++ {
			if num%divisor == 0 {
				sequenceOfMultiples[i] = append(sequenceOfMultiples[i], num)
			}
		}
	}
	sequenceToSum := transformToSingleDimension(sequenceOfMultiples)
	return sum(sequenceToSum)
}

func sum(sequenceToSum []int) int {
	total := 0
	for _, value := range sequenceToSum {
		total += value
	}
	return total
}

func transformToSingleDimension(sequenceOfMultiples [][]int) []int {
	unionMap := make(map[int]bool)

	for divisor := range sequenceOfMultiples {
		sequences := sequenceOfMultiples[divisor]
		for _, element := range sequences {
			unionMap[element] = true
		}
	}

	validSequence := make([]int, len(unionMap))
	for key := range unionMap {
		validSequence = append(validSequence, key)
	}

	return validSequence
}
