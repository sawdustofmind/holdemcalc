package holdemtypes

// Card represents playing card
type Card struct {
	Rank
	Suit
}

// StandardDeck return default 52-cards deks
func StandardDeck() []Card {
	deck := make([]Card, 0, 52)
	for _, rank := range Ranks {
		for _, suit := range Suits {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}
	return deck
}
