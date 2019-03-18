package holdemtypes

// Card represents playing card
type Card struct {
	Rank
	Suit
}

// StandardDeck return default 52-cards deks
func StandardDeck() []Card {
	deck := make([]Card, 0, 52)
	for _, rank := range AllRanks() {
		for _, suit := range AllSuites() {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}
