package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sequenceChannel := make(chan []int)

	validDivisors := make([]int, len(divisors))
	for _, divisor := range divisors {
		if divisor != 0 {
			validDivisors = append(validDivisors, divisor)
		}
	}

	for _, divisor := range validDivisors {
		go generateSequenceOfMultiples(divisor, limit, sequenceChannel)
	}

	for {
		if len(sequenceChannel) == len(validDivisors) {
			break
		}
	}
	sequenceToSum := transformToSingleDimension(sequenceChannel)
	return sum(sequenceToSum)
}

func generateSequenceOfMultiples(divisor int, limit int, sequenceChannel chan []int) {
	sequence := make([]int, limit)
	for num := 0; num < limit; num++ {
		if num%divisor == 0 {
			sequence = append(sequence, num)
		}
	}
	sequenceChannel <- sequence
}

func sum(sequenceToSum []int) int {
	total := 0
	for _, value := range sequenceToSum {
		total += value
	}
	return total
}

func transformToSingleDimension(sequenceChannel chan []int) []int {
	unionMap := make(map[int]bool)

	for sequence := range sequenceChannel {
		for _, value := range sequence {
			unionMap[value] = true
		}
	}

	validSequence := make([]int, len(unionMap))
	for key := range unionMap {
		validSequence = append(validSequence, key)
	}

	return validSequence
}
