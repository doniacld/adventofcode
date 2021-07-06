package bags

// CountBags returns the content of a bag
func (r Bags) CountBags(target string) int {
	_, ok := r[target]
	if !ok {
		return 0
	}

	// browse the children of the target
	// decrease from one because of the target bag that is counted
	return r.countBags(target) - 1
}

// countBags counts the bags recursively
func (r Bags) countBags(target string) int {
	counter := 1
	v := r[target]
	for _, c := range v.children {
		counter += c.nb * r.countBags(c.name)
	}
	return counter
}
