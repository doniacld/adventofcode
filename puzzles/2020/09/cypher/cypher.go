package cypher

// Cypher holds the map for the preamble values and the list of numbers to parse
type Cypher struct {
	Preamble Preamble
	Numbers  []int
}

// Preamble holds pars to sum to get the target value
type Preamble map[int]struct{}

// newCypher returns an initialized Cypher
func newCypher() Cypher {
	return Cypher{newPreamble(), make([]int, 0)}
}

// newPreamble returns an initialized map
func newPreamble() Preamble {
	return make(map[int]struct{}, 0)
}

// insert a new element in Preamble
func (p Preamble) insert(elt int) {
	p[elt] = struct{}{}
}
