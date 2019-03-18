package holdemtypes

// Suit Reprecents first letter of suit
type Suit byte

// All possible Suit values
const (
	Clubs    Suit = 'c'
	Diamonds Suit = 'd'
	Hearts   Suit = 'h'
	Spades   Suit = 's'
)

// AllSuites return all suites
func AllSuites() []Suit {
	return []Suit{Clubs, Diamonds, Hearts, Spades}
}
