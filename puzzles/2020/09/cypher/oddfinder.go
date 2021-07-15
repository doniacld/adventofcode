package cypher

import "math"

// Find the values that is not the sum of two values is the preamble map
// the map contains the x previous values
// x is a range that begins at 0 and ends at the given preamble index
func (c Cypher) FindOdd(idxEndPreamble int) int {
	idxStartPreamble := 0
	for i := idxEndPreamble; i < len(c.Numbers); i++ {

		elt := c.Numbers[i]
		if !c.findPair(elt) {
			return elt
		}

		c.updatePreamble(idxStartPreamble, idxEndPreamble)
		idxStartPreamble++
	}
	// should not happened
	return 0
}

// findPair returns the odd value if there is no pair summing it
func (c Cypher) findPair(elt int) bool {
	for p := range c.Preamble {
		toFind := elt - p

		f := math.Abs(float64(toFind))
		if _, ok := c.Preamble[int(f)]; ok {
			if toFind != p {
				return true
			}
		}
	}
	return false
}

// updatePreamble deletes the first value of the preamble
// and adds the current one to the preamble
func (c Cypher) updatePreamble(idxToRemove, idxPreamble int) {
	// remove the first element
	delete(c.Preamble, c.Numbers[idxToRemove])
	// add the next element
	c.Preamble[c.Numbers[idxToRemove+idxPreamble]] = struct{}{}
}

