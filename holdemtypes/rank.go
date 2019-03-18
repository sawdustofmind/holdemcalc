package holdemtypes

import "unicode/utf8"

// Rank Reprecents suit in one symbol
type Rank byte

// All possible Rank values
const (
	Two   Rank = '2'
	Three Rank = '3'
	Four  Rank = '4'
	Fife  Rank = '5'
	Six   Rank = '6'
	Seven Rank = '7'
	Eight Rank = '8'
	Nine  Rank = '9'
	Ten   Rank = 'T'
	Jack  Rank = 'J'
	Queen Rank = 'Q'
	King  Rank = 'K'
	Ace   Rank = 'A'
)

// Ranks is all possible Ranks
var Ranks = []Rank{Two, Three, Four, Fife, Six, Seven, Eight, Nine,
	Ten, Jack, Queen, King, Ace}

// RunkLess is less function for use in sorting
func RunkLess(r1, r2 Rank) bool {
	if r1 == r2 {
		return false
	}
	for _, r := range Ranks {
		if r1 == r {
			return true
		}
		if r2 == r {
			return false
		}
	}
	return false
}

func (r Rank) String() string {
	rune, _ := utf8.DecodeRune([]byte{byte(r)})
	return string(rune)
}
