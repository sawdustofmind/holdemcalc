package holdemtypes

// CardRange is range of Cards
type CardRange []Card

// Contains return precence of given card
func (r CardRange) Contains(c Card) bool {
	for _, card := range r {
		if card == c {
			return true
		}
	}
	return false
}
