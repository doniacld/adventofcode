package bags

// Bags defines a map with the bag color as key and its family as value
type Bags map[string]bagFamily

// NewBags initializes the Bags map
func NewBags() Bags {
	return make(map[string]bagFamily, 0)
}

// bagFamily defines the list of parent's bag and children ones
type bagFamily struct {
	parents  []bag
	children []bag
}

// a bag is composed of a number and a coloured name
type bag struct {
	nb   int
	name string
}

// newBagFamily returns a new bagFamily
func newBagFamily() bagFamily {
	b := make([]bag, 0)
	return bagFamily{b, b}
}
