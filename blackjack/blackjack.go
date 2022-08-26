package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "ten":
		return 10
	default:
		return 0
	}

}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	myHand := ParseCard(card1) + ParseCard(card2)
	switch {
	case myHand == 22:
		return "P"
	case myHand == 21:
		if ParseCard(dealerCard) >= 10 {
			return "S"
		}
		return "W"
	case myHand >= 17 && myHand <= 20:
		return "S"
	case myHand >= 12 && myHand <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
