package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func addValue(freqMap FreqMap, key rune, value int) FreqMap {
	if _, exists := freqMap[key]; exists {
		freqMap[key] += value
	} else {
		freqMap[key] = value
	}
	return freqMap
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	freqMap := make(FreqMap)
	for _, text := range texts {
		for k, v := range Frequency(text) {
			addValue(freqMap, k, v)
		}
	}

	return freqMap
}
