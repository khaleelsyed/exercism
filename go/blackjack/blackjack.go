package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// isFigureOfTen returns if a card is a Figure of Ten(Ten, Jack, Queen or King).
func isFigureOfTen(card string) bool {
	return ParseCard(card) == 10
}

// calculateSumCards returns the sum value of two cards.
func calculateSumCards(card1 string, card2 string) int {
	return ParseCard(card1) + ParseCard(card2)
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case ParseCard(card1)+ParseCard(card2) == 21:
		if dealerCard != "ace" && !isFigureOfTen(dealerCard) {
			return "W"
		} else {
			return "S"
		}
	case 17 <= calculateSumCards(card1, card2) && calculateSumCards(card1, card2) <= 20:
		return "S"
	case 12 <= calculateSumCards(card1, card2) && calculateSumCards(card1, card2) <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}
	default:
		return "H"
	}
}
