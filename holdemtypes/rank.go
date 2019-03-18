package holdemtypes

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

// AllRanks returns all possible Ranks
func AllRanks() []Rank {
	return []Rank{Two, Three, Four, Fife, Six, Seven, Eight, Nine,
		Ten, Jack, Queen, King, Ace}
}
