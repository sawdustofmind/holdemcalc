package holdemtypes

import "unicode/utf8"

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

func (s Suit) String() string {
	rune, _ := utf8.DecodeRune([]byte{byte(s)})
	return string(rune)
}
